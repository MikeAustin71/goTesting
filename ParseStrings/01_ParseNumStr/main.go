package main

import (
	"fmt"
)

func main() {
	pureNumStr := "12345678901.25"

	testFunc := "delimitThousands05()"
	thousandsSeparator := ','
	intGroupSeq := []uint{2, 3}

	formattedNumStr,
		err :=
		delimitThousands05(
			pureNumStr,
			thousandsSeparator,
			intGroupSeq,
			"main() ")

	fmt.Println()
	fmt.Println("main-test thousands delimiter")
	fmt.Println(testFunc)
	fmt.Println("-----------------------------")
	fmt.Println()

	if err != nil {
		fmt.Printf("Error:\n"+
			"-----------------------\n"+
			"%v\n", err.Error())
		return
	}

	fmt.Printf("Formatted Number String:\n"+
		"'%v'\n", formattedNumStr)
	fmt.Println()

}

// delimitThousands - is designed to delimit or format a pure number string with a thousands
// separator (i.e. ','). Example: Input == 1234567890 -> Output == "1,234,567,890".
// NOTE: This method will not handle number strings containing decimal fractions
// and currency characters. For these options see method ns.DlimDecCurrStr(),
// above.
//
// Example: 1,000,000,000
//
func delimitThousands00(
	pureNumStr string,
	thousandsSeparator rune,
	ePrefix string) (
	numStr string,
	err error) {

	ePrefix += "delimitThousands00()"

	rawNumRunes := []rune(pureNumStr)

	lInStr := len(rawNumRunes)

	if lInStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumStr' is invalid!\n"+
			"pureNumStr is an empty string.\n",
			ePrefix)
		return numStr, err
	}

	pureNumRunes := make([]rune, 0, 256)

	haveFirstNumericDigit := false
	signVal := 0

	for i := 0; i < lInStr; i++ {

		if !haveFirstNumericDigit &&
			signVal == 0 &&
			(rawNumRunes[i] == '+' ||
				rawNumRunes[i] == '-') {

			if rawNumRunes[i] == '+' {
				signVal = 1
			} else {
				signVal = -1
			}

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

	outRunes := make([]rune, 0, 256)
	iCnt := 0

	for i := lInStr - 1; i >= 0; i-- {

		if pureNumRunes[i] >= '0' && pureNumRunes[i] <= '9' {

			iCnt++
			outRunes = append(
				[]rune{pureNumRunes[i]},
				outRunes...)

			if iCnt == 3 && i != 0 {

				iCnt = 0

				outRunes = append(
					[]rune{thousandsSeparator},
					outRunes...)
			}

			continue
		}

	}

	// Check and allow for sign designators
	if signVal != 0 {
		if signVal == -1 {
			outRunes = append(
				[]rune{'-'},
				outRunes...)
		} else {
			outRunes = append(
				[]rune{'+'},
				outRunes...)
		}
	}

	numStr = string(outRunes)

	return numStr, err
}

// delimitThousands - is designed to delimit or format a pure number string with a thousands
// separator (i.e. ','). Example: Input == 1234567890 -> Output == "1,234,567,890".
// NOTE: This method will not handle number strings containing decimal fractions
// and currency characters. For these options see method ns.DlimDecCurrStr(),
// above.
//
// Example: 1,000,000,000
//
// THIS USES A PRE-PENDING ALGORITHM
//
func delimitThousands02(
	pureNumStr string,
	thousandsSeparator rune,
	ePrefix string) (
	numStr string,
	err error) {

	ePrefix += "delimitThousands02()"

	rawNumRunes := []rune(pureNumStr)

	lInStr := len(rawNumRunes)

	if lInStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumStr' is invalid!\n"+
			"pureNumStr is an empty string.\n",
			ePrefix)
		return numStr, err
	}

	pureNumRunes := make([]rune, 0, 256)

	haveFirstNumericDigit := false
	signVal := 0

	for i := 0; i < lInStr; i++ {

		if !haveFirstNumericDigit &&
			signVal == 0 &&
			(rawNumRunes[i] == '+' ||
				rawNumRunes[i] == '-') {

			if rawNumRunes[i] == '+' {
				signVal = 1
			} else {
				signVal = -1
			}

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

	outRunes := make([]rune, 0, 256)
	iCnt := 0

	for i := lInStr - 1; i >= 0; i-- {

		if pureNumRunes[i] >= '0' && pureNumRunes[i] <= '9' {

			iCnt++
			outRunes = append(
				[]rune{pureNumRunes[i]},
				outRunes...)

			if iCnt == 3 && i != 0 {

				iCnt = 0

				outRunes = append(
					[]rune{thousandsSeparator},
					outRunes...)
			}

			continue
		}

	}

	// Check and allow for sign designators
	if signVal != 0 {
		if signVal == -1 {
			outRunes = append(
				[]rune{'-'},
				outRunes...)
		} else {
			outRunes = append(
				[]rune{'+'},
				outRunes...)
		}
	}

	numStr = string(outRunes)

	return numStr, err
}

// delimitThousands - is designed to delimit or format a pure number string with a thousands
// separator (i.e. ','). Example: Input == 1234567890 -> Output == "1,234,567,890".
// NOTE: This method will not handle number strings containing decimal fractions
// and currency characters. For these options see method ns.DlimDecCurrStr(),
// above.
//
// Example: 1,000,000,000
//
//       A VERY GOOD ALTERNATIVE TO PRE-PENDING!!!
//
func delimitThousands03(
	pureNumStr string,
	thousandsSeparator rune,
	ePrefix string) (
	numStr string,
	err error) {

	ePrefix += "delimitThousands03()"

	rawNumRunes := []rune(pureNumStr)

	lInStr := len(rawNumRunes)

	if lInStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumStr' is invalid!\n"+
			"pureNumStr is an empty string.\n",
			ePrefix)
		return numStr, err
	}

	pureNumRunes := make([]rune, 0, 256)

	haveFirstNumericDigit := false
	signVal := 0

	for i := 0; i < lInStr; i++ {

		if !haveFirstNumericDigit &&
			signVal == 0 &&
			(rawNumRunes[i] == '+' ||
				rawNumRunes[i] == '-') {

			if rawNumRunes[i] == '+' {
				signVal = 1
			} else {
				signVal = -1
			}

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

	numOfDelims := lInStr / 3

	remainder := lInStr % 3

	if numOfDelims > 0 &&
		remainder == 0 {
		numOfDelims--
	}

	lenOutRunes := lInStr + numOfDelims

	if signVal != 0 {
		lenOutRunes++
	}

	outRunes := make([]rune, lenOutRunes, 256)
	outIdx := lenOutRunes - 1

	iCnt := 0

	for i := lInStr - 1; i >= 0; i-- {

		if pureNumRunes[i] >= '0' && pureNumRunes[i] <= '9' {

			iCnt++
			outRunes[outIdx] = pureNumRunes[i]
			outIdx--

			if iCnt == 3 && i != 0 {

				iCnt = 0

				outRunes[outIdx] = thousandsSeparator
				outIdx--
			}

			continue
		}

	}

	// Check and allow for sign designators
	if signVal != 0 {

		if signVal == -1 {

			outRunes[0] = '-'

		} else {

			outRunes[0] = '+'

		}
	}

	numStr = string(outRunes)

	return numStr, err
}

// delimitThousands04 - is designed to delimit or format a pure number string with a thousands
// separator (i.e. ','). Example: Input == 1234567890 -> Output == "1,234,567,890".
//
// This method incorporates the concept of integer grouping sequence which allows for
// multiple grouping of integer digits depending on country or culture requires.
// Examples:
//           1,000,000,000
//           10,000,000,00
//
//
// THIS USES A Inter Grouping Sequence with a pre-pending ALGORITHM
//
//
func delimitThousands04(
	pureNumStr string,
	thousandsSeparator rune,
	integerGroupingSequence []uint,
	ePrefix string) (
	numStr string,
	err error) {

	ePrefix += "delimitThousands04()"

	rawNumRunes := []rune(pureNumStr)

	lInStr := len(rawNumRunes)

	if lInStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumStr' is invalid!\n"+
			"pureNumStr is an empty string.\n",
			ePrefix)

		return numStr, err
	}

	if integerGroupingSequence == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerGroupingSequence' is invalid!\n"+
			"'integerGroupingSequence' is a nil pointer.\n",
			ePrefix)

		return numStr, err

	}

	lenIGrpSeq := len(integerGroupingSequence)

	if lenIGrpSeq == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerGroupingSequence' is invalid!\n"+
			"'integerGroupingSequence' is a ZERO length array.\n",
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

	outRunes := make([]rune, 0, 256)
	iCnt := uint(0)
	maxICnt := integerGroupingSequence[0]
	currCntIdx := 0
	lastCntIdx := lenIGrpSeq - 1

	for i := lInStr - 1; i >= 0; i-- {

		if pureNumRunes[i] >= '0' && pureNumRunes[i] <= '9' {

			iCnt++
			outRunes = append(
				[]rune{pureNumRunes[i]},
				outRunes...)

			if iCnt == maxICnt && i != 0 {

				iCnt = 0

				outRunes = append(
					[]rune{thousandsSeparator},
					outRunes...)

				if currCntIdx+1 > lastCntIdx {

					maxICnt = integerGroupingSequence[currCntIdx]

				} else {

					currCntIdx++

					maxICnt = integerGroupingSequence[currCntIdx]

				}

			}

		} // End Of if pureNumRunes[i] >= '0' && pureNumRunes[i] <= '9'

	}

	// Check and allow for sign designators
	if signChar != 0 {

		outRunes = append(
			[]rune{signChar},
			outRunes...)

	}

	numStr = string(outRunes)

	return numStr, err
}

