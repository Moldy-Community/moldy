/*
Copyright © 2021 Moldy Community

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

  $ source <(moldy completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ moldy completion bash > /etc/bash_completion.d/moldy
  # macOS:
  $ moldy completion bash > /usr/local/etc/bash_completion.d/moldy

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ moldy completion zsh > "${fpath[1]}/_moldy"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ moldy completion fish | source

  # To load completions for each session, execute once:
  $ moldy completion fish > ~/.config/fish/completions/moldy.fish

PowerShell:

  PS> moldy completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> moldy completion powershell > moldy.ps1
  # and source this file from your PowerShell profile.

In error case:
  If you have any error report on Github for fix that
  in the next version :D
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Aliases:               []string{"comp"},
	Example:               "moldy comp bash",
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
