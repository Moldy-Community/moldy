package functions

import (
	"os"
	"text/template"

	"github.com/Moldy-Community/moldy/utils"
)

type ReadmeTmpl struct {
	Author      string
	Description string
	Name        string
	Version     string
}

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

func parseTmpl(path string, content ReadmeTmpl) {
	t, err := template.ParseFiles(path)
	utils.CheckErrors(err, "Code 2", "Error in parse the template", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")

	f, err := os.Create(path)
	utils.CheckErrors(err, "Code 2", "Error in write the file Readme Template", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")

	err = t.Execute(f, content)
	utils.CheckErrors(err, "Code 2", "Error in get the content of the Readme Template", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")

	f.Close()
}

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

func ReadmeTmplGenerator(author, description, name, version string) {
	val := ReadmeTmpl{Author: author, Description: description, Name: name, Version: version}
	parseTmpl("./Readme.md", val)
}
