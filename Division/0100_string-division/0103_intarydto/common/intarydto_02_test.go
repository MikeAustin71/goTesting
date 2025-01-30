package common

import (
	"testing"
)


func TestIntAry_MultiplyByTenToPower_01(t *testing.T) {

	nStr := "457.3"
	eNumStr := "45730"
	nRunes := []rune("45730")
	eIAry := []int{4, 5, 7, 3, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(2, false)
	ia.ConvertIntAryToNumStr()

	if ia.NumStr != eNumStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if ia.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia.SignVal)
	}

	if lNRunes != ia.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia.NumRunesLen)
	}

	if lEArray != ia.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_MultiplyByTenToPower_02(t *testing.T) {

	nStr := "457.3"
	power := uint(2)
	eNumStr := "45730"
	nRunes := []rune("45730")
	eIAry := []int{4, 5, 7, 3, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power, true)

	if ia.NumStr != eNumStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if ia.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia.SignVal)
	}

	if lNRunes != ia.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia.NumRunesLen)
	}

	if lEArray != ia.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_MultiplyByTenToPower_03(t *testing.T) {

	nStr := "457.3"
	power := uint(10)
	eNumStr := "4573000000000"
	nRunes := []rune("4573000000000")
	eIAry := []int{4, 5, 7, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power, true)

	if ia.NumStr != eNumStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if ia.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia.SignVal)
	}

	if lNRunes != ia.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia.NumRunesLen)
	}

	if lEArray != ia.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_MultiplyByTenToPower_04(t *testing.T) {

	nStr := "457.3"
	power := uint(0)
	eNumStr := "457.3"
	nRunes := []rune("4573")
	eIAry := []int{4, 5, 7, 3}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power, false)
	ia.ConvertIntAryToNumStr()

	if ia.NumStr != eNumStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if ia.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia.SignVal)
	}

	if lNRunes != ia.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia.NumRunesLen)
	}

	if lEArray != ia.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_MultiplyByTenToPower_05(t *testing.T) {

	nStr := "-457.3"
	power := uint(1)
	eNumStr := "-4573"
	nRunes := []rune("4573")
	eIAry := []int{4, 5, 7, 3}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power, false)
	ia.ConvertIntAryToNumStr()

	if ia.NumStr != eNumStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if ia.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia.SignVal)
	}

	if lNRunes != ia.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia.NumRunesLen)
	}

	if lEArray != ia.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_MultiplyByTenToPower_06(t *testing.T) {

	nStr := "0"
	power := uint(2)
	eNumStr := "000"
	nRunes := []rune("000")
	eIAry := []int{0, 0, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power, true)

	if ia.NumStr != eNumStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if ia.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia.SignVal)
	}

	if lNRunes != ia.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia.NumRunesLen)
	}

	if lEArray != ia.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_MultiplyByTwoToPower_01(t *testing.T) {
	nStr1 := "23"
	expected := "12058624"
	power := uint(19)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.MultiplyByTwoToPower(power, true)

	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= %v. Instead, ia.NumStr= %v", expected, ia.NumStr)
	}

}

func TestIntAry_MultiplyByTwoToPower_02(t *testing.T) {
	nStr1 := "23"
	expected := "23"
	power := uint(0)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.MultiplyByTwoToPower(power, true)

	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= %v. Instead, ia.NumStr= %v", expected, ia.NumStr)
	}

}

func TestIntAry_MultiplyByTwoToPower_03(t *testing.T) {
	nStr1 := "23"
	expected := "46"
	power := uint(1)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.MultiplyByTwoToPower(power, true)

	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= %v. Instead, ia.NumStr= %v", expected, ia.NumStr)
	}

}

func TestIntAry_OptimizeIntArrayLen_01(t *testing.T) {
	nStr1 := "00579.123456000"
	expected := "579.123456000"
	ePrecision:= 9
	eLen := 12

	ia:= IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(false, true)

	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v'. Instead, ia.NumStr= '%v'", expected, ia.NumStr)
	}

	if ePrecision != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v'. Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eLen != ia.IntAryLen {
		t.Errorf("Expected ia.IntAryLen= '%v' .   Instead, ia.IntAryLen= '%v' .", eLen, ia.IntAryLen)
	}

}

func TestIntAry_OptimizeIntArrayLen_02(t *testing.T) {
	nStr1 := "00579.123456000"
	expected := "579.123456"
	ePrecision:= 6
	eLen := 9

	ia:= IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true, true)
	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v'. Instead, ia.NumStr= '%v'", expected, ia.NumStr)
	}

	if ePrecision != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v'. Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eLen != ia.IntAryLen {
		t.Errorf("Expected ia.IntAryLen= '%v' .   Instead, ia.IntAryLen= '%v' .", eLen, ia.IntAryLen)
	}

}

