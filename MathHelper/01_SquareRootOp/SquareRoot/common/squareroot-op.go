package common

import (
	"errors"
	"math/big"
)

// http://www.basic-mathematics.com/square-root-algorithm.html
// https://www.youtube.com/watch?v=uIrjN2Onn8M

type SquareRootOp struct {
	InputPrecision        int
	InputArrayLen					int
	InputFirstIdx					int
	InputLastIdx					int
	InputIntegerLen				int
	BaseNumPairs          [][]int
	BasePairIdx           int
	ResultAry             IntAry
	ResultIdx             int
	ResultPrecision       int
	ResultMaxPrecision		int
	TempSubtractionResult *big.Int
	TempDivisionResult    *big.Int
}

func (sqr *SquareRootOp) Initialize(originalNum *IntAry,  maxPrecision int) error {


	sqr.InputPrecision = originalNum.Precision
	sqr.ResultMaxPrecision = maxPrecision
	originalNum.SetInternalFlags()

	sqr.InputArrayLen = originalNum.IntAryLen
	sqr.InputFirstIdx = originalNum.FirstDigitIdx
	sqr.InputLastIdx =  originalNum.LastDigitIdx
	sqr.InputIntegerLen = sqr.InputArrayLen - sqr.InputPrecision	- sqr.InputFirstIdx
	sqr.ResultAry =  IntAry{}.New()

	if sqr.InputArrayLen < 1 {
		return errors.New("SquareRootOp.Initialize() - ERROR: originalNum Array Length is ZERO!")
	}


	if sqr.InputArrayLen < 1 {
		return errors.New("SquareRootOp.Initialize() - ERROR: originalNum Array Length is ZERO!")
	}

	basePairs := sqr.InputArrayLen / 2
	basePairs += sqr.InputArrayLen - ((sqr.InputArrayLen / 2) * 2)
	sqr.BaseNumPairs = make([][]int, 0)

	for i:=0 ; i < basePairs; i++ {
		arrayPair := []int{0,0}
		sqr.BaseNumPairs = append(sqr.BaseNumPairs, arrayPair )
	}

	basePairIdx := basePairs - 1

	// sqr.BaseNumPairs  = append(sqr.BaseNumPairs, originalNum[0])

	//totalLen := sqr.InputIntegerLen + ((maxPrecision + 1) * 2)
	// 0 1 2 3


	for i:= sqr.InputArrayLen - 1; i >= 0; i-=2 {

		if i == 0 {
			sqr.BaseNumPairs[basePairIdx][1] = originalNum.IntAry[i]
		} else if i-1 >= 0 {
			sqr.BaseNumPairs[basePairIdx][1] = originalNum.IntAry[i]
			sqr.BaseNumPairs[basePairIdx][0] = originalNum.IntAry[i-1]
		}

		basePairIdx--
	}

	maxPrecision ++
	for j:=0; j < maxPrecision; j++ {
		arrayPair := []int {0,0}
		sqr.BaseNumPairs = append(sqr.BaseNumPairs, arrayPair)
	}

	maxPrecision--
	sqr.ResultPrecision = maxPrecision
	sqr.ResultAry = IntAry{}.New()
	sqr.ResultAry.Precision = sqr.ResultPrecision
	return nil
}



