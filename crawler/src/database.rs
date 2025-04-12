use sqlx::{PgPool, Result};

use crate::{config::CONFIG, operation::Operation, provider::Provider, token::Token};

#[derive(Debug)]
pub struct History {
    pub provider: Provider,
    pub token: Token,
    pub operation: Operation,
    pub apr: f32,
}

pub async fn new() -> Result<PgPool> {
    let database = PgPool::connect(&CONFIG.database_url).await?;

    Ok(database)
}

pub async fn insert_history(
    history: History,
    database: &PgPool
) -> Result<()> {
    sqlx::query!(
        r#"
            INSERT INTO histories(
                provider_id,
                token_id,
                operation_id,
                apr
            ) VALUES (
                (SELECT id FROM providers WHERE name = $1),
                (SELECT id FROM tokens WHERE name = $2),
                (SELECT id FROM operations WHERE name = $3),
                $4
            )
        "#,
        history.provider.into_str(),
        history.token.into_str(),
        history.operation.into_str(),
        history.apr,
    )
    .execute(database)
    .await?;

    Ok(())
}
