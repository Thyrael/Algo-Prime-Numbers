package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
	input := os.Args[1]
	// first check if input is a number
	inputAsString := string(input)
	inputAsString = strings.TrimSuffix(inputAsString, "\n")
	inputAsInt, err := strconv.Atoi(inputAsString)

	if err != nil {
		fmt.Println("Erreur : ", err)
		return
	}

	if inputAsInt % 2 == 0 {
		if inputAsInt == 2 {
			fmt.Println("true")
			return
		}
		fmt.Println("false")
		return
	}
	for d := 3; d*d <= inputAsInt; d = d+2 {
		if inputAsInt % d == 0 {
			fmt.Println("false") // Composed number
		}
	}
	fmt.Println("true") // Prime number
}