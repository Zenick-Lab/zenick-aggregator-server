use strum::IntoStaticStr;

#[derive(Clone, Copy, IntoStaticStr)]
pub enum Operation {
    Borrow,
    Lend,
    Stake,
}
