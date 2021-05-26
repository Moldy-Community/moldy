package utils

import (
	"os"

	colors "github.com/fatih/color"
)

func Warn(msg string) {
	yellow := colors.New(colors.FgYellow).Add(colors.Bold).Set().PrintlnFunc()
	yellow(msg)
}

func Error(msg string) {
	red := colors.New(colors.BgHiRed).Add(colors.Bold).Set().PrintlnFunc()
	red(msg)
	os.Exit(1)
}

func Info(msg string) {
	blue := colors.New(colors.FgBlue).Add(colors.Underline).Set().PrintlnFunc()
	blue(msg)
}

func Success(msg string) {
	green := colors.New(colors.FgGreen).Set().PrintlnFunc()
	green(msg)
}