func (sqr *SquareRootOp) ComputeSquareRoot() error {
	const100 := big.NewInt(0).SetInt64(100)
	constTen := big.NewInt(0).SetInt64(10)

	err := sqr.FindNearestSquareLessThan()

	if err != nil {
		return err
	}

	// Compute First Subtraction Result
	twoDigitPair := big.NewInt(0).SetInt64(0)
	sqr.ResultIdx = 0

	// Compute First Square
	divisorBase := big.NewInt(0).SetInt64(int64(sqr.ResultAry.IntAry[0] * sqr.ResultAry.IntAry[0]))
	// Get First Pair
	sqr.BasePairIdx = 0
	twoDigitPair = big.NewInt(0).SetInt64(int64(sqr.BaseNumPairs[sqr.BasePairIdx][0] * 10))
	twoDigitPair = big.NewInt(0).Add(twoDigitPair, big.NewInt(0).SetInt64(int64(sqr.BaseNumPairs[sqr.BasePairIdx][1])) )

	// Subtract First Square from first pair (divisor = 36)
	sqr.TempSubtractionResult = big.NewInt(0).Sub(twoDigitPair, divisorBase)
	sqr.TempSubtractionResult = big.NewInt(0).Mul(sqr.TempSubtractionResult, const100)

	// Advance to Second Pair
	sqr.BasePairIdx++
	divisorBase = big.NewInt(0).SetInt64(int64(sqr.ResultAry.IntAry[0] + sqr.ResultAry.IntAry[0]))
	divisorBase = big.NewInt(0).Mul(divisorBase, constTen) // Now 120
	twoDigitPair = big.NewInt(0).SetInt64(int64(sqr.BaseNumPairs[sqr.BasePairIdx][0] * 10))
	twoDigitPair = big.NewInt(0).Add(twoDigitPair, big.NewInt(0).SetInt64(int64(sqr.BaseNumPairs[sqr.BasePairIdx][1])))
	sqr.TempSubtractionResult = big.NewInt(0).Add(sqr.TempSubtractionResult, twoDigitPair) // 306
	/// End Compute First Subtraction Result

	// This completes the first Subtraction Result
	// saved to sqr.TempSubtractionResult
	//sqr.TempDivisionResult = big.NewInt(0).SetInt64( int64((sqr.ResultAry.IntAry[0] * 2) * 10))
	bigI := big.NewInt(0).SetInt64(0)
	prevBigI := big.NewInt(0).SetInt64(0)
	n1 := big.NewInt(0).SetInt64(0)
	n2 := big.NewInt(0).SetInt64(0)
	n3 := big.NewInt(0).SetInt64(0)
	lenBasePairs := len(sqr.BaseNumPairs)
	for sqr.BasePairIdx < lenBasePairs {

		for i := int64(1); i < 11; i++ {
			bigI = big.NewInt(0).SetInt64(i)
			prevBigI = big.NewInt(0).SetInt64(i-1)
			n1 = big.NewInt(0).Add( divisorBase, bigI) // 120 + i
			n2 = big.NewInt(0).Mul(n1, bigI)

			if n2.Cmp(sqr.TempSubtractionResult) == 1 {
				n2_2 := big.NewInt(0).Add(divisorBase, prevBigI)
				n2_3 := big.NewInt(0).Mul(n2_2, prevBigI)
				n2_4 := big.NewInt(0).Sub(sqr.TempSubtractionResult, n2_3)
				n3 = big.NewInt(0).Mul(n2_4, const100)
				sqr.BasePairIdx++
				if sqr.BasePairIdx >= lenBasePairs {
					sqr.ResultAry.ConvertIntAryToNumStr()
					return nil
				}

				twoDigitPair = big.NewInt(0).SetInt64(int64(sqr.BaseNumPairs[sqr.BasePairIdx][0] * 10))
				twoDigitPair = big.NewInt(0).Add(twoDigitPair, big.NewInt(0).SetInt64(int64(sqr.BaseNumPairs[sqr.BasePairIdx][1])))
				// New Subtraction Result
				sqr.TempSubtractionResult = big.NewInt(0).Add(n3, twoDigitPair)

				sqr.ResultAry.IntAry = append(sqr.ResultAry.IntAry, int(prevBigI.Int64()))
				sqr.ResultIdx = len(sqr.ResultAry.IntAry) - 1
				// New Divisor Base
				divisorBase = big.NewInt(0).Add(n2_2, prevBigI) // 124
				divisorBase = big.NewInt(0).Mul(divisorBase, constTen)
				// Break out of loop
				i=99
			}
		}
	}

	sqr.ResultAry.ConvertIntAryToNumStr()

	return nil

}


func (sqr *SquareRootOp) FindNearestSquareLessThan() (error) {


	if sqr.BaseNumPairs[0][0] < 0 || sqr.BaseNumPairs[0][1] < 0 {
		return errors.New("FindNearestSquareLessThan() - ERROR: twoDigit is less than Zero!")
	}
	// Acquire Pair Zero
	twoDigit := sqr.BaseNumPairs[0][0] * 10
	twoDigit += sqr.BaseNumPairs[0][1]

	if twoDigit == 0 {
		sqr.ResultAry.IntAry = append(sqr.ResultAry.IntAry, 0)
		return nil
	}

	// The square which is less than or equal to Pair Zero
	for i:=1 ; i < 10; i++ {

		if i * i > twoDigit {
			sqr.ResultAry.IntAry = append(sqr.ResultAry.IntAry, i-1)
			return nil
		}
	}

	return errors.New("FindNearestSquareLessThan() - ERROR: Could NOT locate nearest square less than first base pair!")
}