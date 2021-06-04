// Package git provides ...
package git

import (
	"runtime"

	moldyConfig "github.com/Moldy-Community/moldy/core/files/configMoldy"
	"github.com/Moldy-Community/moldy/core/runner"
	"github.com/Moldy-Community/moldy/utils/colors"
)

func CheckIfInitialize() {
	cfg := moldyConfig.Settings().AdminProjects.GitInit
	if cfg {
		InitializeGit()
	}
}

func InitializeGit() {
	platform := runtime.GOOS
	if platform == "linux" || platform == "darwin" {
		runner.ShellRunner("git init", "bash", "-c")
		colors.Success("Succesfuly create the repostory in the current directory")
	} else if platform == "windows" {
		runner.ShellRunner("git init", "cmd", "/k")
		colors.Success("Succesfuly create the repostory in the current directory")

	} else {
		colors.Error("Error in the platform os not supported")
	}
}
