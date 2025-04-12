use std::str::FromStr;

use anyhow::bail;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
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

impl Token {
    pub const fn into_str(self) -> &'static str {
        match self {
            Token::Sui => "sui",
            Token::Usdc => "usdc"
        }
    }
}
