package main

import (
	"encoding/json"
	"fmt"
	"jhaoheng/taiwanlottery/chrome"
	"jhaoheng/taiwanlottery/crawler"
)

func main() {
	chrome_agent := chrome.NewAgent()
	defer chrome_agent.CloseAgent()
	chrome_agent.RunWebDriver()
	//
	web_driver := chrome_agent.GetWebDriver()

	//
	result := crawler.NewLotto649(web_driver).SearchBySerialID("112000018")
	b, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(b))

	//
	var stop bool
	fmt.Scanln(&stop)
}
