package internal

import (
	"fmt"
	"os"
)

func JoinFiles(files []string) (string, error) {
	text := ""

	for _, file := range files {
		fileText, err := os.ReadFile(file)

		if err != nil {
			return "", fmt.Errorf("Error while trying to read file %s.", file)
		}

		text += string(fileText) + "\n"
	}

	return text, nil
}
