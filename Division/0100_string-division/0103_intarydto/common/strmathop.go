package common

import (
	"fmt"
	"errors"
)

// StrMathOp - Used to perform string
// based arithmetic.
//
// Dependencies:
// 	IntAry
//
type StrMathOp struct {
	N1       IntAry
	N2       IntAry
	N3       IntAry
	IntMAry  [][]int
	IFinal   IntAry
	Dividend IntAry
	Divisor  IntAry
	Quotient IntAry
	Modulo   IntAry
}

func (sMathOp StrMathOp) New() StrMathOp {
	iAry := StrMathOp{}
	iAry.N1 = IntAry{}.New()
	iAry.N2 = IntAry{}.New()
	iAry.N3 = IntAry{}.New()
	iAry.IntMAry = make([][]int, 0)
	iAry.IFinal = IntAry{}.New()
	iAry.Dividend = IntAry{}.New()
	iAry.Divisor = IntAry{}.New()
	iAry.Quotient = IntAry{}.New()
	iAry.Modulo = IntAry{}.New()
	return iAry
}

func (sMathOp *StrMathOp) Empty() {

	sMathOp.N1 = IntAry{}.New()
	sMathOp.N2 = IntAry{}.New()
	sMathOp.N3 = IntAry{}.New()
	sMathOp.IntMAry = make([][]int, 0)
	sMathOp.IFinal = IntAry{}.New()
	sMathOp.Dividend = IntAry{}.New()
	sMathOp.Divisor = IntAry{}.New()
	sMathOp.Quotient = IntAry{}.New()
	sMathOp.Modulo = IntAry{}.New()

}

