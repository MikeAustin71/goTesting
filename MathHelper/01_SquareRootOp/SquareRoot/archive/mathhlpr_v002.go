package archive

import (
	"errors"
)

// http://www.basic-mathematics.com/square-root-algorithm.html
// https://www.youtube.com/watch?v=uIrjN2Onn8M

type SquareRootOp struct {
	InputPrecision        int
	InputArrayLen					int
	InputFirstIdx					int
	InputLastIdx					int
	InputIntegerLen				int
	BaseNumPairs          []int
	BasePairIdx           int
	ResultAry             IntAry
	ResultIdx             int
	ResultPrecision       int
	ResultMaxPrecision		int
	TempSubtractionResult int
	TempDivisionResult    int
}

func (sqr *SquareRootOp) Initialize(originalNum *IntAry,  maxPrecision int) error {


	sqr.InputPrecision = originalNum.Precision
	sqr.ResultMaxPrecision = maxPrecision
	originalNum.SetInternalFlags()

	sqr.InputArrayLen = originalNum.IntAryLen
	sqr.InputFirstIdx = originalNum.FirstDigitIdx
	sqr.InputLastIdx =  originalNum.LastDigitIdx
	sqr.InputIntegerLen = sqr.InputArrayLen - sqr.InputPrecision	- sqr.InputFirstIdx

	if sqr.InputArrayLen < 1 {
		return errors.New("SquareRootOp.Initialize() - ERROR: originalNum Array Length is ZERO!")
	}


	if sqr.InputArrayLen < 1 {
		return errors.New("SquareRootOp.Initialize() - ERROR: originalNum Array Length is ZERO!")
	}

	start := 0

	if (sqr.InputIntegerLen / 2) * 2 != sqr.InputIntegerLen {
		//sqr.InputIntegerLen--
		start = 1
		sqr.BaseNumPairs = make([]int, 0, sqr.InputIntegerLen + (maxPrecision * 2) )
		sqr.BaseNumPairs  = append(sqr.BaseNumPairs, originalNum.IntAry[0])
	} else {
		sqr.BaseNumPairs = make([]int, 0, sqr.InputIntegerLen + (maxPrecision * 2) )
	}


	// sqr.BaseNumPairs  = append(sqr.BaseNumPairs, originalNum[0])
	n1:= 0
	totalLen := sqr.InputIntegerLen + ((maxPrecision + 1) * 2)
	// 0 1 2 3 4
	for i:= start; i < totalLen; i+=2 {

		if i < sqr.InputIntegerLen {
			n1 = originalNum.IntAry[i] * 10
		} else {
			n1 = 0
		}


		if i+1 < sqr.InputIntegerLen {
			n1+= originalNum.IntAry[i+1]
		}

		sqr.BaseNumPairs  = append(sqr.BaseNumPairs, n1)

	}

	sqr.ResultPrecision = maxPrecision
	sqr.ResultAry = IntAry{}.New()
	sqr.ResultAry.Precision = sqr.ResultPrecision
	return nil
}


func (sqr *SquareRootOp) ComputeSquareRoot() error {

	err := sqr.FindNearestSquareLessThan()

	if err != nil {
		return err
	}

	sqr.ResultIdx = 0
	n1 := sqr.ResultAry.IntAry[0] * sqr.ResultAry.IntAry[0]
	sqr.BasePairIdx = 0
	sqr.TempSubtractionResult = sqr.BaseNumPairs[sqr.BasePairIdx] - n1
	sqr.TempSubtractionResult = sqr.TempSubtractionResult * 100
	sqr.BasePairIdx++
	sqr.TempSubtractionResult += sqr.BaseNumPairs[sqr.BasePairIdx]

	// This completes the first Subtraction Result
	// saved to sqr.TempSubtractionResult
	sqr.TempDivisionResult = (sqr.ResultAry.IntAry[0] * 2) * 10

	n2 := 0
	n3 := 0
	lenBasePairs := len(sqr.BaseNumPairs)
	for sqr.BasePairIdx < lenBasePairs {

		for i := 1; i < 1000; i++ {
			n1 = sqr.TempDivisionResult + i
			n2 = n1 * i
			if n2 > sqr.TempSubtractionResult {
				n2 = (sqr.TempDivisionResult+(i-1)) * (i - 1)
				n3 = (sqr.TempSubtractionResult - n2) * 100
				sqr.BasePairIdx++
				if sqr.BasePairIdx >= lenBasePairs {
					sqr.ResultAry.ConvertIntAryToNumStr()
					return nil
				}
				n3 += sqr.BaseNumPairs[sqr.BasePairIdx]
				sqr.TempSubtractionResult = n3
				sqr.ResultAry.IntAry = append(sqr.ResultAry.IntAry, i-1)
				sqr.ResultIdx = len(sqr.ResultAry.IntAry) - 1
				sqr.TempDivisionResult = 0
				n1 = sqr.ResultAry.IntAry[0]
				for j := 1; j <= sqr.ResultIdx; j++ {
					n1 = (n1 * 10) + sqr.ResultAry.IntAry[j]
				}
				sqr.TempDivisionResult = (n1 * 2) * 10
				break
			}
		}
	}

	sqr.ResultAry.ConvertIntAryToNumStr()

	return nil

}


func (sqr *SquareRootOp) FindNearestSquareLessThan() (error) {

	twoDigit := sqr.BaseNumPairs[0]

	if twoDigit < 0 {
		return errors.New("FindNearestSquareLessThan() - ERROR: twoDigit is less than Zero!")
	}


	if twoDigit == 0 {
		sqr.ResultAry.IntAry = append(sqr.ResultAry.IntAry, 0)
		return nil
	}

	for i:=1 ; i < 10; i++ {

		if i * i > twoDigit {
			sqr.ResultAry.IntAry = append(sqr.ResultAry.IntAry, i-1)
			return nil
		}
	}

	return errors.New("FindNearestSquareLessThan() - ERROR: Could NOT locate nearest square less than first base pair!")
}