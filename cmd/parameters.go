package cmd

import "flag"

func getParametersAndFlags() (bool, bool, []string) {
	help := flag.Bool("help", false, "Display Help")
	helpShort := flag.Bool("h", false, "Display Help")

	flag.Parse()

	return *help || *helpShort, flag.NArg() > 0, flag.Args()
}
