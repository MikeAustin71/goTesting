package main

import "fmt"

// This example demonstrates use of the
// Short Variable Declaration with a
// variety of variable types.

func main() {

	a := 5
	b := "A String"
	c := 6.914
	d := false
	e := "Goodbye"
	f := `Another string?`
	g := 'x'

	fmt.Printf("a = 5		- Type= %T Value= %v \n", a, a)
	fmt.Printf("b = \"A String\"	- Type= %T Value= %v \n", b, b)
	fmt.Printf("c = 6.914 	- Type= %T Value= %v \n", c, c)
	fmt.Printf("d - false	- Type= %T Value= %v \n", d, d)
	fmt.Printf("e - \"Goodbye\"	- Type= %T Value= %v \n", e, e)
	fmt.Printf("f - `Another String?`  	- Type= %T Value= %v \n", f, f)
	fmt.Printf("g 'x'		- Type= %T Value= %v \n", g, g)
}

/*	Ouput
	$ go run main.go
	a = 5           - Type= int Value= 5
	b = "A String"  - Type= string Value= A String
	c = 6.914       - Type= float64 Value= 6.914
	d - false       - Type= bool Value= false
	e - "Goodbye"   - Type= string Value= Goodbye
	f - `Another String?`   - Type= string Value= Another string?
	g 'x'           - Type= int32 Value= 120

*/
