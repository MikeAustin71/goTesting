package naturalLogCalcs

import (
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

  xValueStr := "1.0001"

  bigFloatHlpr := new(BigFloatHelper)

  expectedResult :=
    "0.00009999500033330833533316668095113106348206440107107551266129432164491607407171907733994721288860975"
  //   12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
  //   ---------1---------2---------3---------4---------5---------6---------7---------8---------9---------0---------0
  //                                                                                                      0         1
  //                                                                                                      1         1
  //
  //     101- Decimal Digits of Accuracy - This is Calculator Result

  _, expectedResultDecimalDigits := bigFloatHlpr.CountDigits(expectedResult, '.')

  workingPrecision := bigFloatHlpr.ComputeBigFloatPrecisionBits(uint(expectedResultDecimalDigits), 1)

  xValue, isOk := new(big.Float).
    SetMode(big.AwayFromZero).
    SetPrec(workingPrecision).
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

  startTime := time.Now()

  actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, workingPrecision)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, workingPrecision)\n"+
      "xValue= '%v'\n"+
      "workingPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      xValue.Text('f', int(workingPrecision)),
      workingPrecision,
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
      "Big Float Precision Bits - workingPrecision = '%v'\n"+
      "Decimal Digits of Precision = '%v'\n"+
      "naturalLogTaylor.lnTaylorDirect() Execution Time: %v\n\n",
      ePrefix,
      expectedResult,
      actualResultStr,
      xValueStr,
      workingPrecision,
      decDigitsOfExpectedPrecision, funcExecutionTime)

    return
  }
  /*
  	decDigitsOfInputAccuracy := new(BigFloatHelper).ComputeBigFloatDecimalDigits(
  		workingPrecision, 0)

  	fmt.Printf("\n\n%v\n"+
  		"Successful Completion !\n"+
  		"expected Result = '%v'\n"+
  		"  actual Result = '%v'\n"+
  		"XValue = '%v'\n"+
  		"     Big Float Precision Bits - workingPrecision = '%v'\n"+
  		"   Decimal Digits of Input Precision = '%v'\n"+
  		"Decimal Digits of Expected Precision = '%v'\n"+
  		"naturalLogTaylor.lnTaylorDirect() Execution Time: %v\n\n",
  		ePrefix,
  		expectedResult,
  		actualResultStr,
  		xValueStr,
  		workingPrecision,
  		decDigitsOfInputAccuracy,
  		decDigitsOfExpectedPrecision,
  		funcExecutionTime)
  */
  return
}
