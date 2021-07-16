package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	cli, api string
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "show status of Moldy",
	Long:    "show  the status of CLI and API",
	Aliases: []string{"s", "st"},
	Run: func(cmd *cobra.Command, args []string) {
		if cli == "view" {

			path, _ := os.Getwd()

			if _, err := os.Stat(path + "/MoldyFile.toml"); os.IsNotExist(err) {

				fmt.Println("No MoldyFile exit pleace make a configfile")

				return
			} else {

				fmt.Println("Directory already exists")
				rf, err := os.ReadFile(path + "/MoldyFile.toml")

				if err != nil {
					log.Fatal(err)
				}

				fmt.Fprintf(os.Stdout, "----- Status -----")
				fmt.Printf("\t %s ", string(rf))
			}

		} else if api == "view" {

			client := http.Client{
				Timeout: 10 * time.Second,
			}
			start := time.Now()
			resp, err := client.Get("https://golangcode.com/http-get-with-timeout/")

			if err != nil {
				log.Fatal(err)
			}
			ms := time.Since(start).Seconds()

			fmt.Println("----- Status API -----")

			fmt.Println("Status code:", resp.StatusCode)
			if int(ms) < 1 {
				fmt.Printf("Time response ms %0.2f\n", ms)
			} else {
				fmt.Printf("Time response seconds %d\n", int(ms))
			}
		}

	},
}

func init() {

	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringVarP(&cli, "cli", "c", " ", "view status CLI")
	statusCmd.Flags().StringVarP(&api, "api", "a", " ", "view staus API")

}
