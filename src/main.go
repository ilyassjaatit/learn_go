package main

import (
	"fmt"
	"math/big"
)

func fibonacci_sequence(limit uint) {
	a := big.NewInt(0)
	b := big.NewInt(1)
	var index uint = 0
	for index < limit {
		a.Add(a, b)
		a, b = b, a
		index = index + 1
		fmt.Println(a)
	}
}

func main() {
	fibonacci_sequence(15)
}
