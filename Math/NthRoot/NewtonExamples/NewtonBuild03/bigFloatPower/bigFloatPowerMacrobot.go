package bigFloatPower

import (
	"math/big"

	ePref "github.com/MikeAustin71/errpref"
)

type bigFloatRaiseToPower struct{}

// raiseToPowerMul01
//
//		Raises a *big.Float value to a power using a multiplication
//		method.
//
//		This returned 'unroundedRaisedToPower' *big.Float has been
//		raised to the power specified by the 'powerInt64' parameter.
//		The decimal precision of the returned value is determined by
//		the 'maxInternalPrecisionUint' parameter.
//
//	 No rounding is performed. However, the returned value is
//	 accurate to the precision of the 'maxInternalPrecisionUint'
//	 parameter, and the rounding mode is set to 'mode'.
func (bigFloatPwr *bigFloatRaiseToPower) raiseToPowerMul01(
	baseBigFloat *big.Float,
	powerInt64 int64,
	maxInternalPrecisionUint uint,
	mode big.RoundingMode,
	errPrefDto *ePref.ErrPrefixDto) (unroundedRaisedToPower *big.Float, err error) {

	unroundedRaisedToPower = new(big.Float).
		SetPrec(maxInternalPrecisionUint).
		SetMode(mode)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigFloatPwr.raiseToPowerMul01",
		"")

	if err != nil {
		return unroundedRaisedToPower, err
	}

	if baseBigFloat == nil {

		return unroundedRaisedToPower,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'BigFloatPower'",
			}
	}

	var ok bool
	_,
		ok = unroundedRaisedToPower.SetString("1.0")

	if !ok {
		return unroundedRaisedToPower,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "_, ok = unroundedRaisedToPower.SetString(\"1.0\")",
				ErrContext: "'ok' == false",
				ErrMessage: "big.Float SetString() FAILED!\n" +
					"Returned 'ok' == false",
			}
	}

	if powerInt64 == 0 {
		return unroundedRaisedToPower, nil
	}

	var newBaseBigFloat = new(big.Float).Copy(baseBigFloat)

	if powerInt64 < 0 {

		unroundedRaisedToPower.Quo(unroundedRaisedToPower, baseBigFloat)

		powerInt64 = -1 * powerInt64
	}

	for i := int64(0); i < powerInt64; i++ {

		unroundedRaisedToPower.Mul(unroundedRaisedToPower, newBaseBigFloat)
	}

	return unroundedRaisedToPower, nil
}

// raiseToPowerIntBySquaring
//
//		Raises a *big.Float value to a power using the exponentiation by
//		squaring algorithm. This method is generally more efficient than
//		the standard multiplication method for large exponents.
//
//		The returned *big.Float value is raised to the power specified by
//		the 'powerInt64' parameter. The internal precision of the
//		calculation and the precision of the returned value are determined
//		by the 'maxInternalPrecisionUint' parameter.
//
//		Negative exponents are supported using the identity:
//		x^(-n) = 1 / (x^n).
//
//		No rounding is performed on the final result, but the calculation
//		accuracy is limited by the 'maxInternalPrecisionUint' value, and
//		the big.Float rounding mode is set to 'roundingMode'.
//
//	 Input parameter 'maxInternalPrecisionUint' describes the number of
//	 bits of mantissa precision that can be used in the calculation.
func (bigFloatPwr *bigFloatRaiseToPower) raiseToPowerIntBySquaring(
	baseBigFloat *big.Float,
	powerInt64 int64,
	maxInternalPrecisionUint uint,
	roundingMode big.RoundingMode,
	errPrefDto *ePref.ErrPrefixDto) (*big.Float, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"bigFloatRaiseToPower.raiseToPowerIntBySquaring",
		"")

	if err != nil {
		return new(big.Float).
				SetPrec(maxInternalPrecisionUint).
				SetMode(roundingMode).
				SetInt64(0),
			err
	}

	if baseBigFloat == nil {

		return new(big.Float).
				SetPrec(maxInternalPrecisionUint).
				SetMode(roundingMode).
				SetInt64(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'baseBigFloat'",
			}
	}

	// x^0 = 1
	if powerInt64 == 0 {
		return new(big.Float).
			SetPrec(maxInternalPrecisionUint).
			SetMode(roundingMode).
			SetInt64(1), nil
	}

	// Prepare result and base
	result := new(big.Float).
		SetPrec(maxInternalPrecisionUint).
		SetMode(roundingMode).
		SetInt64(1)

	base := new(big.Float).
		SetPrec(maxInternalPrecisionUint).
		SetMode(roundingMode).
		Copy(baseBigFloat)

	n := powerInt64

	// Negative exponent: x^(-n) = 1 / x^n
	if n < 0 {
		n = -n
		one := new(big.Float).
			SetPrec(maxInternalPrecisionUint).
			SetMode(roundingMode).
			SetInt64(1)
		_ = base.Quo(one, base) // base = 1/x
	}

	// Exponentiation by squaring
	for n > 0 {
		if n&1 == 1 {
			result.Mul(result, base)
		}
		base.Mul(base, base)
		n >>= 1
	}

	return result, nil
}

