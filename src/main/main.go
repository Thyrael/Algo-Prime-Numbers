package main

import (
	"fmt"
	"math"
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

	if checkIfInputNumberIsPrimeNumber(inputAsInt) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func checkIfInputNumberIsPrimeNumber(input int) bool{
	if  math.Signbit(float64(input)) {
		return false
	}
	if input % 2 == 0 {
		if input == 2 {
			return true
		}
		return false
	}
	for d := 3; d*d <= input; d = d+2 {
		if input % d == 0 {
			return false
		}
	}
	return true
}