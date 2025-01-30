package common

import (
	"errors"
	"fmt"
)

type BackUpIntAry struct {
	NumStr                 string
	NumRunes               []rune
	NumRunesLen            int
	IntAry                 []int
	IntAryLen              int
	IntegerLen             int
	SignificantIntegerLen  int
	SignificantFractionLen int
	FirstDigitIdx          int
	LastDigitIdx           int
	IsZeroValue            bool
	IsIntegerZeroValue     bool
	Precision              int
	SignVal                int
	DecimalSeparator       rune
}

func (iBa BackUpIntAry) New() BackUpIntAry {
	iAry := BackUpIntAry{}

	iAry.NumStr = ""
	iAry.NumRunes = []rune{}
	iAry.NumRunesLen = 0
	iAry.IntAry = []int{}
	iAry.IntAryLen = 0
	iAry.FirstDigitIdx = -1
	iAry.LastDigitIdx = -1
	iAry.IntegerLen = 0
	iAry.SignificantIntegerLen = 0
	iAry.SignificantFractionLen = 0
	iAry.IsZeroValue = true
	iAry.IsIntegerZeroValue = true
	iAry.Precision = 0
	iAry.SignVal = 1
	iAry.DecimalSeparator = '.'

	return iAry
}

func (iBa *BackUpIntAry) Empty() {
	iBa.NumStr = ""
	iBa.NumRunes = []rune{}
	iBa.NumRunesLen = 0
	iBa.IntAry = []int{}
	iBa.IntAryLen = 0
	iBa.IntegerLen = 0
	iBa.SignificantIntegerLen = 0
	iBa.SignificantFractionLen = 0
	iBa.FirstDigitIdx = -1
	iBa.LastDigitIdx = -1
	iBa.IsZeroValue = true
	iBa.IsIntegerZeroValue = true
	iBa.Precision = 0
	iBa.SignVal = 1
	if iBa.DecimalSeparator == 0 {
		iBa.DecimalSeparator = '.'
	}
}

func (iBa *BackUpIntAry) CopyIn(iBa2 *BackUpIntAry) {
	iBa2.SetInternalFlags()
	iBa.Empty()
	iBa.NumStr = iBa2.NumStr
	iBa.NumRunes = make([]rune, iBa2.NumRunesLen)
	for i := 0; i < iBa2.NumRunesLen; i++ {
		iBa.NumRunes[i] = iBa2.NumRunes[i]
	}

	iBa.NumRunesLen = iBa2.NumRunesLen

	iBa.IntAry = make([]int, iBa2.IntAryLen)
	for i := 0; i < iBa2.IntAryLen; i++ {
		iBa.IntAry[i] = iBa2.IntAry[i]
	}

	iBa.IntAryLen = iBa2.IntAryLen
	iBa.IntegerLen = iBa2.IntegerLen
	iBa.SignificantIntegerLen = iBa2.SignificantIntegerLen
	iBa.SignificantFractionLen = iBa2.SignificantFractionLen
	iBa.FirstDigitIdx = iBa2.FirstDigitIdx
	iBa.LastDigitIdx = iBa2.LastDigitIdx
	iBa.IsZeroValue = iBa2.IsZeroValue
	iBa.IsIntegerZeroValue = iBa2.IsIntegerZeroValue
	iBa.Precision = iBa2.Precision
	iBa.SignVal = iBa2.SignVal
	iBa.DecimalSeparator = iBa2.DecimalSeparator
	if iBa.DecimalSeparator == 0 {
		iBa.DecimalSeparator = '.'
	}

}

func (iBa *BackUpIntAry) CopyOut() BackUpIntAry {
	iBa.SetInternalFlags()
	iAry2 := BackUpIntAry{}.New()

	iAry2.NumStr = iBa.NumStr
	iAry2.NumRunes = make([]rune, iBa.NumRunesLen)
	for i := 0; i < iBa.NumRunesLen; i++ {
		iAry2.NumRunes[i] = iBa.NumRunes[i]
	}

	iAry2.NumRunesLen = iBa.NumRunesLen

	iAry2.IntAry = make([]int, iBa.IntAryLen)

	for i := 0; i < iBa.IntAryLen; i++ {
		iAry2.IntAry[i] = iBa.IntAry[i]
	}

	iAry2.IntAryLen = iBa.IntAryLen
	iAry2.IntegerLen = iBa.IntegerLen
	iAry2.SignificantIntegerLen = iBa.SignificantIntegerLen
	iAry2.SignificantFractionLen = iBa.SignificantFractionLen
	iAry2.FirstDigitIdx = iBa.FirstDigitIdx
	iAry2.LastDigitIdx = iBa.LastDigitIdx
	iAry2.IsZeroValue = iBa.IsZeroValue
	iAry2.IsIntegerZeroValue = iBa.IsIntegerZeroValue
	iAry2.Precision = iBa.Precision
	iAry2.SignVal = iBa.SignVal
	iAry2.DecimalSeparator = iBa.DecimalSeparator
	if iAry2.DecimalSeparator == 0 {
		iAry2.DecimalSeparator = '.'
	}

	return iAry2
}

func (iBa *BackUpIntAry) SetInternalFlags() {
	
	iBa.IntAryLen = len(iBa.IntAry)

	if iBa.IntAryLen == iBa.Precision {
		iBa.IntAry = append([]int{0}, iBa.IntAry...)
		iBa.IntAryLen++
	}

	if iBa.IntAryLen < iBa.Precision {

		deltaZeros := iBa.Precision - iBa.IntAryLen + 1
		zeroAry := make([]int,deltaZeros)
		iBa.IntAry = append(zeroAry, iBa.IntAry...)
		iBa.IntAryLen += deltaZeros
	}


	iBa.FirstDigitIdx = -1
	iBa.LastDigitIdx = -1

	lastIntIdx := iBa.IntAryLen - iBa.Precision - 1
	iBa.IsZeroValue = true
	iBa.IsIntegerZeroValue = true
	iBa.IntegerLen = iBa.IntAryLen - iBa.Precision

	for i := 0; i < iBa.IntAryLen; i++ {
		if iBa.IntAry[i] > 0 {
			iBa.IsZeroValue = false

			if i < iBa.IntegerLen {
				iBa.IsIntegerZeroValue = false
			}
		}
		// At minimum, there should be a single
		// leading zero before the decimal point.
		// Example 0.000.
		if i == lastIntIdx && iBa.IntAry[i] == 0 {

			if iBa.FirstDigitIdx == -1 {
				iBa.FirstDigitIdx = i
			}

		}

		if iBa.IntAry[i] > 0 {

			if iBa.FirstDigitIdx == -1 {
				iBa.FirstDigitIdx = i
			}

			iBa.LastDigitIdx = i
		}

	}

	iBa.SignificantIntegerLen = iBa.IntAryLen - iBa.Precision - iBa.FirstDigitIdx

	if iBa.LastDigitIdx >= iBa.IntegerLen {
		iBa.SignificantFractionLen = iBa.Precision - (iBa.LastDigitIdx - iBa.IntegerLen + 1)
	} else {
		iBa.SignificantFractionLen = 0
	}

}

// IntAry - Used to perform string
// based arithmetic.
//
// Dependencies: None
//
type IntAry struct {
	NumStr                 string
	NumRunes               []rune
	NumRunesLen            int
	IntAry                 []int
	IntAryLen              int
	IntegerLen             int
	SignificantIntegerLen  int
	SignificantFractionLen int
	FirstDigitIdx          int
	LastDigitIdx           int
	IsZeroValue            bool
	IsIntegerZeroValue     bool
	Precision              int
	SignVal                int
	DecimalSeparator       rune
	BackUp                 BackUpIntAry
}

func (ia IntAry) New() IntAry {
	iAry := IntAry{}
	iAry.NumStr = ""
	iAry.NumRunes = []rune{}
	iAry.NumRunesLen = 0
	iAry.IntAry = []int{}
	iAry.IntAryLen = 0
	iAry.IntegerLen = 0
	iAry.SignificantIntegerLen = 0
	iAry.SignificantFractionLen = 0
	iAry.FirstDigitIdx = -1
	iAry.LastDigitIdx = -1
	iAry.IsZeroValue = true
	iAry.IsIntegerZeroValue = true
	iAry.Precision = 0
	iAry.SignVal = 1
	iAry.DecimalSeparator = '.'
	iAry.BackUp = BackUpIntAry{}.New()

	return iAry
}

func (ia *IntAry) Empty() {
	ia.NumStr = ""
	ia.NumRunes = []rune{}
	ia.NumRunesLen = 0
	ia.IntAry = []int{}
	ia.IntAryLen = 0
	ia.IntegerLen = 0
	ia.SignificantIntegerLen = 0
	ia.SignificantFractionLen = 0
	ia.FirstDigitIdx = -1
	ia.LastDigitIdx = -1
	ia.IsZeroValue = true
	ia.IsIntegerZeroValue = true
	ia.Precision = 0
	ia.SignVal = 1
	if ia.DecimalSeparator == 0 {
		ia.DecimalSeparator = '.'
	}

}

