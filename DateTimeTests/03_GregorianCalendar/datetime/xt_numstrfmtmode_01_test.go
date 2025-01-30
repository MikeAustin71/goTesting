package datetime

import "testing"

func TestNumStrFmtMode_String_01(t *testing.T) {

	r := PUREINTEGERFMT

	expectedStr := "PureIntegerString"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected PUREINTEGERFMT string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestNumStrFmtMode_String_02(t *testing.T) {

	r := INTSTRDECIMALFMT

	expectedStr := "IntegerDecimalString"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected INTSTRDECIMALFMT string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestNumStrFmtMode_String_03(t *testing.T) {

	r := THOUSANDSNUMSTRFMT

	expectedStr := "ThousandsNumString"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected THOUSANDSNUMSTRFMT string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestNumStrFmtMode_String_04(t *testing.T) {

	r := CURRENCYNUMSTRFMT

	expectedStr := "CurrencyNumString"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected THOUSANDSNUMSTRFMT string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestNumStrFmtMode_Value_01(t *testing.T) {

	var r NumStrFmtMode

	var i int

	r = PUREINTEGERFMT

	i = int(r)

	if i != 0 {
		t.Errorf("Expected 'PUREINTEGERFMT' value = 0. Instead, got %v", i)
	}

}

func TestNumStrFmtMode_Value_02(t *testing.T) {

	var r NumStrFmtMode

	var i int

	r = INTSTRDECIMALFMT

	i = int(r)

	if i != 1 {
		t.Errorf("Expected 'INTSTRDECIMALFMT' value = 1. Instead, got %v", i)
	}

}

func TestNumStrFmtMode_Value_03(t *testing.T) {

	var r NumStrFmtMode

	var i int

	r = THOUSANDSNUMSTRFMT

	i = int(r)

	if i != 2 {
		t.Errorf("Expected 'THOUSANDSNUMSTRFMT' value = 2. Instead, got %v", i)
	}

}

func TestNumStrFmtMode_Value_04(t *testing.T) {

	var r NumStrFmtMode

	var i int

	r = CURRENCYNUMSTRFMT

	i = int(r)

	if i != 3 {
		t.Errorf("Expected 'CURRENCYNUMSTRFMT' value = 3. Instead, got %v", i)
	}

}
