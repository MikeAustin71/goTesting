package datetime

import "testing"

func TestNumStrDto_SetPrecision_01(t *testing.T) {

	nStr := "123456789"
	precision := uint(7)
	roundResult := false
	outPrecision := uint(7)
	expected := "123456789.0000000"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.",
			outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_02(t *testing.T) {

	nStr := "123456789"
	precision := uint(7)
	roundResult := true
	outPrecision := uint(7)
	expected := "123456789.0000000"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_03(t *testing.T) {

	nStr := "-123456789"
	precision := uint(7)
	roundResult := false
	outPrecision := uint(7)
	expected := "-123456789.0000000"
	signVal := -1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_04(t *testing.T) {

	nStr := "-123456789"
	precision := uint(7)
	roundResult := true
	outPrecision := uint(7)
	expected := "-123456789.0000000"
	signVal := -1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_05(t *testing.T) {

	nStr := "123456.789"
	precision := uint(2)
	roundResult := true
	outPrecision := uint(2)
	expected := "123456.79"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "79"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_06(t *testing.T) {

	nStr := "123456.789"
	precision := uint(2)
	roundResult := false
	outPrecision := uint(2)
	expected := "123456.78"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "78"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_07(t *testing.T) {

	nStr := "123456.789"
	precision := uint(5)
	roundResult := false
	outPrecision := uint(5)
	expected := "123456.78900"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "78900"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_08(t *testing.T) {

	nStr := "123.456789"
	precision := uint(1)
	roundResult := false
	outPrecision := uint(1)
	expected := "123.4"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "4"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_09(t *testing.T) {

	nStr := "123.456789"
	precision := uint(1)
	roundResult := true
	outPrecision := uint(1)
	expected := "123.5"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "5"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_10(t *testing.T) {

	nStr := "-123.456789"
	precision := uint(1)
	roundResult := false
	outPrecision := uint(1)
	expected := "-123.4"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "4"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_11(t *testing.T) {

	nStr := "-123.456789"
	precision := uint(1)
	roundResult := true
	outPrecision := uint(1)
	expected := "-123.5"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "5"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_12(t *testing.T) {

	nStr := "123456.789"
	precision := uint(0)
	roundResult := true
	outPrecision := uint(0)
	expected := "123457"
	signVal := 1
	absIntRuneStr := "123457"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_13(t *testing.T) {

	nStr := "-123456.789"
	precision := uint(0)
	roundResult := true
	outPrecision := uint(0)
	expected := "-123457"
	signVal := -1
	absIntRuneStr := "123457"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_14(t *testing.T) {

	nStr := "123456.789"
	precision := uint(0)
	roundResult := false
	outPrecision := uint(0)
	expected := "123456"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_15(t *testing.T) {

	nStr := "-123456.789"
	precision := uint(0)
	roundResult := false
	outPrecision := uint(0)
	expected := "-123456"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_16(t *testing.T) {

	nStr := "123457"
	precision := uint(1)
	roundResult := false
	outPrecision := uint(1)
	expected := "123457.0"
	signVal := 1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_17(t *testing.T) {

	nStr := "123457"
	precision := uint(1)
	roundResult := true
	outPrecision := uint(1)
	expected := "123457.0"
	signVal := 1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_18(t *testing.T) {

	nStr := "-123457"
	precision := uint(1)
	roundResult := false
	outPrecision := uint(1)
	expected := "-123457.0"
	signVal := -1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_19(t *testing.T) {

	nStr := "-123457"
	precision := uint(1)
	roundResult := true
	outPrecision := uint(1)
	expected := "-123457.0"
	signVal := -1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetThisPrecision_01(t *testing.T) {

	nStr := "123456789"
	initialPrecision := uint(0)
	precision := uint(3)
	roundResult := true
	outPrecision := uint(3)
	expected := "123456789.000"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := "000"

	nsDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr). "+
			"Err='%v' ", err.Error())
	}

	if initialPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Error: Expected Initial precision='%v'. Actual Initial precision='%v' ",
			initialPrecision, nsDto.GetPrecisionUint())
	}

	nsDto.SetThisPrecision(precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetThisPrecision(precision, roundResult). "+
			"precision= '%v'. Error= %v", precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	err = nsDto.IsValidInstanceError("Test 'nsDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nsDto.IsValidInstanceError() Error='%v'", err.Error())
	}

	if !nsDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetThisPrecision_02(t *testing.T) {

	nStr := "123456.789"
	initialPrecision := uint(3)
	precision := uint(2)
	roundResult := true
	outPrecision := uint(2)
	expected := "123456.79"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "79"

	nsDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr). "+
			"Err='%v' ", err.Error())
	}

	if initialPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Error: Expected Initial precision='%v'. Actual Initial precision='%v' ",
			initialPrecision, nsDto.GetPrecisionUint())
	}

	nsDto.SetThisPrecision(precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetThisPrecision(precision, roundResult). "+
			"precision= '%v'. Error= %v", precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	err = nsDto.IsValidInstanceError("Test 'nsDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nsDto.IsValidInstanceError() Error='%v'", err.Error())
	}

	if !nsDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_01(t *testing.T) {

	nStr := "123456.789"
	precision := uint(3)
	outPrecision := uint(6)
	expected := "123.456789"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_02(t *testing.T) {

	nStr := "123456.789"
	precision := uint(2)
	outPrecision := uint(5)
	expected := "1234.56789"
	signVal := 1
	absIntRuneStr := "1234"
	absFracRuneStr := "56789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_03(t *testing.T) {

	nStr := "123456.789"
	precision := uint(6)
	outPrecision := uint(9)
	expected := "0.123456789"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "123456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_04(t *testing.T) {

	nStr := "123456789"
	precision := uint(6)
	outPrecision := uint(6)
	expected := "123.456789"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_05(t *testing.T) {

	nStr := "123"
	precision := uint(5)
	outPrecision := uint(5)
	expected := "0.00123"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "00123"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_06(t *testing.T) {

	nStr := "0"
	precision := uint(3)
	outPrecision := uint(3)
	expected := "0.000"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "000"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_07(t *testing.T) {

	nStr := "0.000"
	precision := uint(2)
	outPrecision := uint(5)
	expected := "0.00000"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "00000"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_08(t *testing.T) {

	nStr := "123456.789"
	precision := uint(0)
	outPrecision := uint(3)
	expected := "123456.789"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_09(t *testing.T) {

	nStr := "-123456.789"
	precision := uint(0)
	outPrecision := uint(3)
	expected := "-123456.789"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_10(t *testing.T) {

	nStr := "-123456.789"
	precision := uint(3)
	outPrecision := uint(6)
	expected := "-123.456789"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_11(t *testing.T) {

	nStr := "-123456789"
	precision := uint(6)
	outPrecision := uint(6)
	expected := "-123.456789"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
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
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}
