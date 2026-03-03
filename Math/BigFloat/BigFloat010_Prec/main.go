package main

import (
  "fmt"
  "math"
  "math/big"
  "strings"
)

func main() {

  TestPrec001()
}

func TestPrec001() {

  spacerStr := " "

  lineBreakStr := spacerStr

  lineBreakStr += strings.Repeat("-", 50)

  fmt.Printf("\n%v\n%vRunning TestPrec001()\n%v\n\n",
    lineBreakStr,
    spacerStr,
    lineBreakStr)

  //               Decimal            1         2         3
  //               Digits    12345678901234567890123456789012345
  expectedResultStr := "3020.27536989418403462317254679710941519"

  requiredFractionalDigits := 35

  exponentValue := 6

  baseStr := "3.14159"

  requiredFloatPrecision := ComputeBigFloatPrecisionBits(35, exponentValue)

  baseFloat, isOk :=
    new(big.Float).
      SetPrec(requiredFloatPrecision).
      SetMode(big.AwayFromZero).
      SetString(baseStr)

  if !isOk {
    fmt.Printf("baseFloat SetString(\"3.14159\") Failed!")
    return
  }

  resultBaseToExponent, isOk :=
    new(big.Float).
      SetPrec(requiredFloatPrecision).
      SetMode(big.AwayFromZero).
      SetString(baseStr)

  if !isOk {
    fmt.Printf("resultBaseToExponent SetString(\"3.14159\") Failed!")
    return
  }

  for i := 0; i < exponentValue; i++ {
    resultBaseToExponent.Mul(resultBaseToExponent, baseFloat)
  }

  actualResultStr := resultBaseToExponent.Text('f', requiredFractionalDigits)

  fmt.Printf("\n%v Base: %v\n",
    spacerStr,
    baseStr)

  fmt.Printf("%v Exponent: %v\n",
    spacerStr,
    exponentValue)

  fmt.Printf("Expected Result: %v\n", expectedResultStr)

  fmt.Printf("Actual Result:   %v\n\n%v\n\n",
    actualResultStr,
    lineBreakStr)

}

func ComputeBigFloatPrecisionBits(decDigits int, multCount int) uint {

  // Convert decimal digits → bits
  // baseBits := float64(decDigits) * math.Log2(10)

  //  3.321 928 094 887 362 347 870 319 429 489 390 175 864 831 393 024 6
  // float64 max 15-digits of precision
  baseBits := float64(decDigits) * 3.321928094887362

  // Add safety margin for exponentiation
  safety := 64.0 + (float64(multCount) * 4.0)

  // Round up to nearest whole bit
  return uint(math.Ceil(baseBits + safety))
}
