package datetime

import (
	"math/big"
	"testing"
)

func TestNumStrDto_GetAbsoluteBigInt_01(t *testing.T) {

	nStr := "-123.456"
	absNumStr := "123456"
	expected, isOk := big.NewInt(0).SetString(absNumStr, 10)

	if !isOk {
		t.Errorf("big.SetString(absNumStr,10) Failed!. absNumStr= '%v'", absNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	absBigInt, err := n1.GetAbsoluteBigInt()

	if absBigInt.Cmp(expected) != 0 {
		t.Errorf("Expected absBigInt= %v . Instead got, %v", expected.String(), absBigInt.String())
	}

}

func TestNumStrDto_GetAbsFracRunes_01(t *testing.T) {
	nStr := "-123.456"
	expectedRunes := []rune{'4', '5', '6'}

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr). "+
			"nStr='%v' Error='%v'",
			nStr, err.Error())
	}

	fracRunes := nDto.GetAbsFracRunes()

	if len(expectedRunes) != len(fracRunes) {
		t.Errorf("Expected a rune array length of '%v'.  Instead length='%v' .",
			len(expectedRunes), len(fracRunes))
	}

	for i := 0; i < len(fracRunes); i++ {
		if expectedRunes[i] != fracRunes[i] {
			t.Errorf("Error! Rune Idx='%v'. Expected rune='%v'. Instead rune='%v'. ",
				i, expectedRunes[i], fracRunes[i])
		}
	}

}

func TestNumStrDto_GetAbsFracRunes_02(t *testing.T) {
	nStr := "123"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr). "+
			"nStr='%v' Error='%v'",
			nStr, err.Error())
	}

	fracRunes := nDto.GetAbsFracRunes()

	if len(fracRunes) != 0 {
		t.Errorf("Expected rune array length=0. Instead length='%v' Array String='%v'",
			len(fracRunes), string(fracRunes))
	}

}

func TestNumStrDto_GetAbsIntRunes_01(t *testing.T) {
	nStr := "-123.456"
	expectedRunes := []rune{'1', '2', '3'}

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	intRunes := nDto.GetAbsIntRunes()

	if len(expectedRunes) != len(intRunes) {
		t.Errorf("Error: Expected Rune Array Length='%v'.  Instead, Array Length='%v'",
			len(expectedRunes), len(intRunes))
	}

	for i := 0; i < len(intRunes); i++ {
		if expectedRunes[i] != intRunes[i] {
			t.Errorf("Error: Index='%v' Expected Rune='%v'. Actual Rune='%v'",
				i, expectedRunes[i], intRunes[i])
		}
	}

}

func TestNumStrDto_GetAbsIntRunes_02(t *testing.T) {
	nStr := "-123"
	expectedRunes := []rune{'1', '2', '3'}

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	intRunes := nDto.GetAbsIntRunes()

	if len(expectedRunes) != len(intRunes) {
		t.Errorf("Error: Expected Rune Array Length='%v'.  Instead, Array Length='%v'",
			len(expectedRunes), len(intRunes))
	}

	for i := 0; i < len(intRunes); i++ {
		if expectedRunes[i] != intRunes[i] {
			t.Errorf("Error: Index='%v' Expected Rune='%v'. Actual Rune='%v'",
				i, expectedRunes[i], intRunes[i])
		}
	}

}

func TestNumStrDto_GetAbsIntRunes_03(t *testing.T) {
	nStr := "123456789.7865"
	expectedRunes := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	intRunes := nDto.GetAbsIntRunes()

	if len(expectedRunes) != len(intRunes) {
		t.Errorf("Error: Expected Rune Array Length='%v'.  Instead, Array Length='%v'",
			len(expectedRunes), len(intRunes))
	}

	for i := 0; i < len(intRunes); i++ {
		if expectedRunes[i] != intRunes[i] {
			t.Errorf("Error: Index='%v' Expected Rune='%v'. Actual Rune='%v'",
				i, expectedRunes[i], intRunes[i])
		}
	}

}

func TestNumStrDto_GetBigIntNum_01(t *testing.T) {

	bigI := big.NewInt(int64(123456123456))
	precision := 6

	exStr := "123456.123456"

	expectedBigIntNum := BigIntNum{}.NewBigInt(bigI, uint(precision))

	intAry, err := IntAry{}.NewBigInt(bigI, precision)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewBigInt(bigI, precision). "+
			"Error='%v' ", err.Error())
	}

	bigINum, err := intAry.GetBigIntNum()

	if err != nil {
		t.Errorf("Error returned by intAry.GetBigIntNum(). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBigIntNum.Equal(bigINum) {
		t.Errorf("Error: Expected BigIntNum NOT Equal to Actual BigIntNum! "+
			"expectedBi='%v', expectedPrecision='%v'. actualBi='%v' actualPrecision='%v'",
			expectedBigIntNum.bigInt.Text(10), expectedBigIntNum.precision,
			bigINum.bigInt.Text(10), bigINum.precision)
	}

	actualNumStr := bigINum.GetNumStr()

	if exStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'",
			exStr, actualNumStr)
	}
}

