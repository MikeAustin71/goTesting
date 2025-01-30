package datetime

import (
	"math/big"
	"testing"
)

func TestNumStrDto_GetSignedBigInt_01(t *testing.T) {

	nStr := "-123.456"
	signedNumStr := "-123456"
	expected, isOk := big.NewInt(0).SetString(signedNumStr, 10)

	if !isOk {
		t.Errorf("big.SetString(signedNumStr,10) Failed!. signedNumStr= '%v'", signedNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	signedBigInt, err := n1.GetBigInt()

	if signedBigInt.Cmp(expected) != 0 {
		t.Errorf("Expected signedBigInt= %v . Instead got, %v", expected.String(), signedBigInt.String())
	}

}

func TestNumStrDto_GetThouStr_01(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123,456.97"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouStr_02(t *testing.T) {

	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123.45"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouStr_03(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12,345.29"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouStr_04(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12,345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouStr_05(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12,345.1234"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouStr_06(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-12,345.29"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouStr_07(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-12,345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouStr_08(t *testing.T) {

	nStr := "-123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-123"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouStr_09(t *testing.T) {

	nStr := "-0.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-0.123"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouStr_10(t *testing.T) {

	nStr := "-1234567890.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-1,234,567,890.123"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_01(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123,456.97"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_02(t *testing.T) {

	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123.45"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_03(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12,345.29"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_04(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12,345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_05(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12,345.1234"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_06(t *testing.T) {

	nStr := "1234567890.25"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "1,234,567,890.25"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_07(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(12,345.29)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_08(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(12,345)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_09(t *testing.T) {

	nStr := "-123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(123)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_10(t *testing.T) {

	nStr := "-0.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(0.123)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_11(t *testing.T) {

	nStr := "-1234567890.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(1,234,567,890.123)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}
