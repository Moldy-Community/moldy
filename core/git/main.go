package git

import (
	"runtime"

	log "github.com/Moldy-Community/moldy/core/logger"
	"github.com/Moldy-Community/moldy/core/runner"
)

func InitializeGit() {
	platform := runtime.GOOS
	if platform == "linux" || platform == "darwin" {
		runner.ShellRunner("git init", "bash", "-c")
		log.Logger.Info("Succesfuly create the repostory in the current directory")
	} else if platform == "windows" {
		runner.ShellRunner("git init", "cmd", "/k")
		log.Logger.Info("Succesfuly create the repostory in the current directory")

	} else {
		log.Logger.Error("Error in the platform os not supported")
	}
}
