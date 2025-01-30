package main

import "fmt"

/*
	This example ranges over a string returning only
	the index.
*/

func main() {
	str := "Hello World!"

	for i := range str {
		fmt.Printf("index=%d\n", i)
	}
}

/* Output
$ go run main.go
index=0
index=1
index=2
index=3
index=4
index=5
index=6
index=7
index=8
index=9
index=10
index=11
*/