func TestIntAry_OptimizeIntArrayLen_03(t *testing.T) {
	nStr1 := "00579.000000000"
	expected := "579"
	ePrecision:= 0
	eLen := 3

	ia:= IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true, true)
	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v'. Instead, ia.NumStr= '%v'", expected, ia.NumStr)
	}

	if ePrecision != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v'. Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eLen != ia.IntAryLen {
		t.Errorf("Expected ia.IntAryLen= '%v' .   Instead, ia.IntAryLen= '%v' .", eLen, ia.IntAryLen)
	}

}

func TestIntAry_OptimizeIntArrayLen_04(t *testing.T) {
	nStr1 := "00000.123450000"
	expected := "0.12345"
	ePrecision:= 5
	eLen := 6

	ia:= IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true, true)
	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v'. Instead, ia.NumStr= '%v'", expected, ia.NumStr)
	}

	if ePrecision != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v'. Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eLen != ia.IntAryLen {
		t.Errorf("Expected ia.IntAryLen= '%v' .   Instead, ia.IntAryLen= '%v' .", eLen, ia.IntAryLen)
	}

}

func TestIntAry_OptimizeIntArrayLen_05(t *testing.T) {
	nStr1 := "00000.000123450000"
	expected := "0.00012345"
	ePrecision:= 8
	eLen := 9

	ia:= IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true, true)
	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v'. Instead, ia.NumStr= '%v'", expected, ia.NumStr)
	}

	if ePrecision != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v'. Instead, ia.Precision= '%v' .", ePrecision, ia.Precision)
	}

	if eLen != ia.IntAryLen {
		t.Errorf("Expected ia.IntAryLen= '%v' .   Instead, ia.IntAryLen= '%v' .", eLen, ia.IntAryLen)
	}

}

func TestIntAry_ResetFromBackUp_01(t *testing.T) {
	nStr1 := "99.4564"
	nStr2 := "-5034.123"
	eSign := 1
	ePrecsion := 4
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.CopyToBackUp()
	ia.SetIntAryWithNumStr(nStr2)

	if nStr2 != ia.NumStr {
		t.Errorf("Expected ia.NumStr2= '%v' .  Instead, ia.NumStr2= '%v' .", nStr2, ia.NumStr)
	}

	ia.ResetFromBackUp()

	if nStr1 != ia.NumStr {
		t.Errorf("After Reset - Expected ia.NumStr1= '%v' .  Instead, ia.NumStr1= '%v' .", nStr1, ia.NumStr)
	}

	if ia.SignVal != eSign {
		t.Errorf("After Reset, Expected ia.SignVal= '%v'. Instead, ia.SignVal= '%v' ", eSign, ia.SignVal)
	}

	if ia.Precision != ePrecsion {
		t.Errorf("After Reset, Expected ia.Precision= '%v'. Instead, ia.Precision= '%v' .", ePrecsion, ia.Precision)
	}

}

func TestIntAry_RoundToPrecision_01(t *testing.T) {
	nStr1 := "999.9952"
	expected := "1000.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_RoundToPrecision_02(t *testing.T) {
	nStr1 := "-999.9952"
	expected := "-1000.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_RoundToPrecision_03(t *testing.T) {
	nStr1 := "352.954"
	expected := "352.95"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_RoundToPrecision_04(t *testing.T) {
	nStr1 := "-352.954"
	expected := "-352.95"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_RoundToPrecision_05(t *testing.T) {
	nStr1 := "-352.954"
	expected := "-352.95"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_RoundToPrecision_06(t *testing.T) {
	nStr1 := "999.99"
	expected := "999.99"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_RoundToPrecision_07(t *testing.T) {
	nStr1 := "999.99"
	expected := "999.99000"
	precision := 5
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_RoundToPrecision_08(t *testing.T) {
	nStr1 := "0.000"
	expected := "0.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_RoundToPrecision_09(t *testing.T) {
	nStr1 := "999.995"
	expected := "999.995"
	precision := 3
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_SetEqualArrayLengths_01(t *testing.T) {
	nStr1 := "3536.123456"
	eNStr1 := "3536.123456"
	ePrecision1 := 6
	eSignVal1 := 1

	nStr2 := "12.14"
	eNStr2 := "0012.140000"
	ePrecision2 := 6
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.NumStr != eNStr1 {
		t.Errorf("Error - Expected ia1.NumStr= '%v' . Instead, ia1.NumStr= '%v' .", eNStr1, ia1.NumStr)
	}

	if ia1.SignVal != eSignVal1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", eSignVal1, ia1.SignVal)
	}

	if ia1.Precision != ePrecision1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", ePrecision1, ia1.Precision)
	}

	if ia2.NumStr != eNStr2 {
		t.Errorf("Error - Expected ia2.NumStr= '%v' . Instead, ia2.NumStr= '%v' .", eNStr2, ia2.NumStr)
	}

	if ia2.SignVal != eSignVal2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", eSignVal2, ia2.SignVal)
	}

	if ia2.Precision != ePrecision2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", ePrecision2, ia2.Precision)
	}

}

func TestIntAry_SetEqualArrayLengths_02(t *testing.T) {
	nStr1 := "12.14"
	eNStr1 := "0012.140000"
	ePrecision1 := 6
	eSignVal1 := 1

	nStr2 := "3536.123456"
	eNStr2 := "3536.123456"
	ePrecision2 := 6
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.NumStr != eNStr1 {
		t.Errorf("Error - Expected ia1.NumStr= '%v' . Instead, ia1.NumStr= '%v' .", eNStr1, ia1.NumStr)
	}

	if ia1.SignVal != eSignVal1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", eSignVal1, ia1.SignVal)
	}

	if ia1.Precision != ePrecision1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", ePrecision1, ia1.Precision)
	}

	if ia2.NumStr != eNStr2 {
		t.Errorf("Error - Expected ia2.NumStr= '%v' . Instead, ia2.NumStr= '%v' .", eNStr2, ia2.NumStr)
	}

	if ia2.SignVal != eSignVal2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", eSignVal2, ia2.SignVal)
	}

	if ia2.Precision != ePrecision2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", ePrecision2, ia2.Precision)
	}

}

