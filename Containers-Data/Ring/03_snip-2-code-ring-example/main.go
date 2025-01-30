/*Snip2Code Ring Container Example
  https://www.snip2code.com/Snippet/13705/container-ring-example-golang

*/
package main

import (
	"container/ring"
	"fmt"
	"time"
)

func main() {
	coffee := []string{"kenya", "guatemala", "ethiopia"}

	// create a ring and populate it with some values
	r := ring.New(len(coffee))
	for i := 0; i < r.Len(); i++ {
		r.Value = coffee[i]
		r = r.Next()
	}

	// print all values of the ring, easy done with ring.Do()
	fmt.Println("Print all ring values with ring.Do()")
	r.Do(func(x interface{}) {
		fmt.Println(x)
	})

	// .. or each one by one.
	fmt.Println()
	fmt.Println("=========================================================")
	fmt.Println("Cycling through each ring value one by one for 15-seconds")
	fmt.Println()
	i := 0
	for _ = range time.Tick(time.Second * 1) {

		i++

		if i == 15 {
			fmt.Println()
			fmt.Println("=========================================================")
			fmt.Println("Exiting program after cycling through ring for 15-seconds.")
			break
		}

		r = r.Next()
		fmt.Println(r.Value)
	}
}

/*	Output
	$ go run main.go
	Print all ring values with ring.Do()
	kenya
	guatemala
	ethiopia

	=========================================================
	Cycling through each ring value one by one for 15-seconds

	guatemala
	ethiopia
	kenya
	guatemala
	ethiopia
	kenya
	guatemala
	ethiopia
	kenya
	guatemala
	ethiopia
	kenya
	guatemala
	ethiopia

	=========================================================
	Exiting program after cycling through ring for 15-seconds.
*/
