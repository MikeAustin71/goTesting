package main

import (
	"fmt"
)

/*
	This example ranges over a map and extracts
	only the values. Notice that the values are
	returned in no specific order.
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

	for _, val := range m {
		fmt.Println("Key=", val)
	}

}

/*	Output - Notice that values are presented in specific order.
	$ go run main.go
	Key= str3
	Key= str8
	Key= str9
	Key= str10
	Key= str7
	Key= str1
	Key= str2
	Key= str4
	Key= str5
	Key= str6
*/
