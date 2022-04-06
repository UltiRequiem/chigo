package cmd

import (
	"fmt"
	"github.com/UltiRequiem/chigo/internal"

	chigo "github.com/UltiRequiem/chigo/pkg"

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
			fmt.Println(chigo.Colorize(error.Error()))
			return
		}

		fmt.Println(chigo.Colorize(data))
		return
	}

	internal.StartProcessFromStdin()
}
