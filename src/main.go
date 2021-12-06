package main

import "fmt"

func main() {
	//Constants
	const pi float64 = 3.141592653589793
	const euler = 2.71828
	const hello_words string = "hello world"
	fmt.Println("pi", pi)
	fmt.Println("Euler", euler)
	fmt.Println(hello_words)

	//Sariables
	var one int = 1
	var two int
	three := 3
	two = 2

	fmt.Println("one", one)
	fmt.Println("two", two)
	fmt.Println("three", three)
}
