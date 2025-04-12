mod browser;
mod config;
mod database;
mod operation;
mod provider;
mod token;

use std::sync::Arc;

use anyhow::Result;
use chromiumoxide::Browser;
use database::History;
use operation::Operation;
use provider::navi::Navi;
use tokio::{sync::mpsc, task::JoinSet};

async fn into_task(
    fetch_func: impl AsyncFn(Operation, &Browser) -> Result<Vec<History>>,
    operation: Operation,
    browser: Arc<Browser>,
    history_sender: mpsc::UnboundedSender<History>,
) {
    let histories = fetch_func(operation, &browser).await.unwrap();
    for history in histories {
        history_sender.send(history).unwrap();
    }
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

    let mut join_set = JoinSet::new();

    join_set.spawn(into_task(
        Navi::fetch,
        Operation::Lend,
        browser,
        history_sender.clone(),
    ));
    drop(history_sender);

    join_set.join_all().await;
    history_task.await?;

    Ok(())
}
