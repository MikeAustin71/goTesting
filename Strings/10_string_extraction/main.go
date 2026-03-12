package Strings_string_extraction10

import (
	"fmt"
)

func main() {

	//                          1         2         3
	baseNumStr := "123.123456789012345678901234567890"

	TestCountDigits(baseNumStr, '.')

}

func TestCountDigits(numStr string, decimalSeparator rune) {

	ePrefix := "TestCountDigits()"

	intDigits, decDigits := new(BigFloatHelper).CountDigits(numStr, decimalSeparator)

	fmt.Printf("\n%v\n"+
		"intDigits: %v\n"+
		"decDigits: %v\n",
		ePrefix,
		intDigits,
		decDigits)

}
