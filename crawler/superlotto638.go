package crawler

import (
	"time"

	"github.com/tebeka/selenium"
)

/*
- 威力彩
*/

type Superlotto638Result struct {
	SerialID           string
	Date               string // ex: 2006/02/02
	Num_1              string
	Num_2              string
	Num_3              string
	Num_4              string
	Num_5              string
	Num_6              string
	Num_second_section string // 第二區
}

type ISuperlotto638 interface {
	// 使用`期別`查詢, ex: 112000029
	SearchBySerialID(sid string) Superlotto638Result
}

type Superlotto638 struct {
	URL       string
	WebDriver selenium.WebDriver
}

func NewSuperlotto638(web_driver selenium.WebDriver) ISuperlotto638 {
	return &Superlotto638{
		URL:       "https://www.taiwanlottery.com.tw/lotto/superlotto638/history.aspx",
		WebDriver: web_driver,
	}
}

func (lo *Superlotto638) SearchBySerialID(sid string) Superlotto638Result {
	//
	lo.WebDriver.Get(lo.URL)
	time.Sleep(time.Second * 1)
	//
	input, err := lo.WebDriver.FindElement(selenium.ByCSSSelector, "#SuperLotto638Control_history1_txtNO")
	if err != nil {
		panic(err)
	}
	input.SendKeys(sid)
	input.SendKeys(selenium.EnterKey)
	//

	result := Superlotto638Result{
		SerialID:           lo.get_text("#SuperLotto638Control_history1_dlQuery_DrawTerm_0"),
		Date:               lo.get_text("#SuperLotto638Control_history1_dlQuery_Date_0"),
		Num_1:              lo.get_text("#SuperLotto638Control_history1_dlQuery_No1_0"),
		Num_2:              lo.get_text("#SuperLotto638Control_history1_dlQuery_No2_0"),
		Num_3:              lo.get_text("#SuperLotto638Control_history1_dlQuery_No3_0"),
		Num_4:              lo.get_text("#SuperLotto638Control_history1_dlQuery_No4_0"),
		Num_5:              lo.get_text("#SuperLotto638Control_history1_dlQuery_No5_0"),
		Num_6:              lo.get_text("#SuperLotto638Control_history1_dlQuery_No6_0"),
		Num_second_section: lo.get_text("#SuperLotto638Control_history1_dlQuery_No7_0"),
	}
	return result
}

// -
func (lo *Superlotto638) get_text(key string) string {
	elem, err := lo.WebDriver.FindElement(selenium.ByCSSSelector, key)
	if err != nil {
		panic(err)
	}
	text, _ := elem.Text()
	return text
}
