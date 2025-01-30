/*
	This is an example of a value receiver

*/

package main

import "fmt"

type contact struct {
	first  string
	middle string
	last   string
	age    int
}

func (c contact) fullName() string {
	// Value Receiver = (c contact)
	return c.first + " " + c.middle + " " + c.last
}

func main() {
	p1 := contact{"Puff", "Magic", "Dragon", 208}
	p2 := contact{"Little", "Jackie", "Paper", 8}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}

/*	Output
	$ go run main.go
	Puff Magic Dragon
	Little Jackie Paper
*/
