/*
Copyright © 2021 Moldy-Community

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
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
How contact to Moldy, etc.

Show a beautifull information`,
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
