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
	"github.com/Moldy-Community/moldy/core/config"
	files "github.com/Moldy-Community/moldy/core/files"
	"github.com/Moldy-Community/moldy/core/terminal"
	"github.com/Moldy-Community/moldy/utils/colors"
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
			name := terminal.BasicPrompt("Name of the package", "Example package")
			author := terminal.BasicPrompt("Author name", "none")
			version := terminal.BasicPrompt("Version of the Package", "1.0")
			description := terminal.BasicPrompt("Description of the package", "Example package")
			pattern := terminal.SelectPrompt("Select the desing pattern", []string{"MVC", "DDD"})
			if pattern == "MVC" {
				files.GenerateMVCTemplate()
				files.MoldyCfgFile(author, description, name, version)
				editApa := terminal.BasicPrompt("Edit the aparience preferences ?", "yes")
				if editApa == "yes" {
					files.AparienceChanges()
				} else {
					colors.Info("Aparience not selected bye bye")
				}
				editAdmin := terminal.BasicPrompt("Edit the adminitration project preferences ?", "yes")
				if editAdmin == "yes" {
					files.ProjectChanges()
				} else {
					colors.Info("Edit the Administration project not selected bye bye")
				}
				config.CreateConfigFile()

			} else if pattern == "DDD" {
				files.GenerateDDDTemplate()

				files.MoldyCfgFile(author, description, name, version)
				editApa := terminal.BasicPrompt("Edit the aparience preferences ?", "yes")
				if editApa == "yes" {
					files.AparienceChanges()
				} else {
					colors.Info("Aparience not selected bye bye")
				}
				editAdmin := terminal.BasicPrompt("Edit the administration project preferences ?", "yes")
				if editAdmin == "yes" {
					files.ProjectChanges()
				} else {
					colors.Info("Edit the Administration project not selected bye bye")
				}
				config.CreateConfigFile()
			}
		} else if basicToggle {
			name := terminal.BasicPrompt("Name of the package", "Example package")
			author := terminal.BasicPrompt("Author name", "none")

			version := terminal.BasicPrompt("Version of the Package", "1.0")
			description := terminal.BasicPrompt("Description of the package", "Example package")
			files.GenerateBasicTemplate()
			files.MoldyCfgFile(author, name, version, description)
			editApa := terminal.BasicPrompt("Edit the aparience preferences ?", "yes")
			if editApa == "yes" {
				files.AparienceChanges()
			} else {
				colors.Info("Aparience not selected bye bye")
			}
			editAdmin := terminal.BasicPrompt("Edit the adminitration project preferences ?", "yes")
			if editAdmin == "yes" {
				files.ProjectChanges()
			} else {
				colors.Info("Edit the Administration project not selected bye bye")
			}
			config.CreateConfigFile()
		} else if dotFilesToggle {
			files.CreateDotFiles()
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
