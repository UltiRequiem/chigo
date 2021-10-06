package cmd

func Main() {
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
