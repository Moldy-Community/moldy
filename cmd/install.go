package cmd

import (
	gitF "github.com/Moldy-Community/moldy/core/git"
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/spf13/cobra"
)

var createInstall string

// configCmd represents the config command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Search packages",
	Long: `:
        allow searching and extracting the content of a github packages on github
       `,
	Run: func(cmd *cobra.Command, args []string) {
		if createInstall != "" {
			urlClone := "https://" + createInstall + ".git"
			gitF.CloneRepos(urlClone)
		} else {
			colors.Error("Error in the create install flag")
		}
	},
	Example: "moldy i name-of-pkg or moldy install --url github.com/user/name",
	Aliases: []string{"i", "ins"},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().StringVarP(&createInstall, "url", "u", "github.com/Moldy-Community/moldy", "Clone the repository from a url")
}
