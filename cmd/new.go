/*
Copyright © 2021 Moldy-Community

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/Moldy-Community/moldy/functions"
	"github.com/Moldy-Community/moldy/utils"
	"github.com/spf13/cobra"
)

var (
	patternToogle bool
	basicToggle   bool
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate a new moldy project with a basic structure or a desing pattern",
	Long: `Yey! the moldy templates is here:

You can generate the templates with a pattern desing or not
This generate a readme a MoldyFile.toml and the folder structure

In error case:
  If you have any error report on Github for fix that
  in the next version :D
`,
	Run: func(cmd *cobra.Command, args []string) {
		if patternToogle {
			name := functions.BasicPrompt("Name of the package", "Example package")
			author := functions.BasicPrompt("Author name", "none")
			version := functions.BasicPrompt("Version of the Package", "1.0")
			description := functions.BasicPrompt("Description of the package", "Example package")
			pattern := functions.SelectPrompt("Select the desing pattern", []string{"MVC", "DDD"})
			if pattern == "MVC" {
				functions.GenerateMVCTemplate()
				functions.MoldyCfgFile(author, description, name, version)
				editApa := functions.BasicPrompt("Edit the aparience preferences ?", "yes")
				if editApa == "yes" {
					functions.AparienceChanges()
				} else {
					utils.Info("Aparience not selected bye bye")
				}
				editAdmin := functions.BasicPrompt("Edit the adminitration project preferences ?", "yes")
				if editAdmin == "yes" {
					functions.ProjectChanges()
				} else {
					utils.Info("Edit the Administration project not selected bye bye")
				}
				functions.CreateConfigFile()

			} else if pattern == "DDD" {
				functions.GenerateDDDTemplate()

				functions.MoldyCfgFile(author, description, name, version)
				editApa := functions.BasicPrompt("Edit the aparience preferences ?", "yes")
				if editApa == "yes" {
					functions.AparienceChanges()
				} else {
					utils.Info("Aparience not selected bye bye")
				}
				editAdmin := functions.BasicPrompt("Edit the adminitration project preferences ?", "yes")
				if editAdmin == "yes" {
					functions.ProjectChanges()
				} else {
					utils.Info("Edit the Administration project not selected bye bye")
				}
				functions.CreateConfigFile()
			}
		} else if basicToggle {
			name := functions.BasicPrompt("Name of the package", "Example package")
			author := functions.BasicPrompt("Author name", "none")
			version := functions.BasicPrompt("Version of the Package", "1.0")
			description := functions.BasicPrompt("Description of the package", "Example package")
			functions.GenerateBasicTemplate()
			functions.MoldyCfgFile(author, name, version, description)
			editApa := functions.BasicPrompt("Edit the aparience preferences ?", "yes")
			if editApa == "yes" {
				functions.AparienceChanges()
			} else {
				utils.Info("Aparience not selected bye bye")
			}
			editAdmin := functions.BasicPrompt("Edit the adminitration project preferences ?", "yes")
			if editAdmin == "yes" {
				functions.ProjectChanges()
			} else {
				utils.Info("Edit the Administration project not selected bye bye")
			}
			functions.CreateConfigFile()
		}
	},
	Aliases: []string{"n", "generate"},
	Example: "moldy generate -p or moldy new -b",
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().BoolVarP(&patternToogle, "pattern", "p", false, "Generate a Moldy project with a desing pattern")
	newCmd.Flags().BoolVarP(&basicToggle, "basic", "b", false, "Generate a Moldy project with a basic structure")
}
