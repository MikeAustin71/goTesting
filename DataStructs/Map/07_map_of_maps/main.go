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

	fmt.Println("m[2]['i3']", m[2]["i3"])

	x := m[2]["i9"]

	fmt.Println("Non-Existent Item m[2][\"i9\"]: ", x)

	if m[3] == nil {
		fmt.Println("Value of Non-Existent map key m[3]: ", m[3])
	}

	if m[2]["i10"] == 0 {
		fmt.Println("Test for Non-Existent Item m[2][\"i9\"] is m[2][\"i10\"] == 0: ", m[2]["i10"])

	}

}
