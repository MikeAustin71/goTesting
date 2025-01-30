package main

import (
	"fmt"
	"sort"
)

/*
	This example ranges over a map and prints out the
	key and value sorted by key.
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

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}

}

/*	Output
	$ go run main.go
	Key: 0 Value: str1
	Key: 1 Value: str2
	Key: 2 Value: str3
	Key: 3 Value: str4
	Key: 4 Value: str5
	Key: 5 Value: str6
	Key: 6 Value: str7
	Key: 7 Value: str8
	Key: 8 Value: str9
	Key: 9 Value: str10
*/