/*
This example demonstrates two features. First,
this switch statement switches on type. Second,
a default statement was manually added to handle
the situation where all listed cases evaluate as
false. If default is triggered, the program prints
'unknown'.
*/

package main

import (
	"fmt"
)

type address struct {
	name  string
	addr1 string
	addr2 string
	city  string
	state string
}

type employee struct {
	name       string
	employeeNo int
	dept       string
}

func main() {

	x := 5
	switchOnType(x)

	y := address{"Johh Doe", "221b", "Baker Street", "London", "Texas"}
	switchOnType(y)

	z := "Buenos Dias"
	switchOnType(z)

	e := employee{"John Doe", 5748, "Software Developement"}
	switchOnType(e)

	var f float64 = 5.25
	switchOnType(f)

}

func switchOnType(thing interface{}) {

	switch thing.(type) {
	case int:
		fmt.Println("This is an int")
	case string:
		fmt.Println("This is a string")
	case address:
		fmt.Println("This is an address")
	case employee:
		fmt.Println("This is an employee")
	default:
		fmt.Println("unknown")
	}

}

/*	Output
	$ go run main.go
	This is an int
	This is an address
	This is a string
	This is an employee
	unknown
*/
