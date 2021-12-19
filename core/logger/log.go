package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger = log.New()

func init() {
	Logger.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
		ForceColors:   true,
		ForceQuote:    true,
	})
	Logger.Out = os.Stdout
}
