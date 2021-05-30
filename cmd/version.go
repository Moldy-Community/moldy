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
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of moldy :)",
	Long: `Show the version of Moldy:

In error case:
  If you have any error report on Github for fix that
  in the next version :D`,
	Run: func(cmd *cobra.Command, args []string) {
		colors.Info("Moldy 0.0.1 Alpha")
	},
	Aliases:           []string{"v", "ver"},
	Example:           "moldy version",
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