func (ia *IntAry) EmptyBackUp() {
	ia.BackUp = BackUpIntAry{}.New()
}

// AddToThis - Adds the value of IntAry parameter ia2 to the value
// of the current IntAry object.
//
// Paramters:
//
// ia2 *IntAry - Incoming IntAry object whose value will be subtracted
// 								from this current IntAry value.
//
// convertToNumStr - boolean value determines whether the current IntAry
//                   object will convert the IntAry value to a number string.
//                   Set this parameter to 'false' if this method is called
//                   multiple times in order to improve performance.

func (ia *IntAry) AddToThis(ia2 *IntAry, convertToNumStr bool) error {

	ia.SetEqualArrayLengths(ia2)

	if ia2.IsZeroValue {
		return nil
	}

	compare := ia.CompareAbsoluteValues(ia2)

	newSignVal := ia.SignVal
	doAdd := true
	isZeroResult := false
	doReverseNums := false

	if compare == 1 {
		// compare == + 1
		// Absolute Value: N1 > N2

		if ia.SignVal == 1 && ia2.SignVal == 1 {
			doAdd = true
			newSignVal = 1
		} else if ia.SignVal == -1 && ia2.SignVal == 1 {
			doAdd = false
			newSignVal = -1
		} else if ia.SignVal == -1 && ia2.SignVal == -1 {
			doAdd = true
			newSignVal = -1
		} else {
			// Must Be ia.SignVal == 1 && ia2.SignVal == -1
			doAdd = false
			newSignVal = 1
		}

	} else if compare == -1 {
		// Absolute Values: N2 > N1
		if ia.SignVal == 1 && ia2.SignVal == 1 {
			doAdd = true
			newSignVal = 1
		} else if ia.SignVal == -1 && ia2.SignVal == 1 {
			doAdd = false
			doReverseNums = true
			newSignVal = 1
		} else if ia.SignVal == -1 && ia2.SignVal == -1 {
			doAdd = true
			newSignVal = -1
		} else {
			// Must Be ia.SignVal == 1 && ia2.SignVal == -1
			doAdd = false
			doReverseNums = true
			newSignVal = -1
		}

	} else {
		// Must be compare == 0
		// Absolute Values: N1==N2
		if ia.SignVal == 1 && ia2.SignVal == 1 {
			doAdd = true
			newSignVal = 1
		} else if ia.SignVal == -1 && ia2.SignVal == 1 {
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		} else if ia.SignVal == -1 && ia2.SignVal == -1 {
			doAdd = true
			newSignVal = -1
		} else {
			// Must Be ia.SignVal == 1 && ia2.SignVal == -1
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		}

	}

	return ia.addToSubtractFromThis(ia2, newSignVal, doAdd, isZeroResult, doReverseNums, convertToNumStr)
}

func (ia *IntAry) addToSubtractFromThis(ia2 *IntAry, newSignVal int, doAdd bool, isZeroResult bool, doReverseNums bool, convertToNumStr bool) error {

	if isZeroResult {
		ia.SetIntAryToZero(ia.Precision)
		return nil
	}

	ia.SignVal = newSignVal

	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0

	for j := ia.IntAryLen - 1; j >= 0; j-- {

		if doReverseNums {

			n2 = ia.IntAry[j]
			n1 = ia2.IntAry[j]

		} else {
			n1 = ia.IntAry[j]
			n2 = ia2.IntAry[j]

		}

		if doAdd {
			// doAdd == true
			// Do Addition

			n3 = n1 + n2 + carry

			if n3 > 9 {
				n3 = n1 + n2 + carry - 10
				carry = 1

			} else {
				carry = 0
			}

		} else {
			// doAdd == false
			// Do Subtraction
			n3 = n1 - n2 - carry

			if n3 < 0 {
				n3 = n1 + 10 - n2 - carry
				carry = 1
			} else {
				carry = 0
			}
		}

		ia.IntAry[j] = n3

	}

	if carry > 0 {
		ia.IntAry = append([]int{1}, ia.IntAry...)
		ia.IntAryLen++
	}

	if ia.IntAry[0] == 0 {
		ia.SetSignificantDigitIdxs()
		ia.IntAry = ia.IntAry[ia.FirstDigitIdx:]
	}

	if convertToNumStr {
		ia.ConvertIntAryToNumStr()
	}

	return nil
}

// AddArrayLengthLeft - Adds leading zeros to the IntAry
func (ia *IntAry) AddArrayLengthLeft(addLen int) {

	ia.SetIntAryLength()

	newLen := addLen + ia.IntAryLen
	t := make([]int, newLen)

	for i := 0; i < newLen; i++ {

		if i < addLen {
			t[i] = 0
		} else {
			t[i] = ia.IntAry[i-addLen]
		}

	}

	ia.IntAry = t
	ia.SetIntAryLength()
}

// AddArrayLengthRight - Adds trailing zeros
// to the right of the current IntAry.
func (ia *IntAry) AddArrayLengthRight(addLen int) {
	ia.SetIntAryLength()

	for i := 0; i < addLen; i++ {
		ia.IntAry = append(ia.IntAry, 0)
	}

	ia.SetIntAryLength()
}

func (ia *IntAry) Ceiling() (IntAry, error) {

	err := ia.IsIntAryValid("Ceiling() - ")

	if err != nil {
		return IntAry{}, err
	}

	iAry2 := IntAry{}.New()

	intLen := ia.IntAryLen - ia.Precision

	intIdx := intLen - 1

	hasFracDigits, err := ia.HasFractionalDigits()

	if err != nil {
		return iAry2, err
	}

	if !hasFracDigits {
		iAry2 = ia.CopyOut()
		return iAry2, nil
	}

	if ia.SignVal < 0 && hasFracDigits {

		t := make([]int, ia.IntAryLen)

		for i := 0; i < intLen; i++ {
			t[i] = ia.IntAry[i]
		}

		iAry2.IntAry = t[0:]
		iAry2.IntAryLen = ia.IntAryLen
		iAry2.Precision = ia.Precision
		iAry2.SignVal = ia.SignVal
		iAry2.ConvertIntAryToNumStr()
		return iAry2, nil
	}

	t := make([]int, ia.IntAryLen+1)

	n1 := 0
	n2 := 0
	carry := 0
	adjFac := 1 * ia.SignVal
	for i := intIdx; i >= 0; i-- {

		n1 = ia.IntAry[i]

		if i == intIdx {
			if n1+adjFac < 0 {
				n2 = 10 + n1 + adjFac
				carry = -1
			} else if n1+adjFac > 9 {
				n2 = n1 + adjFac - 10
				carry = 1
			} else {
				n2 = n1 + adjFac
				carry = 0
			}

		} else {

			if n1+carry < 0 {
				n2 = 10 + n1
				carry = -1
			} else if n1+carry > 9 {
				n2 = n1 - 10
				carry = 1
			} else {
				n2 = n1 + carry
				carry = 0
			}

		}

		t[i+1] = n2

	}

	if carry != 0 {
		t[0] = carry
		iAry2.IntAry = t[0 : ia.IntAryLen+1]

	} else {

		iAry2.IntAry = t[1 : ia.IntAryLen+1]
	}

	iAry2.IntAryLen = len(iAry2.IntAry)
	iAry2.Precision = ia.Precision
	iAry2.SignVal = ia.SignVal
	iAry2.ConvertIntAryToNumStr()
	return iAry2, nil
}

func (ia *IntAry) CompareSignedValues(iAry2 *IntAry) int {

	iCompare := ia.CompareAbsoluteValues(iAry2)

	if ia.IsZeroValue && iAry2.IsZeroValue {
		return 0
	}

	if ia.SignVal != iAry2.SignVal {

		if ia.SignVal == 1 {
			return 1
		} else {
			return -1
		}
	}

	// Must be ia.SignVal == iAry2.SignVal

	if ia.SignVal == 1 {
		return iCompare
	}

	// Must be ia.SignVal && iAry2.SignVal == -1

	return iCompare * -1

}

