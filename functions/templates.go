package functions

import (
	"os"

	"github.com/Moldy-Community/moldy/utils"
	"github.com/spf13/viper"
)

var (
	mvcFolders = []string{
		"content",
		"model",
		"view",
		"controller",
	}
	dddFolders = []string{
		"config",
		"app",
		"modules",
		"config",
	}
	basicFolders = []string{
		"src",
		"modules",
		"config",
	}
)

func GenerateMVCTemplate() {
	for _, f := range mvcFolders {
		err := os.Mkdir(f, 0755)
		utils.CheckErrors(err, "Code 2", "Error in create the folders MVC pattern", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")
	}
	utils.Success("Succesfuly generated the MVC template folder")
}

func GenerateDDDTemplate() {
	for _, f := range dddFolders {
		err := os.Mkdir(f, 0755)
		utils.CheckErrors(err, "Code 2", "Error in create the folders DDD pattern", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")
	}
	utils.Success("Succesfuly generated the DDD template folder")
}

func GenerateBasicTemplate() {
	for _, f := range basicFolders {
		err := os.Mkdir(f, 0755)
		utils.CheckErrors(err, "Code 2", "Error in create the folders Basic Moldy pattern", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")
	}
	utils.Success("Succesfuly generated the basic Moldy template folder")
}

func MoldyCfgFile(author, name, version, description string) {
	viper.Set("moldyPackages.name", name)
	viper.Set("moldyPackages.author", author)
	viper.Set("moldyPackages.version", version)
	viper.Set("moldyPackages.description", description)
	utils.Success("Succesfuly set the new values for the MoldyFile.toml")
}
