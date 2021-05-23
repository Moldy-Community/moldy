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
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show the information about Moldy",
	Long: `Print the Information about Moldy
How contact to Moldy the repository ascii art etc :D
Show a beautifull information about the Moldy project

In error case:

If you have any error report on Github for fix that
in the next version :D`,
	Aliases: []string{"inf", "in"},
	Example: "moldy info",
	Run: func(cmd *cobra.Command, args []string) {
		showInfo()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func showInfo() {
	var asciiMoldy string = `

███╗   ███╗ ██████╗ ██╗     ██████╗ ██╗   ██╗
████╗ ████║██╔═══██╗██║     ██╔══██╗╚██╗ ██╔╝
██╔████╔██║██║   ██║██║     ██║  ██║ ╚████╔╝
██║╚██╔╝██║██║   ██║██║     ██║  ██║  ╚██╔╝
██║ ╚═╝ ██║╚██████╔╝███████╗██████╔╝   ██║
╚═╝     ╚═╝ ╚═════╝ ╚══════╝╚═════╝    ╚═╝

	`

	fmt.Printf(asciiMoldy + `
The best tool for initial his project :

Moldy is a Project Starter and Project Administrator
writed in golang that help for start his project and is
100 OPEN SOURCE.

Author: Moldy Community
Contact mail: moldycommunity@gmail.com
Repository: www.github.com/Moldy-Community/Cli
Web Page: www.moldy-community.github.io/web

-----------------------------------------------------

Made with love in Latin America.
	`)
}