func TestIntAry_SetEqualArrayLengths_03(t *testing.T) {
	nStr1 := "3536.123456"
	eNStr1 := "3536.123456"
	ePrecision1 := 6
	eSignVal1 := 1

	nStr2 := "-12.14"
	eNStr2 := "-0012.140000"
	ePrecision2 := 6
	eSignVal2 := -1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.NumStr != eNStr1 {
		t.Errorf("Error - Expected ia1.NumStr= '%v' . Instead, ia1.NumStr= '%v' .", eNStr1, ia1.NumStr)
	}

	if ia1.SignVal != eSignVal1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", eSignVal1, ia1.SignVal)
	}

	if ia1.Precision != ePrecision1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", ePrecision1, ia1.Precision)
	}

	if ia2.NumStr != eNStr2 {
		t.Errorf("Error - Expected ia2.NumStr= '%v' . Instead, ia2.NumStr= '%v' .", eNStr2, ia2.NumStr)
	}

	if ia2.SignVal != eSignVal2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", eSignVal2, ia2.SignVal)
	}

	if ia2.Precision != ePrecision2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", ePrecision2, ia2.Precision)
	}

}

func TestIntAry_SetEqualArrayLengths_04(t *testing.T) {
	nStr1 := "-12.14"
	eNStr1 := "-0012.140000"
	ePrecision1 := 6
	eSignVal1 := -1

	nStr2 := "3536.123456"
	eNStr2 := "3536.123456"
	ePrecision2 := 6
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.NumStr != eNStr1 {
		t.Errorf("Error - Expected ia1.NumStr= '%v' . Instead, ia1.NumStr= '%v' .", eNStr1, ia1.NumStr)
	}

	if ia1.SignVal != eSignVal1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", eSignVal1, ia1.SignVal)
	}

	if ia1.Precision != ePrecision1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", ePrecision1, ia1.Precision)
	}

	if ia2.NumStr != eNStr2 {
		t.Errorf("Error - Expected ia2.NumStr= '%v' . Instead, ia2.NumStr= '%v' .", eNStr2, ia2.NumStr)
	}

	if ia2.SignVal != eSignVal2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", eSignVal2, ia2.SignVal)
	}

	if ia2.Precision != ePrecision2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", ePrecision2, ia2.Precision)
	}

}

func TestIntAry_SetEqualArrayLengths_05(t *testing.T) {
	nStr1 := "-123456.143456"
	eNStr1 := "-123456.143456"
	ePrecision1 := 6
	eSignVal1 := -1

	nStr2 := "353678.123456"
	eNStr2 := "353678.123456"
	ePrecision2 := 6
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.NumStr != eNStr1 {
		t.Errorf("Error - Expected ia1.NumStr= '%v' . Instead, ia1.NumStr= '%v' .", eNStr1, ia1.NumStr)
	}

	if ia1.SignVal != eSignVal1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", eSignVal1, ia1.SignVal)
	}

	if ia1.Precision != ePrecision1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", ePrecision1, ia1.Precision)
	}

	if ia2.NumStr != eNStr2 {
		t.Errorf("Error - Expected ia2.NumStr= '%v' . Instead, ia2.NumStr= '%v' .", eNStr2, ia2.NumStr)
	}

	if ia2.SignVal != eSignVal2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", eSignVal2, ia2.SignVal)
	}

	if ia2.Precision != ePrecision2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", ePrecision2, ia2.Precision)
	}

}

