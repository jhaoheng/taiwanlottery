## example

- 初始化 chrome

```
chrome_agent := chrome.NewAgent()
chrome_agent.Set_EnableWindow()
defer chrome_agent.CloseAgent()
chrome_agent.RunWebDriver()
//
web_driver := chrome_agent.GetWebDriver()
// fmt.Println(web_driver.Capabilities())
```

- 爬資料

```
result, err := NewLotto649(web_driver).SearchBySerialID("112000018")
```