package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // put value on channel
		}
		close(c) // close the channel
	}()

	for n := range c { // range pulls values off the channel
		fmt.Println(n)
	}
}

/*	Output
	$ go run main.go
	0
	1
	2
	3
	4
	5
	6
	7
	8
	9
*/
