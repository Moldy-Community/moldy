package core

import (
	"io/ioutil"

	"github.com/Moldy-Community/moldy/utils"
)

func CreateDotFiles() {
	/* Define the basic questions of dot files */
	ediCfg := BasicPrompt("Create editor config file defaults ?", "yes")
	gitIgno := BasicPrompt("Create a .gitignore file with defaults ?", "yes")
	/* Ask to create the other files  */
	utils.Info("All dot files toggle create other files ? ")
	otherFil := BasicPrompt("Create other files Example: README.md, LICENSE ? ", "yes")
	readme := BasicPrompt("Create a example README.md ? ", "yes")
	licenseFile := BasicPrompt("Create a LICENSE file ? LICENSE: Apache 2.0", "yes")

	/* Write and switch in the dot files */

	switch ediCfg {
	case "yes":
		editcfgCont := []byte(`

root = true

[*]
end_of_line = lf
insert_final_newline = true
indent_style = space
charset = utf-8
trim_trailing_whitespace = true
indent_size = 2

##################################

Made by Moldy CLI

##################################


`)
		err := ioutil.WriteFile("./.editorconfig", editcfgCont, 0666)
		utils.CheckErrors(err, "Code 2", "Error in generate the .editorconfig file", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")
	case "no":
		utils.Info("Ok skipping the creation .editorconfig file")
	default:
		utils.Error("Error option not found :c")
	}
	switch gitIgno {
	case "yes":
		gitignoreCont := []byte(`

    # Ignoring the .env files

    *.env

    # Editor ignore folders

    *.idea/
    *.vscode/
    *.vim/


    ##################################

    Made by Moldy CLI

    ##################################
    `)
		err := ioutil.WriteFile("./.gitignore", gitignoreCont, 0666)
		utils.CheckErrors(err, "Code 2", "Error in generate the .gitignore file", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")

	case "no":
		utils.Info("Ok skipping the creation .gitignore file")
	default:
		utils.Error("Error option not found :c")
	}

	/* Validate the other files input */
	if otherFil == "yes" {
		/* Switch in the other files */
		switch readme {
		case "yes":
			readmeCont := []byte(`

      # Example package

      Example content and example description :p

      # Credits

      **Author:** Jhon Doe
      **Version:** v1.0.0
      **Url:** [here](https://example.com)

      > Made with :heart: by Moldy CLI [here](https://github.com/Moldy-Community/moldy)

      `)
			err := ioutil.WriteFile("./README.md", readmeCont, 0666)
			utils.CheckErrors(err, "Code 2", "Error in generate the README.md file", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")
		case "no":
			utils.Info("Ok skipping the creation README.md file")
		default:
			utils.Error("Error option not found :c")
		}

		switch licenseFile {
		case "yes":
			licenseContent := []byte(`


Copyright © 2021 Example Author

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


      `)

			err := ioutil.WriteFile("./LICENSE", licenseContent, 0666)
			utils.CheckErrors(err, "Code 2", "Error in generate the LICENSE file", "Check the permissions, or re run with the sudo or admin permissions or report the error on github")
		case "no":
			utils.Info("Ok skipping the creation LICENSE file")
		default:
			utils.Error("Error option not found :c")
		}
	}
}
