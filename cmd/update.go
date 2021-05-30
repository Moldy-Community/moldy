package cmd

import (
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions/packages"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a package providing the id",
	Long: `Update a package providing the id.
  Update the URL, or share to the world a new version of your package, since the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			packages.Update(args[0])
		} else {
			colors.Error("Please provide a ID")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
