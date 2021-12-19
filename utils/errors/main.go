package errors

import (
	"fmt"
	"os"

	log "github.com/Moldy-Community/moldy/core/logger"
)

func CheckError(err error, code, msg, solution string) {
	if err != nil {
		err_msg := fmt.Sprintf("ERROR: [ %s ] -> %s || SOLVE: %s", code, msg, solution)
		log.Logger.Error(err_msg)
		os.Exit(2)
	}
}
