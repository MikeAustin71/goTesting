package main

import "fmt"

func main() {

	numStr := "123456789"
	precision := uint(2)
	ePrefix := "main()"
	testFunc := "ParseIntFracFromRuneArray()"

	intNumRunes,
		lenIntNumRunes,
		fracNumRunes,
		lenFracNumRunes,
		err := ParseIntFracFromRuneArray(
		[]rune(numStr),
		precision,
		ePrefix)

	fmt.Println()
	fmt.Println("          main() - Test")
	fmt.Println("Parse Absolute Value Rune Array")
	fmt.Println(testFunc)
	fmt.Println("-------------------------------------")
	fmt.Println()

	if err != nil {
		fmt.Printf("Error:\n"+
			"------------------------\n"+
			"%v\n", err.Error())
		return
	}

	fmt.Printf(" Original Number String: %v\n",
		numStr)

	fmt.Printf("              Precision: '%v'\n",
		precision)

	fmt.Println("-------------------------------------")

	fmt.Printf("Length Int Number Runes: %v\n",
		lenIntNumRunes)

	fmt.Printf("   Integer Number Runes: '%v'\n",
		string(intNumRunes))

	fmt.Printf("Fractional Number Runes: '%v'\n",
		string(fracNumRunes))

	fmt.Printf("  Length Frac Num Runes: %v\n",
		lenFracNumRunes)

}

func ParseIntFracFromRuneArray(
	absNumValRuneArray []rune,
	precision uint,
	ePrefix string) (
	intNumRunes []rune,
	lenIntNumRunes int,
	fracNumRunes []rune,
	lenFracNumRunes int,
	err error) {

	ePrefix += "ParseIntFracFromRuneArray() "
	intNumRunes = make([]rune, 0, 50)
	fracNumRunes = make([]rune, 0, 50)
	intPrecision := int(precision)

	lenAbsValArray := len(absNumValRuneArray)

	if lenAbsValArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absNumValRuneArray' is invalid!\n"+
			"absNumValRuneArray is a ZERO Length Array.\n",
			ePrefix)

		return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	if intPrecision > lenAbsValArray {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'precision' is invalid!\n"+
			"precision is greater than the Length of Array 'absNumValRuneArray'.\n"+
			"precision=='%v'\n"+
			"Length Of Absolute Value Array=='%v'\n",
			ePrefix,
			precision,
			lenAbsValArray)

		return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err

	}

	pureAbsValRuneArray := make([]rune, 0, lenAbsValArray)

	for i := 0; i < lenAbsValArray; i++ {

		if absNumValRuneArray[i] >= '0' &&
			absNumValRuneArray[i] <= '9' {

			pureAbsValRuneArray = append(pureAbsValRuneArray, absNumValRuneArray[i])
		}

	}

	lenAbsValArray = len(pureAbsValRuneArray)

	if lenAbsValArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absNumValRuneArray' is invalid!\n"+
			"absNumValRuneArray contains ZERO Numeric Digits.\n",
			ePrefix)

		return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	if intPrecision == lenAbsValArray {

		fracNumRunes = make([]rune, lenAbsValArray, 10)

		lenFracNumRunes = copy(fracNumRunes, pureAbsValRuneArray)

		if lenFracNumRunes != lenAbsValArray {
			err = fmt.Errorf("%v\n"+
				"Error: 'precision' is equal to the number of numeric\n"+
				"digits in absNumValRuneArray. However, the copy operation\n"+
				"failed. runes copied='%v'  runes to copy='%v'\n",
				ePrefix,
				lenFracNumRunes,
				lenAbsValArray)

		}

		return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	// precision < lenAbsValArray
	lenIAry := lenAbsValArray - intPrecision

	intNumRunes = make([]rune, lenIAry, 50)

	lenIntNumRunes = copy(intNumRunes, pureAbsValRuneArray[0:lenIAry])

	if lenIntNumRunes != lenIAry {
		err = fmt.Errorf("%v\n"+
			"Error: 'precision' is less than the number of numeric digits\n"+
			"in absNumValRuneArray. However, the integer copy operation\n"+
			"failed. integer runes copied='%v'  integer runes to copy='%v'\n",
			ePrefix,
			lenIntNumRunes,
			lenIAry)

		return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	fracNumRunes = make([]rune, intPrecision, 50)

	lenFracNumRunes = copy(fracNumRunes, pureAbsValRuneArray[lenIntNumRunes:])

	if lenFracNumRunes != intPrecision {
		err = fmt.Errorf("%v\n"+
			"Error: 'precision' is less than the number of numeric digits\n"+
			"in absNumValRuneArray. However, the fractional copy operation\n"+
			"failed. fractional runes copied='%v'  fractional runes to copy='%v'\n",
			ePrefix,
			lenFracNumRunes,
			intPrecision)
	}

	return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
}
