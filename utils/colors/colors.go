package colors

import (
	"os"

	moldyConfig "github.com/Moldy-Community/moldy/core/files/configMoldy"
	colors "github.com/fatih/color"
)

func Warn(msg string) {
	colors.NoColor = !moldyConfig.Settings().AparienceOptions.ColorsMode
	yellow := colors.New(colors.FgYellow).Add(colors.Bold).Set().PrintlnFunc()
	yellow(msg)
}

func Error(msg string) {
	colors.NoColor = !moldyConfig.Settings().AparienceOptions.ColorsMode
	red := colors.New(colors.BgHiRed).Add(colors.Bold).Set().PrintlnFunc()
	red(msg)
	os.Exit(1)
}

func Info(msg string) {
	colors.NoColor = !moldyConfig.Settings().AparienceOptions.ColorsMode
	blue := colors.New(colors.FgBlue).Add(colors.Underline).Set().PrintlnFunc()
	blue(msg)
}

func Success(msg string) {
	colors.NoColor = !moldyConfig.Settings().AparienceOptions.ColorsMode
	green := colors.New(colors.FgGreen).Set().PrintlnFunc()
	green(msg)
}
