package main

import "fmt"

/*
	This example ranges over a string returning only
	the runes.
*/

func main() {
	str := "Hello World!"

	for _, val := range str {
		fmt.Printf("character=%q\n", val)
	}
}

/* Output
$ go run main.go
character='H'
character='e'
character='l'
character='l'
character='o'
character=' '
character='W'
character='o'
character='r'
character='l'
character='d'
character='!'
*/
