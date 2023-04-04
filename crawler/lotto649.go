package crawler

import (
	"time"

	"github.com/tebeka/selenium"
)

/*
- 大樂透
*/

type Lotto649Result struct {
	SerialID    string
	Date        string // ex: 2006/02/02
	Num_1       string
	Num_2       string
	Num_3       string
	Num_4       string
	Num_5       string
	Num_6       string
	Num_special string // 特別號
}

type ILotto649 interface {
	// 使用`期別`查詢, ex: 112000029
	SearchBySerialID(sid string) Lotto649Result
}

type Lotto649 struct {
	URL       string
	WebDriver selenium.WebDriver
}

func NewLotto649(web_driver selenium.WebDriver) ILotto649 {
	return &Lotto649{
		URL:       "https://www.taiwanlottery.com.tw/lotto/Lotto649/history.aspx",
		WebDriver: web_driver,
	}
}

// -
func (lo *Lotto649) SearchBySerialID(sid string) Lotto649Result {
	//
	lo.WebDriver.Get(lo.URL)
	time.Sleep(time.Second * 1)
	//
	input, err := lo.WebDriver.FindElement(selenium.ByCSSSelector, "#Lotto649Control_history_txtNO")
	if err != nil {
		panic(err)
	}
	input.SendKeys(sid)
	input.SendKeys(selenium.EnterKey)
	//

	result := Lotto649Result{
		SerialID:    lo.get_text("#Lotto649Control_history_dlQuery_L649_DrawTerm_0"),
		Date:        lo.get_text("#Lotto649Control_history_dlQuery_L649_DDate_0"),
		Num_1:       lo.get_text("#Lotto649Control_history_dlQuery_No1_0"),
		Num_2:       lo.get_text("#Lotto649Control_history_dlQuery_No2_0"),
		Num_3:       lo.get_text("#Lotto649Control_history_dlQuery_No3_0"),
		Num_4:       lo.get_text("#Lotto649Control_history_dlQuery_No4_0"),
		Num_5:       lo.get_text("#Lotto649Control_history_dlQuery_No5_0"),
		Num_6:       lo.get_text("#Lotto649Control_history_dlQuery_No6_0"),
		Num_special: lo.get_text("#Lotto649Control_history_dlQuery_SNo_0"),
	}
	return result
}

// -
func (lo *Lotto649) get_text(key string) string {
	elem, err := lo.WebDriver.FindElement(selenium.ByCSSSelector, key)
	if err != nil {
		panic(err)
	}
	text, _ := elem.Text()
	return text
}
