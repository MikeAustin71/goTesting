package NewtonBuild03

import (
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type BigFloatPower struct {
	lock *sync.Mutex
}

func (bigFloatPwr *BigFloatPower) RaiseToPower(
	base *big.Float,
	power int64,
	mode big.RoundingMode) (*big.Float, error) {

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

	if base == nil {

		return new(big.Float),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	return new(big.Float), nil
}
