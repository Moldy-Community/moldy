/*
Copyright © 2021 Moldy Community

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
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions/packages"
	"github.com/spf13/cobra"
)

// packagesCmd represents the packages command
var packagesCmd = &cobra.Command{
	Use:   "packages [create|delete|update]",
	Short: "Administrate the moldy packages",
	Long: `With this command you can administrate the moldy packages

Valid arguments:
  create: Create a Moldy package in less than 30 seconds
  delete: Delete a Moldy package with a password and a Id
  update: Update a package providing the id

DETALING THE FUNCTIONALITIE AND THE ARGS

Create:
  Launch a prompt for create a command
  ARGS: None

Delete:
  Delete a Moldy Package from the API you need the id and the password
  ARGS: [ <id> <password_package> ]

Update:
  Update the URL, or share to the world a new version of your package, since the terminal.
  ARGS: <id>`,
	ValidArgs: []string{"create", "delete", "update"},
	Args:      cobra.ExactValidArgs(1),
	Aliases:   []string{"pkg"},
	Example:   "moldy pkg create",
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "create":
			packages.Create()
		case "delete":
			if len(args) < 2 {
				colors.Error("The command need two params, the correct way to do this command is:\nmoldy delete <id> <password_package>\n")
			} else {
				packages.Delete(args[1], args[2])
			}
		case "update":
			if len(args) >= 2 {
				packages.Update(args[1])
			} else {
				colors.Error("Please provide a ID")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(packagesCmd)
}
