package Strings_string_extraction10

import (
  "fmt"
  "math"
  "math/big"
)

func main() {

  //                          1         2         3
  baseNumStr := "123.123456789012345678901234567890"

  TestCountDigits(baseNumStr, '.')

}

func TestCountDigits(numStr string, decimalSeparator rune) {

  ePrefix := "TestCountDigits()"

  intDigits, decDigits := CountDigits(numStr, decimalSeparator)

  fmt.Printf("\n%v\n"+
    "intDigits: %v\n"+
    "decDigits: %v\n",
    ePrefix,
    intDigits,
    decDigits)

}

func TestGetDigits(numStr string, decimalSeparator rune, numOfDesiredDecimalDigits int) {

  ePrefix := "TestGetDigits()"

  intDigits, decDigits := CountDigits(numStr, decimalSeparator)

  fmt.Printf("\n%v\n"+
    "intDigits: %v\n"+
    "decDigits: %v\n",
    ePrefix,
    intDigits,
    decDigits)

}
