package main

import "fmt"


func first(x int, _ int) int {
	return x + 1
}

func main() {
	a := 1
	b := 2
	c := first(a, b)

	fmt.Println("Expected value for c is '2'.  c = ", c)

	/* Output
	$ go run main.go
	Expected value for c is '2'.  c =  2
	 */
}
