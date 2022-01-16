package main

import (
	"fmt"
	"time"
)

type Sequence struct {
	Position uint64
	Value    uint64
}

func fibonacci_sequence(limit uint64, ch chan *int64) [20]Sequence {
	var a uint64
	var b uint64
	a = 0
	b = 1
	var index uint64 = 0
	var list_sequence_fibonacci [20]Sequence
	var seq Sequence
	for index < limit {
		a = a + b
		seq = Sequence{Position: index, Value: a}
		list_sequence_fibonacci[index] = seq
		a, b = b, a
		index = index + 1
		time.Sleep(2 * time.Second)
		data := <-ch
		fmt.Printf("end fibonacci_sequence %d \n ", data)
	}

	return list_sequence_fibonacci

}
func printData(c chan *int) {
	time.Sleep(time.Second * 3)
	data := <-c
	fmt.Println("Data in channel is: ", *data)
}

func main() {
	var a int64
	fmt.Println(a)

}
