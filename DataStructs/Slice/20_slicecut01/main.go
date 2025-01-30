package main

import "fmt"

func main() {

	funcName := "20_slicecut01-main"

	targetStr := "How now brown cow."

	targetIndex := len(targetStr)

	var actualCutStr, actualRemainderStr,
		expectedCutStr, expectedRemainderStr string

	var err error

	actualCutStr,
		actualRemainderStr,
		err = cutAtIndex(
		targetStr,
		targetIndex)

	if err != nil {
		fmt.Printf("%v"+
			"Error:\n"+
			"%v",
			funcName,
			err.Error())
		return
	}

	expectedCutStr,
		expectedRemainderStr,
		err = expectedStr(
		targetStr,
		targetIndex)

	fmt.Printf("%v\n"+
		"          Original String: %v\n"+
		"        Actual Cut String: %v\n"+
		"      Expected Cut String: %v\n"+
		"  Actual Remainder String: %v\n"+
		"Expected Remainder String: %v\n"+
		"             Target Index: %v\n"+
		"      Length of TargetStr: %v\n",
		funcName,
		targetStr,
		actualCutStr,
		expectedCutStr,
		actualRemainderStr,
		expectedRemainderStr,
		targetIndex,
		len(targetStr))

	fmt.Printf("\nSuccessful Completion!\n")
	fmt.Printf(funcName + "\n")

}

func expectedStr(
	targetStr string,
	cutAtIndex int) (
	cutStr string,
	remainderStr string,
	err error) {

	lenTargetStr := 0

	if cutAtIndex > lenTargetStr {

		cutStr = targetStr
		remainderStr = ""

		return cutStr, remainderStr, err

	}

	var cutStrRunes []rune
	var remainderRunes []rune
	targetStrRunes := []rune(targetStr)

	for i := 0; i < len(targetStrRunes); i++ {

		if i < cutAtIndex {
			cutStrRunes = append(
				cutStrRunes,
				targetStrRunes[i])
		} else {
			remainderRunes = append(
				remainderRunes,
				targetStrRunes[i])
		}

	}

	cutStr = string(cutStrRunes)
	remainderStr = string(remainderRunes)

	return cutStr, remainderStr, err
}

func cutAtIndex(
	targetStr string,
	targetIndex int) (
	cutStr string,
	remainderStr string,
	err error) {

	if targetIndex > len(targetStr) {

		cutStr = targetStr
		remainderStr = ""

		return cutStr, remainderStr, err
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

	cutStr = targetStr[:targetIndex]
	remainderStr = targetStr[targetIndex:]

	return cutStr, remainderStr, err
}
