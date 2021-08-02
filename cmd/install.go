package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
    "errors"
    "strings"

	// gitF "github.com/Moldy-Community/moldy/core/git"
	// "github.com/Moldy-Community/moldy/core/locks"
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/cavaliercoder/grab"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Search packages",
	Long: `:
        allow searching and extracting the content of a github packages on github
       `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			resp, err := grab.Get(functions.GetCacheDir(), args[0])
			
            functions.CheckErrors(err, "Code 2", "Error in downlaod the cache", "Check your connection or report the error on github")
			 
			var dirfile string
            
             err := filepath.Walk(functions.GetCacheDir(), func(path string, info os.FileInfo, err error) error {

                dirs = strinsg.Split(path,"\n")

                   for i := 0; i < len(dirs); i++ {
		              	if strings.Contains(dirs[i], "cache.json") {
				            dirfile = dirs[i]
		      	           break
                        } 
		              }
                  if dirfile == ""{
                    return errors.New("File not found") 
			       }
		         }
		return nil
	})
            
            
            dat, err2 := ioutil.ReadFile(dirlfile)
			functions.CheckErrors(err2, "Code 2", "Error in read the file check if exists ", "Report the error on github or check if file exists")
			colors.Info(string(dat))
		} else {
			fmt.Print(len(args))
			colors.Error("Error in the create install flag")
		}
	},
	Example: "moldy i name-of-pkg or moldy install --url github.com/user/name",
	Aliases: []string{"i", "ins"},
}

func init() {
	rootCmd.AddCommand(installCmd)
}