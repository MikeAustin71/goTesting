package main

import (
	"fmt"
)

func main() {

	filenameext := "fileName.ext"

	ext := filenameext[8:]

	fmt.Println("   filenameext: ", filenameext)
	fmt.Println("ext from idx 8: ", ext)

}