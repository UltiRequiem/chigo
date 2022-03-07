package chigo

import (
	"bufio"
	"fmt"
	"strings"
)

func PrintWithColors(text string) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	for i := 1.0; scanner.Scan(); i++ {
		rgb := NewRGB(i)

		fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", rgb.Red, rgb.Green, rgb.Blue, scanner.Text())
	}
}
