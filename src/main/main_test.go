package main

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	list := [6]int{2, 3, 5, 7, 11, 13}
	for _, item := range list {
		if checkIfInputNumberIsPrimeNumber(item) != true {
			t.Errorf("%v has to be prime", item)
		}
	}
}

func TestIsNotPrime(t *testing.T) {
	list := [5]int{-1,-4,-7,4,10}
	for _, item := range list {
		if checkIfInputNumberIsPrimeNumber(item) != false {
			t.Errorf("%v is not prime", item)
		}
	}
}

