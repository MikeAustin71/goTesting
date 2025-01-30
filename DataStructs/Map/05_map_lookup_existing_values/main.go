package main

import (
	"fmt"
)

/*
This example will test for the existence of a map key.
*/

func main() {

	m := make(map[int][]string)

	a1 := make([]string, 0, 50)

	a2 := make([]string, 0, 50)

	a3 := make([]string, 0, 50)

	a1 = append(a1, "h1")
	a1 = append(a1, "h2")
	a1 = append(a1, "h3")
	a1 = append(a1, "h4")
	a1 = append(a1, "h5")
	a1 = append(a1, "h6")
	a1 = append(a1, "h7")
	a1 = append(a1, "h8")
	a1 = append(a1, "h9")

	a2 = append(a2, "i1")
	a2 = append(a2, "i2")
	a2 = append(a2, "i3")
	a2 = append(a2, "i4")
	a2 = append(a2, "i5")
	a2 = append(a2, "i6")
	a2 = append(a2, "i7")
	a2 = append(a2, "i8")
	a2 = append(a2, "i9")

	a3 = append(a3, "j1")
	a3 = append(a3, "j2")
	a3 = append(a3, "j3")
	a3 = append(a3, "j4")
	a3 = append(a3, "j5")
	a3 = append(a3, "j6")
	a3 = append(a3, "j7")
	a3 = append(a3, "j8")
	a3 = append(a3, "j9")

	m[1] = a1
	m[2] = a2
	m[3] = a3

	var x []string

	x = m[7]

	if x == nil {
		fmt.Println("x == nil == m[7] does not exist!")
	}

	x = m[2]

	if x != nil {
		fmt.Println("x != nil == m[2] - This does exist!")
		fmt.Println("Value of m[2] = ", x)
	}
}