// delimitThousands05 - is designed to delimit or format a pure number string with a thousands
// separator (i.e. ','). Example: Input == 1234567890 -> Output == "1,234,567,890".
//
// This method incorporates the concept of integer grouping sequence which allows for
// multiple grouping of integer digits depending on country or culture requires.
// Examples:
//           1,000,000,000
//           10,000,000,00
//
//
// THIS USES A Inter Grouping Sequence ALGORITHM with a length
// pre-calculation approach.
//
//
func delimitThousands05(
	pureNumStr string,
	thousandsSeparator rune,
	integerGroupingSequence []uint,
	ePrefix string) (
	numStr string,
	err error) {

	ePrefix += "delimitThousands05()"

	rawNumRunes := []rune(pureNumStr)

	lInStr := len(rawNumRunes)

	if lInStr == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumStr' is invalid!\n"+
			"pureNumStr is an empty string.\n",
			ePrefix)

		return numStr, err
	}

	if integerGroupingSequence == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerGroupingSequence' is invalid!\n"+
			"'integerGroupingSequence' is a nil pointer.\n",
			ePrefix)

		return numStr, err

	}

	lenIGrpSeq := len(integerGroupingSequence)

	if lenIGrpSeq == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerGroupingSequence' is invalid!\n"+
			"'integerGroupingSequence' is a ZERO length array.\n",
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
	maxGroupCnt := integerGroupingSequence[0]
	currGroupCntIdx := 0
	lastGroupCntIdx := lenIGrpSeq - 1

	numOfDelims := 0
	remainder := 0
	numOfPreDelims := 0
	remainingDigits := lInStr

	for j := 0; j < lenIGrpSeq; j++ {

		if j < lastGroupCntIdx {
			remainingDigits -= int(integerGroupingSequence[j])
			if remainingDigits > 0 {
				numOfPreDelims++
			}

		} else {
			// Must be last index
			if remainingDigits > 0 {
				numOfDelims =
					remainingDigits / int(integerGroupingSequence[j])
				remainder =
					remainingDigits % int(integerGroupingSequence[j])
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

				outRunes[outIdx] = thousandsSeparator
				outIdx--

				if currGroupCntIdx+1 > lastGroupCntIdx {

					maxGroupCnt = integerGroupingSequence[currGroupCntIdx]

				} else {

					currGroupCntIdx++

					maxGroupCnt = integerGroupingSequence[currGroupCntIdx]

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
