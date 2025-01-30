package main

import (
	"fmt"
	"math"
	"strings"
)

type NumStrIntSeparator struct {
	intSeparatorChars       []rune // Integer separator characters
	intSeparatorGrouping    uint   // Number of integers in a group
	intSeparatorRepetitions uint   // Number of times this character/group is repeated
	// A zero value signals unlimited repetitions.
	restartIntGroupingSequence bool // If true, the array starts over at index zero.
}

func (intSeparator *NumStrIntSeparator) CopyIn(
	incomingIntSeparator *NumStrIntSeparator) error {

	ePrefix := "NumStrIntSeparator.CopyIn() "

	if incomingIntSeparator == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingIntSeparator' is invalid!\n"+
			"'incomingIntSeparator' is a 'nil' pointer.\n",
			ePrefix)
	}

	if incomingIntSeparator.intSeparatorChars == nil {
		incomingIntSeparator.intSeparatorChars =
			make([]rune, 0, 5)
	}

	lenIncomingIntSep :=
		len(incomingIntSeparator.intSeparatorChars)

	if lenIncomingIntSep == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingIntSeparator.intSeparatorChars' is invalid!\n"+
			"'incomingIntSeparator.intSeparatorChars' is a zero length array.\n",
			ePrefix)
	}

	intSeparator.intSeparatorChars =
		make([]rune, lenIncomingIntSep, 5)

	for i := 0; i < lenIncomingIntSep; i++ {
		intSeparator.intSeparatorChars[i] =
			incomingIntSeparator.intSeparatorChars[i]
	}

	intSeparator.intSeparatorGrouping =
		incomingIntSeparator.intSeparatorGrouping

	intSeparator.intSeparatorRepetitions =
		incomingIntSeparator.intSeparatorRepetitions

	intSeparator.restartIntGroupingSequence =
		incomingIntSeparator.restartIntGroupingSequence

	return nil
}

func (intSeparator *NumStrIntSeparator) IsValidInstanceError(
	ePrefix string) error {

	ePrefix += "NumStrIntSeparator.IsValidInstanceError() "

	if intSeparator.intSeparatorChars == nil {
		intSeparator.intSeparatorChars =
			make([]rune, 0, 5)
	}

	if len(intSeparator.intSeparatorChars) == 0 {
		return fmt.Errorf("%v\n"+
			"Error: 'NumStrIntSeparator.intSeparatorChars' is invalid!\n"+
			"'NumStrIntSeparator.intSeparatorChars' is a zero length rune array.\n",
			ePrefix)
	}

	return nil
}

type NumStrIntSeparatorsDto struct {
	intSeparators []NumStrIntSeparator
}

func (intSepsDto *NumStrIntSeparatorsDto) Add(
	intSeparatorChars []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	restartIntGroupingSequence bool) error {

	ePrefix := "NumStrIntSeparatorsDto.Add() "

	if intSeparatorChars == nil {
		intSeparatorChars =
			make([]rune, 0, 5)
	}

	lenIntSeps := len(intSeparatorChars)

	if lenIntSeps == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is a zero length rune array!\n",
			ePrefix)
	}

	newIntSep := NumStrIntSeparator{}

	newIntSep.intSeparatorChars =
		make([]rune, lenIntSeps, 5)

	for i := 0; i < lenIntSeps; i++ {
		newIntSep.intSeparatorChars[i] =
			intSeparatorChars[i]
	}

	newIntSep.intSeparatorGrouping =
		intSeparatorGrouping

	newIntSep.intSeparatorRepetitions =
		intSeparatorRepetitions

	newIntSep.restartIntGroupingSequence =
		restartIntGroupingSequence

	if intSepsDto.intSeparators == nil {
		intSepsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	intSepsDto.intSeparators =
		append(intSepsDto.intSeparators, newIntSep)

	return nil
}

