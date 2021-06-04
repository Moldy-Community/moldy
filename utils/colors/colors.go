package colors

import (
	"fmt"
	"os"

	moldyConfig "github.com/Moldy-Community/moldy/core/files/configMoldy"
	colors "github.com/gookit/color"
)

func Warn(msg ...string) {
	colors.Enable = !moldyConfig.Settings().AparienceOptions.ColorsMode
	colors.Yellow.Print("WARN:\n")
	fmt.Printf(" %v", msg)
}

func Error(msg ...string) {
	colors.Enable = !moldyConfig.Settings().AparienceOptions.ColorsMode
	colors.Red.Print("ERROR:\n")
	fmt.Printf(" %v", msg)
	os.Exit(1)
}

func Info(msg ...string) {
	colors.Enable = !moldyConfig.Settings().AparienceOptions.ColorsMode
	colors.Cyan.Print("INFO:\n")
	fmt.Printf(" %v", msg)
}

func Success(msg ...string) {
	colors.Enable = !moldyConfig.Settings().AparienceOptions.ColorsMode
	colors.Green.Print("SUCCESS:\n")
	fmt.Printf(" %v", msg)
}
