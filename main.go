package main

import (
	"decoder/utils"

	"fmt"
	"os"
)

const (
	usageText = `art decoder usage:
go run . "[5 #][5 -_]-[5 #]"

output:
#####-_-_-_-_-_-#####`
)

func main() {
	if os.Args[1] == "-h" {
		fmt.Println(usageText)
		return
	}

	if len(os.Args) != 2 {
		fmt.Println(usageText)
		return
	}

	input := os.Args[1]
	if !utils.BalanceBracket(input) {
		fmt.Println("Error")
		return
	}
	fmt.Println()

}
