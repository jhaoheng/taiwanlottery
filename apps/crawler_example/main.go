package main

import (
	"encoding/json"
	"fmt"
	"jhaoheng/taiwanlottery/chrome"
	"jhaoheng/taiwanlottery/crawler"

	"github.com/sirupsen/logrus"
)

/*
- 取得資料庫中, 最新的 serial_id
111000114
112000001 - 112000039

1. 先取得現在網頁中的資料 (期別)
2. 計算出第一期的數字
3. 搜尋資料庫中, 是否有相關資料

應該開發功能列
- 從資料庫中最後一筆搜尋, 如果找不到資料, 則進入到下一個年度
- 從今年開始
- 指定期數
- 範圍期數
*/

func main() {
	chrome.Default_Dir = "../../"
	chrome_agent := chrome.NewAgent()
	// chrome_agent.Set_EnableWindow()
	defer chrome_agent.CloseAgent()
	chrome_agent.RunWebDriver()
	//
	web_driver := chrome_agent.GetWebDriver()

	//
	result, err := crawler.NewLotto649(web_driver).SearchBySerialID("111000114")
	if err != nil {
		logrus.Fatal(err)
	}
	b, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(b))

	// //
	// var stop bool
	// fmt.Scanln(&stop)
}
