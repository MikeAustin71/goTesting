package main

import "fmt"

/*
This example will test the assignment of
slices to a map.
*/

func main() {

	m := make(map[int][]string)
	/*
		This is not required, but for large
		slices there may be a performance factor
		if m[1] == nil {
			m[1] = make([]string, 0, 500)
		}
	*/
	m[1] = append(m[1], "h1")

	fmt.Println("First Value of m[1]", m[1])

	m[1] = append(m[1], "h2")
	m[1] = append(m[1], "h3")
	m[1] = append(m[1], "h4")
	m[1] = append(m[1], "h5")
	m[1] = append(m[1], "h6")
	m[1] = append(m[1], "h7")
	m[1] = append(m[1], "h8")
	m[1] = append(m[1], "h9")
	m[1] = append(m[1], "h10")

	fmt.Println("Subsequent Value of m[1]", m[1])

}