func (ia *IntAry) CompareAbsoluteValues(iAry2 *IntAry) int {

	ia.SetIntAryLength()
	iAry2.SetIntAryLength()

	iAry2.SetIsZeroValue()
	ia.SetIsZeroValue()

	if ia.IsZeroValue && iAry2.IsZeroValue {
		return 0
	}

	iaIntLen := ia.IntAryLen - ia.Precision
	iAry2IntLen := iAry2.IntAryLen - iAry2.Precision

	// Integer Lengths are Equal
	if iaIntLen == iAry2IntLen {
		for i := 0; i < iaIntLen; i++ {
			if ia.IntAry[i] > iAry2.IntAry[i] {
				return 1
			}

			if iAry2.IntAry[i] > ia.IntAry[i] {
				return -1
			}
		}
	}

	deltaStartIdx := 0

	// ia Integer Length is Greater than IAry2 Integer Length
	if iaIntLen > iAry2IntLen {
		deltaStartIdx = iaIntLen - iAry2IntLen

		for j := 0; j < iaIntLen; j++ {

			if j < deltaStartIdx {

				if ia.IntAry[j] > 0 {
					return 1
				}

			} else {
				// i must be >= deltaStartIdx

				if ia.IntAry[j] > iAry2.IntAry[j-deltaStartIdx] {
					return 1
				}

				if iAry2.IntAry[j-deltaStartIdx] > ia.IntAry[j] {
					return -1
				}
			}
		}
	}
	/*
		fmt.Println("   IAry2IntLen = ", iAry2IntLen)
		fmt.Println("  deltaStartIdx= ", deltaStartIdx)
		fmt.Println("              k= ", k)
		fmt.Println("iAry2.IntAry[k]= ", iAry2.IntAry[k])
	*/

	// iAry2 Integer Length is Greater Than ia Integer Length
	if iAry2IntLen > iaIntLen {
		deltaStartIdx = iAry2IntLen - iaIntLen

		for k := 0; k < iAry2IntLen; k++ {

			if k < deltaStartIdx {
				if iAry2.IntAry[k] > 0 {
					return -1
				}

			} else {
				// i must be >= deltaStartIdx

				if iAry2.IntAry[k] > ia.IntAry[k-deltaStartIdx] {
					return -1
				}

				if ia.IntAry[k-deltaStartIdx] > iAry2.IntAry[k] {
					return 1
				}
			}
		}
	}

	// If precision is zero, the IntAry's are equivalent
	if ia.Precision == 0 && iAry2.Precision == 0 {
		return 0
	}

	// Integer Values are Equivalent. Now test
	// digits to the right of the decimal point.

	// Test fractional digits to right of decimal point
	iaFracIdx := iaIntLen
	iAry2FracIdx := iAry2IntLen
	// Test for case of Equal Precision
	if ia.Precision == iAry2.Precision {
		for m := 0; m < ia.Precision; m++ {

			if ia.IntAry[iaFracIdx] > iAry2.IntAry[iAry2FracIdx] {
				return 1
			}

			if iAry2.IntAry[iAry2FracIdx] > ia.IntAry[iaFracIdx] {
				return -1
			}

			iaFracIdx++
			iAry2FracIdx++
		}
	}

	iaFracIdx = iaIntLen
	iAry2FracIdx = iAry2IntLen
	// Test for case where ia Precision Greater than iAry2 Precision
	if ia.Precision > iAry2.Precision {

		for i := 0; i < ia.Precision; i++ {

			if i < iAry2.Precision {

				if ia.IntAry[iaFracIdx] > iAry2.IntAry[iAry2FracIdx] {
					return 1
				}

				if iAry2.IntAry[iAry2FracIdx] > ia.IntAry[iaFracIdx] {
					return -1
				}

				iaFracIdx++
				iAry2FracIdx++

			} else {
				if ia.IntAry[iaFracIdx] > 0 {
					return 1
				}

				iaFracIdx++
			}
		}
	}

	iaFracIdx = iaIntLen
	iAry2FracIdx = iAry2IntLen
	// Test for case where iAry2 Precision Greater than ia Precision
	if iAry2.Precision > ia.Precision {

		for i := 0; i < iAry2.Precision; i++ {

			if i < ia.Precision {

				if ia.IntAry[iaFracIdx] > iAry2.IntAry[iAry2FracIdx] {
					return 1
				}

				if iAry2.IntAry[iAry2FracIdx] > ia.IntAry[iaFracIdx] {
					return -1
				}

				iaFracIdx++
				iAry2FracIdx++

			} else {
				if iAry2.IntAry[iAry2FracIdx] > 0 {
					return -1
				}

				iAry2FracIdx++
			}
		}

	}

	// The two absolute numeric values must be equal
	return 0

}

// ConvertIntAryToNumStr - Converts the
// integer array, IntAry, to an array of
// numeric runes and then converts the
// rune array to a number string.
func (ia *IntAry) ConvertIntAryToNumStr() {
	if ia.DecimalSeparator == 0 {
		ia.DecimalSeparator = '.'
	}

	ia.ConvertIntAryToNumRunes()

	ia.NumStr = ""

	if ia.SignVal < 0 {
		ia.NumStr = "-"
	}

	intLen := ia.IntAryLen - ia.Precision

	ia.NumStr += string(ia.NumRunes[0:intLen])

	if ia.Precision > 0 {
		ia.NumStr += string(ia.DecimalSeparator)
		ia.NumStr += string(ia.NumRunes[intLen:])
	}

	ia.NumRunesLen = len(ia.NumRunes)
}

// ConvertIntAryToNumRunes - Converts the
// IntAry to an array of numeric runes.
func (ia *IntAry) ConvertIntAryToNumRunes() {
	ia.SetSignificantDigitIdxs()

	ia.NumRunes = make([]rune, ia.IntAryLen)
	ia.NumRunesLen = ia.IntAryLen
	for i := 0; i < ia.IntAryLen; i++ {
		ia.NumRunes[i] = rune(ia.IntAry[i] + 48)
	}



	return
}

// ConvertNumRunesToIntAry - Converts the Numeric
// runes array to an array of integers, IntAry.
func (ia *IntAry) ConvertNumRunesToIntAry() {
	ia.NumRunesLen = len(ia.NumRunes)
	ia.IntAry = make([]int, ia.NumRunesLen)

	for i := 0; i < ia.NumRunesLen; i++ {
		ia.IntAry[i] = int(ia.NumRunes[i]) - 48
	}

	ia.IntAryLen = ia.NumRunesLen

	return
}

func (ia *IntAry) CopyIn(iAry2 *IntAry) {
	iAry2.SetInternalFlags()
	ia.Empty()
	ia.NumStr = iAry2.NumStr
	ia.NumRunes = make([]rune, iAry2.NumRunesLen)
	for i := 0; i < iAry2.NumRunesLen; i++ {
		ia.NumRunes[i] = iAry2.NumRunes[i]
	}

	ia.NumRunesLen = iAry2.NumRunesLen

	ia.IntAry = make([]int, iAry2.IntAryLen)
	for i := 0; i < iAry2.IntAryLen; i++ {
		ia.IntAry[i] = iAry2.IntAry[i]
	}

	ia.IntAryLen = iAry2.IntAryLen
	ia.FirstDigitIdx = iAry2.FirstDigitIdx
	ia.LastDigitIdx = iAry2.LastDigitIdx
	ia.IsZeroValue = iAry2.IsZeroValue
	ia.Precision = iAry2.Precision
	ia.SignVal = iAry2.SignVal
	ia.DecimalSeparator = iAry2.DecimalSeparator
	if ia.DecimalSeparator == 0 {
		ia.DecimalSeparator = '.'
	}

}

func (ia *IntAry) CopyOut() IntAry {
	ia.SetIntAryLength()
	iAry2 := IntAry{}.New()

	iAry2.NumStr = ia.NumStr
	iAry2.NumRunes = make([]rune, ia.NumRunesLen)
	for i := 0; i < ia.NumRunesLen; i++ {
		iAry2.NumRunes[i] = ia.NumRunes[i]
	}

	iAry2.NumRunesLen = ia.NumRunesLen

	iAry2.IntAry = make([]int, ia.IntAryLen)

	for i := 0; i < ia.IntAryLen; i++ {
		iAry2.IntAry[i] = ia.IntAry[i]
	}

	iAry2.IntAryLen = ia.IntAryLen
	iAry2.FirstDigitIdx = ia.FirstDigitIdx
	iAry2.LastDigitIdx = ia.LastDigitIdx
	iAry2.IsZeroValue = ia.IsZeroValue
	iAry2.Precision = ia.Precision
	iAry2.SignVal = ia.SignVal
	iAry2.DecimalSeparator = ia.DecimalSeparator
	if iAry2.DecimalSeparator == 0 {
		iAry2.DecimalSeparator = '.'
	}

	return iAry2
}

// CopyToBackUp - Makes a copy of the current
// IntAry object and saves it to BackUp.
//
// See ResetFromBackUp() to retrieve the
// last backup copy.
func (ia *IntAry) CopyToBackUp() {

	ia.ConvertIntAryToNumStr()

	iBa := BackUpIntAry{}.New()

	iBa.NumStr = ia.NumStr
	iBa.NumRunes = make([]rune, ia.NumRunesLen)
	for i := 0; i < ia.NumRunesLen; i++ {
		iBa.NumRunes[i] = ia.NumRunes[i]
	}

	iBa.NumRunesLen = ia.NumRunesLen

	iBa.IntAry = make([]int, ia.IntAryLen)

	for i := 0; i < ia.IntAryLen; i++ {
		iBa.IntAry[i] = ia.IntAry[i]
	}

	iBa.IntAryLen = ia.IntAryLen
	iBa.FirstDigitIdx = ia.FirstDigitIdx
	iBa.LastDigitIdx = ia.LastDigitIdx
	iBa.IsZeroValue = ia.IsZeroValue
	iBa.Precision = ia.Precision
	iBa.SignVal = ia.SignVal
	iBa.DecimalSeparator = ia.DecimalSeparator
	if iBa.DecimalSeparator == 0 {
		iBa.DecimalSeparator = '.'
	}

	ia.BackUp = iBa
}

