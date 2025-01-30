package main

import "fmt"

/*
	This code taken from:
	https://nanxiao.gitbooks.io/golang-101-hacks/content/posts/prepend.html

	Demonstrates an prepend operation. s1 is 'prepended' to s.
*/

func main() {
	var s []int = []int{1, 2}
	fmt.Println(s)

	s1 := make([]int, len(s)+1)
	s1[0] = 0
	copy(s1[1:], s)
	s = s1
	fmt.Println(s)

}

/*	Output
	$ go run main.go
	[1 2]
	[0 1 2]
*/
