package datetime

import (
	"errors"
	"math"
	"sync"
)

type numStrDtoHelper struct {
	lock *sync.Mutex
}


// multiplyNumStrs - Multiplies two NumStrDto instances and returns
// the product as a new NumStrDto instance.
//
// If either of the two input parameters, 'multiplicand' or
// 'multiplier' proves to be invalid, this method will return an
// error.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  multiplicand        *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE the values of internal member variables to
//       achieve the method's objectives.
//
//       The numerical value of 'multiplicand' will be multiplied
//       by input parameter 'multiplier' to generate and return
//       the product as NumStrDto instance.
//
//       If this NumStrDto instance proves to be invalid, this
//       method will return an error.
//
//
//  multiplier          *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE the values of internal member variables to
//       achieve the method's objectives.
//
//       The numerical value of 'multiplier' will be multiplied
//       by input parameter 'multiplicand' to generate and return
//       the product as NumStrDto instance.
//
//       If this NumStrDto instance proves to be invalid, this
//       method will return an error.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  product             NumStrDto
//     - If this method completes successfully, the product obtained from
//       multiplying input parameters 'multiplicand' and 'multiplier' will
//       be returned in a new instance of NumStrDto.
//
//
//  err                error
//     - If this method completes successfully, the returned error Type is
//       set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'. The 'ePrefix' text will be
//       prefixed to the beginning of the returned error message.
//
func (nStrDtoHelper *numStrDtoHelper) multiplyNumStrs(
	multiplicand *NumStrDto,
	multiplier *NumStrDto,
	ePrefix string) (
	product NumStrDto,
	err error) {


	if nStrDtoHelper.lock == nil {
		nStrDtoHelper.lock = new(sync.Mutex)
	}

	nStrDtoHelper.lock.Lock()

	defer nStrDtoHelper.lock.Unlock()

	ePrefix += "numStrDtoHelper.multiplyNumStrs() "

	err = nil

	nStrDtoElectron := numStrDtoElectron{}

	product = nStrDtoElectron.newBaseZeroNumStrDto(0)


	if multiplicand == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'multiplicand' is INVALID!\n" +
			"multiplicand has a 'nil' pointer!\n")
		return product, err
	}

	if multiplier == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'multiplier' is INVALID!\n" +
			"multiplier  has a 'nil' pointer!!\n")
		return product, err
	}

	_,
		err =
		nStrDtoElectron.testNumStrDtoValidity(
			multiplicand,
			ePrefix+ "multiplicand ")

	if err != nil {
		return product, err
	}

	_,
		err =
		nStrDtoElectron.testNumStrDtoValidity(
			multiplier,
			ePrefix+ "multiplier ")

	if err != nil {
		return product, err
	}


	lenN1AbsAllRunes := len(multiplicand.absAllNumRunes)
	lenN2AbsAllRunes := len(multiplier.absAllNumRunes)

	var n1Setup NumStrDto
	var n2Setup NumStrDto

	if lenN2AbsAllRunes > lenN1AbsAllRunes {

		n1Setup,
		err = nStrDtoElectron.copyOut(
					multiplier,
					ePrefix + "multiplier -> n1Setup ")

		if err != nil {
			return product, err
		}

		n2Setup,
			err = nStrDtoElectron.copyOut(
			multiplicand,
			ePrefix + "multiplicand -> n2Setup ")

		if err != nil {
			return product, err
		}

	} else if lenN1AbsAllRunes > lenN2AbsAllRunes {

		n1Setup,
			err = nStrDtoElectron.copyOut(
			multiplicand,
			ePrefix + "multiplicand -> n1Setup ")

		if err != nil {
			return product, err
		}

		n2Setup,
			err = nStrDtoElectron.copyOut(
			multiplier,
			ePrefix + "multiplier -> n2Setup ")

		if err != nil {
			return product, err
		}

	} else {
		// Must be lenN1AbsAllRunes == lenN2AbsAllRunes

		n1Setup,
			err = nStrDtoElectron.copyOut(
			multiplicand,
			ePrefix + "multiplicand -> n1Setup ")

		if err != nil {
			return product, err
		}

		n2Setup,
			err = nStrDtoElectron.copyOut(
			multiplier,
			ePrefix + "multiplier -> n2Setup ")

		if err != nil {
			return product, err
		}

	}

	newPrecision := n1Setup.precision + n2Setup.precision
	newSignVal := 1

	if n1Setup.signVal == n2Setup.signVal {
		newSignVal = 1
	} else {
		// Must be n1Setup.signVal != n2Setup.signVal
		newSignVal = -1
	}

	lenN1AbsAllRunes = len(n1Setup.absAllNumRunes)
	lenN2AbsAllRunes = len(n2Setup.absAllNumRunes)
	lenLevels := lenN2AbsAllRunes
	lenNumPlaces := (lenN1AbsAllRunes + lenN2AbsAllRunes) + 1

	intMAry := make([][]int, lenLevels)

	for i := 0; i < lenLevels; i++ {
		intMAry[i] = make([]int, lenNumPlaces)
	}

	intFinalAry := make([]int, lenNumPlaces+1)

	carry := 0
	levels := 0
	place := 0
	n1 := 0
	n2 := 0
	n3 := 0
	n4 := 0
	for i := lenN2AbsAllRunes - 1; i >= 0; i-- {

		place = (lenNumPlaces - 1) - levels

		for j := lenN1AbsAllRunes - 1; j >= 0; j-- {

			n1 = int(n1Setup.absAllNumRunes[j]) - 48
			n2 = int(n2Setup.absAllNumRunes[i]) - 48
			n3 = (n1 * n2) + carry
			n4 = int(math.Mod(float64(n3), 10.00))

			intMAry[levels][place] = n4

			carry = n3 / 10

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
	for i := 0; i < lenLevels; i++ {
		for j := lenNumPlaces - 1; j >= 0; j-- {

			n1 = intFinalAry[j+1]
			n2 = intMAry[i][j]
			n3 = n1 + n2 + carry
			n4 = 0

			if n3 > 9 {
				n4 = int(math.Mod(float64(n3), 10.0))
				carry = n3 / 10

			} else {
				n4 = n3
				carry = 0
			}

			intFinalAry[j+1] = n4
		}

		if carry > 0 {
			intFinalAry[0] = carry
		}

	}

	nStrDtoMech := numStrDtoMechanics{}

	product, err =
		nStrDtoMech.findIntArraySignificantDigitLimits(
		intFinalAry,
		newPrecision,
		newSignVal,
		ePrefix + "intFinalAry ")

	return product, err
}
