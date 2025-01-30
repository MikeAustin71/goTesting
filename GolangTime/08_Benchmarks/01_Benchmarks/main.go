/*
Performance, map versus slice. A map accelerates lookups: it uses a special hash code to guide a search. Here I compare a map search to a slice search.
And:
For both the map and the slice, we try to locate the index of the value "bird." The map has indexes as its values.
Result:
The map has a clear performance advantage in this case. For larger amounts of data, this will become even greater.
Golang that benchmarks map, slice
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	lookup := map[string]int{
		"cat":  0,
		"dog":  1,
		"fish": 2,
		"possum": 3,
		"wolf": 4,
		"bear": 5,
		"fox" : 6,
		"armadillo": 7,
		"rattle snake": 8,
		"raccoon": 9,
		"bird": 10}

	values := []string{"cat",
						"dog",
						"fish",
						"possum",
						"wolf",
						"bear",
						"fox",
						"armadillo",
						"rattle snake",
						"raccoon",
						"bird"}
	temp := 0

	t0 := time.Now()

	const cycles int = 10000000

	// Version 1: search map with lookup.
	for i := 0; i < cycles; i++ {
		v, ok := lookup["bird"]
		if ok {
			temp = v
		}
	}

	t1 := time.Now()

	// Version 2: search slice with for-loop.
	for i := 0; i < cycles; i++ {
		for x := range values {
			if values[x] == "bird" {
				temp = x
				break
			}
		}
	}

	t2 := time.Now()
	// Benchmark results.
	fmt.Println("Results")
	fmt.Println("")
	fmt.Println(temp)
	fmt.Println(t1.Sub(t0), "map lookup")
	fmt.Println(t2.Sub(t1), "slice for-loop")
}

/* Output
	$ go run main.go
	Results

	10
	156.1324ms map lookup
	258.423ms slice for-loop
*/
