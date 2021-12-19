package gen

import (
	"os"

	log "github.com/Moldy-Community/moldy/core/logger"
	"github.com/Moldy-Community/moldy/utils/errors"
)

func GenerateFolders(folders []string) {
	for _, f := range folders {
		err := os.Mkdir(f, 0o755)
		errors.CheckError(err, "Code 2", "Error in create the folders", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")
	}

	log.Logger.Info("Succesfully created the folders")
}