func (intSepsDto *NumStrIntSeparatorsDto) CopyIn(
	incomingIntSeps *NumStrIntSeparatorsDto) error {

	ePrefix := "NumStrIntSeparatorsDto.CopyIn() "

	if incomingIntSeps == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingIntSeps' is invalid!\n"+
			"'incomingIntSeps' is a 'nil' pointer.\n",
			ePrefix)
	}

	if incomingIntSeps.intSeparators == nil {
		incomingIntSeps.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIncomingIntSeps := len(incomingIntSeps.intSeparators)

	if lenIncomingIntSeps == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingIntSeps' is invalid!\n"+
			"'incomingIntSeps.intSeparators' is a zero length array.\n",
			ePrefix)
	}

	intSepsDto.intSeparators =
		make([]NumStrIntSeparator, lenIncomingIntSeps, 5)

	var err error

	for i := 0; i < lenIncomingIntSeps; i++ {

		err = intSepsDto.intSeparators[i].CopyIn(
			&incomingIntSeps.intSeparators[i])

		if err != nil {
			return fmt.Errorf("%v\n"+
				"Error Returned intSepsDto.intSeparators[%v].CopyIn\n"+
				"%v\n",
				ePrefix,
				i,
				err.Error())
		}

	}

	return nil
}

func (intSepsDto *NumStrIntSeparatorsDto) GetNumOfElements() int {

	if intSepsDto.intSeparators == nil {
		intSepsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	return len(intSepsDto.intSeparators)
}

type NumericSeparators struct {
	decimalSeparator     rune
	integerSeparatorsDto NumStrIntSeparatorsDto
}

func (numSeps *NumericSeparators) SetWithComponents(
	decimalSeparator rune,
	intSeparatorsDto NumStrIntSeparatorsDto,
) error {

	ePrefix := "NumericSeparators.SetWithComponents() "

	numSeps.decimalSeparator = decimalSeparator

	if intSeparatorsDto.intSeparators == nil {
		intSeparatorsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIncomingIntSeps := len(intSeparatorsDto.intSeparators)

	if lenIncomingIntSeps == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorsDto.intSeparators' is invalid!\n"+
			"'intSeparatorsDto.intSeparators' is a zero length array!\n",
			ePrefix)
	}

	numSeps.integerSeparatorsDto.intSeparators =
		make([]NumStrIntSeparator, lenIncomingIntSeps, 5)

	var err error

	for i := 0; i < lenIncomingIntSeps; i++ {

		err =
			numSeps.integerSeparatorsDto.intSeparators[i].
				CopyIn(
					&intSeparatorsDto.intSeparators[i])

		if err != nil {
			return fmt.Errorf("%v\n"+
				"Error returned by numSeps.integerSeparatorsDto.\n"+
				"intSeparators[%v].CopyIn()\n"+
				"Error: %v\n",
				ePrefix,
				i,
				err.Error())
		}

	}

	return nil
}

func main() {
	testDelimiter06()
}

func testDelimiter01() {

	ePrefix := "testDelimiter01()"
	numStr := "123456789012345"

	headerStr := strings.Repeat("-", 60)

	fmt.Println(headerStr)
	fmt.Println("05_DelimNumStr - " + ePrefix)
	fmt.Println(headerStr)

	intSepsDto := NumStrIntSeparatorsDto{}
	var err error

	err = intSepsDto.Add(
		[]rune(","),
		3,
		0,
		false)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	numericSeparators := NumericSeparators{}

	err = numericSeparators.SetWithComponents(
		'.',
		intSepsDto)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var pureNumStr string
	var signChar rune

	pureNumStr,
		signChar,
		err = delimiter_decimalPureNumStr(
		numStr,
		ePrefix)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var finalNumStr string

	finalNumStr,
		err = delimiter_5_1(
		pureNumStr,
		signChar,
		numericSeparators,
		ePrefix)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	fmt.Println("No Errors!")
	fmt.Printf("    Original NumStr: %v\n", numStr)
	fmt.Printf("       Final NumStr: %v\n", finalNumStr)
	fmt.Printf("Length Final NumStr: %v\n",
		len(finalNumStr))
}