// DecrementIntegerOne - Decrements the numeric value of the current
// IntAry by subtracting '1'.
//
// IMPORTANT: This method assumes that SetIntAryLength() and
// SetIsZeroValue() have already been called.

func (ia *IntAry) DecrementIntegerOne() error {

	if ia.IsZeroValue || ia.IsIntegerZeroValue {
		ia.SignVal = -1
	}

	intLen := ia.IntAryLen - ia.Precision
	intIdx := intLen - 1
	lastIdx := ia.IntAryLen - 1

	n1 := 0
	n2 := 0
	carry := 0

	ia.IsZeroValue = true
	ia.IsIntegerZeroValue = true

	for i := lastIdx; i >= 0; i-- {
		n1 = ia.IntAry[i]

		if i > intIdx {
			//  i > intIdx
			// This must be a fractional digit
			// Retain fractional digits

			if n1 != 0 {
				ia.IsZeroValue = false
			}

			continue

		} else if i == intIdx {

			n2 = n1 + (-1 * ia.SignVal)

			if n2 < 0 {
				n2 = n1 + 10 - 1
				carry = 1

			} else if n2 > 9 {
				n2 = n1 + 1 - 10
				carry = 1

			} else {
				carry = 0
			}

		} else {
			// Must be i < intIdx

			n2 = n1 + ((ia.SignVal * carry) * -1)

			if n2 < 0 {
				n2 = n1 + 10 - carry
				carry = 1
			} else if n2 > 9 {
				n2 = n1 - 10 + carry
				carry = 1
			} else {
				carry = 0
			}

		}

		if n2 != 0 {
			ia.IsZeroValue = false
			ia.IsIntegerZeroValue = false
		}

		ia.IntAry[i] = n2

	}

	if ia.IsZeroValue && carry == 0 {
		ia.SignVal = 1
	}

	if carry > 0 {

		ia.IntAry = append([]int{1}, ia.IntAry...)
		ia.IntAryLen++

	} else if ia.IntAry[0] == 0 && intLen > 1 {
		ia.IntAry = ia.IntAry[1:]
		ia.IntAryLen--
	}

	return nil
}

// DivideByTwo - Divides the current value of
// IntAry by 2. If parameter 'convertToNumStr'
// is set to 'true', the result will be converted
// to a number string.
func (ia *IntAry) DivideByTwo(convertToNumStr bool) {

	ia.OptimizeIntArrayLen(false, false)

	if ia.IsZeroValue {

		ia.SetIntAryToZero(ia.Precision)

		if convertToNumStr {
			ia.ConvertIntAryToNumStr()
		}

		return
	}

	n1 := 0
	n2 := 0
	carry := 0

	for i := 0; i < ia.IntAryLen; i++ {

		n1 = ia.IntAry[i] + carry
		n2 = n1 / 2
		carry = (n1 - (n2 * 2)) * 10
		ia.IntAry[i] = n2

	}

	if carry > 0 {
		ia.IntAry = append(ia.IntAry, 5)
		ia.IntAryLen++
		ia.Precision++
	}

	if ia.IntAry[0] == 0 {
		ia.SetSignificantDigitIdxs()
		ia.IntAry = ia.IntAry[ia.FirstDigitIdx:]
		ia.SetIntAryLength()
	}

	if convertToNumStr {
		ia.ConvertIntAryToNumStr()
	}

}

// DivideByInt64 - Divide the current value of the IntAry
// by an int64 'divisor' parameter passed to the method.
func (ia *IntAry) DivideByInt64(divisor int64, maxPrecision uint64, convertToNumStr bool) error {

	if divisor == 0 {
		return errors.New("'divisor' Equals Zero. Cannot Divide By Zero!")
	}

	ia.OptimizeIntArrayLen(false, false)

	if ia.IsZeroValue {

		ia.SetIntAryToZero(ia.Precision)

		if convertToNumStr {
			ia.ConvertIntAryToNumStr()
		}

		return nil
	}

	dSignVal := 1

	if divisor < 0 {
		dSignVal = -1
		divisor = divisor * -1
	}

	ia.SignVal = dSignVal * ia.SignVal

	n1 := int64(0)
	n2 := int64(0)
	carry := int64(0)
	iMaxPrecision := int(maxPrecision) + 1
	newAryLen := ia.IntAryLen
	intAryLen := ia.IntAryLen - ia.Precision
	precisionCnt := 0

	for i := 0; i < newAryLen; i++ {

		if i >= intAryLen {
			precisionCnt++
		}

		if i < ia.IntAryLen {
			n1 = int64(ia.IntAry[i]) + carry
		} else {
			n1 = int64(0) + carry
		}

		n2 = n1 / divisor
		carry = (n1 - (n2 * divisor)) * 10

		if i < ia.IntAryLen {
			ia.IntAry[i] = int(n2)
		} else {
			ia.IntAry = append(ia.IntAry, int(n2))
		}

		if i == newAryLen-1 &&
			carry > 0 && precisionCnt <= iMaxPrecision {

			newAryLen++

		}

	}

	ia.Precision = precisionCnt

	ia.IntAryLen = newAryLen

	if precisionCnt >= iMaxPrecision {
		iMaxPrecision--
		ia.RoundToPrecision(iMaxPrecision)
	}

	if ia.IntAry[0] == 0 {
		ia.SetSignificantDigitIdxs()
		ia.IntAry = ia.IntAry[ia.FirstDigitIdx:]
		ia.SetIntAryLength()
	}

	if convertToNumStr {
		ia.ConvertIntAryToNumStr()
	}

	return nil
}

// DivideByTenToPower - Divide IntAry value by
// 10 raised to the power of the 'power'
// parameter.
//
// If parameter 'convertToNumStr' is set to 'true', the
// result will be automatically converted to a number string.
func (ia *IntAry) DivideByTenToPower(power uint, convertToNumStr bool) {

	if power == 0 {
		return
	}

	ia.Precision += int(power)
	ia.IntAryLen = len(ia.IntAry)
	newLen := ia.Precision + 1

	if ia.IntAryLen < newLen {

		t := make([]int, newLen)

		deltaLen := newLen - ia.IntAryLen

		for i := 0; i < newLen; i++ {

			if i < deltaLen {
				t[i] = 0
			} else {
				t[i] = ia.IntAry[i-deltaLen]
			}

		}

		ia.IntAry = make([]int, newLen)
		for i := 0; i < newLen; i++ {
			ia.IntAry[i] = t[i]
		}

		ia.IntAryLen = newLen
	}

	if convertToNumStr {
		ia.ConvertIntAryToNumStr()
	}

}

// DivideThisBy - Divides the current value of IntAry by the parameter iAry2.
// The result of this division is returned as an IntAry.
//
// Maximum Precision of the division result is controlled by the input
// parameter, 'maxPrecision'.
//
func (ia *IntAry) DivideThisBy(iAry2 *IntAry, maxPrecision int) (IntAry, error) {

	ia.SetInternalFlags()
	iAry2.SetInternalFlags()

	if iAry2.IsZeroValue {
		return IntAry{}.New(), errors.New("Error: Divide By ZERO!")
	}

	quotient := IntAry{}.New()
	quotient.SetIntAryToZero(0)

	if ia.IsZeroValue {
		return quotient, nil
	}

	trialDividend := ia.CopyOut()

	divisor := iAry2.CopyOut()

	tensCount := IntAry{}.New()
	tensCount.SetIntAryToOne(0)

	newSignVal := 1

	if trialDividend.SignVal != divisor.SignVal {
		newSignVal = -1
	}

	if trialDividend.SignVal == -1 {
		trialDividend.SignVal = 1
		trialDividend.ConvertIntAryToNumStr()
	}

	if divisor.SignVal == -1 {
		divisor.SignVal = 1
		trialDividend.ConvertIntAryToNumStr()
	}

	dividendMag := trialDividend.GetMagnitude()
	divisorMag := divisor.GetMagnitude()
	deltaMag := uint(0)
	incrementVal := IntAry{}.New()
	incrementVal = divisor.CopyOut()

	if dividendMag > divisorMag {
		deltaMag = uint(dividendMag - divisorMag)
		tensCount.MultiplyByTenToPower(deltaMag, true)
		incrementVal.MultiplyThisBy(&tensCount, false)

	} else if divisorMag > dividendMag {
		deltaMag = uint(divisorMag - dividendMag)
		trialDividend.MultiplyByTenToPower(deltaMag, true)
		tensCount.DivideByTenToPower(deltaMag, true)

	}

	compare := 0
	precisionCutOff := maxPrecision + dividendMag + 1

	for true {

		if quotient.Precision == precisionCutOff {
			quotient.RoundToPrecision(maxPrecision)
			quotient.OptimizeIntArrayLen(true, false)
			quotient.SignVal = newSignVal
			quotient.ConvertIntAryToNumStr()
			return quotient, nil
		}

		compare = incrementVal.CompareAbsoluteValues(&trialDividend)

		if compare == 0 {
			// incrementalVal is equal to trialDividend
			quotient.AddToThis(&tensCount, false)
			quotient.RoundToPrecision(maxPrecision)
			quotient.OptimizeIntArrayLen(true, false)
			quotient.SignVal = newSignVal
			quotient.ConvertIntAryToNumStr()
			return quotient, nil

		} else if compare == -1 {
			// incrementalVal < trialDividend
			quotient.AddToThis(&tensCount, false)

			// Calc Remainder
			trialDividend.SubtractFromThis(&incrementVal, false)

			continue

		} else {
			// Must Be compare == 1
			// incrementalVal > trialDividend

			tensCount.DivideByTenToPower(1, false)
			incrementVal.DivideByTenToPower(1, false)
		}

	}

	return quotient, nil
}

