package utils

import (
	"os/exec"
	"runtime"

	log "github.com/Moldy-Community/moldy/core/logger"
)

func OpenUrlBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
		log.Logger.Info("Opened the url in the browser")
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		log.Logger.Info("Opened the url in the browser")
	case "darwin":
		err = exec.Command("open", url).Start()
		log.Logger.Info("Opened the url in the browser")

	default:
		log.Logger.Warn("Platform not supported")
	}
	if err != nil {
		log.Logger.Error(err)
	}
}