func testDelimiter01A() {

	ePrefix := "testDelimiter01A()"
	numStr := "-123456789012345"

	headerStr := strings.Repeat("-", 60)

	fmt.Println(headerStr)
	fmt.Println("05_DelimNumStr - " + ePrefix)
	fmt.Println(headerStr)

	intSepsDto := NumStrIntSeparatorsDto{}
	var err error

	err = intSepsDto.Add(
		[]rune(","),
		3,
		0,
		false)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	numericSeparators := NumericSeparators{}

	err = numericSeparators.SetWithComponents(
		'.',
		intSepsDto)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var pureNumStr string
	var signChar rune

	pureNumStr,
		signChar,
		err = delimiter_decimalPureNumStr(
		numStr,
		ePrefix)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var finalNumStr string

	finalNumStr,
		err = delimiter_5_1(
		pureNumStr,
		signChar,
		numericSeparators,
		ePrefix)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	fmt.Println("No Errors!")
	fmt.Printf("    Original NumStr: %v\n", numStr)
	fmt.Printf("       Final NumStr: %v\n", finalNumStr)
	fmt.Printf("Length Final NumStr: %v\n",
		len(finalNumStr))
}

func testDelimiter02() {

	ePrefix := "testDelimiter02()"

	// Indian Numbering System
	numStr := "123456789012345"

	headerStr := strings.Repeat("-", 60)

	fmt.Println(headerStr)
	fmt.Println("05_DelimNumStr - " + ePrefix)
	fmt.Println(headerStr)

	intSepsDto := NumStrIntSeparatorsDto{}
	var err error

	err = intSepsDto.Add(
		[]rune(","),
		3,
		1,
		false)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	err = intSepsDto.Add(
		[]rune(","),
		2,
		0,
		false)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	numericSeparators := NumericSeparators{}

	err = numericSeparators.SetWithComponents(
		'.',
		intSepsDto)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var pureNumStr string
	var signChar rune

	pureNumStr,
		signChar,
		err = delimiter_decimalPureNumStr(
		numStr,
		ePrefix)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var finalNumStr string

	finalNumStr,
		err = delimiter_5_1(
		pureNumStr,
		signChar,
		numericSeparators,
		ePrefix)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	fmt.Println("No Errors!")
	fmt.Printf("    Original NumStr: %v\n", numStr)
	fmt.Printf("       Final NumStr: %v\n", finalNumStr)
	fmt.Printf("Length Final NumStr: %v\n",
		len(finalNumStr))
}

func testDelimiter03() {
	// Indian Numbering System With
	// multiple delimiters

	ePrefix := "testDelimiter03()"
	numStr := "-123456789012345"

	headerStr := strings.Repeat("-", 60)

	fmt.Println(headerStr)
	fmt.Println("05_DelimNumStr - " + ePrefix)
	fmt.Println(headerStr)

	intSepsDto := NumStrIntSeparatorsDto{}
	var err error

	err = intSepsDto.Add(
		[]rune{',', '-'},
		3,
		1,
		false)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	err = intSepsDto.Add(
		[]rune{',', '-'},
		2,
		0,
		false)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	numericSeparators := NumericSeparators{}

	err = numericSeparators.SetWithComponents(
		'.',
		intSepsDto)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var pureNumStr string
	var signChar rune

	pureNumStr,
		signChar,
		err = delimiter_decimalPureNumStr(
		numStr,
		ePrefix)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var finalNumStr string

	finalNumStr,
		err = delimiter_5_1(
		pureNumStr,
		signChar,
		numericSeparators,
		"main() ")

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	fmt.Println("No Errors!")
	fmt.Printf("    Original NumStr: %v\n", numStr)
	fmt.Printf("Original NumStr Len: %v\n", len(numStr))
	fmt.Printf("       Final NumStr: %v\n", finalNumStr)
	fmt.Printf("Length Final NumStr: %v\n",
		len(finalNumStr))
}

func testDelimiter04() {
	// Hexadecimal delimiters

	ePrefix := "testDelimiter04()"
	description := "Hexadecimal Delimiters Test"
	numStr := "0x57abe2ff1abb390557abe2ff1abb3905"

	headerStr := strings.Repeat("-", 60)

	fmt.Println(headerStr)
	fmt.Println("05_DelimNumStr - " + ePrefix)
	fmt.Println(description)
	fmt.Println(headerStr)

	intSepsDto := NumStrIntSeparatorsDto{}
	var err error

	err = intSepsDto.Add(
		[]rune{' '},
		2,
		3,
		false)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	err = intSepsDto.Add(
		[]rune{' ', ' '},
		2,
		1,
		true)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	numericSeparators := NumericSeparators{}

	err = numericSeparators.SetWithComponents(
		'.',
		intSepsDto)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var pureNumStr string
	var signChar rune

	pureNumStr,
		signChar,
		err = delimiter_hexadecimalPureNumStr(
		numStr,
		true,
		ePrefix)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var finalNumStr string

	finalNumStr,
		err = delimiter_5_1(
		pureNumStr,
		signChar,
		numericSeparators,
		"main() ")

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	fmt.Println("No Errors!")
	fmt.Printf("    Original NumStr: %v\n", numStr)
	fmt.Printf("Original NumStr Len: %v\n", len(numStr))
	fmt.Printf("       Final NumStr: %v\n", finalNumStr)
	fmt.Printf("Length Final NumStr: %v\n",
		len(finalNumStr))
}

