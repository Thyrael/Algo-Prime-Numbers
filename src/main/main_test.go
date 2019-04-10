package main

import (
	"testing"
	"fmt"
)

func TestPrimeNumber(t *testing.T) {
	fmt.Println("test method")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	noprimes := [6]int{-1,0,4,-8,10,-2}
	for _, prime := range primes {
		fmt.Println(prime, "is prime :", checkIfInputNumberIsPrimeNumber(prime))
	}

	for _, noprime := range noprimes {
		fmt.Println(noprime, "is prime :", checkIfInputNumberIsPrimeNumber(noprime))
	}

}
