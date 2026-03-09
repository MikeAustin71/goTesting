package naturalLogCalcs

import (
  "fmt"
  "math/big"
  "testing"
  "time"
)

func Test_Natural_Logarithm_Taylor_01(t *testing.T) {

  ePrefix := "Test_Natural_Logarithm_Taylor_01()"

  xValueStr := "3237"

  xValue, isOk := new(big.Float).
    SetMode(big.AwayFromZero).
    SetString(xValueStr)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "xValue, isOk := new(big.Float)\n"+
      ".SetMode(big.AwayFromZero).SetString(xValueStr)\n"+
      "xValueStr= '%v'\n"+
      "Error= 'SetString() FAILED!'\n\n", ePrefix, xValueStr)
    return
  }

  // Fractional Digits ---------1---------2---------3
  //         Accuracy: 123456789012345678901234567890
  expectedResult := "8.082402253926244350924204901779"

  prec := uint(7000)

  actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, prec)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, prec)\n"+
      "xValue= '%v'\n"+
      "prec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      xValue.Text('f', 30),
      prec,
      err.Error())

    return
  }

  actualResultStr := actualResult.Text('f', 30)

  if expectedResult != actualResultStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResult != actualResultStr\n"+
      "Expected actualResultStr = '%v'\n"+
      "  Actual actualResultStr = '%v'\n\n",
      ePrefix, expectedResult, actualResultStr)
  }

  return
}

func Test_Natural_Logarithm_Taylor_02(t *testing.T) {

  ePrefix := "Test_Natural_Logarithm_Taylor_02()"

  xValueStr := "245"

  xValue, isOk := new(big.Float).
    SetMode(big.AwayFromZero).
    SetString(xValueStr)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "xValue, isOk := new(big.Float)\n"+
      ".SetMode(big.AwayFromZero).SetString(xValueStr)\n"+
      "xValueStr= '%v'\n"+
      "Error= 'SetString() FAILED!'\n\n", ePrefix, xValueStr)
    return
  }

  // Fractional Digits ---------1---------2---------3
  //         Accuracy: 123456789012345678901234567890
  expectedResult := "5." +
    "5012582105447269848114648201125470987997708134322400988314281913493387132117963002054508752689226647" +
    "8608969143341784472317150974468024988593153280779150727936782857370573281241801155421927627388228010" +
    "260330798331306088066067384276824977246856395981"

  prec := uint(7000)

  actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, prec)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, prec)\n"+
      "xValue= '%v'\n"+
      "prec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      xValue.Text('f', 30),
      prec,
      err.Error())

    return
  }

  bFloatHlpr := new(BigFloatHelper)

  _, decDigits := bFloatHlpr.CountDigits(expectedResult, '.')

  actualResultStr := actualResult.Text('f', decDigits)

  if expectedResult != actualResultStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResult != actualResultStr\n"+
      "Expected actualResultStr = '%v'\n"+
      "  Actual actualResultStr = '%v'\n\n",
      ePrefix, expectedResult, actualResultStr)
  }

  return
}

func Test_Natural_Logarithm_Taylor_03(t *testing.T) {

  ePrefix := "Test_Natural_Logarithm_Taylor_03()"

  xValueStr := "1.0000001"

  bigFloatHlpr := new(BigFloatHelper)

  //prec := uint(7000)
  // Precision set for approximately
  // 2,108 Decimal Digits of Accuracy

  prec := bigFloatHlpr.ComputeBigFloatPrecisionBits(55, 0)

  xValue, isOk := new(big.Float).
    SetMode(big.AwayFromZero).
    SetPrec(prec).
    SetString(xValueStr)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "xValue, isOk := new(big.Float)\n"+
      ".SetMode(big.AwayFromZero).SetString(xValueStr)\n"+
      "xValueStr= '%v'\n"+
      "Error= 'SetString() FAILED!'\n\n", ePrefix, xValueStr)
    return
  }

  // Copilot expected result "0.000000099999995000000333333341666666708"
  expectedResult :=
    "0.000000099999995000000333333308333335333333166666681"
  //   123456789012345678901234567890123456789012345678901
  //   ---------1---------2---------3---------4---------5
  //     51- Decimal Digits of Accuracy

  startTime := time.Now()

  actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, prec)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, prec)\n"+
      "xValue= '%v'\n"+
      "prec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      xValue.Text('f', 30),
      prec,
      err.Error())

    return
  }

  endTime := time.Now()

  funcExecutionTime := bigFloatHlpr.DurationBreakdown(startTime, endTime)

  _, decDigitsOfExpectedPrecision := bigFloatHlpr.CountDigits(expectedResult, '.')

  actualResultStr := actualResult.Text('f', decDigitsOfExpectedPrecision)

  if expectedResult != actualResultStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResult != actualResultStr\n"+
      "Expected actualResultStr = '%v'\n"+
      "  Actual actualResultStr = '%v'\n"+
      "XValue = '%v'\n"+
      "Big Float Precision Bits - prec = '%v'\n"+
      "Decimal Digits of Precision = '%v'\n"+
      "naturalLogTaylor.lnTaylorDirect() Execution Time: %v\n\n",
      ePrefix,
      expectedResult,
      actualResultStr,
      xValueStr,
      prec,
      decDigitsOfExpectedPrecision, funcExecutionTime)

    return
  }

  decDigitsOfInputAccuracy := new(BigFloatHelper).ComputeBigFloatDecimalDigits(
    prec, 0)

  fmt.Printf("\n\n%v\n"+
    "Successful Completion !\n"+
    "expected Result = '%v'\n"+
    "  actual Result = '%v'\n"+
    "XValue = '%v'\n"+
    "     Big Float Precision Bits - prec = '%v'\n"+
    "   Decimal Digits of Input Precision = '%v'\n"+
    "Decimal Digits of Expected Precision = '%v'\n"+
    "naturalLogTaylor.lnTaylorDirect() Execution Time: %v\n\n",
    ePrefix,
    expectedResult,
    actualResultStr,
    xValueStr,
    prec,
    decDigitsOfInputAccuracy,
    decDigitsOfExpectedPrecision,
    funcExecutionTime)

  return
}
