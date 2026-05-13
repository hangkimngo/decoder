package main

import (
	u "decoder/utils"
	"errors"
	"flag"
	"fmt"
	"os"
)

const (
	usageText = `Usage:
go run . [OPTION] "[encoded text]"

Options:
	-encode    encode mode
	-multi     multi-line mode
	-h         show help`
)

func main() {
	var encode, fileMode bool
	flag.BoolVar(&encode, "encode", false, "encode art instead of decoding")
	flag.BoolVar(&fileMode, "file", false, "input file instead of string")

	flag.Usage = func() {
		fmt.Println(usageText)
	}
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		flag.Usage()
		return
	}

	var input, output string
	var err error

	if fileMode {
		data, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println(errors.New("Error reading file"))
		}
		stringData := string(data)
		output, err = u.DecodeMultiLines(stringData)
	} else {
		input = args[0]
		output, err = u.Decode(input)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(output)

}
