package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jhaoheng/taiwanlottery/chrome"
	"github.com/jhaoheng/taiwanlottery/crawler"
	"github.com/jhaoheng/taiwanlottery/model"

	"github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
	"gorm.io/gorm"
)

/*
flow
- 取得網路上目前最新資料
- 取得資料庫中最後資料
- 驗證儲存數據是否最新
- 爬取新資料
- 寫入資料庫
*/

func init() {
	if err := model.ConnSQLite("../../sql.db"); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	bot := NewBot()
	defer bot.CloseAgent()

	/*
		- 取得網路上目前最新資料
	*/
	the_latest_data, err := bot.GetLatestData()
	if err != nil {
		panic(err)
	}
	fmt.Println("目前網路上最新資料為 ....")
	b, _ := json.MarshalIndent(the_latest_data, "", "	")
	fmt.Println(string(b))
	fmt.Println()

	/*
		- 取得資料庫中最後資料
	*/
	db_latest_result, err := model.NewLottery().SetCategory(model.Lotto649).OrderByDESC("date").Take()
	if err != nil {
		panic(err)
	}
	fmt.Println("資料庫最新資料為 ....")
	b, _ = json.MarshalIndent(db_latest_result, "", "	")
	fmt.Println(string(b))
	fmt.Println()

	// 取得目前資料總數
	fmt.Printf("目前資料庫, %v, 資料庫量為 %v\n", model.Lotto649, func() int {
		results, err := model.NewLottery().SetCategory(model.Lotto649).FindAll()
		if err != nil {
			panic(err)
		}
		return len(results)
	}())

	/*
		- 驗證儲存數據是否最新
	*/
	if db_latest_result.SerialID == the_latest_data.SerialID {
		fmt.Println("資料庫資料與網路最新資料相同, 中止")
		return
	}
	fmt.Println("資料庫資料與網路最新資料不同, 執行爬取資料....")

	/*
		- 爬取資料
	*/
	the_crawl_datas := []crawler.Lotto649Result{}
	var next_serial_id string = ""
	for {
		// 取得下一筆資料
		next_serial_id = func() string {
			if len(next_serial_id) == 0 {
				return GetNestSerialID(db_latest_result.SerialID)
			}
			return GetNestSerialID(next_serial_id)
		}()
		fmt.Printf("=== 預估下一筆 serial_id => %v ===\n", next_serial_id)

		/*
			取得資料, 失敗
			- 連續失敗一次 -> 重複一次
			- 連續失敗兩次 -> 換到下一年的第一筆開始搜尋
			- 連續失敗三次以上 -> 停止

			取得資料, 成功
			- 與最新資料相同 -> 停止
		*/
		continue_failure := 0
	RETRY:
		the_data, err := bot.GetData(next_serial_id)
		if err != nil {
			if continue_failure == 0 {
				fmt.Println("**失敗第一次**")
				continue_failure++
				goto RETRY
			} else if continue_failure == 1 {
				fmt.Println("**失敗第二次**")
				continue_failure++
				// 跳到次年的第一筆(期數)
				year_str := db_latest_result.Date.AddDate(1, 0, 0).Format("2006")
				year_int, _ := strconv.Atoi(year_str)
				ROC_YEAR := year_int - 1911
				next_serial_id = fmt.Sprintf("%v000001", ROC_YEAR)
				fmt.Println("下一年的期數 =>", next_serial_id)
				goto RETRY
			} else {
				panic(err)
			}
		}

		b, _ := json.MarshalIndent(the_data, "", "	")
		fmt.Println(string(b))

		the_crawl_datas = append(the_crawl_datas, *the_data)

		if the_data.SerialID == the_latest_data.SerialID {
			break
		}
	}

	//
	fmt.Printf("爬蟲總共收集 => %v 筆資料\n", len(the_crawl_datas))

	/*
		- 寫入資料庫
	*/
	for _, crawl_data := range the_crawl_datas {
		result, err := model.NewLottery().SetCategory(model.Lotto649).SetSerialID(crawl_data.SerialID).Take()
		if err != nil && err != gorm.ErrRecordNotFound {
			panic(err)
		}
		if len(result.SerialID) != 0 {
			continue
		}
		tx := model.NewLottery().
			SetCategory(model.Lotto649).
			SetSerialID(crawl_data.SerialID).
			SetBallNumbers(func() json.RawMessage {
				nums := model.Lotto649Nums{
					Num_1:      crawl_data.Num_1,
					Num_2:      crawl_data.Num_2,
					Num_3:      crawl_data.Num_3,
					Num_4:      crawl_data.Num_4,
					Num_5:      crawl_data.Num_5,
					Num_6:      crawl_data.Num_6,
					NumSpecial: crawl_data.Num_special,
				}
				b, _ := json.Marshal(nums)
				return b
			}()).
			SetDate(crawl_data.Date)
		if _, err := tx.Create(); err != nil {
			panic(err)
		}
	}

	// 取得目前資料總數
	fmt.Printf("目前資料庫, %v, 資料庫量為 %v\n", model.Lotto649, func() int {
		results, err := model.NewLottery().SetCategory(model.Lotto649).FindAll()
		if err != nil {
			panic(err)
		}
		return len(results)
	}())
}

// 取得下一個 sid
func GetNestSerialID(serial_id string) string {
	i, err := strconv.Atoi(serial_id)
	if err != nil {
		panic(err)
	}
	i = i + 1
	next_serial_id := strconv.Itoa(i)
	return next_serial_id
}

/**/
type Bot struct {
	ChromeAgent chrome.IChromeDriver
	WebDriver   selenium.WebDriver
}

func NewBot() *Bot {
	chrome.Default_Dir = "../../"
	chrome_agent := chrome.NewAgent()
	chrome_agent.RunWebDriver()
	web_driver := chrome_agent.GetWebDriver()
	return &Bot{
		ChromeAgent: chrome_agent,
		WebDriver:   web_driver,
	}
}

// -
func (bot *Bot) GetLatestData() (result *crawler.Lotto649Result, err error) {
	result, err = crawler.NewLotto649(bot.WebDriver).GetLatestData()
	return
}

// -
func (bot *Bot) GetData(serial_id string) (result *crawler.Lotto649Result, err error) {
	result, err = crawler.NewLotto649(bot.WebDriver).SearchBySerialID(serial_id)
	if err != nil {
		return nil, err
	}
	// var stop bool
	// fmt.Scanln(&stop)
	return
}

func (bot *Bot) CloseAgent() {
	bot.ChromeAgent.CloseAgent()
}
