use anyhow::Result;
use chromiumoxide::{Browser, BrowserConfig, Page};
use futures::StreamExt;

const BROWSER_ARGS: [&str; 16] = [
    "--profile-directory=Default",
    "--disable-gpu",
    "--disable-extensions",
    "--disable-dev-shm-usage",
    "--no-sandbox",
    "--disable-setuid-sandbox",
    "--disable-infobars",
    "--disable-notifications",
    "--disable-popup-blocking",
    "--disable-background-timer-throttling",
    "--disable-backgrounding-occluded-windows",
    "--disable-breakpad",
    "--disable-component-extensions-with-background-pages",
    "--disable-features=TranslateUI,BlinkGenPropertyTrees",
    "--disable-ipc-flooding-protection",
    "--disable-renderer-backgrounding",
];

pub async fn new() -> Browser {
    let (browser, mut handler) = Browser::launch(
        BrowserConfig::builder()
            .viewport(None)
            .args(BROWSER_ARGS)
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
