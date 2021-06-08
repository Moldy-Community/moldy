package cmd

import (
	"fmt"
	"os"

	moldyConfig "github.com/Moldy-Community/moldy/core/files/configMoldy"
	"github.com/Moldy-Community/moldy/core/packages"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search packages",
	Long: `Search packages by the name or with the id, and get info about them.
  use -name if you want search many by the name of the packages
  use -id if you want search by id only a especific package`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the params")
		} else if len(args) < 2 {
			fmt.Println("Please provide the search key")
		} else {
			if args[0] == "by-name" || args[0] == "by-n" {
				getData, err := packages.GetSearch(args[1])

				if err != nil {
					fmt.Println(err)
					return
				}

				data := getData.Data

				if moldyConfig.Settings().AparienceOptions.AsciiArt {

					var dataTable [][]string

					for _, value := range data {
						dataTable = append(dataTable, []string{value.Name, value.Url, value.Id})
					}

					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeader([]string{"Name", "URL", "ID"})

					for _, v := range dataTable {
						table.Append(v)
					}

					table.Render()
				} else {
					for i, value := range data {
						fmt.Printf("\n%v.\nName: %v\nURL:%v\nID:%v\n", i+1, value.Name, value.Url, value.Id)
					}
				}
			} else if args[0] == "by-id" {
				getData, err := packages.GetId(args[1])

				if err != nil {
					fmt.Println("This id not was found in any package")
					return
				}

				data := getData.Data
				dataTable := [][]string{{data.Name, data.Url, data.Id}}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Name", "URL", "ID"})

				for _, v := range dataTable {
					table.Append(v)
				}

				table.Render()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
