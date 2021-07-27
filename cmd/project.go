package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/Moldy-Community/moldy/core/git"
	"github.com/Moldy-Community/moldy/core/terminal"
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/spf13/cobra"
	vp "github.com/spf13/viper"
)

var (
	initFlg, gitInitFlg bool
)

var projCmd = &cobra.Command{
	Use:     "project",
	Short:   "Init a project, make commits or change the configuration of moldy",
	Long:    "Init a project with the flag --init\nMake commits in a easier way with the flag --commit\nChange the configuration of moldy with --config\n",
	Aliases: []string{"proj", "prj"},
	Example: "moldy project --init",
	Run: func(cmd *cobra.Command, args []string) {

		if initFlg {
			valueFile := map[string]interface{}{
				"moldyPackageInfo": map[string]string{
					"name":        terminal.BasicPrompt("Moldy Package Info > Name", "none"),
					"version":     terminal.BasicPrompt("Moldy Package Info > Version", "none"),
					"author":      terminal.BasicPrompt("Moldy Package Info > Author", "none"),
					"description": terminal.BasicPrompt("Moldy Package Info > Description", "none"),
					"url":         terminal.BasicPrompt("Moldy Package Info > URL", "none"),
				},
			}

			paths := []string{
				"./",
			}
			configName := "Moldy.pkg.toml"
			configType := "toml"
			_ = os.Remove("./Moldy.pkg.toml")
			for _, p := range paths {
				vp.AddConfigPath(p)
			}
			vp.SetConfigName(configName)
			vp.SetConfigType(configType)
			for k, v := range valueFile {
				vp.SetDefault(k, v)
			}

			if err := vp.SafeWriteConfigAs(configName); err != nil {
				if os.IsNotExist(err) {
					err = vp.WriteConfigAs(configName)
					functions.CheckErrors(err, "Code 2", "Error in write the config file :(", "Report the error on github or retry the command with new permmisions")
				}
			}
		}

		if gitInitFlg {
			if git.IsInstalled("git") {
				execCmd := exec.Command("git", "init")
				err := execCmd.Run()
				functions.CheckErrors(err, "Code 2", "Error in write the git init file", "Report the error on github or retry the command with new permmisions")
			} else {
				colors.Error("Install git before do this command")
			}
		}

		if len(args) > 0 {
			if args[0] == "commit" {
				if git.IsInstalled("git") {
					execCmd := exec.Command("git", "commit", "-m", fmt.Sprintf(`"%v"`, terminal.BasicPrompt("Commit message:", "New commit")))
					err := execCmd.Run()
					functions.CheckErrors(err, "Code 2", "Error doing the commit", "Report the error on github or retry the command with new permmisions")
				} else {
					colors.Error("Install git before do this command")
				}
			} else if args[0] == "add" {
				if git.IsInstalled("git") {
					execCmd := exec.Command("git", "add", terminal.BasicPrompt("Path to add in stage", "."))
					err := execCmd.Run()
					functions.CheckErrors(err, "Code 2", "Error adding the files to stage", "Report the error on github or retry the command with new permmisions")
				} else {
					colors.Error("Install git before do this command")
				}
			}

			if args[0] == "select-branch" || args[0] == "sb" {
				if git.IsInstalled("git") {
					execCmd, err := exec.Command("git", "for-each-ref", "--sort=committerdate", "refs/heads/", "--format='%(refname:short)'").Output()
					functions.CheckErrors(err, "Code 2", "Error reviewing the branches", "Report the error on github or retry the command with new permmisions")
					branchesUnfilter := strings.Split(string(execCmd), "'")
					var branchesFilter []string

					for _, value := range branchesUnfilter {
						match, _ := regexp.MatchString("[a-zA-Z0-9]", value)
						if match {
							branchesFilter = append(branchesFilter, value)
						}
					}
					branchSelected := strings.Replace(terminal.SelectPrompt("Select branch", branchesFilter), "'", "", 2)
					err = exec.Command("git", "checkout", branchSelected).Run()
					functions.CheckErrors(err, "Code 2", "Error changing the branch", "Report the error on github or retry the command with new permmisions")
				} else {
					colors.Error("Install git before do this command")
				}
			}

			if args[0] == "new-branch" || args[0] == "nb" {
				if git.IsInstalled("git") {
					nameBranch := terminal.BasicPrompt("Name of the new branch", "")
					if nameBranch == "" {
						fmt.Println("The name of the branch not is valid")
						return
					}
					execCmd := exec.Command("git", "branch", nameBranch)
					err := execCmd.Run()
					functions.CheckErrors(err, "Code 2", "Error creating the new branch the branch", "Report the error on github or retry the command with new permmisions")
					moveTo := exec.Command("git", "checkout", nameBranch)
					err = moveTo.Run()
					functions.CheckErrors(err, "Code 2", "Branch created but error in the git checkout to this branch", "Report the error on github or retry the command with new permmisions")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(projCmd)
	projCmd.Flags().BoolVarP(&initFlg, "init", "i", false, "Init a moldy project")
	projCmd.Flags().BoolVarP(&gitInitFlg, "git-init", "g", false, "Init git in a folder")
}
