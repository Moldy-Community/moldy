package functions

import (
	"github.com/Moldy-Community/moldy/utils"
	"github.com/spf13/viper"
)

func AparienceChanges() {
	pgb := BasicPrompt("Show a progress bar? OPTIONS: [yes, no]", "yes")
	asciiArt := BasicPrompt("Show a Moldy Ascii art? OPTIONS: [yes, no]", "yes")
	colors := BasicPrompt("Show a color output ? OPTIONS: [ yes, no ]", "yes")
	switch pgb {
	case "yes":
		viper.Set("aparienceOptions.progressBar", true)
		utils.Success("Succesfuly set the progress bar option :p")
	case "no":
		viper.Set("aparienceOptions.progressBar", false)
		utils.Success("Succesfuly set the progress bar option :p")
	default:
		utils.Error("Error option not found :c")
	}

	switch asciiArt {
	case "yes":
		viper.Set("aparienceOptions.asciiArt", true)
		utils.Success("Succesfuly set the ascii art option :p")
	case "no":
		viper.Set("aparienceOptions.asciiArt", false)
		utils.Success("Succesfuly set the ascii art option :p")
	default:
		utils.Error("Error option not found :c")
	}
	switch colors {
	case "yes":
		viper.Set("aparienceOptions.colorsMode", true)
		utils.Success("Succesfuly set the colors option :p")
	case "no":
		viper.Set("aparienceOptions.colorsMode", false)
		utils.Success("Succesfuly set the colors option :p")
	default:
		utils.Error("Error option not found :c")
	}
}
