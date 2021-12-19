package gen

import (
	"os"

	"github.com/Moldy-Community/moldy/utils/errors"
	vp "github.com/spf13/viper"
)

/* Add the default values, paths, aliases and config name and type */
var (
	defaults = map[string]interface{}{
		"adminProjects": map[string]bool{
			"git":                 true,
			"conventionalCommits": true,
			"changelog":           true,
		},
		"aparienceOptions": map[string]bool{
			"progressBar": true,
			"logColors":   true,
		},
		"scripts": map[string]string{
			"test": "echo 'Running a example command'",
		},
	}
	paths = []string{
		"./",
	}
	configName = "MoldyFile.toml"
	configType = "toml"
)

/* Creating the config file */
func NewMoldyFile() {
	for _, p := range paths {
		vp.AddConfigPath(p)
	}
	vp.SetConfigName(configName)
	vp.SetConfigType(configType)
	for k, v := range defaults {
		vp.SetDefault(k, v)
	}

	if err := vp.SafeWriteConfigAs(configName); err != nil {
		if os.IsNotExist(err) {
			err = vp.WriteConfigAs(configName)
			errors.CheckError(err, "Code 2", "Error in write the config file", "Report the error on github or re try the command with new permmisions")
		}
	}
}
