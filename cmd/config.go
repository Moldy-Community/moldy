/*
Copyright © 2021 Moldy

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
	"os"

	"github.com/Moldy-Community/moldy/core/config"
	"github.com/Moldy-Community/moldy/core/terminal"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/spf13/cobra"
	vp "github.com/spf13/viper"
)

var createToggle, editConfig bool

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure moldy for best and custom usage :D",
	Long: `Generate a config file with the basic specifications:

This create a MoldyFile.toml in the current directory
You can change the file manualy or by commands ( RECOMMENDED )

In error case:
  If you have any error report on Github for fix that
  in the next version :D
`,
	Run: func(cmd *cobra.Command, args []string) {
		if createToggle {
			config.CreateConfigFile()
		}

		if editConfig {
			valueFile := map[string]interface{}{
				"adminProjects": map[string]bool{
					"gitInit":               terminal.BasicPrompt("Admin Projects > Git init", "true") == "true",
					"conventionalCommits":   terminal.BasicPrompt("Admin Projects > Conventional Commits", "true") == "true",
					"conventionalWorkflows": terminal.BasicPrompt("Admin Projects > Conventional Workflows", "true") == "true",
					"semverMode":            terminal.BasicPrompt("Admin Projects > Semver Mode", "true") == "true",
					"changelogs":            terminal.BasicPrompt("Admin Projects > Changelogs", "true") == "true",
				},
				"aparienceOptions": map[string]bool{
					"progressBar": terminal.BasicPrompt("Aparience options > Progress bar", "true") == "true",
					"asciiArt":    terminal.BasicPrompt("Aparience options > ASCII Art", "true") == "true",
					"colorsMode":  terminal.BasicPrompt("Aparience options > Colors mode", "true") == "true",
				},
				"moldyRunner": map[string]string{
					"test": terminal.BasicPrompt("Moldy runner > Runner", "echo 'Running a example command'"),
				},
			}

			configName := "MoldyFile.toml"
			configType := "toml"
			_ = os.Remove("./MoldyFile.toml")

			vp.AddConfigPath("utils")

			vp.SetConfigName(configName)
			vp.SetConfigType(configType)
			for k, v := range valueFile {
				vp.SetDefault(k, v)
			}

			if err := vp.SafeWriteConfigAs(configName); err != nil {
				if os.IsNotExist(err) {
					err = vp.WriteConfigAs(configName)
					functions.CheckErrors(err, "Code 2", "Error in write the config file :(", "Report the error on github or re try the command with new permmisions")
				}
			}
		}
	},
	Example: "moldy config --create",
	Aliases: []string{"cfg", "c"},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolVarP(&createToggle, "create", "c", false, "Toggle the flag for create the config file")
	configCmd.Flags().BoolVarP(&editConfig, "edit", "e", false, "Customize the default values of the config")
}
