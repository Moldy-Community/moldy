package cmd

import (
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions/packages"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a package",
	Long: `Delete a package with the command:

  moldy delete <id> <password_package>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			colors.Error("The command need two params, the correct way to do this command is:\nmoldy delete <id> <password_package>\n")
		} else {
			packages.Delete(args[0], args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
