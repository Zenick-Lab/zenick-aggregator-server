use std::str::FromStr;

use anyhow::{Context, Error, Result};
use chromiumoxide::Element;

pub async fn parse<T>(element: Element) -> Result<T>
where
    T: FromStr,
    <T as FromStr>::Err: std::error::Error,
    <T as FromStr>::Err: std::marker::Send,
    <T as FromStr>::Err: std::marker::Sync,
    <T as FromStr>::Err: 'static,
{
    let raw = element.inner_text().await?;
    let raw = raw.context("Missing value")?;

    raw.parse::<T>().map_err(Error::from)
}
