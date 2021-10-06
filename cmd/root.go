package cmd

import (
	"bufio"
	"fmt"
	. "github.com/UltiRequiem/chigo/pkg"
	"io"
	"os"
	"strings"
)

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

func Main() {
	help, isTherFiles, files := getParametersAndFlags()

	if help {
		PrintHelp()
                return
	}

	if isTherFiles {
		text := JoinFilesToString(files)
		scanner := bufio.NewScanner(strings.NewReader(text))

		var j int = 1

		for scanner.Scan() {
			r, g, b := GetRGB(j)

			fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", r, g, b, scanner.Text())

			j++
		}

		return
	}

	startProcessFromStdin()
}
