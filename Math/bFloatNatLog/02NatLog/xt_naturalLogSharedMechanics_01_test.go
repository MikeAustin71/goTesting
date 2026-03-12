package naturalLogCalcs

import "testing"

func Test_getPi_01(t *testing.T) {

  ePrefix := "Test_getPi_01"

  bigFloatPiFull := new(naturalLogSharedMechanics).getPi(DEFAULT_BASE_NUM_DECIMAL_DIGITS)

  bigFloatPiFullStr := bigFloatPiFull.Text('f', 101)

  //                                                                                                                     1
  //                                                                                                                     0
  //                           1         2         3         4         5         6         7         8         9         0
  //                  12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901
  expectedPiStr := "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798"

  if expectedPiStr != bigFloatPiFullStr {
    t.Errorf("%v\n"+
      "Error: Decimal Digits Read Out is INCORRECT!\n"+
      "Test String Name: '%v' to 2,000 Digits of Accuracy.\n"+
      "Expected bigFloatPiFullStr = '%v'\n"+
      "  Actual bigFloatPiFullStr = '%v'\n\n",
      ePrefix, "Pi", expectedPiStr, bigFloatPiFullStr)
  }

  return
}

func Test_getPi_02(t *testing.T) {

  ePrefix := "Test_getPi_02"

  bigFloatPiNumber := new(naturalLogSharedMechanics).getPi(52)

  actualPiNumStr := bigFloatPiNumber.Text('f', 52)

  //                              1         2         3         4         5
  //                     1234567890123456789012345678901234567890123456789012
  expectedPiNumStr := "3.1415926535897932384626433832795028841971693993751058"

  if expectedPiNumStr != actualPiNumStr {
    t.Errorf("%v\n"+
      "Error: Decimal Digits Read Out is INCORRECT!\n"+
      "Test String Name: '%v' to 50 Digits of Accuracy.\n"+
      "Expected actualPiNumStr = '%v'\n"+
      "  Actual actualPiNumStr = '%v'\n\n",
      ePrefix, "Pi", expectedPiNumStr, actualPiNumStr)
  }

  return
}

func Test_getPi_03(t *testing.T) {

  ePrefix := "Test_getPi_03"

  bigFloatPiNumber := new(naturalLogSharedMechanics).getPi(20005)

  actualPiNumberStr := bigFloatPiNumber.Text('f', 19999)

  actualPiIntDigits, actualPiDecDigits := new(BigFloatHelper).CountDigits(actualPiNumberStr, '.')

  expectedPiIntDigits := 1

  expectedPiDecDigits := 19999

  if expectedPiIntDigits != actualPiIntDigits {
    t.Errorf("%v\n"+
      "Error: Number of integer digits in number string is INCORRECT!\n"+
      "Test String Name: '%v' to 19,999 Digits of Accuracy.\n"+
      "Expected actualPiIntDigits = '%v'\n"+
      "  Actual actualPiIntDigits = '%v'\n\n",
      ePrefix, "Pi", expectedPiIntDigits, actualPiIntDigits)
  }

  if expectedPiDecDigits != actualPiDecDigits {
    t.Errorf("%v\n"+
      "Error: Number of decimal digits in number string is INCORRECT!\n"+
      "Test String Name: '%v' to 19,999 Digits of Accuracy.\n"+
      "Expected actualPiDecDigits = '%v'\n"+
      "  Actual actualPiDecDigits = '%v'\n\n",
      ePrefix, "Pi", expectedPiDecDigits, actualPiDecDigits)
  }

  return
}

func Test_getPi_04(t *testing.T) {

  ePrefix := "Test_getPi_04"

  bigFloatPiNumber := new(naturalLogSharedMechanics).getPi(20005)

  actualPiNumberStr := bigFloatPiNumber.Text('f', 19999)

  actualLengthPiNumberStr := len(actualPiNumberStr)

  expectedLengthPiNumberStr := len(pi20kDigitStr)

  if expectedLengthPiNumberStr != actualLengthPiNumberStr {
    t.Errorf("%v\n"+
      "Error: Length of expected vs actual Pi Number String is INCORRECT!\n"+
      "Test String Name: '%v' to 19,999 Digits of Accuracy.\n"+
      "Expected actualLengthPiNumberStr = '%v'\n"+
      "  Actual actualLengthPiNumberStr = '%v'\n\n",
      ePrefix, "Pi", expectedLengthPiNumberStr, actualLengthPiNumberStr)
  }

  return
}

