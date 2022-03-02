package cmd

import (
	"github.com/UltiRequiem/chigo/internal"

	"github.com/mattn/go-colorable"
)

func Main() {
	// Windows Support
	defer colorable.EnableColorsStdout(nil)()

	help, fileArguments, files := parametersAndFlags()

	if help {
		printHelp()
		return
	}

	if fileArguments {
		data, error := internal.JoinFiles(files)

		if error != nil {
			internal.PrintWithScanner(error.Error())
			return
		}

		internal.PrintWithScanner(data)
		return
	}

	internal.StartProcessFromStdin()
}
