package datetime

import "testing"

func TestNumStrDto_ScaleNumStr_01(t *testing.T) {

	ePrefix := "TestNumStrDto_ScaleNumStr_01"

	// "123456.789"				  3						"123.456789"
	nStr := "123456.789"
	scaleMode := SCALEPRECISIONLEFT
	scale := uint(3)
	ePrecision := uint(6)
	expected := "123.456789"
	iStr := "123"
	fracStr := "456789"
	signVal := 1

	nsDto, err := NumStrDto{}.NewPtr().
		ScaleNumStr(nStr,
			scale,
			scaleMode,
			ePrefix)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if ePrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.",
			ePrecision, nsDto.GetPrecisionUint())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nsDto.IsValidInstanceError() Error='%v'", err.Error())
	}

	if !nsDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != iStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", iStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != fracStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", fracStr, s)
	}

}

func TestNumStrDto_ScaleNumStr_02(t *testing.T) {
	nStr := "-123456"
	precision := uint(3)
	expected := "-123456.000"
	iStr := "123456"
	fracStr := "000"
	signVal := -1

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if precision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.",
			precision, nsDto.GetPrecisionUint())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nsDto.IsValidInstanceError() Error='%v'", err.Error())
	}

	if !nsDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != iStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", iStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != fracStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", fracStr, s)
	}

}

func TestNumStrDto_ScaleNumStr_03(t *testing.T) {
	nStr := "123456"
	precision := uint(9)
	expected := "123456.000000000"
	iStr := "123456"
	fracStr := "000000000"
	signVal := 1

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if precision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.",
			precision, nsDto.GetPrecisionUint())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nsDto.IsValidInstanceError() Error='%v'", err.Error())
	}

	if !nsDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != iStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", iStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != fracStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", fracStr, s)
	}

}

func TestNumStrDto_ScaleNumStr_04(t *testing.T) {
	nStr := "123456.7890"
	precision := uint(2)
	expected := "123456.79"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "79"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if precision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.",
			precision, nsDto.GetPrecisionUint())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nsDto.IsValidInstanceError() Error='%v'", err.Error())
	}

	if !nsDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrDto_ScaleNumStr_05(t *testing.T) {
	nStr := "123456.789"
	precision := uint(5)
	expected := "123456.78900"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "78900"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.",
			expected, nsDto.GetNumStr())
	}

	if precision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.",
			precision, nsDto.GetPrecisionUint())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.",
			signVal, nsDto.GetSign())
	}

	err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nsDto.IsValidInstanceError() Error='%v'", err.Error())
	}

	if !nsDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.",
			nsDto.HasNumericDigits())
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrDto_ScaleNumStr_06(t *testing.T) {
	nStr := "-123456.789"
	precision := uint(5)
	expected := "-123456.78900"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := "78900"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.",
			expected, nsDto.GetNumStr())
	}

	if precision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.",
			precision, nsDto.GetPrecisionUint())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.",
			signVal, nsDto.GetSign())
	}

	err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nsDto.IsValidInstanceError() Error='%v'", err.Error())
	}

	if !nsDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.",
			nsDto.HasNumericDigits())
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}
func TestNumStrDto_ScaleNumStr_07(t *testing.T) {
	nStr := "123456.789"
	precision := uint(3)
	expected := "123456.789"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if precision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.",
			precision, nsDto.GetPrecisionUint())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nsDto.IsValidInstanceError() Error='%v'", err.Error())
	}

	if !nsDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}
