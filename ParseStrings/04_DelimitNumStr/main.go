package main

import (
	"fmt"
	"sync"
)

type NumStrIntSeparator struct {
	intSeparatorChar     rune
	intSeparatorGrouping uint
	lock                 *sync.Mutex
}

func (nStrIntSep *NumStrIntSeparator) IsValidInstanceError(
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	ePrefix += "NumStrIntSeparator.IsValidInstanceError() "

	if nStrIntSep.intSeparatorChar == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'nStrIntSep.intSeparatorChar' is invlid!\n"+
			"'nStrIntSep.intSeparatorChar' has a ZERO Value.\n",
			ePrefix)

		return isValid, err
	}

	if nStrIntSep.intSeparatorGrouping == 0 ||
		nStrIntSep.intSeparatorGrouping > 1000 {

		err = fmt.Errorf("%v\n"+
			"Error: 'nStrIntSep.intSeparatorGrouping' is invalid!\n"+
			"nStrIntSep.intSeparatorGrouping='%v'\n",
			ePrefix,
			nStrIntSep.intSeparatorGrouping)

	}

	return isValid, err
}

type NumStrSeparators struct {
	decimalSeparator rune
	intSeparators    []NumStrIntSeparator
}

func main() {

	numSeps := NumStrSeparators{}

	numSeps.decimalSeparator = '.'

	numSeps.intSeparators = []NumStrIntSeparator{
		{
			intSeparatorChar:     ',',
			intSeparatorGrouping: 3,
		},
		{
			intSeparatorChar:     '.',
			intSeparatorGrouping: 2,
		},
	}

	pureNumStr := "-1234567890123456"
	testFunc := "delimiter_4_1()"

	formattedNumStr,
		err :=
		delimiter_4_1(
			pureNumStr,
			numSeps,
			"main() ")

	fmt.Println()
	fmt.Println("main() : Delimit Strings 04")
	fmt.Println(testFunc)
	fmt.Println("-----------------------------")
	fmt.Println()

	fmt.Printf("       Pure Number String: '%v'\n",
		pureNumStr)

	fmt.Println("  numSeps: ")
	for i := 0; i < len(numSeps.intSeparators); i++ {
		fmt.Printf("    Index='%v' numSeps.intSeparatorChar='%v' "+
			"numSeps.intSeparatorGrouping='%v'\n",
			i,
			string(numSeps.intSeparators[i].intSeparatorChar),
			numSeps.intSeparators[i].intSeparatorGrouping)
	}

	fmt.Println()

	if err != nil {
		fmt.Printf("Error:\n"+
			"-----------------------\n"+
			"%v\n", err.Error())
		return
	}

	fmt.Printf("  Formatted Number String: '%v'\n",
		formattedNumStr)
	fmt.Println()

}

func delimiter_4_1(
	pureNumStr string,
	numSeps NumStrSeparators,
	ePrefix string) (
	numStr string,
	err error) {

	rawNumRunes := []rune(pureNumStr)

	lInStr := len(rawNumRunes)

	if lInStr == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumStr' is invalid!\n"+
			"pureNumStr is an empty string.\n",
			ePrefix)

		return numStr, err
	}

	if numSeps.intSeparators == nil {
		numSeps.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIntSepsArray := len(numSeps.intSeparators)

	if lenIntSepsArray == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSeps' is invalid!\n"+
			"'numSeps.intSeparators' is a ZERO length array.\n",
			ePrefix)

		return numStr, err
	}

	pureNumRunes := make([]rune, 0, 256)

	haveFirstNumericDigit := false
	var signChar rune

	for i := 0; i < lInStr; i++ {

		if !haveFirstNumericDigit &&
			signChar == 0 &&
			(rawNumRunes[i] == '+' ||
				rawNumRunes[i] == '-') {

			signChar = rawNumRunes[i]

			continue
		}

		if rawNumRunes[i] >= '0' &&
			rawNumRunes[i] <= '9' {
			pureNumRunes = append(pureNumRunes,
				rawNumRunes[i])
			haveFirstNumericDigit = true
		}
	}

	lInStr = len(pureNumRunes)

	if lInStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumStr' is invalid!\n"+
			"pureNumStr contains no integer digits.\n",
			ePrefix)
	}

	groupCnt := uint(0)
	maxGroupCnt := numSeps.intSeparators[0].intSeparatorGrouping
	currGroupCntIdx := 0
	lastGroupCntIdx := lenIntSepsArray - 1

	numOfDelims := 0
	remainder := 0
	numOfPreDelims := 0
	remainingDigits := lInStr

	for j := 0; j < lenIntSepsArray; j++ {

		_,
			err =
			numSeps.intSeparators[j].IsValidInstanceError(
				ePrefix +
					fmt.Sprintf("Array Index Element No: %v",
						j))

		if err != nil {
			return numStr, err
		}

		if j < lastGroupCntIdx {
			remainingDigits -=
				int(numSeps.intSeparators[j].intSeparatorGrouping)
			if remainingDigits > 0 {
				numOfPreDelims++
			}

		} else {
			// Must be last index
			if remainingDigits > 0 {
				numOfDelims =
					remainingDigits / int(numSeps.intSeparators[j].intSeparatorGrouping)
				remainder =
					remainingDigits % int(numSeps.intSeparators[j].intSeparatorGrouping)
			}
		}
	}

	if numOfDelims > 0 &&
		remainder == 0 {
		numOfDelims--
	}

	lenOutRunes :=
		lInStr +
			numOfPreDelims +
			numOfDelims

	if signChar != 0 {
		lenOutRunes++
	}

	outRunes := make([]rune, lenOutRunes, 256)
	outIdx := lenOutRunes - 1

	for i := lInStr - 1; i >= 0; i-- {

		if pureNumRunes[i] >= '0' && pureNumRunes[i] <= '9' {

			groupCnt++
			outRunes[outIdx] = pureNumRunes[i]
			outIdx--

			if groupCnt == maxGroupCnt && i != 0 {

				groupCnt = 0

				outRunes[outIdx] =
					numSeps.intSeparators[currGroupCntIdx].intSeparatorChar
				outIdx--

				if currGroupCntIdx+1 > lastGroupCntIdx {

					maxGroupCnt =
						numSeps.intSeparators[currGroupCntIdx].intSeparatorGrouping

				} else {

					currGroupCntIdx++

					maxGroupCnt =
						numSeps.intSeparators[currGroupCntIdx].intSeparatorGrouping

				}

			} // End of if groupCnt == maxGroupCnt && i != 0

		} // End Of if pureNumRunes[i] >= '0' && pureNumRunes[i] <= '9'

	}

	// Check and allow for sign designators
	if signChar != 0 {

		outRunes[0] = signChar

	}

	numStr = string(outRunes)

	return numStr, err
}
