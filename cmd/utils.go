package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	. "github.com/UltiRequiem/chigo/pkg"
)

const VERSION = "1.0.0"

func printHelp() {

	helpMessage := ` Chigo %s

 Concatenate FILE(s) or standard input to standard output.
 When no FILE is passed read stardard input.
  
  Examples:
    chigo fOne fTwo   # Output fOne and fTwo contents.
    chigo             # Copy standard input to standard output.
    echo "My Message" # Display "My message".
    fortune | chigo   # Display a rainbow cookie.
            
 If you need more help, found a bug or want to suggest a new feature:
  https://github.com/UltiRequiem/chigo`

	PrintWithScanner(fmt.Sprintf(helpMessage, VERSION))
}

func JoinFilesToString(files []string) string {
	text := make([]string, len(files))

	for index, file := range files {

		fileText, err := os.ReadFile(file)

		if err != nil {
			fmt.Println(fmt.Sprintf("Error while trying to read file %s.", file))
			os.Exit(1)
		}

		text[index] = string(fileText)
	}

	return strings.Join(text, "\n")
}

func PrintWithScanner(text string) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	var j int = 1

	for scanner.Scan() {
		r, g, b := GetRGB(j)

		fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", r, g, b, scanner.Text())

		j++
	}
}

func startProcessFromStdin() {
	reader := bufio.NewReader(os.Stdin)

	var j int = 1

	for {
		input, _, err := reader.ReadRune()

		if err != nil && err == io.EOF {
			break
		}

		r, g, b := GetRGB(j)

		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, input)

		j++
	}
}
