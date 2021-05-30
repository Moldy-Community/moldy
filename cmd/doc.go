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
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	generateFlg bool
	openFlg     bool
	contribFlg  bool
)

// docCmd represents the docu command
var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if generateFlg {
			err := doc.GenMarkdownTree(rootCmd, "./moldyDoc")
			functions.CheckErrors(err, "Code 2", "Error in generate the documentation for moldy", "Report the error on github or re run with sudo")
		} else if openFlg {
			functions.OpenBrowser("https://moldybook.netlify.app/")
		} else if contribFlg {
			functions.OpenBrowser("https://github.com/Moldy-Community/")
		}

	},
}

func init() {
	rootCmd.AddCommand(docCmd)
	docCmd.Flags().BoolVarP(&generateFlg, "generate", "g", false, "Toggle this flag for generate a doc folder with detailed config of moldy commands")
	docCmd.Flags().BoolVarP(&openFlg, "open", "o", false, "Toggle this flag for open the documentation")
	docCmd.Flags().BoolVarP(&contribFlg, "contribute", "c", false, "Toggle this flag for open the moldy org")
}
