package main

import (
	"github.com/Moldy-Community/moldy/utils"
	"github.com/integrii/flaggy"
)

// Commands flags
var (
  InfoOpenBrowser bool
)

var (
	InfoSubCommand   *flaggy.Subcommand
	NewSubCommand    *flaggy.Subcommand
	DocSubCommand    *flaggy.Subcommand
	ConfigSubCommand *flaggy.Subcommand
)

func init() {
	flaggy.SetName("Moldy -> The best project starter and manager of the world")
	flaggy.SetDescription(`Need start a project with a structure Moldy is the solution
Start your project with a template and create a manager for the versions and many other features
Promote your library, language or framework making a template here.
The best speed provide by Golang :D.
`)
	flaggy.SetVersion("1.0")
	flaggy.DefaultParser.AdditionalHelpPrepend = "github.com/Moldy-Community/moldy"

	InfoSubCommand = flaggy.NewSubcommand("info")
	InfoSubCommand.ShortName = "about"
	InfoSubCommand.Description = "Show the information about Moldy"
  InfoSubCommand.Bool(&InfoOpenBrowser, "b", "browser", "Open in the browser the moldy repository")

	NewSubCommand = flaggy.NewSubcommand("new")
	NewSubCommand.ShortName = "init"
	NewSubCommand.Description = "Generate a new moldy project with a basic structure or a desing pattern"

	DocSubCommand = flaggy.NewSubcommand("doc")
	DocSubCommand.ShortName = "dc"
	DocSubCommand.Description = "Generate a documentation and open the doc tools"

	ConfigSubCommand = flaggy.NewSubcommand("config")
	ConfigSubCommand.ShortName = "cfg"
	ConfigSubCommand.Description = "Configure moldy for best and custom usage"

	flaggy.AttachSubcommand(InfoSubCommand, 1)
	flaggy.AttachSubcommand(NewSubCommand, 1)
	flaggy.AttachSubcommand(DocSubCommand, 1)
	flaggy.AttachSubcommand(ConfigSubCommand, 1)

	flaggy.Parse()
}

func main(){
  if InfoOpenBrowser {
    utils.OpenUrlBrowser("https://github.com/Moldy-Community/moldy")
  }
}
