// Internal Chigo Utils
package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	chigo "github.com/UltiRequiem/chigo/pkg"
)

func PrintWithScanner(text string) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	for i := 1.0; scanner.Scan(); i++ {
		rgb := chigo.NewRGB(i)

		fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", rgb.Red, rgb.Green, rgb.Blue, scanner.Text())
	}
}

func StartProcessFromStdin() {
	reader := bufio.NewReader(os.Stdin)

	for i := 1.0; true; i++ {
		input, _, err := reader.ReadRune()

		if err != nil {
			PrintWithScanner(err.Error())
			break
		}

		rgb := chigo.NewRGB(i)

		fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m", rgb.Red, rgb.Green, rgb.Blue, string(input))
	}
}
