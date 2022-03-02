package cmd

import (
	"flag"
	"fmt"

	"github.com/UltiRequiem/chigo/internal"
)

func parametersAndFlags() (bool, bool, []string) {
	help := flag.Bool("help", false, "Display Help")
	helpShort := flag.Bool("h", false, "Display Help")

	flag.Usage = printHelp

	flag.Parse()

	return *help || *helpShort, flag.NArg() > 0, flag.Args()
}

func printHelp() {
	internal.PrintWithScanner(fmt.Sprintf(HELP_MESSAGE, VERSION))
}

const VERSION = "1.0.0"

const HELP_MESSAGE = ` Chigo %s

 Concatenate FILE(s) or standard input to standard output.
 When no FILE is passed read standard input.
  
  Examples:
    chigo fOne fTwo           # Output fOne and fTwo contents.
    chigo                     # Copy standard input to standard output.
    echo "My Message" | chigo # Display "My message".
    fortune | chigo           # Display a rainbow cookie.
            
 If you need more help, found a bug or want to suggest a new feature:
  https://github.com/UltiRequiem/chigo`
