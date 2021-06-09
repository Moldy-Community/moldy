package cmd

import (
	"fmt"

	gitF "github.com/Moldy-Community/moldy/core/git"
	"github.com/Moldy-Community/moldy/core/locks"
	"github.com/Moldy-Community/moldy/core/packages"
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/spf13/cobra"
)

var (
	createInstall string
	nameInstall   bool
)

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
			locks.WriteLockUrl(urlClone)
		} else if nameInstall {
			data, err := packages.GetSearch(args[0])
			functions.CheckErrors(err, "Code 2", "Error in get the api package", "Check if package exists or report the bug on github")
			info := data.Data
			for index, val := range info {
				if index == 0 {
					colors.Info("Cloning the repo: ")
					fmt.Print(val.Name)
					gitF.CloneRepos(val.Url)
					locks.WriteLock(val.Name)
					break
				}
			}
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
	installCmd.Flags().BoolVarP(&nameInstall, "pkg", "p", false, "Insert the name and search in the moldy api the package")
}
