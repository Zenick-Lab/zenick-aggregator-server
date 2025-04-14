use anyhow::Result;
use chromiumoxide::Browser;
use futures::{StreamExt, stream};
use itertools::multizip;

use crate::{browser, database::History, operation::Operation, token::Token, util};

use super::Provider;

const LINK: &str = "https://suilend.fi";

const TOKEN_SELECTOR: &str = "tr.border-b > td:nth-child(1) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > p:nth-child(1)";
const LEND_APR_SELECTOR: &str = "tr.border-b > td:nth-child(5) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1) > p:nth-child(2)";
const BORROW_APR_SELECTOR: &str = "tr.border-b > td:nth-child(6) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1) > p:nth-child(2)";

pub struct Suilend;

impl Suilend {
    pub async fn fetch(browser: &Browser) -> Result<Vec<History>> {
        let page = browser::create_steath_page(browser).await?;

        page.goto(LINK).await?;
        page.wait_for_navigation().await?;

        let tokens = util::find_elements(TOKEN_SELECTOR, &page).await?;
        let lend_aprs = util::find_elements(LEND_APR_SELECTOR, &page).await?;
        let borrow_aprs = util::find_elements(BORROW_APR_SELECTOR, &page).await?;

        let data = stream::iter(multizip((tokens, lend_aprs, borrow_aprs)))
            .map(|(token, lend_apr, borrow_apr)| async move {
                let token = token.inner_text().await.unwrap().unwrap();

                let raw = lend_apr.inner_text().await.unwrap().unwrap();
                let lend_apr = util::parse_float(&raw).unwrap();

                let raw = borrow_apr.inner_text().await.unwrap().unwrap();
                let borrow_apr = util::parse_float(&raw).unwrap();

                (token, lend_apr, borrow_apr)
            })
            .buffer_unordered(20)
            .filter_map(|(token, lend_apr, borrow_apr)| async move {
                match token.parse::<Token>() {
                    Ok(token) => Some(stream::iter([
                        History {
                            provider: Provider::Suilend,
                            operation: Operation::Lend,
                            token,
                            apr: lend_apr,
                        },
                        History {
                            provider: Provider::Suilend,
                            operation: Operation::Borrow,
                            token,
                            apr: borrow_apr,
                        },
                    ])),
                    Err(_) => None,
                }
            })
            .flatten()
            .collect()
            .await;

        eprintln!("DEBUGPRINT[143]: suilend.rs:67: data={:#?}", data);
        Ok(data)
    }
}
