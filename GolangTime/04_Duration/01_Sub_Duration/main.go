package main

import (
	"fmt"
	"time"
)

func main() {

	t0 := time.Now()
	// Do a slow operation.
	count := 0
	for i := 0; i < 100000; i++ {
		for x := 0; x < i; x++ {
			count++
		}
	}

	t1 := time.Now()
	// Display result.
	fmt.Println("Count:", count)
	// Get duration.
	d := t1.Sub(t0)
	fmt.Println("Duration", d)
	// Get miliseconds from duration
	ms := d.Seconds() * 1000.0
	fmt.Println("Miliseconds", ms)
	// Get seconds from duration.
	s := d.Seconds()
	fmt.Println("Seconds", s)
	// Get minutes from duration.
	m := d.Minutes()
	fmt.Println("Minutes", m)
}

/*	Output
	$ go run main.go
	Count: 4999950000
	Duration 1.4048858s
	Miliseconds 1404.8858
	Seconds 1.4048858
	Minutes 0.023414763333333335
*/
