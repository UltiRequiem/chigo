package internal

import "os"

func JoinFiles(files []string) (string, error) {
	text := ""

	for _, file := range files {
		fileText, err := os.ReadFile(file)

		if err != nil {
			return "", err
		}

		text += string(fileText) + "\n"
	}

	return text, nil
}
