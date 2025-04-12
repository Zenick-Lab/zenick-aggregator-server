use strum::IntoStaticStr;

#[derive(IntoStaticStr)]
pub enum Operation {
    Borrow,
    Lend,
    Stake,
}
