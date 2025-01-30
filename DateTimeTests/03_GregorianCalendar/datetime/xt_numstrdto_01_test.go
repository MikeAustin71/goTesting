package datetime

import (
	"testing"
)

func TestNumStrDto_AddNumStrs_01(t *testing.T) {
	nStr1 := "-9589.21"
	nStr2 := "9211.40"
	nStr3 := "-377.81"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
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
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_AddNumStrs_01() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_02(t *testing.T) {
	nStr1 := "9589.21"
	nStr2 := "-9211.40"
	nStr3 := "377.81"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
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
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_AddNumStrs_02() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_03(t *testing.T) {
	nStr1 := "9589.21"
	nStr2 := "9211.40"
	nStr3 := "18800.61"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
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
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_AddNumStrs_03() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_04(t *testing.T) {
	nStr1 := "-9589.21"
	nStr2 := "-9211.40"
	nStr3 := "-18800.61"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
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
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_AddNumStrs_04() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_05(t *testing.T) {
	nStr1 := "2"
	nStr2 := "3"
	nStr3 := "5"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
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
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_AddNumStrs_05() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_06(t *testing.T) {
	nStr1 := "2"
	nStr2 := "0.0"
	nStr3 := "2.0"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
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
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_AddNumStrs_06() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_07(t *testing.T) {
	nStr1 := "0"
	nStr2 := "0.0"
	nStr3 := "0.0"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
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
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_AddNumStrs_07() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_08(t *testing.T) {
	nStr1 := "-6"
	nStr2 := "67.521"
	nStr3 := "61.521"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
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
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_AddNumStrs_08() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_09(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "-6"
	nStr3 := "61.521"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
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
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())

	}

	err = nResult.IsValidInstanceError("TestNumStrDto_AddNumStrs_09() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_10(t *testing.T) {
	nStr1 := "-67.521"
	nStr2 := "67.521"
	nStr3 := "0.000"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
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
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
	}

	err = nResult.IsValidInstanceError("TestNumStrDto_AddNumStrs_10() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_11(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "-67.521"
	nStr3 := "0.000"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
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
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
	}

	err = nResult.IsValidInstanceError("TestNumStrDto_AddNumStrs_10() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}