func testDelimiter05() {
	// Octal delimiters

	ePrefix := "testDelimiter05()"
	description := "Octal Delimiters Test"
	numStr := "123456700123456"

	headerStr := strings.Repeat("-", 60)

	fmt.Println(headerStr)
	fmt.Println("05_DelimNumStr - " + ePrefix)
	fmt.Println(description)
	fmt.Println(headerStr)

	intSepsDto := NumStrIntSeparatorsDto{}
	var err error

	err = intSepsDto.Add(
		[]rune{','},
		3,
		0,
		false)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	numericSeparators := NumericSeparators{}

	err = numericSeparators.SetWithComponents(
		'.',
		intSepsDto)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var pureNumStr string
	var signChar rune

	pureNumStr,
		signChar,
		err = delimiter_octalPureNumStr(
		numStr,
		ePrefix)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var finalNumStr string

	finalNumStr,
		err = delimiter_5_1(
		pureNumStr,
		signChar,
		numericSeparators,
		"main() ")

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	fmt.Println("No Errors!")
	fmt.Printf("    Original NumStr: %v\n", numStr)
	fmt.Printf("        Pure NumStr: %v\n", pureNumStr)
	fmt.Printf("Original NumStr Len: %v\n", len(numStr))
	fmt.Printf("       Final NumStr: %v\n", finalNumStr)
	fmt.Printf("Length Final NumStr: %v\n",
		len(finalNumStr))
}

func testDelimiter06() {

	ePrefix := "testDelimiter06()"
	description := "Chinese Number System Test"

	numStr := "123456789012345"

	headerStr := strings.Repeat("-", 60)

	fmt.Println(headerStr)
	fmt.Println("05_DelimNumStr - " + ePrefix)
	fmt.Println(description)
	fmt.Println(headerStr)

	intSepsDto := NumStrIntSeparatorsDto{}
	var err error

	err = intSepsDto.Add(
		[]rune(","),
		4,
		0,
		false)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	numericSeparators := NumericSeparators{}

	err = numericSeparators.SetWithComponents(
		'.',
		intSepsDto)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var pureNumStr string
	var signChar rune

	pureNumStr,
		signChar,
		err = delimiter_decimalPureNumStr(
		numStr,
		ePrefix)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	var finalNumStr string

	finalNumStr,
		err = delimiter_5_1(
		pureNumStr,
		signChar,
		numericSeparators,
		ePrefix)

	if err != nil {
		fmt.Printf("** Error **\n"+
			"%v\n",
			err.Error())
		return
	}

	fmt.Println("No Errors!")
	fmt.Printf("    Original NumStr: %v\n", numStr)
	fmt.Printf("       Final NumStr: %v\n", finalNumStr)
	fmt.Printf("Length Final NumStr: %v\n",
		len(finalNumStr))
}

func delimiter_octalPureNumStr(
	numStr string,
	ePrefix string) (
	pureOctalNumStr string,
	signChar rune,
	err error) {

	ePrefix += "func delimiter_delimiter_octalPureNumStr() "

	rawNumRunes := []rune(numStr)

	lInStr := len(rawNumRunes)

	if lInStr == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStr' is invalid!\n"+
			"numStr is an empty string.\n",
			ePrefix)

		return pureOctalNumStr, signChar, err
	}

	pureOctalNumRunes := make([]rune, 0, 256)

	haveFirstNumericDigit := false

	for i := 0; i < lInStr; i++ {

		if !haveFirstNumericDigit &&
			signChar == 0 &&
			(rawNumRunes[i] == '+' ||
				rawNumRunes[i] == '-') {

			signChar = rawNumRunes[i]

			continue
		}

		if rawNumRunes[i] >= '0' &&
			rawNumRunes[i] <= '7' {
			pureOctalNumRunes = append(pureOctalNumRunes,
				rawNumRunes[i])
			haveFirstNumericDigit = true
		}
	}

	lInStr = len(pureOctalNumRunes)

	if lInStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumStr' is invalid!\n"+
			"pureNumStr contains no integer digits.\n",
			ePrefix)
		return pureOctalNumStr, signChar, err
	}

	pureOctalNumStr = string(pureOctalNumRunes)

	return pureOctalNumStr, signChar, err
}

