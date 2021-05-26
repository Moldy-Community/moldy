package functions

import (
	"github.com/Moldy-Community/moldy/utils"

	"github.com/manifoldco/promptui"
)

func BasicPrompt(Label, Default string) string {
	prompt := promptui.Prompt{
		Label:   Label,
		Default: Default,
	}
	result, err := prompt.Run()
	utils.CheckErrors(err, "Code 3", "Error in the input of the user", "Re run the command for fix the input with the chars utf-8")
	return result
}

func PasswordPrompt(Label string) string {
	prompt := promptui.Prompt{
		Label: Label,
		Mask:  '#',
	}
	result, err := prompt.Run()
	utils.CheckErrors(err, "Code 3", "Error in the input of the user", "Re run the command for fix the input with utf-8 chars")
	return result
}

func SelectPrompt(Label string, Items []string) string {
	prompt := promptui.Select{
		Label: Label,
		Items: Items,
	}
	_, result, err := prompt.Run()
	utils.CheckErrors(err, "Code 3", "Error in the input of the user", "Re run the command for fix the input with a valid option")
	return result
}
func YesNoQuestion(Label string) string {
	pmp := promptui.Prompt{
		Label:     Label,
		IsConfirm: true,
	}
	result, err := pmp.Run()

	utils.CheckErrors(err, "Code 3", "Error in the yes no question", "Re run the command for fix the input with a valid option")
	return result
}
