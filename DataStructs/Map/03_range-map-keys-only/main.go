package main

import (
	"fmt"
)

/*
	This example ranges over a map and extracts
	only the keys. Notice that the order of the
	keys presented differs with each run of the
	program.
*/

func main() {

	var m = map[int]string{
		0: "str1",
		1: "str2",
		2: "str3",
		3: "str4",
		4: "str5",
		5: "str6",
		6: "str7",
		7: "str8",
		8: "str9",
		9: "str10",
	}

	for key := range m {
		fmt.Println("Key=", key)
	}

}

/*	Output - Notice that key order varies.
	$ go run main.go
	Key= 9
	Key= 1
	Key= 2
	Key= 4
	Key= 6
	Key= 7
	Key= 8
	Key= 0
	Key= 3
	Key= 5
*/
