package utils

import "os"

func ReadFile(fileName string) (string, error) {

	content, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}
	return string(content), nil
}
