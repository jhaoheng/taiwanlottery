```
chrome_agent := chrome.NewAgent()
chrome_agent.Set_EnableWindow()
chrome_agent.Set_AcceptLanguage("zh-TW")
chrome_agent.Set_UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36")
defer chrome_agent.CloseAgent()
chrome_agent.RunWebDriver()
//
web_driver := chrome_agent.GetWebDriver()
// fmt.Println(web_driver.Capabilities())
```