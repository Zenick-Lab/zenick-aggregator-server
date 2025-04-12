#[derive(Debug, Clone, Copy)]
pub enum Operation {
    Borrow,
    Lend,
    Stake,
}

impl Operation {
    pub const fn into_str(self) -> &'static str {
        match self {
            Operation::Borrow => "borrow",
            Operation::Lend => "lend",
            Operation::Stake => "stake"
        }
    }
}
