package main

import "fmt"

func main() {

	pureNumStr := "0.000000"
	decimalSeparator := '.'

	testFunc := "ParseIntFracNumRunes()"

	signChar,
		intNumRunes,
		lenIntNumRunes,
		fracNumRunes,
		lenFracNumRunes,
		err :=
		ParseIntFracNumRunes(
			pureNumStr,
			decimalSeparator,
			"main() ")

	fmt.Println()
	fmt.Println("          main() - Test")
	fmt.Println("Parse Integer-Fractional Number Runes")
	fmt.Println(testFunc)
	fmt.Println("-------------------------------------")
	fmt.Println()

	if err != nil {
		fmt.Printf("Error:\n"+
			"-----------------------\n"+
			"%v\n", err.Error())
		return
	}

	fmt.Printf(" Original Number String: %v\n",
		pureNumStr)

	fmt.Printf("      Decimal Separator: '%c'\n",
		decimalSeparator)
	fmt.Println("-------------------------------------")

	fmt.Printf("    Sign Character Rune: '%c'\n",
		signChar)

	fmt.Printf("Length Int Number Runes: %v\n",
		lenIntNumRunes)

	fmt.Printf("   Integer Number Runes: %v\n",
		string(intNumRunes))

	fmt.Printf("Fractional Number Runes: %v\n",
		string(fracNumRunes))

	fmt.Printf("  Length Frac Num Runes: %v\n",
		lenFracNumRunes)

	fmt.Println()

}

func ParseIntFracNumRunes(
	numStr string,
	decimalSeparator rune,
	ePrefix string) (
	signChar rune,
	intNumRunes []rune,
	lenIntNumRunes int,
	fracNumRunes []rune,
	lenFracNumRunes int,
	err error) {

	ePrefix += "ParseIntFracNumStrs() "

	signChar = 0
	intNumRunes = make([]rune, 0, 256)
	fracNumRunes = make([]rune, 0, 256)

	lInStr := len(numStr)

	if lInStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter numStr is invalid!\n"+
			"'numStr' is an empty or zero length string.\n",
			ePrefix)
		return signChar, intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	if decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparator' is invalid!\n"+
			"'decimalSeparator' has a rune value of zero.\n",
			ePrefix)
		return signChar, intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	rawNumRunes := []rune(numStr)

	lInStr = len(rawNumRunes)

	haveFirstNumericDigit := false
	haveFracSeparator := false

	for i := 0; i < lInStr; i++ {

		if !haveFirstNumericDigit &&
			signChar == 0 &&
			(rawNumRunes[i] == '+' ||
				rawNumRunes[i] == '-') {

			signChar = rawNumRunes[i]

			continue
		}

		if !haveFracSeparator &&
			rawNumRunes[i] == decimalSeparator {

			haveFracSeparator = true

			continue
		}

		if rawNumRunes[i] >= '0' &&
			rawNumRunes[i] <= '9' {

			haveFirstNumericDigit = true

			if !haveFracSeparator {

				intNumRunes = append(intNumRunes,
					rawNumRunes[i])

				lenIntNumRunes++

			} else {

				// Must Have the Fractional Separator
				fracNumRunes = append(fracNumRunes,
					rawNumRunes[i])

				lenFracNumRunes++

			}
		} // End of Numeric Digit If statement

	} // End of 'i' loop

	if lenIntNumRunes == 0 &&
		lenFracNumRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStr' is invalid!\n"+
			"'numStr' contained ZERO Numeric Digits.\n",
			ePrefix)

		signChar = 0
	}

	return signChar, intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
}
