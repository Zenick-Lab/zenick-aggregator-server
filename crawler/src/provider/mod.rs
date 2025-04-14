pub mod navi;
pub mod suilend;

#[derive(Debug, Clone, Copy)]
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

impl Provider {
    pub const fn into_str(self) -> &'static str {
        match self {
            Provider::Suilend => "suilend",
            Provider::Navi => "navi",
            Provider::Cetus => "cetus",
            Provider::Haedal => "haedal",
            Provider::Scallop => "scallop",
            Provider::Bluefin => "bluefin",
            Provider::Bucket => "bucket",
            Provider::AlphaFi => "alpha_fi",
            Provider::AftermathFinance => "aftermath_finance",
            Provider::KaiFinance => "kai_finance",
            Provider::Kriya => "kriya",
            Provider::Volosui => "volosui",
        }
    }
}
