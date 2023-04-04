package chrome

/*
How to use

````
c := chrome.NewAgent()
c.RunWebDriver() or c.RunWebDriverByProxy("YOUR PROXY SERVER IP", PORT)
defer c.CloseAgent()
webDriver := c.GetWebDriver()
````
*/

import (
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"github.com/tebeka/selenium/log"
)

const (
	port = 8180

	LINUX   string = "linux"
	MAC_OSX string = "darwin"

	SHOW_DEBUG_LOG = false
)

type IChromeDriver interface {
	RunWebDriver()
	RunWebDriverByProxy(ip string, port int)
	CloseAgent()
	GetWebDriver() selenium.WebDriver

	//
	Set_ChromeDriverPath(path string)
	Set_UserAgent(userAgent string)
	Set_AcceptLanguage(Language string)
	Set_EnableWindow()
	// Set_AcceptEncoding(Encoding string)
	// Set_Accept(accept string)
}

type ChromeObj struct {

	//
	ENV

	//
	caps      *selenium.Capabilities
	Service   *selenium.Service
	WebDriver selenium.WebDriver
}

var (
	DarwinChromeDriverFile string = "chromedriver-darwin"
	LinuxChromeDriverFile  string = "chromedriver-linux64"
)

func NewAgent() IChromeDriver {

	c := &ChromeObj{}
	c._Default_Env()

	return c
}

// func (c *ChromeObj) SetChrmoeDriverPath(driverPath string) {
// 	c.binPath = driverPath

// pwd, _ := os.Getwd()
// binDir := ""
// if pwd == "/home/app" {
// 	// pro 環境
// 	binDir = "chrome"
// } else if strings.Contains(pwd, "chrome") {
// 	// unit test 環境
// 	binDir = pwd
// } else {
// 	// dev 環境
// 	binDir = "../../chrome"
// }

// if runtime.GOOS == MAC_OSX {
// 	c.binPath = binDir + "/" + DarwinChromeDriverFile
// } else if runtime.GOOS == LINUX {
// 	c.binPath = binDir + "/" + LinuxChromeDriverFile
// } else {
// 	panic("Can not find 'chrmoedriver' for this OS")
// }
// }

func (c *ChromeObj) RunWebDriverByProxy(proxyIp string, proxyPort int) {
	// proxy
	proxy := selenium.Proxy{
		Type:     selenium.Manual,
		HTTP:     proxyIp,
		HTTPPort: proxyPort,
	}
	c.caps.AddProxy(proxy)
	c.RunWebDriver()
}

func (c *ChromeObj) RunWebDriver() {
	// 設定 Selenium Capabilities
	c.setSeleniumCapabilities()
	// 設定 browser
	c.setSeleniumService()
	//
	c.buildWebDriver()
}

func (c *ChromeObj) CloseAgent() {
	c.Service.Stop()
	c.WebDriver.Quit()
}

// 設定 Capabilities
func (c *ChromeObj) setSeleniumCapabilities() {
	//
	caps := selenium.Capabilities{}

	// ChromeDriver ref : `https://sites.google.com/a/chromium.org/chromedriver/capabilities`
	caps.AddChrome(chrome.Capabilities{
		Path: default_Chrome_AppPath,
		Args: c.get_args(),
		Prefs: map[string]interface{}{
			"profile": map[string]interface{}{
				"managed_default_content_settings.images": 2,
				"name": "max",
			},
			"intl.accept_languages": c._AcceptLanguage,
			"web_apps": map[string]string{
				"system_web_app_last_installed_language": c._AcceptLanguage,
				// "system_web_app_last_update":             "0",
			},
		},
		LocalState: map[string]interface{}{
			// "background_mode.enabled": false,
		},
		ExcludeSwitches: []string{
			"--enable-automation", // 關閉警告字樣
			"--allow-file-access-from-files",
			"--default",
		},
	})
	//
	caps.AddLogging(log.Capabilities{
		log.Server:      log.Debug,
		log.Browser:     log.Debug,
		log.Client:      log.Debug,
		log.Driver:      log.Debug,
		log.Performance: log.Debug,
		log.Profiler:    log.Debug,
	})

	c.caps = &caps
}

// Args ref :`https://peter.sh/experiments/chromium-command-line-switches/`
func (c *ChromeObj) get_args() []string {
	tmp_folder := strconv.Itoa(time.Now().Nanosecond())

	args := []string{
		"--disable-in-process-stack-traces",
		"--disable-local-storage",
		"--lang=en",
		"--disable-sync-types=''",
		"--disable-sync",
		"--disabled",

		"--disable-web-security",
		// "ignore-certificate-errors",
		"--disable-crash-reporter",
		"--disable-demo-mode",
		"--disable-cookie-encryption",
		"--disable-component-cloud-policy",
		"--disable-checker-imaging",
		"--disable-bundled-ppapi-flash", // 禁止 flash
		"--disable-internal-flash",      // 禁止 flash
		// "--disable-prompt-on-repost",
		"--disable-logging",
		"--log-level=3",
		"--disable-extensions",
		"--no-sandbox",
		"--user-agent=" + c._UserAgent, // 模擬user-agent
		// "--window-size=375,667",

		"--user-data-dir=/Users/max/Documents/play/seleniumapp/userdata/" + tmp_folder,
		// "--disk-cache-dir=/Users/max/Documents/play/seleniumapp/userdata/" + tmp_folder,
		// "--enable-local-sync-backend",
		// "--local-sync-backend-dir=/Users/max/Documents/play/seleniumapp/userdata/" + tmp_folder,
	}
	if !c._EnableWindow {
		args = append(args, "--headless") // 設定Chrome無頭模式，在linux下執行，需要設定這個引數，否則會報錯
	}
	return args
}
