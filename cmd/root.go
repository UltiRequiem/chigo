package cmd

import "github.com/mattn/go-colorable"

func Main() {
	// Windows Support
	defer colorable.EnableColorsStdout(nil)()

	help, fileArguments, files := parametersAndFlags()

	if help {
		printHelp()
		return
	}

	if fileArguments {
		PrintWithScanner(JoinFilesToString(files))
		return
	}

	startProcessFromStdin()
}
