package datetime

import (
	"testing"
)

func TestNumStrDto_CompareAbsoluteVals_01(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_02(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_03(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "12567.218956"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_04(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "12567.218956"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_05(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "-12567.218956"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_06(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "-567.21"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_07(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "567.21"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_08(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "1567.21"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_01(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_02(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_03(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "9211.40"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_04(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "-9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_05(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "-12567.218956"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_06(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "12567.218956"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_07(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-12567.218956"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_08(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "12567.218956"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedValues(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CopyIn_01(t *testing.T) {

	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := NumStrDto{}.New()

	nDto.CopyIn(n1)

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValidInstanceError("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValidInstanceError() Error='%v'", err.Error())
	}

}

func TestNumStrDto_CopyIn_02(t *testing.T) {

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := NumStrDto{}.New()

	nDto.CopyIn(n1)

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValidInstanceError("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValidInstanceError() Error='%v'", err.Error())
	}

}

func TestNumStrDto_CopyOut_01(t *testing.T) {

	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := n1.CopyOut()

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValidInstanceError() Error='%v'", err.Error())
	}

}

func TestNumStrDto_CopyOut_02(t *testing.T) {

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := n1.CopyOut()

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValidInstanceError() Error='%v'", err.Error())
	}

}

func TestNumStrDto_Divide_01(t *testing.T) {

	nStr1 := "12"
	nStr2 := "3"
	expectedStr := "4.0"

	n1Dto, err := NumStrDto{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error= %v",
			nStr1, err.Error())
	}

	n2Dto, err := NumStrDto{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error= %v",
			nStr2, err.Error())
	}

	err = n1Dto.Divide(n2Dto, 1, 5)

	if err != nil {
		t.Errorf("Error returned from n1Dto.Divide(n2Dto, 1, 5). "+
			" Error= %v", err.Error())
	}

	numStr := n1Dto.GetNumStr()

	if expectedStr != numStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedStr, numStr)
	}

}

func TestNumStrDto_FormatForMathOps_01(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	nStr3 := "-12567.218956"
	nStr4 := "-09211.400000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2, "")

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}

}

func TestNumStrDto_FormatForMathOps_02(t *testing.T) {
	nStr1 := "-9211.40"
	nStr2 := "-12567.218956"
	nStr3 := "-12567.218956"
	nStr4 := "-09211.400000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2, "")

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if true != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", true, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}

}

func TestNumStrDto_FormatForMathOps_03(t *testing.T) {
	nStr1 := "-6"
	nStr2 := "67.521"
	nStr3 := "67.521"
	nStr4 := "-06.000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2, "")

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if true != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", true, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}

}

func TestNumStrDto_FormatForMathOps_04(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "-6"
	nStr3 := "67.521"
	nStr4 := "-06.000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2, "")

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}

}

func TestNumStrDto_FormatForMathOps_05(t *testing.T) {
	nStr1 := "-67.521"
	nStr2 := "6"
	nStr3 := "-67.521"
	nStr4 := "06.000"
	expectedCompare := 1
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2, "")

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}

}

func TestNumStrDto_FormatForMathOps_06(t *testing.T) {
	nStr1 := "-67.521"
	nStr2 := "67.521"
	nStr3 := "-67.521"
	nStr4 := "67.521"
	expectedCompare := 0
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2, "")

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}

}
