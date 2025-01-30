package main

import (
	"fmt"
)

func main() {

	m := make(map[int]map[string]int)

	m[1] = make(map[string]int)
	m[1]["h1"] = 1
	m[1]["h2"] = 2
	m[1]["h3"] = 3
	m[1]["h4"] = 4
	m[1]["h5"] = 5
	m[1]["h6"] = 6
	m[1]["h7"] = 7
	m[1]["h8"] = 8

	m[2] = make(map[string]int)
	m[2]["i1"] = 1
	m[2]["i2"] = 2
	m[2]["i3"] = 3
	m[2]["i4"] = 4
	m[2]["i5"] = 5
	m[2]["i6"] = 6
	m[2]["i7"] = 7
	m[2]["i8"] = 8

	fmt.Println("m[2]", m[2])
	fmt.Println("Starting range over m[2]")
	for key, value := range m[2] {
		fmt.Println("Key: ", key, "  Value: ", value)
	}

}