// Floor - Math 'Floor' function. Finds the
// integer number which is less than or
// equal to the value of the current IntAry.
func (ia *IntAry) Floor() (IntAry, error) {

	err := ia.IsIntAryValid("Floor() - ")

	if err != nil {
		return IntAry{}, err
	}

	iAry2 := IntAry{}.New()

	if ia.IsZeroValue {
		iAry2.SetIntAryToZero(ia.Precision)
	}

	hasFracDigits, err := ia.HasFractionalDigits()

	if err != nil {
		return iAry2, err
	}

	if !hasFracDigits {
		// There are NO non-zero digits to the
		// right of the decimal place
		iAry2 = ia.CopyOut()
		return iAry2, nil
	}

	intLen := ia.IntAryLen - ia.Precision
	intIdx := intLen - 1

	if hasFracDigits && ia.SignVal > 0 {
		// There ARE non-zero digits to the right of the
		// decimal place
		t := make([]int, ia.IntAryLen)

		for i := 0; i < intLen; i++ {
			t[i] = ia.IntAry[i]
		}

		iAry2.IntAry = t[0:]
		iAry2.IntAryLen = ia.IntAryLen
		iAry2.Precision = ia.Precision
		iAry2.SignVal = ia.SignVal
		iAry2.ConvertIntAryToNumStr()
		return iAry2, nil
	}

	// The number has non-zero digits to
	// the right of the decimal place and
	// the number sign is minus (- or ia.SignVal = -1)

	t := make([]int, ia.IntAryLen+1)

	n1 := 0
	n2 := 0
	carry := 0

	for i := intIdx; i >= 0; i-- {

		n1 = ia.IntAry[i]

		if i == intIdx {

			if n1+1 > 9 {
				n2 = n1 + 1 - 10
				carry = 1
			} else {
				n2 = n1 + 1
				carry = 0
			}

		} else {

			if n1+carry > 9 {
				n2 = n1 + carry - 10
				carry = 1
			} else {
				n2 = n1 + carry
				carry = 0
			}
		}

		t[i+1] = n2

	}

	if carry != 0 {
		t[0] = carry
		iAry2.IntAry = t[0 : ia.IntAryLen+1]

	} else {

		iAry2.IntAry = t[1 : ia.IntAryLen+1]
	}

	iAry2.Precision = ia.Precision
	iAry2.SignVal = ia.SignVal
	iAry2.IntAryLen = len(iAry2.IntAry)
	iAry2.SetIsZeroValue()
	iAry2.ConvertIntAryToNumStr()

	return iAry2, nil
}

// GetFractionalDigits - Examines the current IntAry and
// returns another IntAry consisting of the fractional
// digits to the right of the decimal point from the
// current IntAry object.
//
// Note: The Sign Value of the returned int Ary is always
// positive or +1.
//
// The return IntAry will display fractional digits with
// a leading integer digit of zero. Example '0.5678'
func (ia *IntAry) GetFractionalDigits() (IntAry, error) {

	err := ia.IsIntAryValid("GetFractionalDigits() - ")

	if err != nil {
		return IntAry{}, err
	}

	iAry2 := IntAry{}.New()

	iAry2.SetIntAryToZero(0)

	if ia.Precision == 0 {
		return iAry2, nil
	}

	fracIdx := ia.IntAryLen - ia.Precision

	iAry2.IntAry = make([]int, ia.Precision+1)
	idx := 1
	for i := fracIdx; i < ia.IntAryLen; i++ {
		iAry2.IntAry[idx] = ia.IntAry[i]
		idx++
	}

	iAry2.Precision = ia.Precision
	iAry2.SignVal = 1
	iAry2.ConvertIntAryToNumStr()

	return iAry2, nil
}

// GetIntegerDigits - Examines the current IntAry object
// and returns a new IntAry consisting of only the integer
// digits to the left of the decimal point in the current
// IntAry object.
func (ia *IntAry) GetIntegerDigits() (IntAry, error) {

	err := ia.IsIntAryValid("GetFractionalDigits() - ")

	if err != nil {
		return IntAry{}, err
	}

	iAry2 := IntAry{}.New()

	if ia.IsZeroValue {
		iAry2.SetIntAryToZero(0)
		return iAry2, nil
	}

	intLen := ia.IntAryLen - ia.Precision

	iAry2.IntAry = make([]int, intLen)

	for i := 0; i < intLen; i++ {
		iAry2.IntAry[i] = ia.IntAry[i]
	}

	iAry2.SignVal = ia.SignVal
	iAry2.Precision = 0

	iAry2.SetInternalFlags()

	if iAry2.IsZeroValue {
		iAry2.SignVal = 1
	}

	iAry2.ConvertIntAryToNumStr()

	return iAry2, nil

}

func (ia *IntAry) GetMagnitude() int {
	ia.SetSignificantDigitIdxs()
	return ia.IntAryLen - ia.Precision - ia.FirstDigitIdx

}

// HasFractionalDigits - This method examines the
// current IntAry object to determine if there
// are non-zero digits to the right of the decimal
// place. If all digits to the right of the decimal
// place are zero, this method returns 'false'
//
// If non-zero digits are present to the right of the
// decimal place, the method returns 'true'.
func (ia *IntAry) HasFractionalDigits() (bool, error) {

	err := ia.IsIntAryValid("HasFractionalDigits() - ")

	if err != nil {
		return false, err
	}

	if ia.Precision == 0 {
		return false, nil
	}

	ia.SetIntAryLength()

	intLen := ia.IntAryLen - ia.Precision

	if intLen < 1 {
		return false, fmt.Errorf("HasFractionalDigits() Error - Int Array integer length is less than 1. intLen= '%v'", intLen)
	}

	for i := intLen; i < ia.IntAryLen; i++ {
		if ia.IntAry[i] > 0 {
			return true, nil
		}
	}

	return false, nil
}

// IncrementIntegerOne - Increment the value of the
// current IntAry by adding '1'
func (ia *IntAry) IncrementIntegerOne() error {

	if ia.IsZeroValue || ia.IsIntegerZeroValue {
		ia.SignVal = 1
	}

	intLen := ia.IntAryLen - ia.Precision
	intIdx := intLen - 1
	lastIdx := ia.IntAryLen - 1

	n1 := 0
	n2 := 0
	carry := 0

	ia.IsZeroValue = true
	ia.IsIntegerZeroValue = true

	for i := lastIdx; i >= 0; i-- {
		n1 = ia.IntAry[i]

		if i > intIdx {
			//  i > intIdx
			// This must be a fractional digit
			// Retain fractional digits

			if n1 != 0 {
				ia.IsZeroValue = false
			}

			continue

		} else if i == intIdx {

			n2 = n1 + (1 * ia.SignVal)

			if n2 < 0 {
				n2 = n1 + 10 - 1
				carry = 1

			} else if n2 > 9 {
				n2 = n1 + 1 - 10
				carry = 1

			} else {
				carry = 0
			}

		} else {
			// Must be i < intIdx

			n2 = n1 + ((ia.SignVal * carry) * 1)

			if n2 < 0 {
				n2 = n1 + 10 - carry
				carry = 1
			} else if n2 > 9 {
				n2 = n1 - 10 + carry
				carry = 1
			} else {
				carry = 0
			}

		}

		if n2 != 0 {
			ia.IsZeroValue = false
			ia.IsIntegerZeroValue = false
		}

		ia.IntAry[i] = n2

	}

	if ia.IsZeroValue && carry == 0 {
		ia.SignVal = 1
	}

	if carry > 0 {

		ia.IntAry = append([]int{1}, ia.IntAry...)
		ia.IntAryLen++

	} else if ia.IntAry[0] == 0 && intLen > 1 {
		ia.IntAry = ia.IntAry[1:]
		ia.IntAryLen--
	}

	return nil
}

/*
func (ia *IntAry) IncrementIntegerOne() error {

	ia.SetIntAryLength()
	intIdx := ia.IntAryLen - ia.Precision - 1
	lastIdx := ia.IntAryLen - 1

	n1 := 0
	n2 := 0
	carry := 0

	for i := lastIdx; i >= 0; i-- {

		n1 = ia.IntAry[i]

		if i == intIdx {

			if n1+1 > 9 {
				n2 = n1 + 1 - 10
				carry = 1
			} else {
				n2 = n1 + 1
				carry = 0
			}

		} else if i < intIdx {

			if n1+carry > 9 {
				n2 = n1 + carry - 10
				carry = 1
			} else {
				n2 = n1 + carry
				carry = 0
			}
		} else {
			// Must be i > intIdx
			// Retain Fractional Digits
			//n2 = ia.IntAry[i]
			continue
		}

		ia.IntAry[i] = n2

	}

	if carry != 0 {

		ia.IntAry = append([]int{1}, ia.IntAry...)
		ia.IntAryLen++

	}

	return nil

}
*/

