package naturalLogCalcs

// ********************************************
// **    xt_BigFloatMath_IntExpo_01_test.go  **
// **    013Fix — IntExpoBigFloat tests      **
// *********************************************

import (
  "math/big"
  "testing"
)

func Test_BigFloatMath_IntExpo_Basic_Positive_01(t *testing.T) {

  ePrefix := "Test_BigFloatMath_IntExpo_Basic_Positive_01"

  bFloatHlpr := new(BigFloatHelper)

  // We only need a few decimal digits here
  precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(20, 4)

  x, ok := new(big.Float).
    SetMode(big.AwayFromZero).
    SetPrec(precBits).
    SetString("2")

  if !ok {
    t.Fatalf("%v\nSetString FAILED for base '2'\n", ePrefix)
  }

  result, err := new(BigFloatMath).IntExpoBigFloat(x, 10, precBits)

  if err != nil {
    t.Fatalf("%v\nIntExpoBigFloat error: %v\n", ePrefix, err)
  }

  resultStr := result.Text('f', 0)
  expectedStr := "1024"

  if resultStr != expectedStr {
    t.Errorf("%v\nUnexpected result for 2^10\nExpected: %v\nActual:   %v\n",
      ePrefix, expectedStr, resultStr)
  }
}

func Test_BigFloatMath_IntExpo_NegativeExponent_01(t *testing.T) {

  ePrefix := "Test_BigFloatMath_IntExpo_NegativeExponent_01"

  bFloatHlpr := new(BigFloatHelper)

  precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(30, 4)

  x, ok := new(big.Float).
    SetMode(big.AwayFromZero).
    SetPrec(precBits).
    SetString("2")

  if !ok {
    t.Fatalf("%v\nSetString FAILED for base '2'\n", ePrefix)
  }

  result, err := new(BigFloatMath).IntExpoBigFloat(x, -3, precBits)
  if err != nil {
    t.Fatalf("%v\nIntExpoBigFloat error: %v\n", ePrefix, err)
  }

  // 2^-3 = 1/8 = 0.125
  resultStr := result.Text('f', 10)
  expectedStr := "0.1250000000"

  if resultStr != expectedStr {
    t.Errorf("%v\nUnexpected result for 2^-3\nExpected: %v\nActual:   %v\n",
      ePrefix, expectedStr, resultStr)
  }
}

func Test_BigFloatMath_IntExpo_ZeroExponent_01(t *testing.T) {

  ePrefix := "Test_BigFloatMath_IntExpo_ZeroExponent_01"

  bFloatHlpr := new(BigFloatHelper)

  precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(20, 2)

  x, ok := new(big.Float).
    SetMode(big.AwayFromZero).
    SetPrec(precBits).
    SetString("5.75")

  if !ok {
    t.Fatalf("%v\nSetString FAILED for base '5.75'\n", ePrefix)
  }

  result, err := new(BigFloatMath).IntExpoBigFloat(x, 0, precBits)
  if err != nil {
    t.Fatalf("%v\nIntExpoBigFloat error: %v\n", ePrefix, err)
  }

  resultStr := result.Text('f', 5)
  expectedStr := "1.00000"

  if resultStr != expectedStr {
    t.Errorf("%v\nUnexpected result for x^0\nExpected: %v\nActual:   %v\n",
      ePrefix, expectedStr, resultStr)
  }
}

// naiveIntPowBigFloat is a simple repeated-multiplication reference
// used only for testing IntExpoBigFloat.
func naiveIntPowBigFloat(x *big.Float, n int64, precBits uint) *big.Float {

  newF := func() *big.Float {
    return new(big.Float).
      SetMode(big.AwayFromZero).
      SetPrec(precBits)
  }

  if n == 0 {
    return newF().SetFloat64(1.0)
  }

  neg := n < 0
  if neg {
    n = -n
  }

  result := newF().SetFloat64(1.0)
  base := newF().Set(x)

  for i := int64(0); i < n; i++ {
    result.Mul(result, base)
  }

  if neg {
    one := newF().SetFloat64(1.0)
    result.Quo(one, result)
  }

  return result
}

func Test_ExpoBigFloat_HighPrecision_Consistency_01(t *testing.T) {

  ePrefix := "Test_ExpoBigFloat_HighPrecision_Consistency_01"

  bFloatHlpr := new(BigFloatHelper)

  // Ask for ~200 decimal digits
  precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(200, 10)

  x, ok := new(big.Float).
    SetMode(big.AwayFromZero).
    SetPrec(precBits).
    SetString("1.0001")

  if !ok {
    t.Fatalf("%v\nSetString FAILED for base '1.0001'\n", ePrefix)
  }

  exponent := int64(500)

  resultFast, err := new(BigFloatMath).IntExpoBigFloat(x, exponent, precBits)
  if err != nil {
    t.Fatalf("%v\nIntExpoBigFloat error: %v\n", ePrefix, err)
  }

  resultNaive := naiveIntPowBigFloat(x, exponent, precBits)

  // Compare as strings with a generous number of decimal digits
  decDigits := bFloatHlpr.ComputeBigFloatDecimalDigits(precBits, 0)

  resultFastStr := resultFast.Text('f', int(decDigits/2))
  resultNaiveStr := resultNaive.Text('f', int(decDigits/2))

  if resultFastStr != resultNaiveStr {
    t.Errorf("%v\nHigh-precision mismatch between IntExpoBigFloat and naiveIntPowBigFloat\nFast:  %v\nNaive: %v\n",
      ePrefix, resultFastStr, resultNaiveStr)
  }
}