func Test_getLn2_01(t *testing.T) {

  ePrefix := "Test_getLn2_01"

  bigFloatLn2Full := new(naturalLogSharedMechanics).getLn2(DEFAULT_BASE_NUM_DECIMAL_DIGITS)

  actualLn2NumberStr := bigFloatLn2Full.Text('f', 53)

  //                                  1         2         3         4         5
  //                         12345678901234567890123456789012345678901234567890123
  expectedLn2NumberStr := "0.69314718055994530941723212145817656807550013436025525"

  if expectedLn2NumberStr != actualLn2NumberStr {
    t.Errorf("%v\n"+
      "Error: Decimal Digits Read Out is INCORRECT!\n"+
      "Test String Name: '%v' to 2,000 Digits of Accuracy.\n"+
      "Expected actualLn2NumberStr = '%v'\n"+
      "  Actual actualLn2NumberStr = '%v'\n\n",
      ePrefix, "Natural Log of 2", expectedLn2NumberStr, actualLn2NumberStr)
  }

  return
}

func Test_getLn2_02(t *testing.T) {

  ePrefix := "Test_getLn2_02"

  bigFloatLn2 := new(naturalLogSharedMechanics).getLn2(53)

  actualLn2NumberStr := bigFloatLn2.Text('f', 53)

  //                                  1         2         3         4         5
  //                         12345678901234567890123456789012345678901234567890123
  expectedLn2NumberStr := "0.69314718055994530941723212145817656807550013436025525"

  if expectedLn2NumberStr != actualLn2NumberStr {
    t.Errorf("%v\n"+
      "Error: Decimal Digits Read Out is INCORRECT!\n"+
      "Test String Name: '%v' to 53 Digits of Accuracy.\n"+
      "Expected actualLn2NumberStr = '%v'\n"+
      "  Actual actualLn2NumberStr = '%v'\n\n",
      ePrefix, "Natural Log of 2", expectedLn2NumberStr, actualLn2NumberStr)
  }

  return
}

func Test_getLn2_03(t *testing.T) {

  ePrefix := "Test_getLn2_03"

  bigFloatLn2 := new(naturalLogSharedMechanics).getLn2(20003)

  actualLn2NumberStr := bigFloatLn2.Text('f', -1)

  actualIntDigits, actualDecDigits := new(BigFloatHelper).CountDigits(actualLn2NumberStr, '.')

  expectedIntDigits := 1

  expectedDecDigits := 20001

  if expectedIntDigits != actualIntDigits {
    t.Errorf("%v\n"+
      "Error: Number of integer digits in number string is INCORRECT!\n"+
      "Test String Name: '%v' to 20,002 Digits of Accuracy.\n"+
      "Expected actualIntDigits = '%v'\n"+
      "  Actual actualIntDigits = '%v'\n\n",
      ePrefix, "Natural Log of 2", expectedIntDigits, actualIntDigits)
  }

  if expectedDecDigits != actualDecDigits {
    t.Errorf("%v\n"+
      "Error: Number of decimal digits in number string is INCORRECT!\n"+
      "Test String Name: '%v' to 20,001 Digits of Accuracy.\n"+
      "Expected actualDecDigits = '%v'\n"+
      "  Actual actualDecDigits = '%v'\n\n",
      ePrefix, "Natural Log of 2", expectedDecDigits, actualDecDigits)
  }

  return
}

func Test_getLn2_04(t *testing.T) {

  ePrefix := "Test_getLn2_04"

  bigFloatLn2 := new(naturalLogSharedMechanics).getLn2(20003)

  actualLn2NumberStr := bigFloatLn2.Text('f', 20001)

  actualLengthLn2NumberStr := len(actualLn2NumberStr)

  expectedLenNatLog2DigitStr := len(natLog2Str20kDigits)

  if expectedLenNatLog2DigitStr != actualLengthLn2NumberStr {
    t.Errorf("%v\n"+
      "Error: Length of expected vs actual Ln2 Number String is INCORRECT!\n"+
      "Test String Name: '%v' to 20,001 Digits of Accuracy.\n"+
      "Expected actualLengthLn2NumberStr = '%v'\n"+
      "  Actual actualLengthLn2NumberStr = '%v'\n\n",
      ePrefix, "Natural Log of 2", expectedLenNatLog2DigitStr, actualLengthLn2NumberStr)

  }

  return
}

