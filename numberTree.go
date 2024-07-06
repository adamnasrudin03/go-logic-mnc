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
		"{": "}",
		"[": "]",
	}

	openCharValid = map[string]bool{
		"<": true,
		"{": true,
		"[": true,
	}

	closeCharValid = map[string]bool{
		">": true,
		"}": true,
		"]": true,
	}
)

func main() {
	var input string

	fmt.Scan(&input)
	input = strings.ReplaceAll(input, " ", "") // remove space(input)

	// check valid input and count char
	countOpenChar := map[string]int{}
	countCloseChar := map[string]int{}
	for i, v := range input {
		val := string(v)
		// Check char first
		if i == 0 && !openCharValid[val] {
			fmt.Println("False")
			return
		}
		// Check char last
		if i == len(input)-1 && !closeCharValid[val] {
			fmt.Println("False")
			return
		}

		// check valid char
		if !charValid[val] {
			fmt.Println("False")
			return
		}

		if openCharValid[val] {
			countOpenChar[val]++
		} else if closeCharValid[val] {
			countCloseChar[val]++
		}
	}

	// validate char is not closed
	for _, v := range input {
		val := string(v)
		openCharIsNotClosed := openCharValid[val] && countOpenChar[val] != countCloseChar[openCloseCharValid[val]]
		if openCharIsNotClosed {
			fmt.Println("False")
			return
		}
	}

	fmt.Println("True")
}
