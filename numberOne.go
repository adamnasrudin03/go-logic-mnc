package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		input  string
		N      int64
		inputN []string
		err    error
	)

	fmt.Println("Contoh input:")
	fmt.Scan(&input)
	N, err = strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("N Hanya menerima bilangan bulat")
		return
	}
	// scan input sebanyak N
	for i := 0; i < int(N); i++ {
		fmt.Scan(&input)
		input = strings.TrimSpace(strings.ToLower(input))
		inputN = append(inputN, input)
	}

	tmp := map[string]int64{}
	key := ""
	for _, v := range inputN {
		tmp[v] = tmp[v] + 1
		if tmp[v] == 2 {
			key = v
			break
		}
	}
	if key == "" {
		fmt.Println("Contoh output: false")
		return
	}

	indeksKey := []int{}
	for i, v := range inputN {
		if v == key {
			indeksKey = append(indeksKey, i+1)
		}
	}

	txtIndex := ""
	for _, v := range indeksKey {
		txtIndex = txtIndex + strconv.Itoa(v) + " "
	}
	fmt.Println("Contoh output:", txtIndex)
}
