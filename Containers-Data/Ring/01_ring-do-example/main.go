/*This example creates a ring container and
  executes the Ring.Do() function
*/
package main

import (
	"container/ring"
	"fmt"
)

func main() {

	r := ring.New(5) // Create a new ring

	i := 0
	f := func(v interface{}) {
		v = i
		i++
		fmt.Printf("v = %d\n", v)
	}

	r.Do(f) // Ring.Do(f) calls the passed function
	// on every element in the ring.
}

/*	Output
	$ go run main.go
	v = 0
	v = 1
	v = 2
	v = 3
	v = 4
*/
