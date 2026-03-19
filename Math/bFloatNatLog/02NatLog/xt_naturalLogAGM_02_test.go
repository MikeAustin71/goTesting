package naturalLogCalcs

import (
  "math/big"
  "testing"
)

func Test_Natural_Logarithm_AGM_1000_01(t *testing.T) {

  // Test Natural Logarithm of 5 calculated to 1004 digits of accuracy.
  // Expected Results (expectedResultStr) value was calculated using
  // Python's 'mpmath' library and function mp.ln(5).

  ePrefix := "Test_Natural_Logarithm_AGM_1000_01()"

  xValueStr := "5"

  expectedResultStr :=
    "1.60943791243410037460075933322618763952560135426851772191264789147417898770765776463013387809317961079996630302171556289972400522932467619963361661746370572755217963749718324565349285620234152505727015519360087977738972568819354071276615473122180952794852129282135805972256767228528724046158944817836467132867399842463775959318942384393435345105097505445419474050136598708786738321313057297204065948538383872366275387654556271816151165993091524320736491167786390067587258577876639158383682395042548795623948403100198269711749099374149848095762101691101437886240335432151272312573458846155978729198088657068402006659984447269997321768118651760122020297840810901964752669976506892065897891338157171620722777455977053432038778749682937536113380094676404833028500217477487970800714346561635989704125890133764152401700588859873494848778797105431417832020145306209061578879067159812516749904143417212203694470009268559375978549211896178643287587684832205835596199679131669650932338853158539898237098543146311008"

  bFloatHlpr := new(BigFloatHelper)

  var workingPrecision uint

  var sourceExpectedDecimalDigits, sourceInputDecimalDigits, sourceCalculatedDecimalDigits int

  _, sourceInputDecimalDigits = bFloatHlpr.CountDigits(xValueStr, '.')

  _, sourceExpectedDecimalDigits = bFloatHlpr.CountDigits(expectedResultStr, '.')

  if sourceInputDecimalDigits > sourceExpectedDecimalDigits {
    sourceCalculatedDecimalDigits = sourceInputDecimalDigits
  } else {
    sourceCalculatedDecimalDigits = sourceExpectedDecimalDigits
  }

  sourceCalculatedDecimalDigits += 3

  workingPrecision = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceCalculatedDecimalDigits), 1)

  xValue, isOk := new(big.Float).
    SetMode(big.AwayFromZero).
    SetPrec(workingPrecision).
    SetString(xValueStr)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "xValue, isOk := new(big.Float)\n"+
      ".SetMode(big.AwayFromZero).SetPrec(workingPrecision).SetString(xValueStr)\n"+
      "xValueStr= '%v'\n"+
      "Working Precision Bits = '%v'\n"+
      "Error= 'SetString() FAILED!'\n\n", ePrefix, workingPrecision, xValueStr)
    return
  }

  actualResult, err := new(naturalLogAGM).lnAGMDirect(xValue, workingPrecision)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualResult, err := new(naturalLogAGM).lnAGMDirect(xValue, workingPrecision)\n"+
      "xValue= '%v'\n"+
      "workingPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      xValue.Text('f', sourceInputDecimalDigits),
      workingPrecision,
      err.Error())

    return
  }

  actualResultStr := actualResult.Text('f', sourceExpectedDecimalDigits)

  if expectedResultStr != actualResultStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result From AGM!\n"+
      "Because expectedResult != actualResultStr\n"+
      "Expected actualResultStr = '%v'\n"+
      "  Actual actualResultStr = '%v'\n\n",
      ePrefix, expectedResultStr, actualResultStr)
  }

  return
}
