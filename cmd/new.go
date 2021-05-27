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
	"github.com/Moldy-Community/moldy/core"
	"github.com/Moldy-Community/moldy/utils"
	"github.com/spf13/cobra"
)

var (
	patternToogle  bool
	basicToggle    bool
	dotFilesToggle bool
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
			name := core.BasicPrompt("Name of the package", "Example package")
			author := core.BasicPrompt("Author name", "none")
			version := core.BasicPrompt("Version of the Package", "1.0")
			description := core.BasicPrompt("Description of the package", "Example package")
			pattern := core.SelectPrompt("Select the desing pattern", []string{"MVC", "DDD"})
			if pattern == "MVC" {
				core.GenerateMVCTemplate()
				core.MoldyCfgFile(author, description, name, version)
				editApa := core.BasicPrompt("Edit the aparience preferences ?", "yes")
				if editApa == "yes" {
					core.AparienceChanges()
				} else {
					utils.Info("Aparience not selected bye bye")
				}
				editAdmin := core.BasicPrompt("Edit the adminitration project preferences ?", "yes")
				if editAdmin == "yes" {
					core.ProjectChanges()
				} else {
					utils.Info("Edit the Administration project not selected bye bye")
				}
				core.CreateConfigFile()

			} else if pattern == "DDD" {
				core.GenerateDDDTemplate()

				core.MoldyCfgFile(author, description, name, version)
				editApa := core.BasicPrompt("Edit the aparience preferences ?", "yes")
				if editApa == "yes" {
					core.AparienceChanges()
				} else {
					utils.Info("Aparience not selected bye bye")
				}
				editAdmin := core.BasicPrompt("Edit the administration project preferences ?", "yes")
				if editAdmin == "yes" {
					core.ProjectChanges()
				} else {
					utils.Info("Edit the Administration project not selected bye bye")
				}
				core.CreateConfigFile()
			}
		} else if basicToggle {
			name := core.BasicPrompt("Name of the package", "Example package")
			author := core.BasicPrompt("Author name", "none")
			version := core.BasicPrompt("Version of the Package", "1.0")
			description := core.BasicPrompt("Description of the package", "Example package")
			core.GenerateBasicTemplate()
			core.MoldyCfgFile(author, name, version, description)
			editApa := core.BasicPrompt("Edit the aparience preferences ?", "yes")
			if editApa == "yes" {
				core.AparienceChanges()
			} else {
				utils.Info("Aparience not selected bye bye")
			}
			editAdmin := core.BasicPrompt("Edit the adminitration project preferences ?", "yes")
			if editAdmin == "yes" {
				core.ProjectChanges()
			} else {
				utils.Info("Edit the Administration project not selected bye bye")
			}
			core.CreateConfigFile()
		} else if dotFilesToggle {
			core.CreateDotFiles()
		}
	},
	Aliases: []string{"n", "generate"},
	Example: "moldy generate -p or moldy new -b",
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().BoolVarP(&patternToogle, "pattern", "p", false, "Generate a Moldy project with a desing pattern")
	newCmd.Flags().BoolVarP(&basicToggle, "basic", "b", false, "Generate a Moldy project with a basic structure")
	newCmd.Flags().BoolVarP(&dotFilesToggle, "dotfiles", "d", false, "Create a dot files for git, editor config etc")
}
