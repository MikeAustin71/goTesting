package bigIntEx

import "math/big"

type BigIntMathUtility struct{}

func (bigINumUtil *BigIntMathUtility) BigIntInverse(numBigInt *big.Int) (*big.Int, *big.Int, error) {

  ePrefix := "BigIntMathUtility.BigIntInverse() "

  if numBigInt == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix,
        ParameterName: "'numBigInt'",
      }
  }

  numDivisorBigInt := new(big.Int).Set(numBigInt)

  oneDivisorBigInt := big.NewInt(1)
  quotientBigInt := big.NewInt(0)
  remainderBigInt := big.NewInt(0)

  quotientBigInt, remainderBigInt = big.NewInt(0).QuoRem(oneDivisorBigInt, numDivisorBigInt, remainderBigInt)

  return quotientBigInt, remainderBigInt, nil
}
