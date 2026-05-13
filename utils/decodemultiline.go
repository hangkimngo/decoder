package utils

import (
	"strings"
)

func DecodeMultiLines(text string) (string, error) {
	lines := strings.Split(text, "\n")

	var decodedLines []string

	for _, line := range lines {
		rowOutput, err := Decode(line)
		if err != nil {
			return "", err
		}
		decodedLines = append(decodedLines, rowOutput)
	}
	output := strings.Join(decodedLines, "\n")

	return output, nil
}
