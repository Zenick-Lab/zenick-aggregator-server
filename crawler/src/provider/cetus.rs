use anyhow::Result;
use serde::Deserialize;

use crate::{database::LiquidityPoolHistory, provider::Provider, token::Token, util};

const LINK: &str = "https://api-sui.cetus.zone/v2/sui/stats_pools?is_vaults=false&display_all_pools=false&has_mining=true&has_farming=true&no_incentives=true&order_by=-vol&limit=20&offset=0&coin_type=&pool=";

#[derive(Debug, Deserialize)]
struct Coin {
    symbol: String,
}

#[derive(Debug, Deserialize)]
struct LiquidityPool {
    coin_a: Coin,
    coin_b: Coin,
    total_apr: String,
}

#[derive(Debug, Deserialize)]
struct LiquidityPoolList {
    lp_list: Vec<LiquidityPool>,
}

#[derive(Debug, Deserialize)]
struct ApiResponse {
    data: LiquidityPoolList,
}

pub struct Cetus;

impl Cetus {
    pub async fn fetch() -> Result<impl Iterator<Item = LiquidityPoolHistory>> {
        let response = reqwest::get(LINK).await?.json::<ApiResponse>().await?;

        let data = response.data.lp_list.into_iter().filter_map(|pool| {
            match (
                pool.coin_a.symbol.parse::<Token>(),
                pool.coin_b.symbol.parse::<Token>(),
            ) {
                (Ok(token_a), Ok(token_b)) => Some(LiquidityPoolHistory {
                    provider: Provider::Cetus,
                    token_a,
                    token_b,
                    apr: util::parse_float(&pool.total_apr).ok()? * 100.,
                }),
                _ => None,
            }
        });

        Ok(data)
    }
}
