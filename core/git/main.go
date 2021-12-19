package git

import (
	"os"
	"runtime"

	log "github.com/Moldy-Community/moldy/core/logger"
	"github.com/Moldy-Community/moldy/core/runner"
	"github.com/Moldy-Community/moldy/utils/errors"
	"github.com/go-git/go-git/v5"
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

func CloneRepos(url string) {
	path, err := os.Getwd()

	errors.CheckError(err, "Code 2", "Error in get the directory", "Check the permissions or report the error on github")

	log.Logger.Info("Clonning the repository now...\n")
	_, err2 := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	errors.CheckError(err2, "Code 2", "Error in clone the template", "Check the permissions, check if are installed git")
	log.Logger.Info("Succesfuly cloned the repository in the current directory")
}
