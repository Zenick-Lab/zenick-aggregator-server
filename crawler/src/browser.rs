use anyhow::Result;
use chromiumoxide::{Browser, BrowserConfig, Page};
use futures::StreamExt;

pub async fn new() -> Browser {
    let (browser, mut handler) = Browser::launch(
        BrowserConfig::builder()
            .viewport(None)
            .with_head()
            .build()
            .unwrap(),
    )
    .await
    .unwrap();

    tokio::spawn(async move {
        loop {
            let _ = handler.next().await;
        }
    });

    browser
}

pub async fn create_steath_page(browser: &Browser) -> Result<Page> {
    let page = browser.new_page("about:blank").await?;
    page.enable_stealth_mode().await?;

    Ok(page)
}
