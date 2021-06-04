package git

import (
	"os"

	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/go-git/go-git/v5"
)

func CloneRepos(url string) {

	path, err := os.Getwd()

	functions.CheckErrors(err, "Code 2", "Error in get the directory", "Check the permissions or report the error on github")

	colors.Info("Clonning the repository now...\n")
	_, err2 := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	functions.CheckErrors(err2, "Code 2", "Error in clone the template", "Check the permissions, check if are installed git")
	colors.Success("\nSuccesfuly cloned the repository in the current directory")
}
