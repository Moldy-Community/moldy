package utils

import (
	colors "github.com/fatih/color/"
)

func Warn(msg string) {
	yellow := colors.New(colors.FgYellow).Add(colors.Bold).PrintfFunc()
	yellow(msg)
}

func Error(msg string) {
	red := colors.New(colors.BgHiRed).Add(colors.Bold).PrintfFunc()
	red(msg)
}

func Info(msg string) {
	blue := colors.New(colors.FgBlue).Add(colors.Underline).PrintfFunc()
	blue(msg)
}

func Success(msg string) {
	green := colors.New(colors.FgGreen).PrintfFunc()
	green(msg)
}
