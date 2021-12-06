package main

import "fmt"

func add_int(a int, b int) int {
	return a + b
}

func subtract_float(a float64, b float64) float64 {
	return a - b
}

func main() {
	fmt.Println("Add: 10 + 15", add_int(10, 15))
	fmt.Println("subtract: 10 + 15", subtract_float(10, 15))
}
