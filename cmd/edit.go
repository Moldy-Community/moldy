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

var (
	aparienceToggle bool
	adminToggle     bool
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit Moldy File configuration of aparience and projects",
	Long: `This command help you to edit the moldy config options

Run the edit command for change the aparience or the project configs
with a beautiful terminal UI :D

In error case:
  If you have any error report on Github for fix that
  in the next version :p`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")
	},
	Aliases: []string{"e", "ed"},
	Example: "moldy e -a or moldy ed -p",
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().BoolVarP(&aparienceToggle, "aparience", "a", false, "Toggle this flag for chage the aparience options")
	editCmd.Flags().BoolVarP(&adminToggle, "project", "p", false, "Toggle this flag for chage the adminProject option :p")
}