// IsIntAryValid - Examines the current IntAry and returns
// an error if the IntAry object is found to be invalid.
func (ia *IntAry) IsIntAryValid(errName string) error {

	if ia.SignVal != -1 && ia.SignVal != 1 {
		return fmt.Errorf("%v Sign Value is INVALID! Sign Value= '%v'", errName, ia.SignVal)
	}

	if ia.Precision < 0 {
		return fmt.Errorf("%v Precision Value is INVALID! Sign Value= '%v'", errName, ia.Precision)
	}

	ia.SetSignificantDigitIdxs()

	if ia.Precision >= ia.IntAryLen {
		return fmt.Errorf("%v ERROR: Precision is Greater Than or Equal To IntArray Length! ia.Precision= %v  ia.IntAryLen= %v ", errName, ia.Precision, ia.IntAryLen)
	}

	if ia.IntegerLen == 0 {
		return fmt.Errorf("%v ERROR: Integer Length is Zero. Missing Leading Integer Zero.", errName)
	}

	return nil
}

// MultiplyByTwoToPower Multiply the existing value
// of the IntAry by 2 to the power of the passed in
// parameter.
//
func (ia *IntAry) MultiplyByTwoToPower(power uint, convertToNumStr bool) {

	ia.SetIntAryLength()

	if power == 0 {
		return
	}

	for h := uint(0); h < power; h++ {
		n1 := 0
		carry := 0

		for i := ia.IntAryLen - 1; i >= 0; i-- {

			n1 = (ia.IntAry[i] * 2) + carry

			if n1 > 9 {
				n1 = n1 - 10
				carry = 1
			} else {
				carry = 0
			}

			ia.IntAry[i] = n1
		}

		if carry > 0 {
			ia.IntAry = append([]int{1}, ia.IntAry...)
			ia.IntAryLen++
		}

	}

	if convertToNumStr {
		ia.ConvertIntAryToNumStr()
	}

}

// MultiplyByTenToPower - The value of IntAry is multiplied
// by 10 to the power of the passed in parameter.
func (ia *IntAry) MultiplyByTenToPower(power uint, convertToNumStr bool) {

	if power == 0 {
		return
	}
	for i := uint(0); i < power; i++ {

		if ia.Precision > 0 {
			ia.Precision--
			continue
		}

		ia.IntAry = append(ia.IntAry, 0)
	}

	ia.IntAryLen = len(ia.IntAry)

	if ia.Precision < 0 {
		ia.Precision = 0
	}

	if convertToNumStr {
		ia.ConvertIntAryToNumStr()
	}

}

func (ia *IntAry) MultiplyThisBy(ia2 *IntAry, convertToNumStr bool) error {

	ia.SetInternalFlags()
	ia2.SetInternalFlags()

	if ia.IsZeroValue {
		return nil
	}

	if ia2.IsZeroValue {
		ia.SetIntAryToZero(ia.Precision)
		return nil
	}

	newSignVal := 1

	if ia.SignVal != ia2.SignVal {
		newSignVal = -1
	}

	newPrecision := ia.Precision + ia2.Precision

	lenLevels := ia2.IntAryLen
	lenNumPlaces := (ia.IntAryLen + ia2.IntAryLen) + 1

	intMAry := make([][]int, lenLevels)

	for i := 0; i < lenLevels; i++ {
		intMAry[i] = make([]int, lenNumPlaces)
	}

	carry := 0
	levels := 0
	place := 0
	n1 := 0
	n2 := 0
	n3 := 0
	n4 := 0
	x := 0

	for i := ia2.IntAryLen - 1; i >= 0; i-- {

		place = (lenNumPlaces - 1) - levels

		for j := ia.IntAryLen - 1; j >= 0; j-- {

			n1 = ia.IntAry[j]
			n2 = ia2.IntAry[i]
			n3 = (n1 * n2) + carry
			x = n3 / 10
			n4 = n3 - (x * 10)
			// n4 = int(math.Mod(float64(n3), float64(10.00)))

			intMAry[levels][place] = n4

			carry = x

			place--
		}

		intMAry[levels][place] = carry
		carry = 0
		levels++
	}

	carry = 0
	n1 = 0
	n2 = 0
	n3 = 0
	n4 = 0
	x = 0

	ia.IntAryLen = lenNumPlaces
	ia.IntAry = make([]int, ia.IntAryLen)

	for i := 0; i < lenLevels; i++ {
		for j := lenNumPlaces - 1; j >= 0; j-- {

			n1 = ia.IntAry[j]
			n2 = intMAry[i][j]
			n3 = n1 + n2 + carry
			n4 = 0

			if n3 > 9 {
				x = n3 / 10
				n4 = n3 - (x * 10)
				carry = x

			} else {
				n4 = n3
				carry = 0
			}

			ia.IntAry[j] = n4
		}

	}

	if carry > 0 {
		ia.IntAry = append([]int{1}, ia.IntAry...)
	}

	ia.Precision = newPrecision
	ia.SignVal = newSignVal

	if ia.IntAry[0] == 0 {
		ia.SetSignificantDigitIdxs()
		ia.IntAry = ia.IntAry[ia.FirstDigitIdx:]
	}

	if convertToNumStr {
		ia.ConvertIntAryToNumStr()
	}

	return nil

}

// OptimizeIntArrayLen - Eliminates Leading
// zeros from the front or integer portion
// of the integer string.
//
// If parameter 'optimizeFracDigits' is set
// equal to 'true', trailing zeros to the
// right of the decimal place will also be
// eliminated.
//
func (ia *IntAry) OptimizeIntArrayLen(optimizeFracDigits bool, convertToNumStr bool) {
	ia.SetSignificantDigitIdxs()

	integerLen := ia.IntAryLen - ia.Precision - ia.FirstDigitIdx

	if optimizeFracDigits {

		ia.IntAry = ia.IntAry[ia.FirstDigitIdx : ia.LastDigitIdx+1]
		ia.IntAryLen = ia.LastDigitIdx - ia.FirstDigitIdx + 1
	} else {
		ia.IntAry = ia.IntAry[ia.FirstDigitIdx:]
		ia.IntAryLen = ia.IntAryLen - ia.FirstDigitIdx
	}

	ia.Precision = ia.IntAryLen - integerLen

	if convertToNumStr {
		ia.ConvertIntAryToNumStr()
	}
}

// ResetFromBackUp - Retrieves data from the
// last saved backup and populates the current
// IntAry object.
func (ia *IntAry) ResetFromBackUp() {
	ia.BackUp.SetInternalFlags()

	ia.NumStr = ia.BackUp.NumStr
	ia.NumRunes = make([]rune, ia.BackUp.NumRunesLen)
	for i := 0; i < ia.BackUp.NumRunesLen; i++ {
		ia.NumRunes[i] = ia.BackUp.NumRunes[i]
	}

	ia.NumRunesLen = ia.BackUp.NumRunesLen

	ia.IntAry = make([]int, ia.BackUp.IntAryLen)
	for i := 0; i < ia.BackUp.IntAryLen; i++ {
		ia.IntAry[i] = ia.BackUp.IntAry[i]
	}

	ia.IntAryLen = ia.BackUp.IntAryLen
	ia.FirstDigitIdx = ia.BackUp.FirstDigitIdx
	ia.LastDigitIdx = ia.BackUp.LastDigitIdx
	ia.IsZeroValue = ia.BackUp.IsZeroValue
	ia.Precision = ia.BackUp.Precision
	ia.SignVal = ia.BackUp.SignVal
	ia.DecimalSeparator = ia.BackUp.DecimalSeparator
	if ia.DecimalSeparator == 0 {
		ia.DecimalSeparator = '.'
	}

}

