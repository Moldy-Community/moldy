package cmd

import (


	//_"github.com/Moldy-Community/moldy/core/config"

	"github.com/spf13/cobra"
)


var createInstall bool

// configCmd represents the config command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Search packages",
	Long: `:
        allow searching and extracting the content of a github packages on github
       `,
	Run: func(cmd *cobra.Command, args []string) {
		if createInstall {
			if flags,_:= cmd.Flags().GetBool("url");flags{
				//do something
			}else{
				  //packages.GetId().go
				}
		   }
	      },
	Example: "moldy i name-of-pkg or moldy install --url github.com/user/name",
	Aliases: []string{"i", "ins"},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&createInstall, "url", "u",false, "a command for download package")
}

