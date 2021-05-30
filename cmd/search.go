package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Moldy-Community/moldy/utils/functions/packages"
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
					fmt.Println("This name not match with any package")
					return
				}
				data := getData.Data
				var dataTable [][]string
				for i, value := range data {
					dataTable = append(dataTable, []string{strconv.Itoa(i + 1), value.Author, value.Name, value.Url, value.Version, value.Id})
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Number", "Author", "Name", "URL", "Version", "ID"})

				for _, v := range dataTable {
					table.Append(v)
				}

				table.Render()
			} else if args[0] == "by-id" {
				getData, err := packages.GetId(args[1])
				if err != nil {
					fmt.Println("This id not was found in any package")
					return
				}
				data := getData.Data
				dataTable := [][]string{{"1", data.Author, data.Name, data.Url, data.Version, data.Id}}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Number", "Author", "Name", "URL", "Version", "ID"})
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
