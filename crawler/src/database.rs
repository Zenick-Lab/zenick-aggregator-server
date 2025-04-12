use chrono::{DateTime, Utc};
use sqlx::{PgPool, Result};

use crate::{operation::Operation, provider::Provider, token::Token};

pub struct Database {
    pool: PgPool,
}

impl Database {
    pub async fn new(database_url: &str) -> Result<Self> {
        let pool = PgPool::connect(database_url).await?;

        Ok(Self { pool })
    }

    pub async fn insert_history(
        &self,
        provider: Provider,
        token: Token,
        operation: Operation,
        apr: f32,
        created_at: DateTime<Utc>,
    ) -> Result<()> {
        let provider: &'static str = provider.into();
        let token: &'static str = token.into();
        let operation: &'static str = operation.into();

        sqlx::query!(
            r#"
                INSERT INTO histories(
                    provider_id,
                    token_id,
                    operation_id,
                    apr,
                    created_at
                ) VALUES (
                    (SELECT id FROM providers WHERE name = $1),
                    (SELECT id FROM tokens WHERE name = $2),
                    (SELECT id FROM operations WHERE name = $3),
                    $4,
                    $5
                )
            "#,
            provider,
            token,
            operation,
            apr,
            created_at.naive_utc()
        )
        .execute(&self.pool)
        .await?;

        Ok(())
    }
}