func TestIntAry_SetEqualArrayLengths_06(t *testing.T) {
	nStr1 := "0.00"
	eNStr1 := "0.00"
	ePrecision1 := 2
	eSignVal1 := 1

	nStr2 := "0"
	eNStr2 := "0.00"
	ePrecision2 := 2
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.NumStr != eNStr1 {
		t.Errorf("Error - Expected ia1.NumStr= '%v' . Instead, ia1.NumStr= '%v' .", eNStr1, ia1.NumStr)
	}

	if ia1.SignVal != eSignVal1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", eSignVal1, ia1.SignVal)
	}

	if ia1.Precision != ePrecision1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", ePrecision1, ia1.Precision)
	}

	if ia2.NumStr != eNStr2 {
		t.Errorf("Error - Expected ia2.NumStr= '%v' . Instead, ia2.NumStr= '%v' .", eNStr2, ia2.NumStr)
	}

	if ia2.SignVal != eSignVal2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", eSignVal2, ia2.SignVal)
	}

	if ia2.Precision != ePrecision2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", ePrecision2, ia2.Precision)
	}

}

func TestIntAry_SetEqualArrayLengths_07(t *testing.T) {
	nStr1 := "0"
	eNStr1 := "0.00"
	ePrecision1 := 2
	eSignVal1 := 1

	nStr2 := "0.00"
	eNStr2 := "0.00"
	ePrecision2 := 2
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.NumStr != eNStr1 {
		t.Errorf("Error - Expected ia1.NumStr= '%v' . Instead, ia1.NumStr= '%v' .", eNStr1, ia1.NumStr)
	}

	if ia1.SignVal != eSignVal1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", eSignVal1, ia1.SignVal)
	}

	if ia1.Precision != ePrecision1 {
		t.Errorf("Error - Expected ia1.SignVal= '%v' . Instead, ia1.SignVal= '%v' .", ePrecision1, ia1.Precision)
	}

	if ia2.NumStr != eNStr2 {
		t.Errorf("Error - Expected ia2.NumStr= '%v' . Instead, ia2.NumStr= '%v' .", eNStr2, ia2.NumStr)
	}

	if ia2.SignVal != eSignVal2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", eSignVal2, ia2.SignVal)
	}

	if ia2.Precision != ePrecision2 {
		t.Errorf("Error - Expected ia2.SignVal= '%v' . Instead, ia2.SignVal= '%v' .", ePrecision2, ia2.Precision)
	}

}

func TestIntAry_SetIntAryWithInt64_01(t *testing.T) {

	num := uint64(123456789)

	eNumStr := "123456.789"
	ePrecision := uint(3)
	eSignVal := 1

	ia := IntAry{}.New()

	err:= ia.SetIntAryWithInt64(num, ePrecision, eSignVal)
	
	if err!=nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetIntAryWithInt64_02(t *testing.T) {

	num := uint64(123456789)

	eNumStr := "-12345.6789"
	ePrecision := uint(4)
	eSignVal := -1

	ia := IntAry{}.New()

	err:= ia.SetIntAryWithInt64(num, ePrecision, eSignVal)

	if err!=nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetIntAryWithInt64_03(t *testing.T) {

	num := uint64(0)

	eNumStr := "0.0000"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err:= ia.SetIntAryWithInt64(num, ePrecision, eSignVal)

	if err!=nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetIntAryWithInt64_04(t *testing.T) {

	num := uint64(32)

	eNumStr := "0.0032"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err:= ia.SetIntAryWithInt64(num, ePrecision, eSignVal)

	if err!=nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetIntAryWithInt64_05(t *testing.T) {

	num := uint64(32)

	eNumStr := "-32"
	ePrecision := uint(0)
	eSignVal := -1

	ia := IntAry{}.New()

	err:= ia.SetIntAryWithInt64(num, ePrecision, eSignVal)

	if err!=nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetIntAryWithInt64_06(t *testing.T) {

	num := uint64(32)

	eNumStr := "0.32"
	ePrecision := uint(2)
	eSignVal := 1

	ia := IntAry{}.New()

	err:= ia.SetIntAryWithInt64(num, ePrecision, eSignVal)

	if err!=nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetIntAryWithIntAry_01(t *testing.T) {
	iAry := []int{1,2,3,4,5,6,7,8,9}
	eNumStr := "123456.789"
	ePrecision := uint(3)
	eSignVal := 1

	ia := IntAry{}.New()

	err:= ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal)

	if err!=nil {
		t.Errorf("Error returned from ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetIntAryWithIntAry_02(t *testing.T) {
	iAry := []int{1,2,3,4,5,6,7,8,9}
	eNumStr := "-12345.6789"
	ePrecision := uint(4)
	eSignVal := -1

	ia := IntAry{}.New()

	err:= ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal)

	if err!=nil {
		t.Errorf("Error returned from ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetIntAryWithIntAry_03(t *testing.T) {
	iAry := []int{3,2}
	eNumStr := "0.0032"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err:= ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal)

	if err!=nil {
		t.Errorf("Error returned from ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetIntAryWithIntAry_04(t *testing.T) {
	iAry := []int{3,2}
	eNumStr := "0.32"
	ePrecision := uint(2)
	eSignVal := 1

	ia := IntAry{}.New()

	err:= ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal)

	if err!=nil {
		t.Errorf("Error returned from ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetWithNumStr_01(t *testing.T) {
	nStr := "123.456"
	nRunes := []rune("123456")
	eIAry := []int{1, 2, 3, 4, 5, 6}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.SetIntAryLength()

	if ia.NumStr != nStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", nStr, ia.NumStr)
	}

	if ia.Precision != 3 {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", 3, ia.Precision)
	}

	if ia.SignVal != 1 {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", 1, ia.SignVal)
	}

	if lNRunes != ia.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia.NumRunesLen)
	}

	if lEArray != ia.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_SetWithNumStr_02(t *testing.T) {
	nStr := "-12345.9"
	nRunes := []rune("123459")
	eIAry := []int{1, 2, 3, 4, 5, 9}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := -1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.SetIntAryLength()

	if ia.NumStr != nStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", nStr, ia.NumStr)
	}

	if ia.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia.SignVal)
	}

	if lNRunes != ia.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia.NumRunesLen)
	}

	if lEArray != ia.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_SetWithNumStr_03(t *testing.T) {
	nStrRaw := "-123  45.9"
	nStr := "-12345.9"
	nRunes := []rune("123459")
	eIAry := []int{1, 2, 3, 4, 5, 9}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := -1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStrRaw)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.SetIntAryLength()

	if ia.NumStr != nStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", nStr, ia.NumStr)
	}

	if ia.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia.SignVal)
	}

	if lNRunes != ia.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia.NumRunesLen)
	}

	if lEArray != ia.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_SetPrecision_01(t *testing.T) {
	nStr1 := "99.995"
	expected := "99.99"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, false)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_SetPrecision_02(t *testing.T) {
	nStr1 := "99.995"
	expected := "100.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, true)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_SetPrecision_03(t *testing.T) {
	nStr1 := "-0"
	expected := "0.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, true)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_SetPrecision_04(t *testing.T) {
	nStr1 := "-999.995"
	expected := "-1000.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, true)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_SetPrecision_05(t *testing.T) {
	nStr1 := "-999.995"
	expected := "-999.995"
	precision := 3
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, true)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_SetPrecision_06(t *testing.T) {
	nStr1 := "-999.995"
	expected := "-999.995000"
	precision := 6
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, true)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if ia.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.Precision)
	}

}

