package main

import "fmt"

// This rune example adapted from a Todd Mcleod 'package'
// example.
func main() {
	s := "Hello"
	fmt.Println("Before processing s=", s)
	s = reverseTwo(s)
	fmt.Println("After processing s=", s)
}

func reverseTwo(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
