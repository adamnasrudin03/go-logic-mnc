package main

import (
	"fmt"
	"strings"
)

var (
	charValid = map[string]bool{
		"<": true,
		">": true,
		"{": true,
		"}": true,
		"[": true,
		"]": true,
	}
	openCloseCharValid = map[string]string{
		"<": ">",
		">": "<",
		"{": "}",
		"}": "{",
		"[": "]",
		"]": "[",
	}
)

func main() {
	var input string

	fmt.Scan(&input)
	input = strings.ReplaceAll(input, " ", "") // remove space(input)
	for _, v := range input {
		if !charValid[string(v)] {
			fmt.Println("False")
			return
		}

	}

	fmt.Println("True")
}
