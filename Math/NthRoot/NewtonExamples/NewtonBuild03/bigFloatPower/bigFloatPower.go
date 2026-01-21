package bigFloatPower

import (
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type BigFloatPower struct {
	lock *sync.Mutex
}

func (bigFloatPwr *BigFloatPower) RaiseToPower(
	baseBigFloat *big.Float,
	powerInt64 int64,
	maxInternalPrecisionUint uint,
	roundingMode big.RoundingMode) (*big.Float, error) {

	if bigFloatPwr.lock == nil {
		bigFloatPwr.lock = new(sync.Mutex)
	}

	bigFloatPwr.lock.Lock()

	defer bigFloatPwr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigFloatPower.RaiseToPower",
		"")

	if err != nil {
		return new(big.Float), err
	}

	if baseBigFloat == nil {

		return new(big.Float),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'baseBigFloat'",
			}
	}

	return new(bigFloatRaiseToPower).raiseToPowerIntBySquaring(
		baseBigFloat,
		powerInt64,
		maxInternalPrecisionUint,
		roundingMode,
		ePrefix)
}
