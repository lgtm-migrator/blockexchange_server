const {Builder} = require('selenium-webdriver');

(async function example() {
  let driver = await new Builder()
    .usingServer('http://blockexchange_selenium:4444/wd/hub')
    .forBrowser('chrome')
    .build();
  try {
    await driver.get('http://blockexchange_server:8080/');
    //await driver.findElement(By.name('q')).sendKeys('webdriver', Key.RETURN);
    //await driver.wait(until.titleIs('webdriver - Google Search'), 1000);
  } finally {
    await driver.quit();
    console.log("e2e tests done!");
  }
})();