// AddN1N2 - Adds the values in the N1 and N2 arrays
// and returns the sum in the IFinal Array. Arrays
// N1 and N2 must be first correctly populated before
// calling this method.
func (sMathOp *StrMathOp) AddN1N2() error {

	sMathOp.N1.SetEqualArrayLengths(&sMathOp.N2)

	if sMathOp.N1.IsZeroValue && sMathOp.N2.IsZeroValue {
		sMathOp.IFinal.SetIntAryToZero(sMathOp.N1.Precision)
		return nil
	}

	compare := sMathOp.N1.CompareAbsoluteValues(&sMathOp.N2)

	newSignVal := sMathOp.N1.SignVal
	doAdd := true
	isZeroResult := false
	doReverseNums := false

	if compare == 1 {
		// compare == + 1
		// Absolute Value: N1 > N2

		if sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == 1 {
			doAdd = true
			newSignVal = 1
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == 1 {
			doAdd = false
			newSignVal = -1
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == -1 {
			doAdd = true
			newSignVal = -1
		} else {
			// Must Be sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == -1
			doAdd = false
			newSignVal = 1
		}

	} else if compare == -1 {
		// Absolute Values: N2 > N1
		if sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == 1 {
			doAdd = true
			newSignVal = 1
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == 1 {
			doAdd = false
			doReverseNums = true
			newSignVal = 1
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == -1 {
			doAdd = true
			newSignVal = -1
		} else {
			// Must Be sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == -1
			doAdd = false
			doReverseNums = true
			newSignVal = -1
		}

	} else {
		// Must be compare == 0
		// Absolute Values: N1==N2
		if sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == 1 {
			doAdd = true
			newSignVal = 1
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == 1 {
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == -1 {
			doAdd = true
			newSignVal = -1
		} else {
			// Must Be sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == -1
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		}

	}

	if isZeroResult {
		sMathOp.IFinal.SetIntAryToZero(sMathOp.N1.Precision)
		return nil
	}
	
	sMathOp.IFinal.IntAry = make([]int, sMathOp.N1.IntAryLen)
	sMathOp.IFinal.SignVal = newSignVal
	sMathOp.IFinal.Precision = sMathOp.N1.Precision
	// Array Lengths of N1 and N2 are now equal

	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0

	for j := sMathOp.N1.IntAryLen - 1; j >= 0; j-- {

		if doReverseNums {

			n2 = sMathOp.N1.IntAry[j]
			n1 = sMathOp.N2.IntAry[j]

		} else {
			n1 = sMathOp.N1.IntAry[j]
			n2 = sMathOp.N2.IntAry[j]

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

		sMathOp.IFinal.IntAry[j] = n3

	}

	if carry > 0 {
		sMathOp.IFinal.IntAry = append([]int{1}, sMathOp.IFinal.IntAry...)
		sMathOp.IFinal.IntAryLen++
	}

	if sMathOp.IFinal.IntAry[0] == 0 {
		sMathOp.IFinal.SetSignificantDigitIdxs()
		sMathOp.IFinal.IntAry = sMathOp.IFinal.IntAry[sMathOp.IFinal.FirstDigitIdx:]
	}

	sMathOp.IFinal.ConvertIntAryToNumStr()

	return nil
}

// RaiseToPower - Raises the value of sMathOp.N1 Int Array
// to the power specified by the 'power' parameter passed
// to this method.
//								N1 ^ power
// Before calling this method, sMathOp.N1 Int Array must
// be set to the desired value.
func (sMathOp *StrMathOp) RaiseToPower(power int) error {

	if power < 0 {
		return fmt.Errorf("Error: Power is less than ZERO! power= '%v'", power)
	}

	newPrecision := sMathOp.N1.Precision * power
	newSignVal := sMathOp.N1.SignVal

	if power == 0 {

		sMathOp.IFinal.SetIntAryToOne(sMathOp.N1.Precision)

		return nil
	}

	sMathOp.N1.SetIntAryLength()
	sMathOp.N1.SetIsZeroValue()

	if sMathOp.N1.IsZeroValue {
		sMathOp.IFinal.SetIntAryToZero(sMathOp.N1.Precision)
		return nil
	}

	if power == 1 {

		sMathOp.IFinal = sMathOp.N1.CopyOut()
		return nil
	}

	power--

	sMathOp.N2.IntAry = make([]int, sMathOp.N1.IntAryLen)

	for i := 0; i < sMathOp.N1.IntAryLen; i++ {
		sMathOp.N2.IntAry[i] = sMathOp.N1.IntAry[i]
	}

	for i := 0; i < power; i++ {

		sMathOp.MultiplyN1N2()

		sMathOp.N1 = sMathOp.IFinal

	}

	sMathOp.IFinal.Precision = newPrecision
	sMathOp.IFinal.SignVal = newSignVal
	sMathOp.IFinal.SetIntAryLength()
	sMathOp.IFinal.SetIsZeroValue()
	sMathOp.IFinal.ConvertIntAryToNumStr()
	return nil
}

func (sMathOp *StrMathOp) MultiplyN1N2() error {

	sMathOp.N1.SetIntAryLength()
	sMathOp.N1.SetIsZeroValue()
	sMathOp.N2.SetIntAryLength()
	sMathOp.N2.SetIsZeroValue()

	newPrecision := sMathOp.N1.Precision + sMathOp.N2.Precision

	if sMathOp.N1.IsZeroValue || sMathOp.N2.IsZeroValue {
		sMathOp.IFinal.SetIntAryToZero(newPrecision)
		return nil
	}

	newSignVal := 1

	if sMathOp.N1.SignVal != sMathOp.N2.SignVal {
		newSignVal = -1
	}
	lenLevels := sMathOp.N2.IntAryLen
	lenNumPlaces := (sMathOp.N1.IntAryLen + sMathOp.N2.IntAryLen) + 1

	sMathOp.IntMAry = make([][]int, lenLevels)

	for i := 0; i < lenLevels; i++ {
		sMathOp.IntMAry[i] = make([]int, lenNumPlaces)
	}

	sMathOp.IFinal.IntAry = make([]int, lenNumPlaces+1)

	carry := 0
	levels := 0
	place := 0
	n1 := 0
	n2 := 0
	n3 := 0
	n4 := 0
	x := 0

	for i := sMathOp.N2.IntAryLen - 1; i >= 0; i-- {

		place = (lenNumPlaces - 1) - levels

		for j := sMathOp.N1.IntAryLen - 1; j >= 0; j-- {

			n1 = sMathOp.N1.IntAry[j]
			n2 = sMathOp.N2.IntAry[i]
			n3 = (n1 * n2) + carry
			x = n3 / 10
			n4 = n3 - (x * 10)
			// n4 = int(math.Mod(float64(n3), float64(10.00)))

			sMathOp.IntMAry[levels][place] = n4

			carry = x

			place--
		}

		sMathOp.IntMAry[levels][place] = carry
		carry = 0
		levels++
	}

	carry = 0
	n1 = 0
	n2 = 0
	n3 = 0
	n4 = 0
	x = 0

	for i := 0; i < lenLevels; i++ {
		for j := lenNumPlaces - 1; j >= 0; j-- {

			n1 = sMathOp.IFinal.IntAry[j+1]
			n2 = sMathOp.IntMAry[i][j]
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

			sMathOp.IFinal.IntAry[j+1] = n4
		}

		if carry > 0 {
			sMathOp.IFinal.IntAry[0] = carry
		}

	}

	sMathOp.IFinal.SignVal = newSignVal
	sMathOp.IFinal.Precision = newPrecision
	sMathOp.IFinal.OptimizeIntArrayLen(false, true)

	return nil
}


// Divide - Divides the Dividend IntAry
// field by the Divisor IntAry field. The results are
// stored int IntAry fields Quotient and Modulo
func (sMathOp *StrMathOp) Divide(maxPrecision int) error {

	sMathOp.Quotient.SetIntAryToZero(0)
	sMathOp.Modulo.SetIntAryToZero(0)
	tensCount := IntAry{}.New()
	tensCount.SetIntAryToOne(0)

	newSignVal := 1

	if sMathOp.Divisor.SignVal != sMathOp.Dividend.SignVal {
		newSignVal = -1
	}

	if sMathOp.Divisor.SignVal == -1 {
		sMathOp.Divisor.SignVal = 1
		sMathOp.Divisor.ConvertIntAryToNumStr()
	}

	if sMathOp.Dividend.SignVal == -1 {
		sMathOp.Dividend.SignVal = 1
		sMathOp.Dividend.ConvertIntAryToNumStr()
	}

	sMathOp.Divisor.SetIsZeroValue()
	if sMathOp.Divisor.IsZeroValue {
		return errors.New("Divisor is ZERO. Cannot Divide By Zero!!")
	}

	sMathOp.Dividend.SetIsZeroValue()

	if sMathOp.Dividend.IsZeroValue {
		return nil
	}

	trialDividend := sMathOp.Dividend.CopyOut()

	dividendMag := sMathOp.Dividend.GetMagnitude()
	divisorMag := sMathOp.Divisor.GetMagnitude()
	deltaMag := uint(0)
	incrementVal := IntAry{}.New()
	incrementVal.SetIntAryWithNumStr(sMathOp.Divisor.NumStr)

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
	doLoop := true
	precisionCutOff := maxPrecision + dividendMag + 1

	for doLoop {

		if sMathOp.Quotient.Precision == precisionCutOff {
			sMathOp.Quotient.SignVal = newSignVal
			sMathOp.Quotient.RoundToPrecision(maxPrecision)
			return nil
		}

		compare = incrementVal.CompareAbsoluteValues(&trialDividend)

		if compare == 0 {
			// incrementalVal is equal to trialDividend
			sMathOp.Quotient.AddToThis(&tensCount, false)
			sMathOp.Quotient.SignVal = newSignVal
			sMathOp.Quotient.ConvertIntAryToNumStr()
			return nil

		} else if compare == -1 {
			// incrementalVal < trialDividend
			sMathOp.Quotient.AddToThis(&tensCount, false)

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

	return nil
}

// DivideDividendByDivisor - Divides the Dividend IntAry
// field by the Divisor IntAry field. The results are
// stored int IntAry fields Quotient and
func (sMathOp *StrMathOp) DivideDividendByDivisor() error {

	sMathOp.N1 = sMathOp.Divisor
	compare := -1
	quotient := 0
	sN := fmt.Sprintf("%v", quotient)
	sMathOp.N2.SetIntAryWithNumStr(sN)

	for compare < 1 {
		quotient++
		sMathOp.N2.IncrementIntegerOne()
		sMathOp.N3 = sMathOp.IFinal
		sMathOp.MultiplyN1N2()
		compare = sMathOp.IFinal.CompareSignedValues(&sMathOp.Dividend)

	}
	quotient = quotient - 1
	sN = fmt.Sprintf("%v", quotient)
	sMathOp.Quotient.SetIntAryWithNumStr(sN)
	sMathOp.Quotient.ConvertIntAryToNumStr()

	if compare == 0 {
		sMathOp.Modulo.SetIntAryToZero(0)
		return nil
	}

	sMathOp.N1 = sMathOp.Dividend
	sMathOp.N2 = sMathOp.N3
	sMathOp.SubtractN1N2()
	sMathOp.Modulo = sMathOp.IFinal
	sMathOp.Modulo.ConvertIntAryToNumStr()
	return nil
}

func (sMathOp *StrMathOp) DivideBySubtraction() {

	sMathOp.Modulo = sMathOp.Dividend.CopyOut()
	sMathOp.Quotient.SetIntAryToZero(0)
	compare := 1

	for compare >= 0 {
		sMathOp.Modulo.SubtractFromThis(&sMathOp.Divisor, false)
		sMathOp.Quotient.IncrementIntegerOne()
		compare = sMathOp.Modulo.CompareSignedValues(&sMathOp.Divisor)
	}

	sMathOp.Modulo.ConvertIntAryToNumStr()
	sMathOp.Quotient.ConvertIntAryToNumStr()

	return
}

func (sMathOp *StrMathOp) SubtractDivArys()  {



}

// SubtractN1N2 - Subtracts N2 IntAry from
// N1 IntAry. The result is returned in the IFinal
// IntAry.
//
func (sMathOp *StrMathOp) SubtractN1N2() error {

	sMathOp.N1.SetEqualArrayLengths(&sMathOp.N2)

	if sMathOp.N1.IsZeroValue && sMathOp.N2.IsZeroValue {
		sMathOp.IFinal.SetIntAryToZero(sMathOp.N1.Precision)
		return nil
	}

	compare := sMathOp.N1.CompareAbsoluteValues(&sMathOp.N2)
	isZeroResult := false

	// Largest Value in now in N1 slot
	newSignVal := sMathOp.N1.SignVal
	doSubtract := true
	doReverseNums := false

	if compare == 1 {
		// compare == + 1
		// Absolute Value: N1 > N2

		if sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == 1 {
			doSubtract = true
			newSignVal = 1
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == 1 {
			doSubtract = false
			newSignVal = -1
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == -1 {
			doSubtract = true
			newSignVal = -1
		} else {
			// Must Be sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == -1
			doSubtract = false
			newSignVal = 1
		}

	} else if compare == -1 {
		// Absolute Values: N2 > N1
		if sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == 1 {
			doSubtract = true
			doReverseNums = true
			newSignVal = -1
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == 1 {
			doSubtract = false
			newSignVal = -1
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == -1 {
			doSubtract = true
			doReverseNums = true
			newSignVal = 1
		} else {
			// Must Be sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == -1
			doSubtract = false
			newSignVal = 1
		}

	} else {
		// Must be compare == 0
		// Absolute Values: N1==N2
		if sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == 1 {
			doSubtract = true
			newSignVal = 1
			isZeroResult = true
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == 1 {
			doSubtract = false
			newSignVal = -1
		} else if sMathOp.N1.SignVal == -1 && sMathOp.N2.SignVal == -1 {
			doSubtract = true
			newSignVal = 1
			isZeroResult = true
		} else {
			// Must Be sMathOp.N1.SignVal == 1 && sMathOp.N2.SignVal == -1
			doSubtract = false
			newSignVal = 1
		}

	}

	if isZeroResult {
		sMathOp.IFinal.SetIntAryToZero(sMathOp.N1.Precision)
		return nil
	}

	sMathOp.IFinal.IntAry = make([]int, sMathOp.N1.IntAryLen)
	sMathOp.IFinal.SignVal = newSignVal
	sMathOp.IFinal.Precision = sMathOp.N1.Precision
	// Array Lengths of N1 and N2 are now equal

	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0
	for j := sMathOp.N1.IntAryLen - 1; j >= 0; j-- {

		if doReverseNums {
			n1 = sMathOp.N2.IntAry[j]
			n2 = sMathOp.N1.IntAry[j]

		} else {
			n1 = sMathOp.N1.IntAry[j]
			n2 = sMathOp.N2.IntAry[j]

		}

		if !doSubtract {
			// doSubtract == false
			// Do Addition

			n3 = n1 + n2 + carry

			if n3 > 9 {
				n3 = n1 + n2 + carry - 10
				carry = 1

			} else {
				carry = 0
			}

		} else {
			// doSubtract == true
			// Do Subtraction
			n3 = n1 - n2 - carry

			if n3 < 0 {
				n3 = n1 + 10 - n2 - carry
				carry = 1
			} else {
				carry = 0
			}
		}

		sMathOp.IFinal.IntAry[j] = n3

	}

	if carry > 0 {
		sMathOp.IFinal.IntAry = append([]int{1}, sMathOp.IFinal.IntAry...)
		sMathOp.IFinal.IntAryLen++
	}

	if sMathOp.IFinal.IntAry[0] == 0 {
		sMathOp.IFinal.SetSignificantDigitIdxs()
		sMathOp.IFinal.IntAry = sMathOp.IFinal.IntAry[sMathOp.IFinal.FirstDigitIdx:]
	}

	sMathOp.IFinal.ConvertIntAryToNumStr()

	return nil
}
