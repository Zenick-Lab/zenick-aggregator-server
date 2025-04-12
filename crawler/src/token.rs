use strum::IntoStaticStr;

#[derive(IntoStaticStr)]
pub enum Token {
    Sui,
    USDC,
}
