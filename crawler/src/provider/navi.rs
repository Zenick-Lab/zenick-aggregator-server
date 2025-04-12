use anyhow::Result;
use chromiumoxide::Page;
use futures::{StreamExt, stream};

use crate::{operation::Operation, token::Token, util};

use super::Data;

const TOKEN_SELECTOR: &str = "tr.MuiTableRow-root > td:nth-child(1) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > p:nth-child(1)";
const APR_SELECTOR: &str = ".MuiTableBody-root > tr > td:nth-child(3) p:nth-child(1)";

pub struct Navi;

impl Navi {
    const fn get_link(operation: Operation) -> &'static str {
        match operation {
            Operation::Borrow => "https://app.naviprotocol.io/borrow",
            Operation::Lend => "https://app.naviprotocol.io",
            Operation::Stake => unreachable!(),
        }
    }

    pub async fn fetch(operation: Operation, page: &Page) -> Result<Vec<Data>> {
        let link = Self::get_link(operation);
        page.goto(link).await?;

        let (tokens, aprs) = tokio::try_join!(
            page.find_elements(TOKEN_SELECTOR),
            page.find_elements(APR_SELECTOR),
        )?;

        let tokens = stream::iter(tokens)
            .map(util::parse::<String>)
            .buffer_unordered(20)
            .map(|token| token.unwrap());
        let aprs = stream::iter(aprs)
            .map(util::parse::<f32>)
            .buffer_unordered(20)
            .map(|apr| apr.unwrap());

        let data = tokens
            .zip(aprs)
            .filter_map(|(token, apr)| async move {
                token.parse::<Token>().map(|token| Data { token, apr }).ok()
            })
            .collect()
            .await;

        Ok(data)
    }
}
