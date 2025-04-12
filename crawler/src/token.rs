use std::str::FromStr;

use anyhow::bail;
use serde::{Deserialize, Serialize};
use strum::IntoStaticStr;

#[derive(IntoStaticStr, Serialize, Deserialize)]
pub enum Token {
    Sui,
    Usdc,
}

impl FromStr for Token {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s.to_lowercase().as_str() {
            "sui" => Ok(Token::Sui),
            "usdc" => Ok(Token::Usdc),
            _ =>  bail!("invalid token name")
        }
    }
}
