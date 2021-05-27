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
	"github.com/Moldy-Community/moldy/core"
	"github.com/spf13/cobra"
)

var createToogle bool

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
		if createToogle {
			core.CreateConfigFile()
		}
	},
	Example: "moldy config --create",
	Aliases: []string{"cfg", "c"},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolVarP(&createToogle, "create", "c", false, "Toggle the flag for create the config file")
}
