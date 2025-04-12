mod config;
mod database;
mod provider;
mod token;
mod operation;
mod util;

use anyhow::Result;
use config::CONFIG;
use database::Database;

#[tokio::main]
async fn main() -> Result<()> {
    let database = Database::new(&CONFIG.database_url);

    Ok(())
}
