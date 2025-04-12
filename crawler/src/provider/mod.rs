mod navi;

use serde::{Deserialize, Serialize};
use strum::IntoStaticStr;

use crate::token::Token;

#[derive(Serialize, Deserialize)]
struct Data {
    pub token: Token,
    pub apr: f32,
}

#[derive(IntoStaticStr)]
pub enum Provider {
    Suilend,
    Navi,
    Cetus,
    Haedal,
    Scallop,
    Bluefin,
    Bucket,
    AlphaFi,
    AftermathFinance,
    KaiFinance,
    Kriya,
    Volosui,
}
