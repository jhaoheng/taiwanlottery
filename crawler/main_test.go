package crawler

import (
	"os"
	"testing"

	"github.com/jhaoheng/taiwanlottery/chrome"

	"github.com/tebeka/selenium"
)

var web_driver selenium.WebDriver

func TestMain(m *testing.M) {
	chrome_agent := chrome.NewAgent()
	defer chrome_agent.CloseAgent()
	chrome_agent.RunWebDriver()
	//
	web_driver = chrome_agent.GetWebDriver()
	//
	m.Run()
	os.Exit(0)
}
