package cmd

import (
	"os"
	"strings"

	"github.com/Moldy-Community/moldy/core/terminal"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
)

var (
	initFlg bool
)

var projCmd = &cobra.Command{
	Use:     "project",
	Short:   "Init a project, make commits or change the configuration of moldy",
	Long:    "Init a project with the flag --init\nMake commits in a easier way with the flag --commit\nChange the configuration of moldy with --config\n",
	Aliases: []string{"proj", "prj"},
	Example: "moldy project --init",
	Run: func(cmd *cobra.Command, args []string) {

		if initFlg {
			content := map[string]interface{}{
				"moldyPackageInfo": map[string]string{
					"name":        terminal.BasicPrompt("Moldy Package Info > Name", "none"),
					"version":     terminal.BasicPrompt("Moldy Package Info > Version", "none"),
					"author":      terminal.BasicPrompt("Moldy Package Info > Author", "none"),
					"description": terminal.BasicPrompt("Moldy Package Info > Description", "none"),
					"url":         terminal.BasicPrompt("Moldy Package Info > URL", "none"),
				},
			}

			dir, _ := os.Getwd()
			homeDir, _ := os.UserHomeDir()
			points := ""
			dirArr := strings.Split(strings.Replace(dir, "/", "", 1), "/")
			homeDirArr := strings.Split(strings.Replace(homeDir, "/", "", 1), "/")

			difference := len(dirArr) - len(homeDirArr)

			for i := 1; i <= difference*2; i++ {
				if i%3 == 0 {
					points += "/"
				} else {
					points += "."
				}
			}

			lastDir := strings.Split(strings.Replace(dir, homeDir, "", 1), "/")
			lastDir[1] = ""
			lastDirStr := strings.Replace(strings.Join(lastDir, "/"), "/", "", 1)
			file, err := os.Create(points + lastDirStr + "/Moldy.pkg.toml")

			functions.CheckErrors(err, "1", "Something bad happened with the path", "Unknown exact solution. Leave please the issue in github.com/Moldy-Community/moldy")

			err = toml.NewEncoder(file).Encode(content)

			functions.CheckErrors(err, "1", "Something bad happened at the moment when the file was attempted to be created", "Unknown exact solution. Leave please the issue in github.com/Moldy-Community/moldy")

		}
	},
}

func init() {
	rootCmd.AddCommand(projCmd)
	projCmd.Flags().BoolVarP(&initFlg, "init", "i", false, "Init a moldy project")
}
