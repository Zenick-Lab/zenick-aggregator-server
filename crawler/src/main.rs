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
use provider::{navi::Navi, suilend::Suilend};
use tokio::{sync::mpsc, task::JoinSet};

async fn into_task(
    fetch_func: impl AsyncFn(&Browser) -> Result<Vec<History>>,
    browser: Arc<Browser>,
    history_sender: mpsc::UnboundedSender<History>,
) {
    let histories = fetch_func(&browser).await.unwrap();
    for history in histories {
        history_sender.send(history).unwrap();
    }
}

async fn spawn_tasks(browser: Arc<Browser>, history_sender: mpsc::UnboundedSender<History>) {
    let mut join_set = JoinSet::new();

    join_set.spawn(into_task(
        Navi::fetch,
        browser.clone(),
        history_sender.clone(),
    ));

    join_set.spawn(into_task(
        Suilend::fetch,
        browser.clone(),
        history_sender.clone(),
    ));

    join_set.join_all().await;
}

#[tokio::main]
async fn main() -> Result<()> {
    let database = database::new().await?;
    let (history_sender, mut history_receiver) = mpsc::unbounded_channel::<History>();
    let history_task = tokio::spawn(async move {
        while let Some(history) = history_receiver.recv().await {
            if let Err(error) = database::insert_history(history, &database).await {
                eprint!("{:?}", error);
            }
        }
    });

    let browser = Arc::new(browser::new().await);

    spawn_tasks(browser, history_sender).await;

    history_task.await?;

    Ok(())
}
