package main

import (
	"fmt"
)

func main() {

	outputs := make([]chan int, 0, 10)

	out := make(chan int)

	go gen(out)

	outputs = append(outputs, out)

	for _, n := range outputs {

		for c := range n {
			fmt.Println(c)
		}
	}

}

func gen(c chan<- int) {

	for i := 0; i < 25; i++ {
		for j := 3; j < 13; j++ {
			c <- j
		}
	}

	close(c)
}
