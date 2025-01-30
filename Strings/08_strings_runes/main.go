package main

import (
	"fmt"
	"strings"
)

func main() {
	CountRunesBytes01()
}

func CountRunesBytes01() {

	funcName := "CountRunesBytes01"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	expectedStr := "    £ -1,234.56"

	lenExpectedStr := len(expectedStr)

	expectedRunes := []rune(expectedStr)

	lenExpectedRunes := len(expectedRunes)

	fmt.Printf("%v\n"+
		"Currency Number String\n"+
		"Expected String: '%v'\n"+
		"Expected String Length: %v\n\n"+
		"Expected Runes: '%v'\n"+
		"Expected Rune Length: %v\n\n",
		funcName,
		expectedStr,
		lenExpectedStr,
		expectedRunes,
		lenExpectedRunes)

	expectedRunes = []rune{'£'}
	lenExpectedRunes = len(expectedRunes)

	expectedStr = string(expectedRunes)
	lenExpectedStr = len(expectedStr)
	byteArray := []byte(expectedStr)
	lenByteArray := len(byteArray)

	fmt.Printf("\n%v\n"+
		"Pound Sign - £\n"+
		"Expected String: '%v'\n"+
		"Expected String Length: %v\n\n"+
		"Expected Runes: '%v'\n"+
		"Expected Rune Length: %v\n"+
		"Expected Bytes: '%v'\n"+
		"Expected Bytes Length: %v\n",
		funcName,
		expectedStr,
		lenExpectedStr,
		expectedRunes,
		lenExpectedRunes,
		byteArray,
		lenByteArray)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}
