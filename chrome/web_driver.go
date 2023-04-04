package chrome

import (
	"fmt"

	"github.com/tebeka/selenium"
)

/*
WebDriver
*/
func (c *ChromeObj) buildWebDriver() {
	if (*c).WebDriver != nil {
		return
	}

	webDriver, err := selenium.NewRemote(*c.caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}

	(*c).WebDriver = webDriver
}

func (c *ChromeObj) GetWebDriver() selenium.WebDriver {
	return c.WebDriver
}
