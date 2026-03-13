package naturalLogCalcs

import (
  "math/big"
  "testing"
)

// 008Fix01, 007Fix
func Test_Natural_Logarithm_Taylor_NearOne(t *testing.T) {

  ePrefix := "Test_Natural_Logarithm_Taylor_NearOne()"

  tests := []struct {
    xStr     string
    expected string
  }{
    // Very close to 1
    //	{"1.0000001", "0.000000099999995000000333333341666666708"},
    {"1.0000001", "0.00000009999999500000033333330833333533333316666668095237970238106349205349206440115431782107551337480"},
    //
    //{"1.000001", "0.000000999999500000333333416666708333208"},
    {"1.000001", "0.00000099999950000033333308333353333316666680952368452392063482063501154392821075513368370525037178787"},
    //
    //{"1.00001", "0.000009999950000333341666708333208333208"},
    {"1.00001", "0.00000999995000033333083335333316666809522559534920534921544003210755133040854707452208102937253322237"},
    //
    // {"1.0001", "0.000099995000333308335333166680951"},
    {"1.0001", "0.00009999500033330833533316668095113106348206440107107551266129432164491607407171907733994721288860975"},
    //
    //{"1.001", "0.000999500333083423333416708333208"},
    {"1.001", "0.00099950033308353316680939892053501146075506239316655199701966682890032495765871955429625476220091215"},
  }

  var workingPrecision uint

  taylor := new(naturalLogTaylor)

  bFloatHlpr := new(BigFloatHelper)

  var sourceExpectedDecimalDigits, sourceInputDecimalDigits, sourceCalculatedDecimalDigits int

  for idx, tc := range tests {

    tc.expected = bFloatHlpr.CleanNumberString(tc.expected)

    _, sourceExpectedDecimalDigits = bFloatHlpr.CountDigits(tc.expected, '.')

    _, sourceInputDecimalDigits = bFloatHlpr.CountDigits(tc.xStr, '.')

    if sourceInputDecimalDigits > sourceExpectedDecimalDigits {
      sourceCalculatedDecimalDigits = sourceInputDecimalDigits
    } else {
      sourceCalculatedDecimalDigits = sourceExpectedDecimalDigits
    }

    sourceCalculatedDecimalDigits += 3

    workingPrecision = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceCalculatedDecimalDigits), 1)

    bFloatXValue, ok := new(big.Float).
      SetMode(big.AwayFromZero).
      SetPrec(workingPrecision).
      SetString(tc.xStr)

    if !ok {
      t.Errorf("%v\n"+
        "FAILED: SetString(%v)\n"+
        "Structure Index: %v\n",
        ePrefix, tc.xStr, idx)
      continue
    }

    actual, err := taylor.lnTaylorDirect(bFloatXValue, workingPrecision)

    if err != nil {
      t.Errorf("%v\n"+
        "Error computing ln(%v): %v\n",
        ePrefix,
        tc.xStr,
        err)
      continue
    }

    actualStr := actual.Text('f', sourceExpectedDecimalDigits)

    if actualStr != tc.expected {
      t.Errorf("%v\n"+
        "ln(%v) mismatch\n"+
        "Expected: %v\n"+
        "  Actual: %v\n"+
        "sourceCalculatedDecimalDigits = %v\n"+
        "Big Float Precision Bits - workingPrecision = %v\n",
        ePrefix,
        tc.xStr,
        tc.expected,
        actualStr,
        sourceCalculatedDecimalDigits,
        workingPrecision)
    }

  } // End of 'for' loop
}
