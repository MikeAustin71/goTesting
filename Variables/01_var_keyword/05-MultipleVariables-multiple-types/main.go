package main

import "fmt"

// This example demonstrates the use of the var
// keyword in declaring multiple variables of different
// types.
var (
	a = 10
	b = 20.702
	c = true
	d = "Hello World"
)

func main() {
	fmt.Printf("a = 10     - Type: %T  - Value= %v\n", a, a)
	fmt.Printf("b = 20.702 - Type: %T  - Value= %v\n", b, b)
	fmt.Printf("c = true   - Type: %T  - Value= %v\n", c, c)
	fmt.Printf("d = \"Hello World\" - Type: %T  - Value= %v\n", d, d)
}

/*	Output
	$ go run main.go
	a = 10     - Type: int  - Value= 10
	b = 20.702 - Type: float64  - Value= 20.702
	c = true   - Type: bool  - Value= true
	d = "Hello World" - Type: string  - Value= Hello World
 */
