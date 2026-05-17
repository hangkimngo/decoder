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
	-file      input file mode	
	-h         show help`
)

func main() {
	var encode, multiline bool
	var file string
	flag.BoolVar(&encode, "encode", false, "encode art instead of decoding")
	flag.StringVar(&file, "file", "", "decide input file contents instead of string")
	flag.BoolVar(&multiline, "multi", false, "decode multiple lines from argument")
	flag.Usage = func() {
		fmt.Println(usageText)
	}
	flag.Parse()

	args := flag.Args()

	var input, output string
	var err error

	if file != "" {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(errors.New("Error reading file"))
			return
		}
		stringData := string(data)
		output, err = u.Multiline(stringData, encode)
	} else if multiline {
		input := u.ScanLines()
		output, err = u.Multiline(input, encode)
	} else {
		if len(args) != 1 {
			flag.Usage()
			return
		}

		input = args[0]
		if encode {
			output, err = u.Encode(input)
		} else {
			output, err = u.Decode(input)
		}
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(output)

}
