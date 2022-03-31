// Internal Chigo Utils
package internal

import (
	"bufio"
	"fmt"
	"os"

	chigo "github.com/UltiRequiem/chigo/pkg"
)

func StartProcessFromStdin() {
	reader := bufio.NewReader(os.Stdin)

	for i := 1.0; true; i++ {
		input, _, err := reader.ReadRune()

		if err != nil {
			chigo.PrintWithColors(err.Error())
			break
		}

		rgb := chigo.NewRGB(i)

		fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m", rgb.Red, rgb.Green, rgb.Blue, string(input))
	}
}
