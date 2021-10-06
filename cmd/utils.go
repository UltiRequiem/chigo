package cmd

import (
	"fmt"
	"os"
	"strings"
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

	return strings.Join(text, "")
}

func PrintHelp() {
	fmt.Println("TODO")
}
