package chigo

import (
	"bufio"
	"fmt"
	"strings"
)

func Colorize(text string) string {
	scanner := bufio.NewScanner(strings.NewReader(text))

	result := ""

	for i := 1.0; scanner.Scan(); i++ {
		rgb := NewRGB(i)

		result += fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m\n", rgb.Red, rgb.Green, rgb.Blue, scanner.Text())
	}

	return result
}
