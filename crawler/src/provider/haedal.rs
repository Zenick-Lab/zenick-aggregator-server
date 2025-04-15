use anyhow::{Context, Result};
use chromiumoxide::Browser;

use crate::{browser, database::History, operation::Operation, token::Token, util};

use super::Provider;

const LINK: &str = "https://haedal.xyz/stake";

const STAKE_APR_SELECTOR: &str =
    "div.rounded-xl:nth-child(4) > div:nth-child(3) > span:nth-child(2)";

pub struct Haedal;

impl Haedal {
    pub async fn fetch(browser: &Browser) -> Result<impl Iterator<Item = History>> {
        let page = browser::create_steath_page(browser).await?;

        page.goto(LINK).await?;
        page.wait_for_navigation().await?;

        let elements = util::find_elements(STAKE_APR_SELECTOR, &page).await?;

        let stake_apr = &elements[0];
        let stake_apr = stake_apr.inner_text().await?.context("No apr")?;
        let stake_apr = util::parse_float(&stake_apr)?;

        Ok([History {
            provider: Provider::Haedal,
            token: Token::Sui,
            operation: Operation::Stake,
            apr: stake_apr,
        }]
        .into_iter())
    }
}
