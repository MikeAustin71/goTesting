package main

import "fmt"

func main() {

	testStr := "HelloxWorldx"
	charToReplace := 'x'
	replacementChar := ' '

	testReplace(
		testStr,
		charToReplace,
		replacementChar)
}

func testReplace(
	testStr string,
	charToReplace rune,
	replacementChar rune) {

	fmt.Printf("charToReplace: %v\n",
		charToReplace)
	fmt.Printf("replacementChar: %v\n",
		replacementChar)

	fmt.Println("Original String: " + testStr)

	runeArray := []rune(testStr)

	for i, r := range runeArray {

		if r == charToReplace {
			runeArray[i] = replacementChar
		}
	}

	testStr = string(runeArray)
	fmt.Println("Final String: " + testStr)

}
