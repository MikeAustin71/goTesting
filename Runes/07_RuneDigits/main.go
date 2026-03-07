package main

import "fmt"

func main() {

  numStr := " 1 2 3 45.123456 78901"

  inDigits, decDigits := CountDigits(numStr, '.')

  fmt.Printf("\n\n%v\n"+
    "Integer Digits: %v\n"+
    "Decimal Digits: %v\n\n",
    "CountDigits()",
    inDigits,
    decDigits)

}

func CountDigits(numStr string, decimalSeparator rune) (intDigits int, decDigits int) {

  intDigits = 0
  decDigits = 0

  var isInt bool

  isInt = true

  for _, v := range numStr {

    if v == decimalSeparator {
      isInt = false
      continue
    }

    if v >= '0' && v <= '9' {

      if isInt {
        // Integer digit to left of decimal point
        intDigits++
      } else {
        // Decimal digit to right of decimal point
        decDigits++
      }
    }

  }

  return intDigits, decDigits
}
