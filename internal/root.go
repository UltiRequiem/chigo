// Internal Chigo Utils
package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/UltiRequiem/chigo/pkg"
)

func PrintWithScanner(text string) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	var j int = 1

	for scanner.Scan() {
		r, g, b := chigo.GetRGB(j)

		fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", r, g, b, scanner.Text())

		j++
	}
}

func StartProcessFromStdin() {
	reader := bufio.NewReader(os.Stdin)

	var j int = 1

	for {
		input, _, err := reader.ReadRune()

		if err != nil && err == io.EOF {
			break
		}

		r, g, b := chigo.GetRGB(j)

		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, input)

		j++
	}
}
