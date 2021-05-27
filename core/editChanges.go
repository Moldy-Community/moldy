package core

import (
	"github.com/Moldy-Community/moldy/utils"
	"github.com/spf13/viper"
)

func AparienceChanges() {
	/* Define the prompts questions */

	pgb := BasicPrompt("Show a progress bar? OPTIONS: [yes, no]", "yes")
	asciiArt := BasicPrompt("Show a Moldy Ascii art? OPTIONS: [yes, no]", "yes")
	colors := BasicPrompt("Show a color output ? OPTIONS: [ yes, no ]", "yes")

	/* Switch in the options  */

	switch pgb {
	case "yes":
		viper.Set("aparienceOptions.progressBar", true)
		utils.Success("Succesfuly set the progress bar option :p")
	case "no":
		viper.Set("aparienceOptions.progressBar", false)
		utils.Success("Succesfuly set the progress bar option :D")
	default:
		utils.Error("Error option not found :c")
	}

	switch asciiArt {
	case "yes":
		viper.Set("aparienceOptions.asciiArt", true)
		utils.Success("Succesfuly set the ascii art option :p")
	case "no":
		viper.Set("aparienceOptions.asciiArt", false)
		utils.Success("Succesfuly set the ascii art option :D")
	default:
		utils.Error("Error option not found :c")
	}

	switch colors {
	case "yes":
		viper.Set("aparienceOptions.colorsMode", true)
		utils.Success("Succesfuly set the colors option :p")
	case "no":
		viper.Set("aparienceOptions.colorsMode", false)
		utils.Success("Succesfuly set the colors option :D")
	default:
		utils.Error("Error option not found :c")
	}
}

func ProjectChanges() {

	/* Define the prompts questions */

	iniGit := BasicPrompt("Initialize git ? OPTIONS: [yes, no]", "yes")
	convenCommits := BasicPrompt("Use the conventional commits tool ? OPTIONS: [yes, no]", "yes")
	convenWorkflows := BasicPrompt("Use the conventional workflows ? OPTIONS: [yes, no]", "yes")
	semverUsage := BasicPrompt("Usage the semantic versioning ? OPTIONS: [yes, no]", "yes")
	changeLogGen := BasicPrompt("Usage the change log Generator ? OPTIONS: [yes, no]", "yes")

	/* Switch in the options  */

	switch iniGit {
	case "yes":
		viper.Set("adminProjects.gitInit", true)
		utils.Success("Succesfuly set the gitInit option :p")
	case "no":
		viper.Set("adminProjects.gitInit", false)
		utils.Success("Succesfuly set the gitInit option :D")
	default:
		utils.Error("Error option not found :c")
	}

	switch convenCommits {
	case "yes":
		viper.Set("adminProjects.conventionalCommits", true)
		utils.Success("Succesfuly set the conventionalCommits option :p")
	case "no":
		viper.Set("adminProjects.conventionalCommits", false)
		utils.Success("Succesfuly set the conventionalCommits option :p")
	default:
		utils.Error("Error option not found :c")
	}

	switch convenWorkflows {
	case "yes":
		viper.Set("adminProjects.conventionalWorkflows", true)
		utils.Success("Succesfuly set the conventionalWorkflows option :p")
	case "no":
		viper.Set("adminProjects.conventionalWorkflows", false)
		utils.Success("Succesfuly set the conventionalWorkflows option :p")
	default:
		utils.Error("Error option not found :c")
	}

	switch semverUsage {
	case "yes":
		viper.Set("adminProjects.semverMode", true)
		utils.Success("Succesfuly set the semverMode option :p")
	case "no":
		viper.Set("adminProjects.semverMode", false)
		utils.Success("Succesfuly set the semverMode option :p")
	default:
		utils.Error("Error option not found :c")
	}

	switch changeLogGen {
	case "yes":
		viper.Set("adminProjects.changelogs", true)
		utils.Success("Succesfuly set the changelogs option :p")
	case "no":
		viper.Set("adminProjects.changelogs", false)
		utils.Success("Succesfuly set the changelogs option :p")
	default:
		utils.Error("Error option not found :c")
	}
}
