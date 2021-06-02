package runner

import (
	"fmt"
	"os"

	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions"
	toml "github.com/pelletier/go-toml"
)

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
