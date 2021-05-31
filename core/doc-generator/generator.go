package docgenerator

import (
	"bytes"
	"os"

	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func GenDocTree(cmd *cobra.Command) {
	_, err := os.Stat("moldyDoc")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("moldyDoc", 0755)
		functions.CheckErrors(errDir, "Code 2", "Error in create the moldyDoc folder", "Report the error on github or re run with sudo")
	}
	errDoc := doc.GenMarkdownTree(cmd, "./moldyDoc")
	functions.CheckErrors(errDoc, "Code 2", "Error in generate the documentation for moldy", "Report the error on github or re run with sudo")
}

func DocCommand(cmd *cobra.Command) {
	out := new(bytes.Buffer)
	err := doc.GenMarkdown(cmd, out)
	functions.CheckErrors(err, "Code 2", "Error in print out the doc for the command", "Report the error on github or re run with sudo")
}
