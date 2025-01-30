package main

import (
	"fmt"
)

func main() {

	s := make([]string, 0)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

}

/* 	Output
$ go run main.go
apd: [d e f]
*/
