mod browser;
mod config;
mod database;
mod operation;
mod provider;
mod token;
mod util;

use std::sync::Arc;

use anyhow::Result;
use chromiumoxide::Browser;
use database::History;
use futures::{Stream, StreamExt};
use provider::{haedal::Haedal, navi::Navi, scallop::Scallop, suilend::Suilend};
use tokio::{sync::mpsc, task::JoinSet};
use tracing::level_filters::LevelFilter;
use tracing_subscriber::{
    EnvFilter, Layer, filter, fmt::time::ChronoLocal, layer::SubscriberExt, util::SubscriberInitExt,
};

fn send_sync<Item>(
    histories: impl Iterator<Item = Item>,
    history_sender: mpsc::UnboundedSender<Item>,
) {
    for history in histories {
        if let Err(error) = history_sender.send(history) {
            tracing::error!("{:?}", error);
        }
    }
}

async fn send<Item>(
    histories: impl Stream<Item = Item>,
    history_sender: mpsc::UnboundedSender<Item>,
) {
    histories
        .for_each(move |history| {
            let history_sender = history_sender.clone();
            async move {
                if let Err(error) = history_sender.send(history) {
                    tracing::error!("{:?}", error);
                }
            }
        })
        .await;
}

async fn spawn_tasks(browser: Arc<Browser>, history_sender: mpsc::UnboundedSender<History>) {
    let mut join_set = JoinSet::new();

    let b = browser.clone();
    let tx = history_sender.clone();
    join_set.spawn(async move {
        let histories = Suilend::fetch(&b).await?;
        send(histories, tx).await;

        Ok::<_, anyhow::Error>(())
    });

    let b = browser.clone();
    let tx = history_sender.clone();
    join_set.spawn(async move {
        let histories = Navi::fetch(&b).await?;
        send(histories, tx).await;

        Ok::<_, anyhow::Error>(())
    });

    let b = browser.clone();
    let tx = history_sender.clone();
    join_set.spawn(async move {
        let histories = Haedal::fetch(&b).await?;
        send_sync(histories, tx);

        Ok::<_, anyhow::Error>(())
    });

    let tx = history_sender.clone();
    join_set.spawn(async move {
        let histories = Scallop::fetch().await?;
        send_sync(histories, tx);

        Ok::<_, anyhow::Error>(())
    });

    join_set.join_all().await;
}

#[tokio::main]
async fn main() -> Result<()> {
    let filter = filter::filter_fn(|metadata| !metadata.target().contains("chromiumoxide"));

    tracing_subscriber::registry()
        .with(
            tracing_subscriber::fmt::layer()
                .pretty()
                .with_timer(ChronoLocal::rfc_3339())
                .with_filter(filter),
        )
        .with(
            EnvFilter::builder()
                .with_default_directive(LevelFilter::INFO.into())
                .from_env_lossy(),
        )
        .init();

    let database = database::new().await?;
    let (history_sender, mut history_receiver) = mpsc::unbounded_channel::<History>();
    let history_task = tokio::spawn(async move {
        while let Some(history) = history_receiver.recv().await {
            if let Err(error) = database::insert_history(history, &database).await {
                tracing::error!("{:?}", error);
            }
        }
    });

    let browser = Arc::new(browser::new().await);

    spawn_tasks(browser, history_sender).await;

    history_task.await?;

    Ok(())
}
