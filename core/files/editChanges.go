package files

import (
	"github.com/Moldy-Community/moldy/core/terminal"
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/spf13/viper"
)

func AparienceChanges() {
	/* Define the prompts questions */

	pgb := terminal.BasicPrompt("Show a progress bar? OPTIONS: [yes, no]", "yes")
	asciiArt := terminal.BasicPrompt("Show a Moldy Ascii art? OPTIONS: [yes, no]", "yes")
	colors := terminal.BasicPrompt("Show a color output ? OPTIONS: [ yes, no ]", "yes")

	/* Switch in the options  */

	switch pgb {
	case "yes":
		viper.Set("aparienceOptions.progressBar", true)
		colors.Success("Succesfuly set the progress bar option :p")
	case "no":
		viper.Set("aparienceOptions.progressBar", false)
		colors.Success("Succesfuly set the progress bar option :D")
	default:
		colors.Error("Error option not found :c")
	}

	switch asciiArt {
	case "yes":
		viper.Set("aparienceOptions.asciiArt", true)
		colors.Success("Succesfuly set the ascii art option :p")
	case "no":
		viper.Set("aparienceOptions.asciiArt", false)
		colors.Success("Succesfuly set the ascii art option :D")
	default:
		colors.Error("Error option not found :c")
	}

	switch colors {
	case "yes":
		viper.Set("aparienceOptions.colorsMode", true)
		colors.Success("Succesfuly set the colors option :p")
	case "no":
		viper.Set("aparienceOptions.colorsMode", false)
		colors.Success("Succesfuly set the colors option :D")
	default:
		colors.Error("Error option not found :c")
	}
}

func ProjectChanges() {

	/* Define the prompts questions */

	iniGit := terminal.BasicPrompt("Initialize git ? OPTIONS: [yes, no]", "yes")
	convenCommits := terminal.BasicPrompt("Use the conventional commits tool ? OPTIONS: [yes, no]", "yes")
	convenWorkflows := terminal.BasicPrompt("Use the conventional workflows ? OPTIONS: [yes, no]", "yes")
	semverUsage := terminal.BasicPrompt("Usage the semantic versioning ? OPTIONS: [yes, no]", "yes")
	changeLogGen := terminal.BasicPrompt("Usage the change log Generator ? OPTIONS: [yes, no]", "yes")

	/* Switch in the options  */

	switch iniGit {
	case "yes":
		viper.Set("adminProjects.gitInit", true)
		colors.Success("Succesfuly set the gitInit option :p")
	case "no":
		viper.Set("adminProjects.gitInit", false)
		colors.Success("Succesfuly set the gitInit option :D")
	default:
		colors.Error("Error option not found :c")
	}

	switch convenCommits {
	case "yes":
		viper.Set("adminProjects.conventionalCommits", true)
		colors.Success("Succesfuly set the conventionalCommits option :p")
	case "no":
		viper.Set("adminProjects.conventionalCommits", false)
		colors.Success("Succesfuly set the conventionalCommits option :p")
	default:
		colors.Error("Error option not found :c")
	}

	switch convenWorkflows {
	case "yes":
		viper.Set("adminProjects.conventionalWorkflows", true)
		colors.Success("Succesfuly set the conventionalWorkflows option :p")
	case "no":
		viper.Set("adminProjects.conventionalWorkflows", false)
		colors.Success("Succesfuly set the conventionalWorkflows option :p")
	default:
		colors.Error("Error option not found :c")
	}

	switch semverUsage {
	case "yes":
		viper.Set("adminProjects.semverMode", true)
		colors.Success("Succesfuly set the semverMode option :p")
	case "no":
		viper.Set("adminProjects.semverMode", false)
		colors.Success("Succesfuly set the semverMode option :p")
	default:
		colors.Error("Error option not found :c")
	}

	switch changeLogGen {
	case "yes":
		viper.Set("adminProjects.changelogs", true)
		colors.Success("Succesfuly set the changelogs option :p")
	case "no":
		viper.Set("adminProjects.changelogs", false)
		colors.Success("Succesfuly set the changelogs option :p")
	default:
		colors.Error("Error option not found :c")
	}
}
