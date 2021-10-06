package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	. "github.com/UltiRequiem/chigo/pkg"
)

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

func PrintHelp() {
	fmt.Println("TODO")
}

func PrintFilesWithScanner(files []string) {
	text := JoinFilesToString(files)
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


