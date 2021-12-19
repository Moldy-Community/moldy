package new

import (
	"os"

	"github.com/Moldy-Community/moldy/utils/errors"
  log 	"github.com/Moldy-Community/moldy/core/logger"
)

func GenerateFolders(folders []string) {
	for _, f := range folders {
		err := os.Mkdir(f, 0755)
    errors.CheckError(err, "Code 2", "Error in create the folders", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")
	}

  log.Logger.Info("Succesfully created the folders")
}
