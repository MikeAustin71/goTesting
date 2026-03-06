package main

import (
  "math/big"
  "testing"
)

func Test_Natural_Logarithm_AGM_01(t *testing.T) {

  ePrefix := "Test_Natural_Logarithm_AGM_01() "

  x, isOk := new(big.Float).SetMode(big.AwayFromZero).SetString("2")

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "x, isOk := new(big.Float).SetMode(big.AwayFromZero).SetString(\"2\")\n"+
      "Error= 'Failed to parse '2' as big.Float!'\n\n", ePrefix)
    return
  }

  prec := uint(256)

  result := bFloatLnAGM(x, prec)

  //                    1         2         3         4         5
  //         0.12345678901234567890123456789012345678901234567890
  // ln(2) = 0.69314718055994530941723212145817656807550013436026
  resultStr := result.Text('f', 50)

  if resultStr != "0.69314718055994530941723212145817656807550013436026" {
    t.Errorf("%v\n"+
      "Error:   Actual resultStr = '%v'\n"+
      "       Expected resultStr = '0.69314718055994530941723212145817656807550013436026'\n\n",
      ePrefix, resultStr)
  }

  return
}
