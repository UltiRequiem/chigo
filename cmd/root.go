package cmd

import "github.com/mattn/go-colorable"

func Main() {
	// Windows Support
	defer colorable.EnableColorsStdout(nil)()

	help, thereFileArguments, files := getParametersAndFlags()

	if help {
		PrintHelp()
		return
	}

	if thereFileArguments {
		text := JoinFilesToString(files)
		PrintWithScanner(text)
		return
	}

	startProcessFromStdin()
}
