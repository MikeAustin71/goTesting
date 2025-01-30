package main

import "fmt"

func main() {

	baseStr := "Hello"
	runeArray:= []rune(baseStr)

	byteArray := []byte(string(runeArray))

	fmt.Println("Printing the byte array")
	fmt.Println("-------------------------")
	fmt.Printf("%s\n", byteArray)
	fmt.Println(string(byteArray))

}
