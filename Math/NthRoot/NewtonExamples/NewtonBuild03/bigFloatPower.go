package NewtonBuild03

import (
	"math/big"
	"sync"
)

type BigFloatPower struct {
	lock *sync.Mutex
}

func (bigFloatPwr *BigFloatPower) RaiseToPower(base *big.Float, power int64) (*big.Float, error) {

	return new(big.Float), nil
}
