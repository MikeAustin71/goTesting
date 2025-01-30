package main

/*
	This example demonstrates import aliasing.
	The 'fmt' package is aliased as 'f'. Thereafter
	the package is referenced with the f. notation.
*/

import f "fmt" // fmt is aliased as 'f'

func main() {

	f.Println("Hello World!")

}

/*	Output
	$ go run main.go
	Hello World!
*/
