use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use strum::IntoStaticStr;

#[derive(Serialize, Deserialize)]
struct History {
    pub apr: f32,
    pub created_at: DateTime<Utc>,
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
