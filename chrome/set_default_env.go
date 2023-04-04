package chrome

import "runtime"

type ENV struct {
	_ChromeDriverPath string
	_UserAgent        string
	_AcceptLanguage   string
	_AcceptEncoding   string
	_Accept           string

	_EnableWindow bool
}

var (
	Default_Dir                    string = "../"
	default_Chrome_AppPath         string = ""
	default_Chrome_OsxDriverPath   string = "./chrome/chromedriver-darwin"
	default_Chrome_LinuxDriverPath string = "./chrome/chromedriver-linux64"
	default_UserAgent              string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"
	default_AcceptLanguage         string = "en-US, en;q=0.9, zh-TW;q=0.8, zh;q=0.7"
	default_AcceptEncoding         string = "gzip, deflate, br"
	default_Accept                 string = "text/html, application/xhtml+xml, application/xml;q=0.9, image/webp, image/apng, */*;q=0.8, application/signed-exchange;v=b3;q=0.9"
	default_EnableWindow           bool   = false
)

func (c *ChromeObj) _Default_Env() {
	if runtime.GOOS == MAC_OSX {
		c._ChromeDriverPath = Default_Dir + "/" + default_Chrome_OsxDriverPath
	} else if runtime.GOOS == LINUX {
		c._ChromeDriverPath = Default_Dir + "/" + default_Chrome_LinuxDriverPath
	} else {
		panic("Can not find 'chrmoedriver' for this OS")
	}

	c._UserAgent = default_UserAgent
	c._AcceptLanguage = default_AcceptLanguage
	c._AcceptEncoding = default_AcceptEncoding
	c._Accept = default_Accept
	c._EnableWindow = default_EnableWindow
}

func (c *ChromeObj) Set_ChromeDriverPath(path string) {
	c._ChromeDriverPath = path
}

func (c *ChromeObj) Set_UserAgent(userAgent string) {
	c._UserAgent = userAgent
}

func (c *ChromeObj) Set_AcceptLanguage(Language string) {
	c._AcceptLanguage = Language
}

// func (c *ChromeObj) Set_AcceptEncoding(Encoding string) {
// 	c._AcceptEncoding = Encoding
// }

// func (c *ChromeObj) Set_Accept(accept string) {
// 	c._Accept = accept
// }

func (c *ChromeObj) Set_EnableWindow() {
	c._EnableWindow = true
}
