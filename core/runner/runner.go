package runner

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions"
	toml "github.com/pelletier/go-toml"
)

func ShellRunner(command, shell, runner string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shell, runner, command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func RunnerWorker(name string) {
	if _, err := os.Stat("./MoldyFile.toml"); err == nil {
		file, err := toml.LoadFile("MoldyFile.toml")
		functions.CheckErrors(err, "Code 2", "Error in read the MoldyFile toml con run the command", "Check the moldy or user permissions")
		keysCmd := file.Get("moldyrunner").(*toml.Tree).Keys()
		for _, k := range keysCmd {
			if k == name {
				colors.Info("Command name to run:")
				fmt.Println(k)

				val := file.Get("moldyrunner").(*toml.Tree).Get(k).(string)

				colors.Info("Command value to run:")
				fmt.Println(val)

				platform := runtime.GOOS
				if platform == "linux" || platform == "darwin" {
					err, stout, sterr := ShellRunner("bash", "-c", val)
					functions.CheckErrors(err, "Code 2", "Error in run the command", "Check the permissions or report the error on github")
					colors.Info("StdOut of the command: ")
					fmt.Println(stout)
					colors.Info("Sterr of the command:")
					fmt.Println(sterr)
				} else if platform == "windows " {
					err, stout, sterr := ShellRunner("cmd", "/k", val)
					functions.CheckErrors(err, "Code 2", "Error in run the command", "Check the permissions or report the error on github")
					colors.Info("StdOut of the command: ")
					fmt.Println(stout)
					colors.Info("Sterr of the command:")
					fmt.Println(sterr)
				} else {
					colors.Error("Platform not supported :(")
				}

				break
			} else {
				colors.Error("Error command not found exit :(")
			}
		}

	} else if os.IsNotExist(err) {
		colors.Error("Moldy file not exist run **moldy config -g** for generate a config")
	} else {
		colors.Error("WTF How you are here? :o")
	}
}
