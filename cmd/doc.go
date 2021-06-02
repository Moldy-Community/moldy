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
	docGen "github.com/Moldy-Community/moldy/core/doc-generator"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/spf13/cobra"
)

var (
	generateFlg bool
	openFlg     bool
	contribFlg  bool
)

// docCmd represents the docu command
var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Generate a documentation and open the doc tools",
	Long: `Yey! Moldy Doc generator is here:

If you run --generate this create a folder called moldyDoc here is the documentation for
alll commands of moldy.

If you run --open this open the moldy documentation online.

If you run --contribute open the moldy org on github`,
	Aliases: []string{"docu", "dc", "d"},
	Example: "moldy doc --generate or moldy dc --open",
	Run: func(cmd *cobra.Command, args []string) {
		if generateFlg {
			docGen.GenDocTree(rootCmd)
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