func TestIntAry_SetSignificantDigitIdxs_01(t *testing.T) {
	nStr := ".7770"
	eAryLen := 5
	eNumStr := "0.7770"
	eIntegerLen := 1
	eSigIntegerLen := 1
	eSigFractionLen := 3
	eIsZeroValue := false
	eIsIntegerZeroValue := true
	eFirstDigitIdx := 0
	eLastDigitIdx := 3
	eSignVal := 1
	ePrecision := 4

	ia := IntAry{}.New()
	err:=ia.SetIntAryWithNumStr(nStr)

	if err!=nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}


	if eFirstDigitIdx != ia.FirstDigitIdx {
		t.Errorf("Expected ia.FirstDigitIdx= '%v' .  Instead, ia.FirstDigitIdx= '%v' .", eFirstDigitIdx, ia.FirstDigitIdx)
	}

	if eLastDigitIdx != ia.LastDigitIdx {
		t.Errorf("Expected ia.LastDigitIdx= '%v' .  Instead, ia.LastDigitIdx= '%v' .", eLastDigitIdx, ia.LastDigitIdx)
	}


	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != ia.IsZeroValue {
		t.Errorf("Expected ia.IsZeroValue= '%v' .  Instead, ia.IsZeroValue= '%v' .", eIsZeroValue, ia.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != ia.IsIntegerZeroValue {
		t.Errorf("Expected ia.IsIntegerZeroValue= '%v' .  Instead, ia.IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, ia.IsIntegerZeroValue)
	}


	if eAryLen != ia.IntAryLen {
		t.Errorf("Expected ia.IntAryLen= '%v' .  Instead, ia.IntAryLen= '%v' .", eAryLen, ia.IntAryLen)
	}

	if eIntegerLen != ia.IntegerLen {
		t.Errorf("Expected ia.IntegerLen= '%v' .  Instead, ia.IntegerLen= '%v' .", eIntegerLen, ia.IntegerLen)
	}

	if eSigFractionLen != ia.SignificantFractionLen {
		t.Errorf("Expected ia.SignificantFractionLen= '%v' .  Instead, ia.SignificantFractionLen= '%v' .", eSigFractionLen, ia.SignificantFractionLen)
	}

	if eSigIntegerLen!= ia.SignificantIntegerLen {
		t.Errorf("Expected ia.SignificantIntegerLen= '%v' .  Instead, ia.SignificantIntegerLen= '%v' .", eSigIntegerLen, ia.SignificantIntegerLen)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetSignificantDigitIdxs_02(t *testing.T) {
	nStr := "000123456.123456000"
	eAryLen := 18
	eNumStr := "000123456.123456000"
	eIntegerLen := 9
	eSigIntegerLen := 6
	eSigFractionLen := 6
	eIsZeroValue := false
	eIsIntegerZeroValue := false
	eFirstDigitIdx := 3
	eLastDigitIdx := 14
	eSignVal := 1
	ePrecision := 9

	ia := IntAry{}.New()
	err:=ia.SetIntAryWithNumStr(nStr)

	if err!=nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}


	if eFirstDigitIdx != ia.FirstDigitIdx {
		t.Errorf("Expected ia.FirstDigitIdx= '%v' .  Instead, ia.FirstDigitIdx= '%v' .", eFirstDigitIdx, ia.FirstDigitIdx)
	}

	if eLastDigitIdx != ia.LastDigitIdx {
		t.Errorf("Expected ia.LastDigitIdx= '%v' .  Instead, ia.LastDigitIdx= '%v' .", eLastDigitIdx, ia.LastDigitIdx)
	}


	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != ia.IsZeroValue {
		t.Errorf("Expected ia.IsZeroValue= '%v' .  Instead, ia.IsZeroValue= '%v' .", eIsZeroValue, ia.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != ia.IsIntegerZeroValue {
		t.Errorf("Expected ia.IsIntegerZeroValue= '%v' .  Instead, ia.IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, ia.IsIntegerZeroValue)
	}


	if eAryLen != ia.IntAryLen {
		t.Errorf("Expected ia.IntAryLen= '%v' .  Instead, ia.IntAryLen= '%v' .", eAryLen, ia.IntAryLen)
	}

	if eIntegerLen != ia.IntegerLen {
		t.Errorf("Expected ia.IntegerLen= '%v' .  Instead, ia.IntegerLen= '%v' .", eIntegerLen, ia.IntegerLen)
	}

	if eSigFractionLen != ia.SignificantFractionLen {
		t.Errorf("Expected ia.SignificantFractionLen= '%v' .  Instead, ia.SignificantFractionLen= '%v' .", eSigFractionLen, ia.SignificantFractionLen)
	}

	if eSigIntegerLen!= ia.SignificantIntegerLen {
		t.Errorf("Expected ia.SignificantIntegerLen= '%v' .  Instead, ia.SignificantIntegerLen= '%v' .", eSigIntegerLen, ia.SignificantIntegerLen)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetSignificantDigitIdxs_03(t *testing.T) {
	nStr := "-000123456.123456000"
	eAryLen := 18
	eNumStr := "-000123456.123456000"
	eIntegerLen := 9
	eSigIntegerLen := 6
	eSigFractionLen := 6
	eIsZeroValue := false
	eIsIntegerZeroValue := false
	eFirstDigitIdx := 3
	eLastDigitIdx := 14
	eSignVal := -1
	ePrecision := 9

	ia := IntAry{}.New()
	err:=ia.SetIntAryWithNumStr(nStr)

	if err!=nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}


	if eFirstDigitIdx != ia.FirstDigitIdx {
		t.Errorf("Expected ia.FirstDigitIdx= '%v' .  Instead, ia.FirstDigitIdx= '%v' .", eFirstDigitIdx, ia.FirstDigitIdx)
	}

	if eLastDigitIdx != ia.LastDigitIdx {
		t.Errorf("Expected ia.LastDigitIdx= '%v' .  Instead, ia.LastDigitIdx= '%v' .", eLastDigitIdx, ia.LastDigitIdx)
	}


	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != ia.IsZeroValue {
		t.Errorf("Expected ia.IsZeroValue= '%v' .  Instead, ia.IsZeroValue= '%v' .", eIsZeroValue, ia.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != ia.IsIntegerZeroValue {
		t.Errorf("Expected ia.IsIntegerZeroValue= '%v' .  Instead, ia.IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, ia.IsIntegerZeroValue)
	}


	if eAryLen != ia.IntAryLen {
		t.Errorf("Expected ia.IntAryLen= '%v' .  Instead, ia.IntAryLen= '%v' .", eAryLen, ia.IntAryLen)
	}

	if eIntegerLen != ia.IntegerLen {
		t.Errorf("Expected ia.IntegerLen= '%v' .  Instead, ia.IntegerLen= '%v' .", eIntegerLen, ia.IntegerLen)
	}

	if eSigFractionLen != ia.SignificantFractionLen {
		t.Errorf("Expected ia.SignificantFractionLen= '%v' .  Instead, ia.SignificantFractionLen= '%v' .", eSigFractionLen, ia.SignificantFractionLen)
	}

	if eSigIntegerLen!= ia.SignificantIntegerLen {
		t.Errorf("Expected ia.SignificantIntegerLen= '%v' .  Instead, ia.SignificantIntegerLen= '%v' .", eSigIntegerLen, ia.SignificantIntegerLen)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetSignificantDigitIdxs_04(t *testing.T) {
	nStr := "000.123456000"
	eAryLen := 12
	eNumStr := "000.123456000"
	eIntegerLen := 3
	eSigIntegerLen := 1
	eSigFractionLen := 6
	eIsZeroValue := false
	eIsIntegerZeroValue := true
	eFirstDigitIdx := 2
	eLastDigitIdx := 8
	eSignVal := 1
	ePrecision := 9

	ia := IntAry{}.New()
	err:=ia.SetIntAryWithNumStr(nStr)

	if err!=nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}


	if eFirstDigitIdx != ia.FirstDigitIdx {
		t.Errorf("Expected ia.FirstDigitIdx= '%v' .  Instead, ia.FirstDigitIdx= '%v' .", eFirstDigitIdx, ia.FirstDigitIdx)
	}

	if eLastDigitIdx != ia.LastDigitIdx {
		t.Errorf("Expected ia.LastDigitIdx= '%v' .  Instead, ia.LastDigitIdx= '%v' .", eLastDigitIdx, ia.LastDigitIdx)
	}


	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != ia.IsZeroValue {
		t.Errorf("Expected ia.IsZeroValue= '%v' .  Instead, ia.IsZeroValue= '%v' .", eIsZeroValue, ia.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != ia.IsIntegerZeroValue {
		t.Errorf("Expected ia.IsIntegerZeroValue= '%v' .  Instead, ia.IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, ia.IsIntegerZeroValue)
	}


	if eAryLen != ia.IntAryLen {
		t.Errorf("Expected ia.IntAryLen= '%v' .  Instead, ia.IntAryLen= '%v' .", eAryLen, ia.IntAryLen)
	}

	if eIntegerLen != ia.IntegerLen {
		t.Errorf("Expected ia.IntegerLen= '%v' .  Instead, ia.IntegerLen= '%v' .", eIntegerLen, ia.IntegerLen)
	}

	if eSigFractionLen != ia.SignificantFractionLen {
		t.Errorf("Expected ia.SignificantFractionLen= '%v' .  Instead, ia.SignificantFractionLen= '%v' .", eSigFractionLen, ia.SignificantFractionLen)
	}

	if eSigIntegerLen!= ia.SignificantIntegerLen {
		t.Errorf("Expected ia.SignificantIntegerLen= '%v' .  Instead, ia.SignificantIntegerLen= '%v' .", eSigIntegerLen, ia.SignificantIntegerLen)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetSignificantDigitIdxs_05(t *testing.T) {
	nStr := "256"
	eAryLen := 3
	eNumStr := "256"
	eIntegerLen := 3
	eSigIntegerLen := 3
	eSigFractionLen := 0
	eIsZeroValue := false
	eIsIntegerZeroValue := false
	eFirstDigitIdx := 0
	eLastDigitIdx := 2
	eSignVal := 1
	ePrecision := 0

	ia := IntAry{}.New()
	err:=ia.SetIntAryWithNumStr(nStr)

	if err!=nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}


	if eFirstDigitIdx != ia.FirstDigitIdx {
		t.Errorf("Expected ia.FirstDigitIdx= '%v' .  Instead, ia.FirstDigitIdx= '%v' .", eFirstDigitIdx, ia.FirstDigitIdx)
	}

	if eLastDigitIdx != ia.LastDigitIdx {
		t.Errorf("Expected ia.LastDigitIdx= '%v' .  Instead, ia.LastDigitIdx= '%v' .", eLastDigitIdx, ia.LastDigitIdx)
	}


	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != ia.IsZeroValue {
		t.Errorf("Expected ia.IsZeroValue= '%v' .  Instead, ia.IsZeroValue= '%v' .", eIsZeroValue, ia.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != ia.IsIntegerZeroValue {
		t.Errorf("Expected ia.IsIntegerZeroValue= '%v' .  Instead, ia.IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, ia.IsIntegerZeroValue)
	}


	if eAryLen != ia.IntAryLen {
		t.Errorf("Expected ia.IntAryLen= '%v' .  Instead, ia.IntAryLen= '%v' .", eAryLen, ia.IntAryLen)
	}

	if eIntegerLen != ia.IntegerLen {
		t.Errorf("Expected ia.IntegerLen= '%v' .  Instead, ia.IntegerLen= '%v' .", eIntegerLen, ia.IntegerLen)
	}

	if eSigFractionLen != ia.SignificantFractionLen {
		t.Errorf("Expected ia.SignificantFractionLen= '%v' .  Instead, ia.SignificantFractionLen= '%v' .", eSigFractionLen, ia.SignificantFractionLen)
	}

	if eSigIntegerLen!= ia.SignificantIntegerLen {
		t.Errorf("Expected ia.SignificantIntegerLen= '%v' .  Instead, ia.SignificantIntegerLen= '%v' .", eSigIntegerLen, ia.SignificantIntegerLen)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SetSignificantDigitIdxs_06(t *testing.T) {
	nStr := "000256"
	eAryLen := 6
	eNumStr := "000256"
	eIntegerLen := 6
	eSigIntegerLen := 3
	eSigFractionLen := 0
	eIsZeroValue := false
	eIsIntegerZeroValue := false
	eFirstDigitIdx := 3
	eLastDigitIdx := 5
	eSignVal := 1
	ePrecision := 0

	ia := IntAry{}.New()
	err:=ia.SetIntAryWithNumStr(nStr)

	if err!=nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}


	if eFirstDigitIdx != ia.FirstDigitIdx {
		t.Errorf("Expected ia.FirstDigitIdx= '%v' .  Instead, ia.FirstDigitIdx= '%v' .", eFirstDigitIdx, ia.FirstDigitIdx)
	}

	if eLastDigitIdx != ia.LastDigitIdx {
		t.Errorf("Expected ia.LastDigitIdx= '%v' .  Instead, ia.LastDigitIdx= '%v' .", eLastDigitIdx, ia.LastDigitIdx)
	}


	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != ia.IsZeroValue {
		t.Errorf("Expected ia.IsZeroValue= '%v' .  Instead, ia.IsZeroValue= '%v' .", eIsZeroValue, ia.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != ia.IsIntegerZeroValue {
		t.Errorf("Expected ia.IsIntegerZeroValue= '%v' .  Instead, ia.IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, ia.IsIntegerZeroValue)
	}


	if eAryLen != ia.IntAryLen {
		t.Errorf("Expected ia.IntAryLen= '%v' .  Instead, ia.IntAryLen= '%v' .", eAryLen, ia.IntAryLen)
	}

	if eIntegerLen != ia.IntegerLen {
		t.Errorf("Expected ia.IntegerLen= '%v' .  Instead, ia.IntegerLen= '%v' .", eIntegerLen, ia.IntegerLen)
	}

	if eSigFractionLen != ia.SignificantFractionLen {
		t.Errorf("Expected ia.SignificantFractionLen= '%v' .  Instead, ia.SignificantFractionLen= '%v' .", eSigFractionLen, ia.SignificantFractionLen)
	}

	if eSigIntegerLen!= ia.SignificantIntegerLen {
		t.Errorf("Expected ia.SignificantIntegerLen= '%v' .  Instead, ia.SignificantIntegerLen= '%v' .", eSigIntegerLen, ia.SignificantIntegerLen)
	}

	if eNumStr != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v'", eNumStr, ia.NumStr)
	}

	if int(ePrecision) != ia.Precision {
		t.Errorf("Expected ia.Precision= '%v' . Instead, ia.Precision= '%v'", ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v' . Instead, ia.SignVal= '%v'", eSignVal, ia.SignVal)
	}

}

func TestIntAry_SubtractFromThis_01(t *testing.T) {
	nStr1 := "900.777"
	nStr2 := "901.000"
	eNumStr := "-0.223"
	ePrecision := 3
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_02(t *testing.T) {
	nStr1 := "350"
	nStr2 := "122"
	eNumStr := "228"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_03(t *testing.T) {
	nStr1 := "-350"
	nStr2 := "122"
	eNumStr := "-472"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_04(t *testing.T) {
	nStr1 := "-350"
	nStr2 := "-122"
	eNumStr := "-228"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_05(t *testing.T) {
	nStr1 := "350"
	nStr2 := "-122"
	eNumStr := "472"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_06(t *testing.T) {
	nStr1 := "350"
	nStr2 := "0"
	eNumStr := "350"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_07(t *testing.T) {
	nStr1 := "-350"
	nStr2 := "0"
	eNumStr := "-350"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_08(t *testing.T) {
	nStr1 := "122"
	nStr2 := "350"
	eNumStr := "-228"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_09(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "350"
	eNumStr := "-472"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_10(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "-350"
	eNumStr := "228"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_11(t *testing.T) {
	nStr1 := "122"
	nStr2 := "-350"
	eNumStr := "472"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_12(t *testing.T) {
	nStr1 := "0"
	nStr2 := "350"
	eNumStr := "-350"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_13(t *testing.T) {
	nStr1 := "0"
	nStr2 := "-350"
	eNumStr := "350"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_14(t *testing.T) {
	nStr1 := "122"
	nStr2 := "122"
	eNumStr := "0"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_15(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "122"
	eNumStr := "-244"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_16(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "-122"
	eNumStr := "0"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_17(t *testing.T) {
	nStr1 := "122"
	nStr2 := "-122"
	eNumStr := "244"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_18(t *testing.T) {
	nStr1 := "0"
	nStr2 := "0"
	eNumStr := "0"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

func TestIntAry_SubtractFromThis_19(t *testing.T) {
	nStr1 := "1.122"
	nStr2 := "4.5"
	eNumStr := "-3.378"
	ePrecision := 3
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2, true)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.NumStr {
		t.Errorf("Error - Expected IFinal.NumStr= '%v' .  Instead, IFinal.NumStr= '%v' .", eNumStr, ia1.NumStr)
	}

	if ePrecision != ia1.Precision {
		t.Errorf("Error - Expected IFinal.Precision= '%v' .  Instead, IFinal.Precision= '%v' .", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error - Expected IFinal.SignVal= '%v' .  Instead, IFinal.SignVal= '%v' .", eSignVal, ia1.SignVal)
	}

}

