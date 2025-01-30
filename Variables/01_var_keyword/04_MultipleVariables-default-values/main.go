package main

import "fmt"

// This example is designed to show the default values
// assigned to various types when the var keyword is used.

func main() {

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("a int     type= %T 	- %v \n", a, a)
	fmt.Printf("b string  type= %T  - %v \n", b, b)
	fmt.Printf("b string length = %v\n\n", len(b))

	fmt.Printf("c float64  type= %T  - %v \n", c, c)
	fmt.Printf("d bool     type= %T  - %v \n", d, d)

	fmt.Println()
}

/*	Output
	$ go run main.go
	a int     type= int     - 0
	b string  type= string  -
	b string length = 0

	c float64  type= float64  - 0
	d bool     type= bool  - false
*/
