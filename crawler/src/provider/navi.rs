use anyhow::Result;
use chromiumoxide::Browser;
use futures::{StreamExt, stream};

use crate::{browser, database::History, operation::Operation, token::Token};

use super::Provider;

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

    pub async fn fetch(operation: Operation, browser: &Browser) -> Result<Vec<History>> {
        let page = browser::create_steath_page(browser).await?;

        let link = Self::get_link(operation);
        page.goto(link).await?;
        page.wait_for_navigation().await?;

        let tokens = page.find_elements(TOKEN_SELECTOR).await?;
        let aprs = page.find_elements(APR_SELECTOR).await?;

        let tokens = stream::iter(tokens)
            .map(|token| async move { token.inner_text().await.unwrap().unwrap() })
            .buffer_unordered(20);
        let aprs = stream::iter(aprs)
            .map(|apr| async move {
                let apr_raw = apr.inner_text().await.unwrap().unwrap();
                apr_raw[0..apr_raw.len() - 1].parse::<f32>().unwrap()
            })
            .buffer_unordered(20);

        let data = tokens
            .zip(aprs)
            .filter_map(|(token, apr)| async move {
                token
                    .parse::<Token>()
                    .map(|token| History {
                        provider: Provider::Navi,
                        operation,
                        token,
                        apr,
                    })
                    .ok()
            })
            .collect()
            .await;

        Ok(data)
    }
}
