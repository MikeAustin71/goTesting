package common

import (
"testing"
)

func TestIntAry_AddToThis_01(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "457.3"
	nStr2 := "22.2"
	expected := "479.5"
	nRunes := []rune("4795")
	eIAry := []int{4, 7, 9, 5}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_02(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "457.325"
	nStr2 := "-22.2"
	expected := "435.125"
	nRunes := []rune("435125")
	eIAry := []int{4, 3, 5, 1, 2, 5}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_03(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-457.325"
	nStr2 := "22.2"
	expected := "-435.125"
	nRunes := []rune("435125")
	eIAry := []int{4, 3, 5, 1, 2, 5}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_04(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-457.325"
	nStr2 := "-22.2"
	expected := "-479.525"
	nRunes := []rune("479525")
	eIAry := []int{4, 7, 9, 5, 2, 5}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_05(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "0.000"
	nStr2 := "-22.2"
	expected := "-22.200"
	nRunes := []rune("22200")
	eIAry := []int{2, 2, 2, 0, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_06(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "0.000"
	nStr2 := "0"
	expected := "0.000"
	nRunes := []rune("0000")
	eIAry := []int{0, 0, 0, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_07(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "99.225"
	nStr2 := "-99.1"
	expected := "0.125"
	nRunes := []rune("0125")
	eIAry := []int{0, 1, 2, 5}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_08(t *testing.T) {
	// N1 > N2 + and +
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "350"
	nStr2 := "122"
	expected := "472"
	nRunes := []rune("472")
	eIAry := []int{4, 7, 2}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_09(t *testing.T) {
	// N1 > N2 - and +
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-350"
	nStr2 := "122"
	expected := "-228"
	nRunes := []rune("228")
	eIAry := []int{2, 2, 8}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_10(t *testing.T) {
	// N1 > N2 - and -
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-350"
	nStr2 := "-122"
	expected := "-472"
	nRunes := []rune("472")
	eIAry := []int{4, 7, 2}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_11(t *testing.T) {
	// N1 > N2 + and -
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "350"
	nStr2 := "-122"
	expected := "228"
	nRunes := []rune("228")
	eIAry := []int{2, 2, 8}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_12(t *testing.T) {
	// N1 > N2  350 + 0
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "350"
	nStr2 := "0"
	expected := "350"
	nRunes := []rune("350")
	eIAry := []int{3, 5, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_13(t *testing.T) {
	// N1 > N2  -350 + 0
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-350"
	nStr2 := "0"
	expected := "-350"
	nRunes := []rune("350")
	eIAry := []int{3, 5, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_14(t *testing.T) {
	// N2 > N1  + and +
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "122"
	nStr2 := "350"
	expected := "472"
	nRunes := []rune("472")
	eIAry := []int{4, 7, 2}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_15(t *testing.T) {
	// N2 > N1  - and +
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-122"
	nStr2 := "350"
	expected := "228"
	nRunes := []rune("228")
	eIAry := []int{2, 2, 8}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_16(t *testing.T) {
	// N2 > N1  - and -
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-122"
	nStr2 := "-350"
	expected := "-472"
	nRunes := []rune("472")
	eIAry := []int{4, 7, 2}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_17(t *testing.T) {
	// N2 > N1  + and -
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "122"
	nStr2 := "-350"
	expected := "-228"
	nRunes := []rune("228")
	eIAry := []int{2, 2, 8}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_18(t *testing.T) {
	// N2 > N1  0 and +350
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "0"
	nStr2 := "350"
	expected := "350"
	nRunes := []rune("350")
	eIAry := []int{3, 5, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_19(t *testing.T) {
	// N2 > N1  0 and -350
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "0"
	nStr2 := "-350"
	expected := "-350"
	nRunes := []rune("350")
	eIAry := []int{3, 5, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_20(t *testing.T) {
	// N1 == N2  +122 and +122
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "122"
	nStr2 := "122"
	expected := "244"
	nRunes := []rune("244")
	eIAry := []int{2, 4, 4}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_21(t *testing.T) {
	// N1 == N2  -122 and +122
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-122"
	nStr2 := "122"
	expected := "0"
	nRunes := []rune("0")
	eIAry := []int{0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_22(t *testing.T) {
	// N1 == N2  -122 and -122
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-122"
	nStr2 := "-122"
	expected := "-244"
	nRunes := []rune("244")
	eIAry := []int{2, 4, 4}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_23(t *testing.T) {
	// N1 == N2  +122 and -122
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "122"
	nStr2 := "-122"
	expected := "0"
	nRunes := []rune("0")
	eIAry := []int{0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_AddToThis_24(t *testing.T) {
	// N1 == N2  0 and 0
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "0"
	nStr2 := "0"
	expected := "0"
	nRunes := []rune("0")
	eIAry := []int{0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2, true)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.NumStr

	if s != expected {
		t.Errorf("Expected IFinal.NumStr= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.Precision != ePrecision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", ePrecision, ia1.Precision)
	}

	if eSignVal != ia1.SignVal {
		t.Errorf("Error: Expected SignVal= '%v'. Instead received SignVal= '%v'", eSignVal, ia1.SignVal)
	}

	if lNRunes != ia1.NumRunesLen {
		t.Errorf("Error: Expected NumRunes Length= '%v'. Instead received NumRunes Length= '%v'", lNRunes, ia1.NumRunesLen)
	}

	if lEArray != ia1.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.IntAryLen)
	}

	for i := 0; i < lNRunes; i++ {

		if nRunes[i] != ia1.NumRunes[i] {
			t.Error("Error: Expected nRunes Array does NOT match ia.NumRunes Array! ")
			return
		}

	}

	for i := 0; i < lEArray; i++ {
		if eIAry[i] != ia1.IntAry[i] {

			t.Error("Error: Expected IntAry Array does NOT match ia.IntAry! ")
			return

		}
	}

}

func TestIntAry_Ceiling_01(t *testing.T) {
	nStr1 := "0.925"
	expected := "1.000"
	precision := 3

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_Ceiling_02(t *testing.T) {
	nStr1 := "-2.7"
	expected := "-2.0"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_Ceiling_03(t *testing.T) {
	nStr1 := "2.9"
	expected := "3.0"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}
func TestIntAry_Ceiling_04(t *testing.T) {
	nStr1 := "2.0"
	expected := "2.0"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_Ceiling_05(t *testing.T) {
	nStr1 := "2.4"
	expected := "3.0"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_Ceiling_06(t *testing.T) {
	nStr1 := "2.9"
	expected := "3.0"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_Ceiling_07(t *testing.T) {
	nStr1 := "-2"
	expected := "-2"
	precision := 0

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_CompareSignedValues_01(t *testing.T) {

	nStr1 := "451.3"
	nStr2 := "451.2"
	expectedCompare := 1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_02(t *testing.T) {

	nStr1 := "45.13"
	nStr2 := "451.2"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_03(t *testing.T) {

	nStr1 := "45.13975"
	nStr2 := "-451.21"
	expectedCompare := 1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_04(t *testing.T) {

	nStr1 := "0"
	nStr2 := "0.00"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_05(t *testing.T) {

	nStr1 := "-625.414"
	nStr2 := "-625.413"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_06(t *testing.T) {

	nStr1 := "625.414"
	nStr2 := "625.413000"
	expectedCompare := 1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_07(t *testing.T) {

	nStr1 := "625.413"
	nStr2 := "625.413001"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_08(t *testing.T) {

	nStr1 := "625.413"
	nStr2 := "625.413000"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_09(t *testing.T) {

	nStr1 := "625.413"
	nStr2 := "00625.413000"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_10(t *testing.T) {

	nStr1 := "625.413"
	nStr2 := "-00625.413000"
	expectedCompare := 1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_11(t *testing.T) {

	nStr1 := "-625.413"
	nStr2 := "-00625.413000"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_12(t *testing.T) {

	nStr1 := "-625.413"
	nStr2 := "00625.413000"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_01(t *testing.T) {

	nStr1 := "45.13"
	nStr2 := "451.2"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_02(t *testing.T) {

	nStr1 := "45.13"
	nStr2 := "-451.2"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_03(t *testing.T) {

	nStr1 := "-45.13"
	nStr2 := "45.13"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_04(t *testing.T) {

	nStr1 := "-45.13000"
	nStr2 := "45.13"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_05(t *testing.T) {

	nStr1 := "-45.13001"
	nStr2 := "45.13"
	expectedCompare := 1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_06(t *testing.T) {

	nStr1 := "-45.1300"
	nStr2 := "452"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CopyToBackUp_01(t *testing.T) {
	nStr1 := "100.123"
	nStr2 := "-52.795813"
	expected := "100.123"

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.CopyToBackUp()
	ia.SetIntAryWithNumStr(nStr2)

	if nStr2 != ia.NumStr {
		t.Errorf("Error - Expected ia.NumStr= '%v' . Instead, ia.NumStr= '%v' .", nStr2, ia.NumStr)
	}

	str := ia.BackUp.NumStr

	if expected != str {
		t.Errorf("Error - Expected ia.BackUp.NumStr= '%v' .  Instead, ia.BackUp.NumStr= '%v' .", expected, str)
	}

}

func TestIntAry_DecrementIntegerOne_01(t *testing.T) {
	nStr1 := "100.123"
	expected := "-100.123"
	cycles := 200
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.DecrementIntegerOne()
	}

	ia.ConvertIntAryToNumStr()

	if expected != ia.NumStr {
		t.Errorf("Error - Expected NumStr= '%v'. Instead, NumStr= '%v'", expected, ia.NumStr)
	}

}

func TestIntAry_DecrementIntegerOne_02(t *testing.T) {
	nStr1 := "2000"
	expected := "-2000"
	cycles := 4000
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.DecrementIntegerOne()
	}

	ia.ConvertIntAryToNumStr()

	if expected != ia.NumStr {
		t.Errorf("Error - Expected NumStr= '%v'. Instead, NumStr= '%v'", expected, ia.NumStr)
	}

}

func TestIntAry_DecrementIntegerOne_03(t *testing.T) {
	nStr1 := "2000.123"
	expected := "-2000.123"
	cycles := 4000
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.DecrementIntegerOne()
	}

	ia.ConvertIntAryToNumStr()

	if expected != ia.NumStr {
		t.Errorf("Error - Expected NumStr= '%v'. Instead, NumStr= '%v'", expected, ia.NumStr)
	}

}

func TestIntAry_DivideByInt64_01(t *testing.T) {
	nStr1 := "579"
	expected := "17.545454545454545454545454545455"
	maxPrecision := uint64(30)
	divisor := int64(33)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.DivideByInt64(divisor, maxPrecision, true )

	if err != nil {
		t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
	}

	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v'. Instead, ia.Numstr= '%v'", expected, ia.NumStr)
	}

	if ia.Precision != int(maxPrecision) {
		t.Errorf("Expected ia.Precision= '%v'. Instead, ia.Precision= '%v'.",maxPrecision, ia.Precision)
	}

}

func TestIntAry_DivideByInt64_02(t *testing.T) {
	nStr1 := "579"
	maxPrecision := uint64(0)
	divisor := int64(0)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err:=ia.DivideByInt64(divisor, maxPrecision, true )

	if err == nil {
		t.Errorf("Expected Divide By Zero Error. No Error Presented. Error = %v", err)
	}

}

func TestIntAry_DivideByInt64_03(t *testing.T) {
	nStr1 := "4"
	expected := "2"
	maxPrecision := uint64(0)
	divisor := int64(2)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.DivideByInt64(divisor, maxPrecision, true )

	if err != nil {
		t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
	}

	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v'. Instead, ia.Numstr= '%v'", expected, ia.NumStr)
	}

	if ia.Precision != int(maxPrecision) {
		t.Errorf("Expected ia.Precision= '%v'. Instead, ia.Precision= '%v'.",maxPrecision, ia.Precision)
	}

}

func TestIntAry_DivideByInt64_04(t *testing.T) {
	nStr1 := "476"
	expected := "-14"
	maxPrecision := uint64(10)
	ePrecision := 0
	eSignVal := -1
	divisor := int64(-34)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.DivideByInt64(divisor, maxPrecision, true )

	if err != nil {
		t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
	}

	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v'. Instead, ia.Numstr= '%v'", expected, ia.NumStr)
	}

	if ia.Precision != ePrecision {
		t.Errorf("Expected ia.Precision= '%v'. Instead, ia.Precision= '%v'.",ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v'. Instead, ia.Precision= '%v'.",eSignVal, ia.SignVal)
	}

}

func TestIntAry_DivideByInt64_05(t *testing.T) {
	nStr1 := "-476"
	expected := "-14"
	maxPrecision := uint64(10)
	ePrecision := 0
	eSignVal := -1
	divisor := int64(34)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.DivideByInt64(divisor, maxPrecision, true )

	if err != nil {
		t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
	}

	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v'. Instead, ia.Numstr= '%v'", expected, ia.NumStr)
	}

	if ia.Precision != ePrecision {
		t.Errorf("Expected ia.Precision= '%v'. Instead, ia.Precision= '%v'.",ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v'. Instead, ia.Precision= '%v'.",eSignVal, ia.SignVal)
	}

}

func TestIntAry_DivideByInt64_06(t *testing.T) {
	nStr1 := "-476"
	expected := "14"
	maxPrecision := uint64(10)
	ePrecision := 0
	eSignVal := 1
	divisor := int64(-34)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.DivideByInt64(divisor, maxPrecision, true )

	if err != nil {
		t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
	}

	if expected != ia.NumStr {
		t.Errorf("Expected ia.NumStr= '%v'. Instead, ia.Numstr= '%v'", expected, ia.NumStr)
	}

	if ia.Precision != ePrecision {
		t.Errorf("Expected ia.Precision= '%v'. Instead, ia.Precision= '%v'.",ePrecision, ia.Precision)
	}

	if eSignVal != ia.SignVal {
		t.Errorf("Expected ia.SignVal= '%v'. Instead, ia.Precision= '%v'.",eSignVal, ia.SignVal)
	}

}

func TestIntAry_DivideByTwo_01(t *testing.T) {

	nStr1 := "1959"
	expected := "979.5"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.DivideByTwo(true)

	if expected != ia.NumStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", expected, ia.NumStr)
	}

	if precision != ia.Precision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", precision, ia.Precision)
	}

}

func TestIntAry_DivideByTwo_02(t *testing.T) {

	nStr1 := "1"
	expected := "0.5"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.DivideByTwo(true)

	if expected != ia.NumStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", expected, ia.NumStr)
	}

	if precision != ia.Precision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", precision, ia.Precision)
	}

}

func TestIntAry_DivideByTwo_03(t *testing.T) {

	nStr1 := "0"
	expected := "0"
	precision := 0

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.DivideByTwo(true)

	if expected != ia.NumStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", expected, ia.NumStr)
	}

	if precision != ia.Precision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", precision, ia.Precision)
	}

}

func TestIntAry_DivideByTwo_04(t *testing.T) {

	nStr1 := "-2959"
	expected := "-1479.5"
	precision := 1
	signVal := -1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.DivideByTwo(true)

	if expected != ia.NumStr {
		t.Errorf("Error: Expected NumStr= '%v'. Instead received NumStr= '%v'", expected, ia.NumStr)
	}

	if precision != ia.Precision {
		t.Errorf("Error: Expected Precision= '%v'. Instead received Precision= '%v'", precision, ia.Precision)
	}

	if signVal != ia.SignVal {
		t.Errorf("Error: Expected Sign Value= '%v'. Instead, Sign Value= '%v'. ", signVal, ia.SignVal)
	}

}

func TestIntAry_DivideByTenToPower_01(t *testing.T) {

	nStr := "457.3"
	power := uint(1)
	eNumStr := "45.73"
	nRunes := []rune("4573")
	eIAry := []int{4, 5, 7, 3}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 2
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power, false)
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

func TestIntAry_DivideByTenToPower_02(t *testing.T) {

	nStr := "457.3"
	power := uint(3)
	eNumStr := "0.4573"
	nRunes := []rune("04573")
	eIAry := []int{0, 4, 5, 7, 3}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 4
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power, false)
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

func TestIntAry_DivideByTenToPower_03(t *testing.T) {

	nStr := "457.3"
	power := uint(7)
	eNumStr := "0.00004573"
	nRunes := []rune("000004573")
	eIAry := []int{0, 0, 0, 0, 0, 4, 5, 7, 3}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 8
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power, false)
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

func TestIntAry_DivideByTenToPower_04(t *testing.T) {

	nStr := "-457.3"
	power := uint(7)
	eNumStr := "-0.00004573"
	nRunes := []rune("000004573")
	eIAry := []int{0, 0, 0, 0, 0, 4, 5, 7, 3}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 8
	eSignVal := -1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power, false)
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

func TestIntAry_DivideByTenToPower_05(t *testing.T) {

	nStr := "0"
	power := uint(2)
	eNumStr := "0.00"
	nRunes := []rune("000")
	eIAry := []int{0, 0, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 2
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power, true)

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

func TestIntAry_DivideByTenToPower_06(t *testing.T) {

	nStr := "-4573"
	power := uint(1)
	eNumStr := "-457.3"
	nRunes := []rune("4573")
	eIAry := []int{4, 5, 7, 3}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := -1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power, true)

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

func TestIntAry_DivideThisBy_01(t *testing.T) {
	dividend := "56234369384300"
	divisor :=  "24"
	eQuotient := "2343098724345.833333333333333333333"
	eSignVal := 1
	maxPrecision := 21
	ePrecision := 21

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.NumStr {
		t.Errorf("Expected quotient.NumStr= '%v' .  Instead, quotient.NumStr= '%v'  .", eQuotient, quotient.NumStr)
	}

	if ePrecision != quotient.Precision {
		t.Errorf("Expected quotient.Precision= '%v' .  Instead, quotient.Precision= '%v'  .", ePrecision, quotient.Precision)
	}

	if eSignVal != quotient.SignVal {
		t.Errorf("Error - Expected smop.Quotient.SignVal= '%v'. Instead, smop.Quotient.SignVal= '%v' .", eSignVal, quotient.SignVal)
	}

}

func TestIntAry_DivideThisBy_02(t *testing.T) {
	dividend := "48"
	divisor :=  "24"
	eSignVal := 1
	eQuotient := "2"
	maxPrecision := 21
	ePrecision := 0

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.NumStr {
		t.Errorf("Expected quotient.NumStr= '%v' .  Instead, quotient.NumStr= '%v'  .", eQuotient, quotient.NumStr)
	}

	if ePrecision != quotient.Precision {
		t.Errorf("Expected quotient.Precision= '%v' .  Instead, quotient.Precision= '%v'  .", ePrecision, quotient.Precision)
	}

	if eSignVal != quotient.SignVal {
		t.Errorf("Error - Expected smop.Quotient.SignVal= '%v'. Instead, smop.Quotient.SignVal= '%v' .", eSignVal, quotient.SignVal)
	}

}

func TestIntAry_DivideThisBy_03(t *testing.T) {
	dividend := "24"
	divisor :=  "24"
	eQuotient := "1"
	eSignVal := 1
	maxPrecision := 21
	ePrecision := 0

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.NumStr {
		t.Errorf("Expected quotient.NumStr= '%v' .  Instead, quotient.NumStr= '%v'  .", eQuotient, quotient.NumStr)
	}

	if ePrecision != quotient.Precision {
		t.Errorf("Expected quotient.Precision= '%v' .  Instead, quotient.Precision= '%v'  .", ePrecision, quotient.Precision)
	}

	if eSignVal != quotient.SignVal {
		t.Errorf("Error - Expected smop.Quotient.SignVal= '%v'. Instead, smop.Quotient.SignVal= '%v' .", eSignVal, quotient.SignVal)
	}

}

func TestIntAry_DivideThisBy_04(t *testing.T) {
	dividend := "0.05"
	divisor :=  "24"
	eQuotient := "0.00208333333333333333333333333333"
	eSignVal := 1
	maxPrecision := 32
	ePrecision := 32

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.NumStr {
		t.Errorf("Expected quotient.NumStr= '%v' .  Instead, quotient.NumStr= '%v'  .", eQuotient, quotient.NumStr)
	}

	if ePrecision != quotient.Precision {
		t.Errorf("Expected quotient.Precision= '%v' .  Instead, quotient.Precision= '%v'  .", ePrecision, quotient.Precision)
	}

	if eSignVal != quotient.SignVal {
		t.Errorf("Error - Expected smop.Quotient.SignVal= '%v'. Instead, smop.Quotient.SignVal= '%v' .", eSignVal, quotient.SignVal)
	}

}

func TestIntAry_DivideThisBy_05(t *testing.T) {
	dividend := "0"
	divisor :=  "24"
	eQuotient := "0"
	eSignVal := 1
	maxPrecision := 7
	ePrecision := 0

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.NumStr {
		t.Errorf("Expected quotient.NumStr= '%v' .  Instead, quotient.NumStr= '%v'  .", eQuotient, quotient.NumStr)
	}

	if ePrecision != quotient.Precision {
		t.Errorf("Expected quotient.Precision= '%v' .  Instead, quotient.Precision= '%v'  .", ePrecision, quotient.Precision)
	}

	if eSignVal != quotient.SignVal {
		t.Errorf("Error - Expected smop.Quotient.SignVal= '%v'. Instead, smop.Quotient.SignVal= '%v' .", eSignVal, quotient.SignVal)
	}

}

func TestIntAry_DivideThisBy_06(t *testing.T) {
	dividend := "48"
	divisor :=  "0"

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	_, err := ia1.DivideThisBy(&ia2, 15)

	if err == nil {
		t.Error("Expected an error from Divideby Zero. No Error Received!")
	}

}

func TestIntAry_DivideThisBy_07(t *testing.T) {
	dividend := "-9360"
	divisor :=  "24.48"
	eQuotient := "-382.35294117647058823529411764706"
	eSignVal := -1
	maxPrecision := 29
	ePrecision := 29

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.NumStr {
		t.Errorf("Expected quotient.NumStr= '%v' .  Instead, quotient.NumStr= '%v'  .", eQuotient, quotient.NumStr)
	}

	if ePrecision != quotient.Precision {
		t.Errorf("Expected quotient.Precision= '%v' .  Instead, quotient.Precision= '%v'  .", ePrecision, quotient.Precision)
	}

	if eSignVal != quotient.SignVal {
		t.Errorf("Error - Expected smop.Quotient.SignVal= '%v'. Instead, smop.Quotient.SignVal= '%v' .", eSignVal, quotient.SignVal)
	}
}

func TestIntAry_DivideThisBy_08(t *testing.T) {
	dividend := "-9360"
	divisor :=  "-24.48"
	eQuotient := "382.35294117647058823529411764706"
	eSignVal := 1
	maxPrecision := 29
	ePrecision := 29

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.NumStr {
		t.Errorf("Expected quotient.NumStr= '%v' .  Instead, quotient.NumStr= '%v'  .", eQuotient, quotient.NumStr)
	}

	if ePrecision != quotient.Precision {
		t.Errorf("Expected quotient.Precision= '%v' .  Instead, quotient.Precision= '%v'  .", ePrecision, quotient.Precision)
	}

	if eSignVal != quotient.SignVal {
		t.Errorf("Error - Expected smop.Quotient.SignVal= '%v'. Instead, smop.Quotient.SignVal= '%v' .", eSignVal, quotient.SignVal)
	}
}

func TestIntAry_DivideThisBy_09(t *testing.T) {
	dividend := "9360"
	divisor :=  "-24.48"
	eQuotient := "-382.35294117647058823529411764706"
	eSignVal := -1
	maxPrecision := 29
	ePrecision := 29

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.NumStr {
		t.Errorf("Expected quotient.NumStr= '%v' .  Instead, quotient.NumStr= '%v'  .", eQuotient, quotient.NumStr)
	}

	if ePrecision != quotient.Precision {
		t.Errorf("Expected quotient.Precision= '%v' .  Instead, quotient.Precision= '%v'  .", ePrecision, quotient.Precision)
	}

	if eSignVal != quotient.SignVal {
		t.Errorf("Error - Expected smop.Quotient.SignVal= '%v'. Instead, smop.Quotient.SignVal= '%v' .", eSignVal, quotient.SignVal)
	}
}

func TestIntAry_DivideThisBy_10(t *testing.T) {
	dividend := "-19260.549"
	divisor :=  "-246.483"
	eQuotient := "78.141490488187826340964691276883"
	eSignVal := 1
	maxPrecision := 30
	ePrecision := 30

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.NumStr {
		t.Errorf("Expected quotient.NumStr= '%v' .  Instead, quotient.NumStr= '%v'  .", eQuotient, quotient.NumStr)
	}

	if ePrecision != quotient.Precision {
		t.Errorf("Expected quotient.Precision= '%v' .  Instead, quotient.Precision= '%v'  .", ePrecision, quotient.Precision)
	}

	if eSignVal != quotient.SignVal {
		t.Errorf("Error - Expected smop.Quotient.SignVal= '%v'. Instead, smop.Quotient.SignVal= '%v' .", eSignVal, quotient.SignVal)
	}
}

func TestIntAry_Floor_01(t *testing.T) {
	nStr1 := "99.925"
	expected := "99.000"
	precision := 3

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_Floor_02(t *testing.T) {
	nStr1 := "-99.925"
	expected := "-100.000"
	precision := 3

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_Floor_03(t *testing.T) {
	nStr1 := "0.925"
	expected := "0.000"
	precision := 3

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_Floor_04(t *testing.T) {
	nStr1 := "2.0"
	expected := "2.0"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_Floor_05(t *testing.T) {
	nStr1 := "-2.7"
	expected := "-3.0"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_Floor_06(t *testing.T) {
	nStr1 := "-2"
	expected := "-2"
	precision := 0

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_Floor_07(t *testing.T) {
	nStr1 := "2.9"
	expected := "2.0"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

}

func TestIntAry_GetFractionalDigits_01(t *testing.T) {

	nStr1 := "2.7894"
	expected := "0.7894"
	precision := 4
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetFractionalDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	if false != iAry2.IsZeroValue {
		t.Errorf("Error. Expected IsZeroValue= '%v'. Instead, got IsZeroValue='%v'\n", false, iAry2.IsZeroValue)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

	if signVal != iAry2.SignVal {
		t.Errorf("Error. Expected SignVal= '%v'. Instead, got SignVal='%v'\n", signVal, iAry2.SignVal)
	}

}

func TestIntAry_GetFractionalDigits_02(t *testing.T) {

	nStr1 := "2"
	expected := "0"
	precision := 0
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetFractionalDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	if true != iAry2.IsZeroValue {
		t.Errorf("Error. Expected IsZeroValue= '%v'. Instead, got IsZeroValue='%v'\n", true, iAry2.IsZeroValue)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

	if signVal != iAry2.SignVal {
		t.Errorf("Error. Expected SignVal= '%v'. Instead, got SignVal='%v'\n", signVal, iAry2.SignVal)
	}

}

func TestIntAry_GetFractionalDigits_03(t *testing.T) {

	nStr1 := "2.00"
	expected := "0.00"
	precision := 2
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetFractionalDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	if true != iAry2.IsZeroValue {
		t.Errorf("Error. Expected IsZeroValue= '%v'. Instead, got IsZeroValue='%v'\n", true, iAry2.IsZeroValue)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

	if signVal != iAry2.SignVal {
		t.Errorf("Error. Expected SignVal= '%v'. Instead, got SignVal='%v'\n", signVal, iAry2.SignVal)
	}

}

func TestIntAry_GetFractionalDigits_04(t *testing.T) {

	nStr1 := "-2.978562154907"
	expected := "0.978562154907"
	precision := 12
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetFractionalDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	if false != iAry2.IsZeroValue {
		t.Errorf("Error. Expected IsZeroValue= '%v'. Instead, got IsZeroValue='%v'\n", false, iAry2.IsZeroValue)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

	if signVal != iAry2.SignVal {
		t.Errorf("Error. Expected SignVal= '%v'. Instead, got SignVal='%v'\n", signVal, iAry2.SignVal)
	}

}

func TestIntAry_GetIntegerDigits_01(t *testing.T) {

	nStr1 := "997562.4692"
	expected := "997562"
	precision := 0
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetIntegerDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	if false != iAry2.IsZeroValue {
		t.Errorf("Error. Expected IsZeroValue= '%v'. Instead, got IsZeroValue='%v'\n", false, iAry2.IsZeroValue)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

	if signVal != iAry2.SignVal {
		t.Errorf("Error. Expected SignVal= '%v'. Instead, got SignVal='%v'\n", signVal, iAry2.SignVal)
	}

}

func TestIntAry_GetIntegerDigits_02(t *testing.T) {

	nStr1 := "0.4692"
	expected := "0"
	precision := 0
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetIntegerDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	if true != iAry2.IsZeroValue {
		t.Errorf("Error. Expected IsZeroValue= '%v'. Instead, got IsZeroValue='%v'\n", true, iAry2.IsZeroValue)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

	if signVal != iAry2.SignVal {
		t.Errorf("Error. Expected SignVal= '%v'. Instead, got SignVal='%v'\n", signVal, iAry2.SignVal)
	}

}

func TestIntAry_GetIntegerDigits_03(t *testing.T) {

	nStr1 := "-987.4692"
	expected := "-987"
	precision := 0
	signVal := -1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetIntegerDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	if false != iAry2.IsZeroValue {
		t.Errorf("Error. Expected IsZeroValue= '%v'. Instead, got IsZeroValue='%v'\n", false, iAry2.IsZeroValue)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'\n", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.Precision)
	}

	if signVal != iAry2.SignVal {
		t.Errorf("Error. Expected SignVal= '%v'. Instead, got SignVal='%v'\n", signVal, iAry2.SignVal)
	}

}

func TestIntAry_GetIntegerDigits_04(t *testing.T) {

	nStr1 := "-0.4692"
	expected := "0"
	precision := 0
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetIntegerDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	if true != iAry2.IsZeroValue {
		t.Errorf("Error. Expected IsZeroValue= '%v'. Instead, got IsZeroValue='%v'", true, iAry2.IsZeroValue)
	}

	s := iAry2.NumStr
	if expected != s {
		t.Errorf("Error. Expected NumStr= '%v'. Instead, got NumStr='%v'", expected, s)
	}

	if iAry2.Precision != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'", precision, iAry2.Precision)
	}

	if signVal != iAry2.SignVal {
		t.Errorf("Error. Expected SignVal= '%v'. Instead, got SignVal='%v'", signVal, iAry2.SignVal)
	}

}

func TestIntAry_IncrementIntegerOne_01(t *testing.T) {
	expected := "100.123"
	nStr1 := "-100.123"
	cycles := 200
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.IncrementIntegerOne()
	}

	ia.ConvertIntAryToNumStr()

	if expected != ia.NumStr {
		t.Errorf("Error - Expected NumStr= '%v'. Instead, NumStr= '%v'", expected, ia.NumStr)
	}

}

func TestIntAry_IncrementIntegerOne_02(t *testing.T) {
	nStr1 := "-2000"
	expected := "2000"
	cycles := 4000
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.IncrementIntegerOne()
	}

	ia.ConvertIntAryToNumStr()

	if expected != ia.NumStr {
		t.Errorf("Error - Expected NumStr= '%v'. Instead, NumStr= '%v'", expected, ia.NumStr)
	}

}

func TestIntAry_IncrementIntegerOne_03(t *testing.T) {
	nStr1 := "-2000.123"
	expected := "2000.123"
	cycles := 4000
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.IncrementIntegerOne()
	}

	ia.ConvertIntAryToNumStr()

	if expected != ia.NumStr {
		t.Errorf("Error - Expected NumStr= '%v'. Instead, NumStr= '%v'", expected, ia.NumStr)
	}

}

func TestIntAry_IncrementIntegerOne_04(t *testing.T) {
	nStr1 := "0"
	expected := "40"
	cycles := 40
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.IncrementIntegerOne()
	}

	ia.ConvertIntAryToNumStr()

	if expected != ia.NumStr {
		t.Errorf("Error - Expected NumStr= '%v'. Instead, NumStr= '%v'", expected, ia.NumStr)
	}

}

