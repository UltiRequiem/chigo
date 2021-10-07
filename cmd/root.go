package cmd

import (
	"github.com/mattn/go-colorable"
)

func Main() {
    defer colorable.EnableColorsStdout(nil)()

	help, isThereFileArguments, files := getParametersAndFlags()

	if help {
		PrintHelp()
		return
	}

	if isThereFileArguments {
		text := JoinFilesToString(files)
		PrintWithScanner(text)
		return
	}

	startProcessFromStdin()
}
