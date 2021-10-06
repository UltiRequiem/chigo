package cmd

func Main() {
	help, isThereFileArguments, files := getParametersAndFlags()

	if help {
		PrintHelp()
		return
	}

	if isThereFileArguments {
		PrintFilesWithScanner(files)
		return
	}

	startProcessFromStdin()
}
