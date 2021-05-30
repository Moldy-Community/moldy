package cmd

import (
	"github.com/Moldy-Community/moldy/utils/functions/packages"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a package since the terminal.",
	Long: `Create a package since the terminal, be fast.
  In 30 seconds you can create your own package`,
	Run: func(cmd *cobra.Command, args []string) {
		packages.Create()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
