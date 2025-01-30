package main

import "fmt"

func main() {

	funcName := "19_slicesplice01-main"

	fmt.Printf("\n\n" + funcName + "\n\n")

	origStr := "The cow jumped over the moon"

	fmt.Printf("Original String\n")
	fmt.Printf(origStr + "\n")

	insertString := "XXX"

	insertIndex := len(origStr)

	outputStr := InsertStringAtIndex(
		origStr,
		insertIndex,
		insertString)

	fmt.Printf(outputStr + "\n")

	fmt.Printf("\nSuccessful Completion!\n")
	fmt.Printf(funcName + "\n")
}

// Slice Examples
//arr := []int{1,2,3,4,5}
//
//fmt.Println(arr[:2])        // [1,2]
//
//fmt.Println(arr[2:])        // [3,4,5]
//
//fmt.Println(arr[2:3])        // [3]
//
//fmt.Println(arr[:])            // [1,2,3,4,5]

//https://github.com/golang/go/wiki/SliceTricks

func InsertStringAtIndex(
	targetStr string,
	targetIndex int,
	insertString string) string {

	var outputStr string

	outputStr = targetStr[:targetIndex] + insertString + targetStr[targetIndex:]

	return outputStr
}
