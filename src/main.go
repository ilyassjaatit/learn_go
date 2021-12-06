package main

import (
	"fmt"
	"math"
)

const pi float32 = 3.141592654
const square float64 = 2

func area_circle(radius float64) float64 {
	var area float64 = math.Pow(radius, square)
	return area
}

func main() {
	const stop_for uint8 = 10
	var index uint8
	for index = 0; index <= stop_for; index++ {
		var radius float64 = float64(index)
		var area float64 = area_circle(radius)
		var message string = fmt.Sprintf("var type area %T", area)
		fmt.Println(message)

	}
}
