package datetime

import "testing"

func TestNumStrDto_IsNumStrDtoValid_01(t *testing.T) {

	nDto := NumStrDto{}

	absAllNumRunes := nDto.GetAbsAllNumRunes()

	lenAbsAllNumRunes := len(absAllNumRunes)

	if lenAbsAllNumRunes != 0 {
		t.Errorf("Expected zero Length absAllNumRunes array. Array length was '%v'",
			lenAbsAllNumRunes)
	}

	nDto.IsValidInstanceError("")

	absAllNumRunes = nDto.GetAbsAllNumRunes()

	lenAbsAllNumRunes = len(absAllNumRunes)

	if lenAbsAllNumRunes != 0 {
		t.Errorf("Expected after validation absAllNumRunes length='0'. Instead, length='%v'",
			lenAbsAllNumRunes)
	}

}

func TestNumStrDto_Multiply_01(t *testing.T) {

	nStr1 := "35.123456"
	nStr2 := "47.9876514"
	nStr3 := "1685.4921624912384"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	err := n1.Multiply(n2)

	if err != nil {
		t.Errorf("n1.Multiply(n2) returned an error. Error= %v", err)
	}

	s := n1.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

}

func TestNumStrDto_MultiplyNumStrs_01(t *testing.T) {
	nStr1 := "35.123456"
	nStr2 := "47.9876514"
	nStr3 := "1685.4921624912384"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_02(t *testing.T) {
	nStr1 := "35.123456"
	nStr2 := "-47.9876514"
	nStr3 := "-1685.4921624912384"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_03(t *testing.T) {
	nStr1 := "-35.123456"
	nStr2 := "-47.9876514"
	nStr3 := "1685.4921624912384"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_04(t *testing.T) {
	nStr1 := "57"
	nStr2 := "123"
	nStr3 := "7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_05(t *testing.T) {
	nStr1 := "57"
	nStr2 := "-123"
	nStr3 := "-7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_06(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "123"
	nStr3 := "-7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_07(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "-123"
	nStr3 := "7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_08(t *testing.T) {
	nStr1 := "0"
	nStr2 := "123"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_09(t *testing.T) {
	nStr1 := "57"
	nStr2 := "0"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_10(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "0"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_11(t *testing.T) {
	nStr1 := "57"
	nStr2 := "0.123"
	nStr3 := "7.011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_12(t *testing.T) {
	nStr1 := "62.1234567890123"
	nStr2 := "3.12345678901234"
	nStr3 := "194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_13(t *testing.T) {
	nStr1 := "-62.1234567890123"
	nStr2 := "3.12345678901234"
	nStr3 := "-194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_14(t *testing.T) {
	nStr1 := "-62.1234567890123"
	nStr2 := "-3.12345678901234"
	nStr3 := "194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}