// RoundToPrecision - Rounds the value of the IntAry to a precision
// specified by the 'roundToPrecision' parameter.
func (ia *IntAry) RoundToPrecision(roundToPrecision int) error {

	if roundToPrecision < 0 {
		fmt.Errorf("RoundToPrecision() - Error: roundToPrecision is less than ZERO! roundToPrecision= '%v'", roundToPrecision)
	}

	if ia.Precision == 0 {
		return nil
	}

	err := ia.IsIntAryValid("RoundToPrecision() - ")

	if err != nil {
		return err
	}

	if ia.IsZeroValue {
		ia.SetIntAryToZero(roundToPrecision)
		return nil
	}

	if roundToPrecision == ia.Precision {
		return nil
	}

	if roundToPrecision > ia.Precision {
		deltaPrecision := roundToPrecision - ia.Precision

		for i := 0; i < deltaPrecision; i++ {
			ia.IntAry = append(ia.IntAry, 0)
		}

		ia.IntAryLen = len(ia.IntAry)
		ia.Precision = roundToPrecision
		ia.ConvertIntAryToNumStr()
		return nil
	}

	// roundToPrecision must be < ia.Precision

	intLen := ia.IntAryLen - ia.Precision
	newIntAryLen := intLen + roundToPrecision
	fracIdx := intLen

	fracRoundIdx := fracIdx + roundToPrecision

	t := make([]int, ia.IntAryLen+1)
	n1 := 0
	n2 := 0

	carry := 0
	for i := fracRoundIdx; i >= 0; i-- {

		n1 = ia.IntAry[i]

		if i == fracRoundIdx {
			n2 = n1 + 5
		} else {
			n2 = n1 + carry
		}

		if n2 > 9 {
			carry = 1
			n2 = n2 - 10
		} else {
			carry = 0
		}

		t[i+1] = n2
	}

	ia.IntAry = []int{}

	if carry > 0 {
		t[0] = carry
		ia.IntAry = t[0 : newIntAryLen+1]
	} else {
		ia.IntAry = t[1 : newIntAryLen+1]
	}
	ia.Precision = roundToPrecision
	ia.IntAryLen = len(ia.IntAry)
	ia.ConvertIntAryToNumStr()
	return nil
}

// SetEqualArrayLengths - Compares an IntAry object
// to the current IntAry and ensures that the lengths
// of both IntArrays are equal.
func (ia *IntAry) SetEqualArrayLengths(iAry2 *IntAry) {
	iAry2.SetIntAryLength()
	ia.SetIntAryLength()

	iaIntLen := ia.IntAryLen - ia.Precision
	iAry2IntLen := iAry2.IntAryLen - iAry2.Precision

	if iaIntLen > iAry2IntLen {
		iAry2.AddArrayLengthLeft(iaIntLen - iAry2IntLen)
	}

	if iAry2IntLen > iaIntLen {
		ia.AddArrayLengthLeft(iAry2IntLen - iaIntLen)
	}

	if ia.Precision > iAry2.Precision {
		iAry2.AddArrayLengthRight(ia.Precision - iAry2.Precision)
		iAry2.Precision = ia.Precision
	}

	if iAry2.Precision > ia.Precision {
		ia.AddArrayLengthRight(iAry2.Precision - ia.Precision)
		ia.Precision = iAry2.Precision
	}

	ia.ConvertIntAryToNumStr()
	iAry2.ConvertIntAryToNumStr()
	return
}


// SetIntAryWithInt64 - Sets the value of the current IntAry
// object to that of the input parameter 'intDigits', a 64-bit
// integer.
//
// Note: Input parameter 'precision' to indicate the number of
// digits to the right of the decimal place.
//
// Example:
//  intDigits  precision  signVal		result
//  946254  			3					1				946.254
//  946254				0					1				946254
//  946254  			3					-1			-946.254
//  946254				0					-1			-946254
//
// The Sign of the resulting
func (ia *IntAry) SetIntAryWithInt64(intDigits uint64, precision uint, signVal int) error {

	if signVal != 1 && signVal != -1 {
		return fmt.Errorf("ERROR - Input parameter must be equal to +1 or -1. Input signVal= %v", signVal)
	}

	ia.SignVal = signVal

	if intDigits == 0 {
		ia.SetIntAryToZero(int(precision))
		return nil
	}

	ia.Precision = int(precision)

	newIntDigits := intDigits
	quotient := uint64(0)
	mod := uint64(0)
	ten := uint64(10)

	ia.IntAry = []int{}
	ia.IntAryLen = 0
	for true {

		if newIntDigits==0 {
			break
		}

		quotient = newIntDigits / ten
		mod = newIntDigits - (quotient * ten)

		ia.IntAry = append(ia.IntAry, int(mod))
		ia.IntAryLen++

		newIntDigits = quotient

	}

	n1:=0
	lastIdx := ia.IntAryLen -1
	totalLen := ia.IntAryLen /2
	for i:=0; i < totalLen; i++ {
		n1 = ia.IntAry[i]
		ia.IntAry[i] = ia.IntAry[lastIdx]
		ia.IntAry[lastIdx] = n1
		lastIdx--
	}

	ia.ConvertIntAryToNumStr()

	return nil
}


// SetIntAryWithIntAry - Sets the value of the current IntAry based on []int, precision
// and Sign Value arguments passed to this method. Note: If signVal is not equal to
// +1 or -1, an error is generated.
func (ia *IntAry) SetIntAryWithIntAry(iAry2 []int, precision uint, signVal int) error {


	if signVal != 1 && signVal != -1 {
		return fmt.Errorf("SetIntAryWithIntAry() - Error: signVal parameter is INVALID! signVal must be -1 or +1. signVal='%v'", signVal)
	}

	lIAry2 := len(iAry2)
	ia.IntAry = make([]int, lIAry2)
	for i := 0; i < lIAry2; i++ {
		ia.IntAry[i] = iAry2[i]
	}

	ia.IntAryLen = lIAry2
	ia.Precision = int(precision)
	ia.SignVal = signVal
	ia.ConvertIntAryToNumStr()

	return nil
}

// SetIntAryWithNumStr - receives a raw number string and sets the
// fields of the internal IntAry structure to the appropriate
// values.
func (ia *IntAry) SetIntAryWithNumStr(str string) error {

	if len(str) == 0 {
		return errors.New("SetIntAryWithNumStr() Received zero length number string!")
	}

	ia.Empty()

	if ia.DecimalSeparator == 0 {
		ia.DecimalSeparator = '.'
	}

	ia.SignVal = 1
	baseRunes := []rune(str)
	lBaseRunes := len(baseRunes)
	isStartRunes := false
	isEndRunes := false
	isFractionalValue := false

	for i := 0; i < lBaseRunes && isEndRunes == false; i++ {

		if baseRunes[i] == '+' ||
			baseRunes[i] == ' ' ||
			baseRunes[i] == ',' ||
			baseRunes[i] == '$' {

			continue

		}

		if isStartRunes == true &&
			isEndRunes == false &&
			isFractionalValue &&
			baseRunes[i] == ia.DecimalSeparator {

			continue
		}

		if baseRunes[i] == '-' &&
			isStartRunes == false && isEndRunes == false &&
			i+1 < lBaseRunes &&
			((baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9') ||
				baseRunes[i+1] == ia.DecimalSeparator) {

			ia.SignVal = -1
			isStartRunes = true
			continue

		} else if isEndRunes == false &&
			baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			ia.NumRunes = append(ia.NumRunes, baseRunes[i])
			ia.IntAry = append(ia.IntAry, int(baseRunes[i])-48)
			isStartRunes = true

			if isFractionalValue {
				ia.Precision++
			}

		} else if isEndRunes == false &&
			i+1 < lBaseRunes &&
			baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' &&
			baseRunes[i] == ia.DecimalSeparator {

			isFractionalValue = true
			continue

		} else if isStartRunes && !isEndRunes {

			isEndRunes = true

		}

	}

	ia.SetSignificantDigitIdxs()

	if ia.IntAryLen == 0 || ia.IsZeroValue {
		ia.SetIntAryToZero(ia.Precision)
		return nil
	}

	ia.NumStr = ""

	if ia.SignVal < 0 {
		ia.NumStr = "-"
	}

	intRunes := ia.NumRunesLen - ia.Precision

	ia.NumStr += string(ia.NumRunes[0:intRunes])

	if ia.Precision > 0 {

		ia.NumStr += string(ia.DecimalSeparator)
		ia.NumStr += string(ia.NumRunes[intRunes:])
	}

	// Validate IntAry object
	err := ia.IsIntAryValid("SetIntAryWithNumStr() - ")

	if err != nil {
		return err
	}

	if ia.IntAryLen != ia.NumRunesLen {
		ia.ConvertIntAryToNumStr()
	}



	return nil
}

func (ia *IntAry) SetIntAryLength() {
	ia.IntAryLen = len(ia.IntAry)
	ia.NumRunesLen = len(ia.NumRunes)
}

// SetIsZeroValue - Analyzes the value
// of the IntAry and sets a flag
// if the value of IntAry evaluates
// to zero.
func (ia *IntAry) SetIsZeroValue() {
	ia.IntAryLen = len(ia.IntAry)

	ia.IsZeroValue = true
	ia.IsIntegerZeroValue = true

	intLen := ia.IntAryLen - ia.Precision

	for i := 0; i < ia.IntAryLen; i++ {

		if i < intLen && ia.IntAry[i] > 0 {
			ia.IsIntegerZeroValue = false
		}

		if ia.IntAry[i] > 0 {
			ia.IsZeroValue = false
			return
		}
	}

	// ia.IsZeroValue == true
	// SignVal must be 1
	ia.SignVal = 1
}

// SetIntAryToOne - Sets the value of the IntAry object to one ('1').
func (ia *IntAry) SetIntAryToOne(precision int) error {

	if precision < 0 {
		return fmt.Errorf("SetIntAryToOne() - Error: Precision is less than ZERO! precision= '%v'", precision)
	}

	ia.IntAryLen = 1 + precision
	ia.Precision = precision
	ia.IntAry = make([]int, ia.IntAryLen)
	ia.IntAry[0] = 1
	ia.SignVal = 1
	ia.IsZeroValue = false
	ia.IsIntegerZeroValue = false
	ia.FirstDigitIdx = 0
	ia.LastDigitIdx = 0

	if ia.DecimalSeparator == 0 {
		ia.DecimalSeparator = '.'
	}

	ia.ConvertIntAryToNumStr()

	return nil
}

