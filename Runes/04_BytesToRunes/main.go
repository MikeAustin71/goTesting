package main

import "fmt"

func main() {
	str := "Hello Bytes to Runes"

	bArray := []byte(str)
	newStr := string(bArray)
	fmt.Println("Byte Array to Rune Array")
	fmt.Println("--------------------------------------")
	fmt.Println("     Original String: ", str)
	fmt.Println("Byte Array as String: ", newStr)

	runeArray := []rune(newStr)

	fmt.Println("Rune Array as String: ", string(runeArray))




}
