package main

import (
	"fmt"
)

/*
	This example ranges over a map and prints
	out the key and value. However, the values
	are NOT printed in Key number order.
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

	for key, value := range m {
		fmt.Println("Key=", key, " Value=", value)
	}

}

/*	Output
	$ go run main.go
	Key= 1  Value= str2
	Key= 3  Value= str4
	Key= 4  Value= str5
	Key= 5  Value= str6
	Key= 8  Value= str9
	Key= 0  Value= str1
	Key= 2  Value= str3
	Key= 6  Value= str7
	Key= 7  Value= str8
	Key= 9  Value= str10
*/