func delimiter_hexadecimalPureNumStr(
	numStr string,
	applyCapitalLetters bool,
	ePrefix string) (
	pureHexNumStr string,
	signChar rune,
	err error) {

	ePrefix += "func delimiter_hexadecimalPureNumStr() "

	rawNumRunes := []rune(numStr)

	lInStr := len(rawNumRunes)

	if lInStr == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStr' is invalid!\n"+
			"numStr is an empty string.\n",
			ePrefix)

		return pureHexNumStr, signChar, err
	}

	pureNumRunes := make([]rune, 0, 256)

	haveFirstNumericDigit := false
	var numRune rune

	for i := 0; i < lInStr; i++ {

		if !haveFirstNumericDigit &&
			signChar == 0 &&
			(rawNumRunes[i] == '+' ||
				rawNumRunes[i] == '-') {

			signChar = rawNumRunes[i]

			continue
		}

		if rawNumRunes[i] == '0' &&
			i+1 < lInStr {
			// Skip 0x
			if rawNumRunes[i+1] == 'x' ||
				rawNumRunes[i+1] == 'X' {
				i += 1
				continue
			}
		}

		if rawNumRunes[i] >= '0' &&
			rawNumRunes[i] <= '9' {
			pureNumRunes = append(pureNumRunes,
				rawNumRunes[i])
			haveFirstNumericDigit = true
		}

		if rawNumRunes[i] >= 'A' &&
			rawNumRunes[i] <= 'F' {

			numRune = rawNumRunes[i]

			if !applyCapitalLetters {
				numRune += 32
			}

			pureNumRunes = append(pureNumRunes,
				numRune)
			haveFirstNumericDigit = true
		}

		if rawNumRunes[i] >= 'a' &&
			rawNumRunes[i] <= 'f' {

			numRune = rawNumRunes[i]

			if applyCapitalLetters {
				numRune -= 32
			}

			pureNumRunes = append(pureNumRunes,
				numRune)
			haveFirstNumericDigit = true
		}

	}

	pureHexNumStr = string(pureNumRunes)

	return pureHexNumStr, signChar, err
}

// delimiter_decimalPureNumStr - Receives a string of decimal numbers
// and returns a pure decimal number string as well as the numeric
// sign associated with the pure decimal number string.
//
func delimiter_decimalPureNumStr(
	numStr string,
	ePrefix string) (
	pureNumStr string,
	signChar rune,
	err error) {

	ePrefix += "func delimiter_decimalPureNumStr() "

	rawNumRunes := []rune(numStr)

	lInStr := len(rawNumRunes)

	if lInStr == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStr' is invalid!\n"+
			"numStr is an empty string.\n",
			ePrefix)

		return pureNumStr, signChar, err
	}

	pureNumRunes := make([]rune, 0, 256)

	haveFirstNumericDigit := false

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
		return pureNumStr, signChar, err
	}

	pureNumStr = string(pureNumRunes)

	return pureNumStr, signChar, err
}

