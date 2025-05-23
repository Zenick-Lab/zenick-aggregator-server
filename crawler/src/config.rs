use std::sync::LazyLock;

use serde::Deserialize;

#[derive(Deserialize)]
pub struct Config {
    pub database_url: String,
}

pub static CONFIG: LazyLock<Config> = LazyLock::new(|| {
    ::config::Config::builder()
        .add_source(::config::Environment::default().try_parsing(true))
        .build()
        .unwrap()
        .try_deserialize()
        .unwrap()
});
