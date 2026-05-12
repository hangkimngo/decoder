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
		parts := strings.SplitN(content, " ", 2)
		fmt.Println(parts)
		if len(parts) != 2 {
			decodeErr = errors.New("invalid format")
			return ""
		}

		count, err := strconv.Atoi(parts[0])
		fmt.Println(count)
		if err != nil {
			decodeErr = errors.New("The first argument is not a number.")
			return ""
		}
		chars := parts[1]
		return strings.Repeat(chars, count)
	})

	if decodeErr != nil {
		return "", decodeErr
	}

	return output, nil
}
