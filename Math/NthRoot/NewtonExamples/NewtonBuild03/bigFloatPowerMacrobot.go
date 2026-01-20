package NewtonBuild03

import (
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigFloatRaiseToPower struct {
	lock *sync.Mutex
}

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

	if bigFloatPwr.lock == nil {
		bigFloatPwr.lock = new(sync.Mutex)
	}

	bigFloatPwr.lock.Lock()

	defer bigFloatPwr.lock.Unlock()

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
				ParameterName: "'bNum'",
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

	var newBaseBigFloat *big.Float = new(big.Float).Copy(baseBigFloat)

	if powerInt64 < 0 {

		unroundedRaisedToPower.Quo(unroundedRaisedToPower, baseBigFloat)

		powerInt64 = -1 * powerInt64
	}

	for i := int64(0); i < powerInt64; i++ {

		unroundedRaisedToPower.Mul(unroundedRaisedToPower, newBaseBigFloat)
	}

	return unroundedRaisedToPower, nil
}

func (bigFloatPwr *bigFloatRaiseToPower) raiseToPowerBySquaring() {

}