// mikeRaiseToPowerIntBySquaring
//
//		Raises a *big.Float value to a power using the exponentiation by
//		squaring algorithm. This method is generally more efficient than
//		the standard multiplication method for large exponents.
//
//		The returned *big.Float value is raised to the power specified by
//		the 'powerInt64' parameter. The internal precision of the
//		calculation and the precision of the returned value are determined
//		by the 'maxInternalPrecisionUint' parameter.
//
//		Negative exponents are supported using the identity:
//		x^(-n) = 1 / (x^n).
//
//		No rounding is performed on the final result, but the calculation
//		accuracy is limited by the 'maxInternalPrecisionUint' value, and
//		the big.Float rounding mode is set to 'roundingMode'.
//
//	 Input parameter 'maxInternalPrecisionUint' describes the number of
//	 bits of mantissa precision that can be used in the calculation.
func (bigFloatPwr *bigFloatRaiseToPower) mikeRaiseToPowerIntBySquaring(
	baseBigFloat *big.Float,
	powerInt64 int64,
	maxInternalPrecisionUint uint,
	roundingMode big.RoundingMode,
	errPrefDto *ePref.ErrPrefixDto) (*big.Float, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"bigFloatRaiseToPower.mikeRaiseToPowerIntBySquaring",
		"")

	if err != nil {
		return new(big.Float).
				SetPrec(maxInternalPrecisionUint).
				SetMode(roundingMode).
				SetInt64(0),
			err
	}

	if baseBigFloat == nil {

		return new(big.Float).
				SetPrec(maxInternalPrecisionUint).
				SetMode(roundingMode).
				SetInt64(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'baseBigFloat'",
			}
	}

	// x^0 = 1
	if powerInt64 == 0 {
		return new(big.Float).
			SetPrec(maxInternalPrecisionUint).
			SetMode(roundingMode).
			SetInt64(1), nil
	}

	// Prepare result and base
	result := new(big.Float).
		SetPrec(maxInternalPrecisionUint).
		SetMode(roundingMode).
		SetInt64(1)

	base := new(big.Float).
		SetPrec(maxInternalPrecisionUint).
		SetMode(roundingMode).
		Copy(baseBigFloat)

	n := powerInt64

	// Negative exponent: x^(-n) = 1 / x^n
	if n < 0 {

		n = -n

		one := new(big.Float).
			SetPrec(maxInternalPrecisionUint).
			SetMode(roundingMode).
			SetInt64(1)

		_ = base.Quo(one, base) // base = 1/x

		base.SetPrec(base.MinPrec())
	}

	// Exponentiation by squaring
	for n > 1 {

		if n%2 != 0 {
			// n is odd
			result.Mul(result, base)
		}

		base.Mul(base, base)

		n = n / 2

	}

	result.Mul(result, base)

	result.SetPrec(result.MinPrec())

	return result, nil
}

// mike2RaiseToPowerIntBySquaring
//
//		Raises a *big.Float value to a power using the exponentiation by
//		squaring algorithm. This method is generally more efficient than
//		the standard multiplication method for large exponents.
//
//		The returned *big.Float value is raised to the power specified by
//		the 'powerInt64' parameter. The internal precision of the
//		calculation and the precision of the returned value are determined
//		by the 'maxInternalPrecisionUint' parameter.
//
//		Negative exponents are supported using the identity:
//		x^(-n) = 1 / (x^n).
//
//		No rounding is performed on the final result, but the calculation
//		accuracy is limited by the 'maxInternalPrecisionUint' value, and
//		the big.Float rounding mode is set to 'roundingMode'.
//
//	 Input parameter 'maxInternalPrecisionUint' describes the number of
//	 bits of mantissa precision that can be used in the calculation.
func (bigFloatPwr *bigFloatRaiseToPower) mike2RaiseToPowerIntBySquaring(
	baseBigFloat *big.Float,
	powerInt64 int64,
	maxInternalPrecisionUint uint,
	roundingMode big.RoundingMode,
	errPrefDto *ePref.ErrPrefixDto) (*big.Float, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"bigFloatRaiseToPower.mikeRaiseToPowerIntBySquaring",
		"")

	if err != nil {
		return new(big.Float).
				SetPrec(maxInternalPrecisionUint).
				SetMode(roundingMode).
				SetInt64(0),
			err
	}

	if baseBigFloat == nil {

		return new(big.Float).
				SetPrec(maxInternalPrecisionUint).
				SetMode(roundingMode).
				SetInt64(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'baseBigFloat'",
			}
	}

	// x^0 = 1
	if powerInt64 == 0 {
		return new(big.Float).
			SetPrec(maxInternalPrecisionUint).
			SetMode(roundingMode).
			SetInt64(1), nil
	}

	// Prepare result and base
	result := new(big.Float).
		SetPrec(maxInternalPrecisionUint).
		SetMode(roundingMode).
		SetInt64(1)

	base := new(big.Float).
		SetPrec(maxInternalPrecisionUint).
		SetMode(roundingMode).
		Copy(baseBigFloat)

	n := powerInt64

	// Negative exponent: x^(-n) = 1 / x^n
	if n < 0 {

		n = -n

		one := new(big.Float).
			SetPrec(maxInternalPrecisionUint).
			SetMode(roundingMode).
			SetInt64(1)

		_ = base.Quo(one, base) // base = 1/x

		//base.SetPrec(base.MinPrec())
	}

	// Exponentiation by squaring
	for n > 1 {

		if n%2 != 0 {
			// n is odd
			result.Mul(result, base)
		}

		base.Mul(base, base)

		n = n / 2

	}

	result.Mul(result, base)

	// result.SetPrec(result.MinPrec())

	return result, nil
}
