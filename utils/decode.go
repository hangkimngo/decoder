package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Decode(input string) (string, error) {
	// TODO: implement decoder

	// {	re := regexp.MustCompile(`\[(\d+) ([^\]]+)\]`)

	re := regexp.MustCompile(`\[(.*?)\]`)
	var decodeErr error

	output := re.ReplaceAllStringFunc(input, func(match string) string {
		content := match[1 : len(match)-1]

		spaceIndex := strings.Index(content, " ")
		if spaceIndex == -1 {
			decodeErr = errors.New("Error")
		}
		parts := strings.SplitN(content, " ", 2)
		fmt.Println(parts)
		if len(parts) != 2 {
			decodeErr = errors.New("Error")
			return ""
		}
		firstArg := content[:spaceIndex]
		secondArg := content[spaceIndex+1:]

		if secondArg == "" {
			decodeErr = errors.New("Error")
			return ""
		}

		count, err := strconv.Atoi(firstArg)
		if err != nil {
			decodeErr = errors.New("Error")
			return ""
		}

		return strings.Repeat(secondArg, count)
	})

	if decodeErr != nil {
		return "", decodeErr
	}

	return output, nil
}
