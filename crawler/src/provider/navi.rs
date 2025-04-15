use anyhow::{Context, Result, ensure};
use chromiumoxide::{Browser, Element};
use futures::{stream::{self, FuturesUnordered}, Stream, StreamExt};

use crate::{browser, database::History, operation::Operation, token::Token, util};

use super::Provider;

const LINK: &str = "https://app.naviprotocol.io/market";

const ROW_SELECTOR: &str = "a.MuiTypography-root > tr";
const TOKEN_SELECTOR: &str =
    "td:nth-child(1) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > p:nth-child(1)";
const LEND_APR_SELECTOR: &str = "td:nth-child(3)  div:not(:has(div))";
const BORROW_APR_SELECTOR: &str = "td:nth-child(5) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1), td:nth-child(5) > p:nth-child(1)";

pub struct Navi;

impl Navi {
    async fn get_row_data(row: Element) -> Result<impl Stream<Item = History>> {
        let token = row
            .find_element(TOKEN_SELECTOR)
            .await?
            .inner_text()
            .await?
            .context("No token")?;
        let token = token.parse::<Token>()?;

        let lend_apr = row
            .find_element(LEND_APR_SELECTOR)
            .await?
            .inner_text()
            .await?
            .context("No lend apr")?;
        ensure!(!lend_apr.is_empty(), "No lend apr");
        let lend_apr = util::parse_float(&lend_apr)?;

        let borrow_apr = row
            .find_element(BORROW_APR_SELECTOR)
            .await?
            .inner_text()
            .await?
            .context("No borrow apr")?;
        ensure!(!borrow_apr.is_empty(), "No borrow apr");
        let borrow_apr = util::parse_float(&borrow_apr)?;

        Ok(stream::iter([
            History {
                provider: Provider::Navi,
                token,
                operation: Operation::Lend,
                apr: lend_apr,
            },
            History {
                provider: Provider::Navi,
                token,
                operation: Operation::Borrow,
                apr: borrow_apr,
            },
        ]))
    }

    pub async fn fetch(browser: &Browser) -> Result<impl Stream<Item = History>> {
        let page = browser::create_steath_page(browser).await?;

        page.goto(LINK).await?;
        page.wait_for_navigation().await?;

        let rows = util::find_elements(ROW_SELECTOR, &page).await?;
        let data = rows
            .into_iter()
            .map(Self::get_row_data)
            .collect::<FuturesUnordered<_>>()
            .filter_map(|x| async move { x.ok() })
            .flatten_unordered(None);

        Ok(data)
    }
}