func delimiter_5_1(
	pureNumStr string,
	signChar rune,
	numSeps NumericSeparators,
	ePrefix string) (
	numStr string,
	err error) {

	ePrefix += "func delimiter_5_1() "

	pureNumRunes := []rune(pureNumStr)

	lInPureNumRunes := len(pureNumRunes)

	if lInPureNumRunes == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumStr' is invalid!\n"+
			"pureNumStr is an empty string.\n",
			ePrefix)

		return numStr, err
	}

	if numSeps.integerSeparatorsDto.intSeparators == nil {
		numSeps.integerSeparatorsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIntSepsArray := len(numSeps.integerSeparatorsDto.intSeparators)

	if lenIntSepsArray == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSeps' is invalid!\n"+
			"'numSeps.integerSeparatorsDto.intSeparators' is a ZERO length array.\n",
			ePrefix)

		return numStr, err
	}

	intSepCharMaxLen := 0
	intGroupMinLen := math.MaxInt32
	lenIntSeps := 0
	lenIntGroupLen := 0

	for j := 0; j < lenIntSepsArray; j++ {

		err =
			numSeps.integerSeparatorsDto.intSeparators[j].
				IsValidInstanceError(
					fmt.Sprintf("%v - numSeps.integerSeparatorsDto.intSeparators[%v]",
						ePrefix,
						j))

		if err != nil {
			return numStr, err
		}

		lenIntSeps = len(
			numSeps.integerSeparatorsDto.intSeparators[j].intSeparatorChars)

		if lenIntSeps > intSepCharMaxLen {
			intSepCharMaxLen = lenIntSeps
		}

		lenIntGroupLen =
			int(numSeps.integerSeparatorsDto.intSeparators[j].intSeparatorGrouping)

		if lenIntGroupLen < intGroupMinLen {
			intGroupMinLen = lenIntGroupLen
		}

	}

	if intGroupMinLen < 1 {
		intGroupMinLen = 1
	}

	if intSepCharMaxLen < 2 {
		intSepCharMaxLen = 2
	}

	lenOutRunes :=
		(lInPureNumRunes + 2) +
			(((lInPureNumRunes + 1) / intGroupMinLen) * intSepCharMaxLen) +
			5

	outRunes := make([]rune, lenOutRunes, lenOutRunes+10)

	currOutRuneIdx := lenOutRunes - 2

	currGroupIdx := 0
	lastGroupIdx := lenIntSepsArray - 1
	currGroupDigitCount := uint(0)
	groupRepetitionsCount := uint(0)

	for i := lInPureNumRunes - 1; i >= 0; i-- {

		// This is a pure number runes array.
		// Assume the digit is valid numeric digit
		currOutRuneIdx--

		outRunes[currOutRuneIdx] = pureNumRunes[i]

		currGroupDigitCount++

		if currGroupDigitCount ==
			numSeps.integerSeparatorsDto.
				intSeparators[currGroupIdx].
				intSeparatorGrouping &&
			i != 0 {

			lenIntSepChars := len(
				numSeps.integerSeparatorsDto.
					intSeparators[currGroupIdx].
					intSeparatorChars)

			targetLenIntSepCharsIdx :=
				currOutRuneIdx - lenIntSepChars

			copy(outRunes[targetLenIntSepCharsIdx:],
				numSeps.integerSeparatorsDto.
					intSeparators[currGroupIdx].
					intSeparatorChars)

			groupRepetitionsCount++

			currOutRuneIdx = targetLenIntSepCharsIdx

			currGroupDigitCount = 0

			if groupRepetitionsCount <
				numSeps.integerSeparatorsDto.
					intSeparators[currGroupIdx].intSeparatorRepetitions {
				// Group Repetitions less than max
				// Go around again
				continue

			} else {
				// Group Repetitions >= to max
				// Time for next group
				groupRepetitionsCount = 0

				if currGroupIdx == lastGroupIdx {

					if numSeps.integerSeparatorsDto.
						intSeparators[currGroupIdx].
						restartIntGroupingSequence == true {
						// Restart the grouping sequence
						currGroupIdx = 0
					} else {
						// Continue indefinitely with current
						// group separators
						continue
					}
				} else {
					// This is NOT the last group in the
					// Grouping Sequence
					currGroupIdx++
					// Start next group
					continue
				}
			}
		}
	}

	if signChar != 0 {
		currOutRuneIdx--
		outRunes[currOutRuneIdx] = signChar
	}

	charStartIdx := currOutRuneIdx
	lenOutRunes = len(outRunes)
	newLenOfOutRunes := len(outRunes[charStartIdx:])

	fmt.Printf("Length Of Processed OutRunes: %v\n"+
		"Final Length of OutRunes: %v\n",
		lenOutRunes,
		newLenOfOutRunes)

	numStr = string(outRunes[charStartIdx:])

	return numStr, err
}
