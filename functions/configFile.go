package functions

import (
	"github.com/Moldy-Community/CLI/utils"
	vp "github.com/spf13/viper"
)

/* Add the default values, paths, aliases and config name and type */
var (
	defaults = map[string]interface{}{
		"progressBar": true,
		"asciiArt":    true,
		"colorsMode":  true,
	}
	aliases = map[string]string{
		"prgB": "progressbar",
		"ascA": "asciiArt",
		"colM": "colorsMode",
	}
	configName = "MoldyFile"
	configType = "toml"
)

/* Creating the config file */
func CreateConfigFile() {
	vp.SetConfigName(configName)
	vp.SetConfigType(configType)
	for k, v := range defaults {
		vp.SetDefault(k, v)
	}
	for a, k := range aliases {
		vp.RegisterAlias(a, k)
	}
	err := vp.WriteConfig()
	utils.CheckErrors(err, "Code 6", "Error writing the config File", "Report the error in github :D")
}

func ReadCfgFile() (bool, bool, bool) {
	progressBar := vp.GetBool("prgB")
	asciiArt := vp.GetBool("ascA")
	colM := vp.GetBool("colM")
	return progressBar, asciiArt, colM
}
