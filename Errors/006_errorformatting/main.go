package main

import (
	"fmt"
	"strings"
)

func main() {

	// Benchmark said testMsg04 was best.
	//testMsg := "eMsg\n\n\n"

	testMsg := ""

	fmt.Printf("Original Test Message:\n"+
		"%v", []rune(testMsg))

	fmt.Printf("\n")

	testMsg = testMsg04(testMsg)

	fmt.Printf("Printing Final String from testMsg01:\n"+
		"%s", testMsg)

	fmt.Printf("\n\nPrinting testMsg02 as runes:\n%v", []rune(testMsg))

	return
}

func testMsg01(errMsg *string) error {

	ePrefix := "testMsg01()"

	runeStr := []rune(*errMsg)

	lenErrMsg := len(runeStr)

	if lenErrMsg == 0 {
		return fmt.Errorf("%s\n",
			"errMsg is an empty string!\n"+
				"Length of errMsg is %d\n", ePrefix, 0)
	}

	lastStrIdx := lenErrMsg - 1

	lastNewLineIdx := 0

	previousNewLineIdx := -1

	for i := lastStrIdx; i >= 0; i-- {

		if runeStr[i] == '\n' {

			if lastNewLineIdx > 0 {
				previousNewLineIdx = lastNewLineIdx
			}

			lastNewLineIdx = i

		} else {

			break
		}

	}

	if previousNewLineIdx > 0 {

		runeStr = runeStr[:previousNewLineIdx]

		*errMsg = string(runeStr)

	} else {
		return nil
	}

	return nil
}

func testMsg02(errMsg *string) error {
	ePrefix := "testMsg02()"

	if len(*errMsg) == 0 {
		return fmt.Errorf("%s\n"+
			"%s\n",
			ePrefix, "errMsg is an empty string!\n")
	}

	for strings.HasSuffix(*errMsg, "\n") {
		*errMsg = strings.TrimSuffix(*errMsg, "\n")
	}

	*errMsg = *errMsg + "\n"

	return nil
}

func testMsg03(errMsg *string) error {

	ePrefix := "testMsg03()"

	if errMsg == nil {
		return fmt.Errorf("%s\n%s\n",
			ePrefix, "Input parameter 'errMsg' is a nil pointer!\n")
	}

	*errMsg = strings.TrimRight(*errMsg, "\n")

	*errMsg = *errMsg + "\n"

	return nil
}

func testMsg04(errMsg string) string {

	errMsg = strings.TrimRight(errMsg, "\n")

	return errMsg + "\n"

}
