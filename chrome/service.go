package chrome

import (
	"os"

	"github.com/tebeka/selenium"
)

func (c *ChromeObj) setSeleniumService() {
	opts := []selenium.ServiceOption{
		// selenium.StartFrameBuffer(), // Enable fake XWindow session.
		// selenium.Output(os.Stderr), // Output debug information to STDERR
	}

	SHOW_DEBUG_LOG := false
	if SHOW_DEBUG_LOG {
		opts = append(opts, selenium.Output(os.Stderr))
	}

	// Enable debug info.
	selenium.SetDebug(false)

	// Starts a ChromeDriver instance in the background. (This is browser)
	ChromeDriverService, err := selenium.NewChromeDriverService((*c)._ChromeDriverPath, port, opts...)
	if err != nil {
		panic(err)
	}
	(*c).Service = ChromeDriverService
}