func Test_getEulersNumber_Default_01(t *testing.T) {

  ePrefix := "Test_getEulersNumber_Default_01"

  bigFloatEulersFull := new(naturalLogSharedMechanics).getEulersNum(DEFAULT_BASE_NUM_DECIMAL_DIGITS)

  bigFloatEulersNumFullStr := bigFloatEulersFull.Text('f', 100)

  expectedEulersNumberStr := "2.7182818284590452353602874713526624977572470936999595749669676277240766303535475945713821785251664274"

  if expectedEulersNumberStr != bigFloatEulersNumFullStr {
    t.Errorf("%v\n"+
      "Error: Decimal Digits Read Out is INCORRECT!\n"+
      "Test String Name: '%v' to 2,000 Digits of Accuracy.\n"+
      "Expected bigFloatEulersNumFullStr = '%v'\n"+
      "  Actual bigFloatEulersNumFullStr = '%v'\n\n",
      ePrefix, "Euler's Number", expectedEulersNumberStr, bigFloatEulersNumFullStr)
  }

  return
}

func Test_getEulersNumber_Default_02(t *testing.T) {

  ePrefix := "Test_getEulersNumber_Default_02"

  bigFloatEulersNum := new(naturalLogSharedMechanics).getEulersNum(50)

  bigFloatEulersNumStr := bigFloatEulersNum.Text('f', 50)

  expectedEulersNumberStr := "2.71828182845904523536028747135266249775724709369995"

  if expectedEulersNumberStr != bigFloatEulersNumStr {
    t.Errorf("%v\n"+
      "Error: Decimal Digits Read Out is INCORRECT!\n"+
      "Test String Name: '%v' to 50 Digits of Accuracy.\n"+
      "Expected bigFloatEulersNumFullStr = '%v'\n"+
      "  Actual bigFloatEulersNumFullStr = '%v'\n\n",
      ePrefix, "Euler's Number", expectedEulersNumberStr, bigFloatEulersNumStr)
  }

  return
}

func Test_getEulersNumber_Default_03(t *testing.T) {

  ePrefix := "Test_getEulersNumber_Default_03"

  bigFloatEulersNum := new(naturalLogSharedMechanics).getEulersNum(50005)

  bigFloatEulersNumStr := bigFloatEulersNum.Text('f', -1)

  actualIntDigits, actualDecDigits := new(BigFloatHelper).CountDigits(bigFloatEulersNumStr, '.')

  expectedIntDigits := 1

  expectedDecDigits := 49999

  if expectedIntDigits != actualIntDigits {
    t.Errorf("%v\n"+
      "Error: Number of integer digits in number string is INCORRECT!\n"+
      "Test String Name: '%v' to 50,000 Digits of Accuracy.\n"+
      "Expected actualIntDigits = '%v'\n"+
      "  Actual actualIntDigits = '%v'\n\n",
      ePrefix, "Euler's Number", expectedIntDigits, actualIntDigits)
  }

  if expectedDecDigits != actualDecDigits {
    t.Errorf("%v\n"+
      "Error: Number of decimal digits in number string is INCORRECT!\n"+
      "Test String Name: '%v' to 50,000 Digits of Accuracy.\n"+
      "Expected actualDecDigits = '%v'\n"+
      "  Actual actualDecDigits = '%v'\n\n",
      ePrefix, "Euler's Number", expectedDecDigits, actualDecDigits)
  }

  return
}

func Test_getEulersNumber_Default_04(t *testing.T) {

  ePrefix := "Test_getEulersNumber_Default_04"

  bigFloatEulersNum := new(naturalLogSharedMechanics).getEulersNum(50005)

  actualEulersNumStr := bigFloatEulersNum.Text('f', 49999)

  lengthActualEulersNumStr := len(actualEulersNumStr)

  expectedLengthEulersNumStr := len(eulersConstNumStr)

  if expectedLengthEulersNumStr != lengthActualEulersNumStr {
    t.Errorf("%v\n"+
      "Error: Length of expected vs actual Euler's Number String is INCORRECT!\n"+
      "Test String Name: '%v' to 49,999 Digits of Accuracy.\n"+
      "Expected lengthActualEulersNumStr = '%v'\n"+
      "  Actual lengthActualEulersNumStr = '%v'\n\n",
      ePrefix, "Euler's Number", expectedLengthEulersNumStr, lengthActualEulersNumStr)

  }

  return
}
