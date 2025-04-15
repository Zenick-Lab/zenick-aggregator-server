use anyhow::Result;
use serde::Deserialize;

use crate::{database::History, operation::Operation, provider::Provider, token::Token};

const LINK: &str = "https://sdk.api.scallop.io/api/market/migrate";

#[derive(Debug, Deserialize)]
#[serde(rename_all = "camelCase")]
struct Coin {
    coin_name: String,
    supply_apr: f32,
    borrow_apr: f32,
}

#[derive(Debug, Deserialize)]
#[serde(rename_all = "camelCase")]
struct Market {
    pools: Vec<Coin>,
}

pub struct Scallop;

impl Scallop {
    pub async fn fetch() -> Result<impl Iterator<Item = History>> {
        let market_data = reqwest::get(LINK).await?.json::<Market>().await?;

        let data = market_data
            .pools
            .into_iter()
            .filter_map(
                |coin| match coin.coin_name.to_lowercase().parse::<Token>() {
                    Ok(token) => Some(
                        [
                            History {
                                provider: Provider::Scallop,
                                token,
                                operation: Operation::Borrow,
                                apr: coin.borrow_apr * 100.,
                            },
                            History {
                                provider: Provider::Scallop,
                                token,
                                operation: Operation::Lend,
                                apr: coin.supply_apr * 100.,
                            },
                        ]
                        .into_iter(),
                    ),
                    Err(_) => None,
                },
            )
            .flatten();

        Ok(data)
    }
}