// SetIntAryToTen - Sets the value of the IntAry object to Ten ('10')
func (ia *IntAry) SetIntAryToTen(precision int) error {

	if precision < 0 {
		return fmt.Errorf("SetIntAryToOne() - Error: Precision is less than ZERO! precision= '%v'", precision)
	}

	ia.IntAryLen = 2 + precision
	ia.Precision = precision
	ia.IntAry = make([]int, ia.IntAryLen)
	ia.IntAry[0] = 1
	ia.SignVal = 1
	ia.IsZeroValue = false
	ia.IsIntegerZeroValue = false
	ia.FirstDigitIdx = 0
	ia.LastDigitIdx = 0

	if ia.DecimalSeparator == 0 {
		ia.DecimalSeparator = '.'
	}

	ia.ConvertIntAryToNumStr()

	return nil
}

// SetIntAryToZero - Sets the value of the IntAry object to Zero ('0').
func (ia *IntAry) SetIntAryToZero(precision int) error {

	if precision < 0 {
		return fmt.Errorf("SetIntAryToOne() - Error: Precision is less than ZERO! precision= '%v'", precision)
	}

	ia.IntAryLen = 1 + precision
	ia.Precision = precision
	ia.IntAry = make([]int, ia.IntAryLen)
	ia.SignVal = 1
	ia.IsZeroValue = true
	ia.IsIntegerZeroValue = true
	ia.FirstDigitIdx = 0
	ia.LastDigitIdx = 0

	if ia.DecimalSeparator == 0 {
		ia.DecimalSeparator = '.'
	}

	ia.ConvertIntAryToNumStr()

	return nil
}

// SetInternalFlags - Sets Array Lengths and
// test for zero values
func (ia *IntAry) SetInternalFlags() {
	ia.SetSignificantDigitIdxs()
}

func (ia *IntAry) SetNumRunesLength() {
	ia.NumRunesLen = len(ia.NumRunes)
}

// SetTruncateToPrecision - Truncates the existing
// value to precision specified by the 'precision'
// parameter. No rounding is performed.
//
// If 'precision' is zero, the Int Ary value will
// be truncated to an integer value with no
// fractional digits.
//
// If 'precision' is set to a value less than zero,
// an error will be returned.
//
// If 'precision' is greater than the existing precision,
// trailing zeros will be added

func (ia *IntAry) SetPrecision(precision int, roundResult bool) error {

	if precision < 0 {
		return fmt.Errorf("SetPrecision() - Error: 'precision' value is less than ZERO! precision= '%v'", precision)
	}

	err := ia.IsIntAryValid("SetPrecision() - ")

	if err != nil {
		return err
	}

	if ia.IsZeroValue {
		ia.SetIntAryToZero(precision)
		return nil
	}

	if precision == ia.Precision {
		return nil
	}

	if precision > ia.Precision {
		deltaPrecision := precision - ia.Precision

		for i := 0; i < deltaPrecision; i++ {
			ia.IntAry = append(ia.IntAry, 0)
		}

		ia.Precision = precision
		ia.IntAryLen = len(ia.IntAry)
		ia.ConvertIntAryToNumStr()
		return nil
	}

	// Must ia.Precision > precision

	if roundResult {
		ia.RoundToPrecision(precision)
		return nil
	}

	intLen := ia.IntAryLen - ia.Precision
	newAryLen := intLen + precision
	ia.IntAry = ia.IntAry[0:newAryLen]

	ia.IntAryLen = newAryLen
	ia.Precision = precision
	ia.ConvertIntAryToNumStr()

	return nil
}

// SetSignificantDigitIdxs - Finds the first
// significant digit (the first numeric digit
// greater than zero) and sets index value in
// the local field variable, 'FirstDigitIdx'.
//
// In addition, this method also identifies the
// Last Significant Digit (the last non-zero value
// in the IntAry) and records that index in the
// local field variable, 'LastDigitIdx'.
func (ia *IntAry) SetSignificantDigitIdxs() {

	ia.IntAryLen = len(ia.IntAry)
	ia.NumRunesLen = len(ia.NumRunes)

	if ia.IntAryLen == ia.Precision {
		ia.IntAry = append([]int{0}, ia.IntAry...)
		ia.IntAryLen++
	}

	if ia.IntAryLen < ia.Precision {

		deltaZeros := ia.Precision - ia.IntAryLen + 1
		zeroAry := make([]int,deltaZeros)
		ia.IntAry = append(zeroAry, ia.IntAry...)
		ia.IntAryLen += deltaZeros
	}


	ia.FirstDigitIdx = -1
	ia.LastDigitIdx = -1

	ia.IntegerLen = 0
	ia.SignificantIntegerLen = 0
	ia.SignificantFractionLen = 0

	lastIntIdx := ia.IntAryLen - ia.Precision - 1
	ia.IsZeroValue = true
	ia.IsIntegerZeroValue = true
	ia.IntegerLen = ia.IntAryLen - ia.Precision

	for i := 0; i < ia.IntAryLen; i++ {
		if ia.IntAry[i] > 0 {
			ia.IsZeroValue = false

			if i < ia.IntegerLen {
				ia.IsIntegerZeroValue = false
			}
		}


		// At minimum, there should be a single
		// leading zero before the decimal point.
		// Example 0.000.
		if i == lastIntIdx && ia.IntAry[i] == 0 {

			if ia.FirstDigitIdx == -1 {
				ia.FirstDigitIdx = i
			}

		}

		if ia.IntAry[i] > 0 {

			if ia.FirstDigitIdx == -1 {
				ia.FirstDigitIdx = i
			}

			ia.LastDigitIdx = i
		}

	}

	ia.SignificantIntegerLen = ia.IntAryLen - ia.Precision - ia.FirstDigitIdx

	if ia.LastDigitIdx >= ia.IntegerLen {
		ia.SignificantFractionLen = ia.LastDigitIdx - ia.IntegerLen + 1
	} else {
		ia.SignificantFractionLen = 0
	}
}

// SubtractFromThis - Subtracts the value of parameter
// 'ia2' from the current IntAry object.
// Paramters:
//
// ia2 *IntAry - Incoming IntAry object whose value will be subtracted
// 								from this current IntAry value.
//
// convertToNumStr - boolean value determines whether the current IntAry
//                   object will convert the IntAry value to a number string.
//                   Set this parameter to 'false' if this method is called
//                   multiple times in order to improve performance.
func (ia *IntAry) SubtractFromThis(ia2 *IntAry, convertToNumStr bool) error {

	ia.SetEqualArrayLengths(ia2)

	if ia.IsZeroValue && ia2.IsZeroValue {
		ia.SetIntAryToZero(ia.Precision)
		return nil
	}

	compare := ia.CompareAbsoluteValues(ia2)
	isZeroResult := false

	// Largest Value in now in N1 slot
	newSignVal := ia.SignVal
	doAdd := false
	doReverseNums := false

	if compare == 1 {
		// compare == + 1
		// Absolute Value: N1 > N2

		if ia.SignVal == 1 && ia2.SignVal == 1 {
			doAdd = false
			newSignVal = 1
		} else if ia.SignVal == -1 && ia2.SignVal == 1 {
			doAdd = true
			newSignVal = -1
		} else if ia.SignVal == -1 && ia2.SignVal == -1 {
			doAdd = false
			newSignVal = -1
		} else {
			// Must Be ia.SignVal == 1 && ia2.SignVal == -1
			doAdd = true
			newSignVal = 1
		}

	} else if compare == -1 {
		// Absolute Values: N2 > N1
		if ia.SignVal == 1 && ia2.SignVal == 1 {
			doAdd = false
			doReverseNums = true
			newSignVal = -1
		} else if ia.SignVal == -1 && ia2.SignVal == 1 {
			doAdd = true
			newSignVal = -1
		} else if ia.SignVal == -1 && ia2.SignVal == -1 {
			doAdd = false
			doReverseNums = true
			newSignVal = 1
		} else {
			// Must Be ia.SignVal == 1 && ia2.SignVal == -1
			doAdd = true
			newSignVal = 1
		}

	} else {
		// Must be compare == 0
		// Absolute Values: N1==N2
		if ia.SignVal == 1 && ia2.SignVal == 1 {
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		} else if ia.SignVal == -1 && ia2.SignVal == 1 {
			doAdd = true
			newSignVal = -1
		} else if ia.SignVal == -1 && ia2.SignVal == -1 {
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		} else {
			// Must Be ia.SignVal == 1 && ia2.SignVal == -1
			doAdd = true
			newSignVal = 1
		}

	}

	return ia.addToSubtractFromThis(ia2, newSignVal, doAdd, isZeroResult, doReverseNums, convertToNumStr)
}
