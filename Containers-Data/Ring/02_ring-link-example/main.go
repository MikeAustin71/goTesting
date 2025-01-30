/*This example executes the Ring.Link() function
  which will link two rings together.
*/

package main

import (
	"container/ring"
	"fmt"
)

func main() {

	r := ring.New(10)

	// populate our ring
	for i := 0; i < 10; i++ {
		r.Value = i
		r = r.Next()
	}

	r2 := ring.New(10)

	// populate our ring
	for t := 10; t < 20; t++ {
		r2.Value = t
		r2 = r2.Next()
	}

	f := func(v interface{}) {
		fmt.Printf("%d ", v)
	}

	// Values in r before link
	fmt.Println("Values in r BEFORE link")
	r.Do(f)

	r = r.Prev() // Try commenting this line an observing result

	fmt.Println()
	fmt.Println()
	fmt.Println("Values in r2 BEFORE link")
	r2.Do(f)

	r3 := r.Link(r2)

	// Values in r AFTER link
	fmt.Println()
	fmt.Println()
	fmt.Println("Values in r3 AFTER link")
	r3.Do(f)
	fmt.Println()

}

/*	Output
	$ go run main.go
	Values in r BEFORE link
	0 1 2 3 4 5 6 7 8 9

	Values in r2 BEFORE link
	10 11 12 13 14 15 16 17 18 19

	Values in r3 AFTER link
	0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19
*/
