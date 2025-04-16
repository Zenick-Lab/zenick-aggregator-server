use sqlx::{PgPool, Result};

use crate::{config::CONFIG, operation::Operation, provider::Provider, token::Token};

pub async fn new() -> Result<PgPool> {
    let database = PgPool::connect(&CONFIG.database_url).await?;

    Ok(database)
}

#[derive(Debug)]
pub struct History {
    pub provider: Provider,
    pub token: Token,
    pub operation: Operation,
    pub apr: f32,
}

pub async fn insert_history(history: History, database: &PgPool) -> Result<()> {
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

#[derive(Debug)]
pub struct LiquidityPoolHistory {
    pub provider: Provider,
    pub token_a: Token,
    pub token_b: Token,
    pub apr: f32,
}

pub async fn insert_liquidity_pool_history(
    history: LiquidityPoolHistory,
    database: &PgPool,
) -> Result<()> {
    sqlx::query!(
        r#"
            INSERT INTO liquidity_pool_histories(
                provider_id,
                token_a_id,
                token_b_id,
                apr
            ) VALUES (
                (SELECT id FROM providers WHERE name = $1),
                (SELECT id FROM tokens WHERE name = $2),
                (SELECT id FROM tokens WHERE name = $3),
                $4
            )
        "#,
        history.provider.into_str(),
        history.token_a.into_str(),
        history.token_b.into_str(),
        history.apr
    )
    .execute(database)
    .await?;

    Ok(())
}