func TestNumStrDto_GetCurrencyStr_01(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$123,456.97"

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

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_02(t *testing.T) {

	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$123.45"

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

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_03(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$12,345.29"

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

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_04(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$12,345"

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

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_05(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$12,345.1234"

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

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_06(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-$12,345.29"

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

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_07(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-$12,345"

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

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_08(t *testing.T) {

	nStr := "-12345"
	// '\U000020ac', // Euro €
	currencySymbol := '€'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-€12,345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetCurrencySymbol('\U000020ac')

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

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_09(t *testing.T) {

	nStr := "12345.12"
	// '\U000020ac', // Euro €
	currencySymbol := '€'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "€12,345.12"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetCurrencySymbol('\U000020ac')

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

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_01(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$123,456.97"

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

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_02(t *testing.T) {

	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$123.45"

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

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_03(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$12,345.29"

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

	actualStr := nDto.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_04(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$12,345"

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

	actualStr := nDto.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_05(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "($12,345.29)"

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

	actualStr := nDto.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_06(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "($12,345)"

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

	actualStr := nDto.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_07(t *testing.T) {

	nStr := "-12345"
	// '\U000020ac', // Euro €
	currencySymbol := '€'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(€12,345)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetCurrencySymbol('\U000020ac')

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

	actualStr := nDto.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_08(t *testing.T) {

	nStr := "12345.12"
	// '\U000020ac', // Euro €
	currencySymbol := '€'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "€12,345.12"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetCurrencySymbol('\U000020ac')

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

	actualStr := nDto.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetDecimal_01(t *testing.T) {

	numStr := "198649257.12345678"

	controlDecimal, err := Decimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by controlDecimal = Decimal{}.NewNumStr(numStr). "+
			"numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	controlNDto, err := NumStrDto{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by controlNDto = NumStrDto{}.NewNumStr(numStr). "+
			"numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	decActual, err := controlNDto.GetDecimal()

	if err != nil {
		t.Errorf("Error returned by decActual = controlNDto.GetDecimal(). "+
			"Error='%v'", err.Error())
	}

	if numStr != decActual.GetNumStr() {
		t.Errorf("Error: Expected decActual.GetNumStr()='%v'. Instead, "+
			"decActual.GetNumStr()='%v'.",
			numStr, decActual.GetNumStr())
	}

	if !controlDecimal.Equal(decActual) {
		t.Errorf("Error: controlDecimal NOT EQUAL to decActual! "+
			"controlDecimal='%v' decActual='%v'",
			controlDecimal.GetNumStr(), decActual.GetNumStr())
	}
}

func TestNumStrDto_GetIntAry(t *testing.T) {

	numStr := "589627.123456"

	controlNDto, err := NumStrDto{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStr). "+
			"numStr='%v'  Error='%v'", numStr, err.Error())
	}

	controlIa, err := IntAry{}.NewNumStr(numStr)

	actualIa, err := controlNDto.GetIntAry()

	if err != nil {
		t.Errorf("Error returned by controlNDto.GetIntAryElements(). "+
			"numStr='%v'  Error='%v'", numStr, err.Error())
	}

	if numStr != actualIa.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			numStr, actualIa.GetNumStr())
	}

	if !controlIa.Equal(actualIa) {
		t.Errorf("Error: controlIa NOT EQUAL to actual actualIa! "+
			"controlNDto='%v' nDto='%v'",
			controlNDto.GetNumStr(), actualIa.GetNumStr())
	}

}

func TestNumStrDto_GetNumStr_01(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123456.97"

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

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_02(t *testing.T) {

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

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_03(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.29"

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

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_04(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345"

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

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_05(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.1234"

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

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_06(t *testing.T) {

	nStr := "1234567890.25"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "1234567890.25"

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

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_07(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-12345.29"

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

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_08(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-12345"

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

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_09(t *testing.T) {

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

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_10(t *testing.T) {

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

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_11(t *testing.T) {

	nStr := "-1234567890.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-1234567890.123"

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

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_01(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123456.97"

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

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_02(t *testing.T) {

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

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_03(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.29"

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

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_04(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345"

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

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_05(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.1234"

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

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_06(t *testing.T) {

	nStr := "1234567890.25"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "1234567890.25"

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

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_07(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(12345.29)"

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

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_08(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(12345)"

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

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_09(t *testing.T) {

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

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_10(t *testing.T) {

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

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_11(t *testing.T) {

	nStr := "-1234567890.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(1234567890.123)"

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

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}
