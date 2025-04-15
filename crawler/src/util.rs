use std::{num::ParseFloatError, time::Duration};

use anyhow::{Result, bail};
use backon::{ExponentialBuilder, Retryable};
use chromiumoxide::{Element, Page};

pub fn parse_float(raw: &str) -> Result<f32, ParseFloatError> {
    let filtered: String = raw
        .chars()
        .filter(|&c| c.is_numeric() || c == '.')
        .collect();

    filtered.parse()
}

pub async fn find_elements(selector: &str, page: &Page) -> Result<Vec<Element>> {
    let find = || async move {
        let elements = page.find_elements(selector).await?;

        if elements.is_empty() {
            bail!("Empty elements")
        }

        Ok(elements)
    };

    find.retry(ExponentialBuilder::default().without_max_times())
        .sleep(tokio::time::sleep)
        .notify(|err, dur: Duration| {
            eprintln!("retrying {:?} after {:?}", err, dur);
        })
        .await
}
