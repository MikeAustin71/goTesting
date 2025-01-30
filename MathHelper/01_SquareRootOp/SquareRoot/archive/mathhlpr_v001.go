package archive

/*
import (
	"errors"
)

// http://www.basic-mathematics.com/square-root-algorithm.html
// https://www.youtube.com/watch?v=uIrjN2Onn8M

type SquareRootOp struct {
	InputArray            []int
	InputPrecision        int
	InputArrayLen					int
	InputFirstIdx					int
	InputLastIdx					int
	InputIntegerLen				int
	BaseNumPairs          []int
	BasePairIdx           int
	ResultAry             []int
	ResultIdx             int
	ResultPrecision       int
	ResultMaxPrecision		int
	TempSubtractionResult int
	TempDivisionResult    int
}

func (sqr *SquareRootOp) Initialize(originalNum []int, originalPrecision int,  maxPrecision int) error {


	sqr.InputArray = originalNum[0:]
	sqr.InputPrecision = 0
	sqr.ResultMaxPrecision = maxPrecision
	sqr.CalcInputArrayStats()

	if sqr.InputArrayLen < 1 {
		return errors.New("SquareRootOp.Initialize() - ERROR: originalNum Array Length is ZERO!")
	}


	if sqr.InputArrayLen < 1 {
		return errors.New("SquareRootOp.Initialize() - ERROR: originalNum Array Length is ZERO!")
	}

	sqr.BaseNumPairs = make([]int, 0, sqr.InputIntegerLen + (maxPrecision * 2) )
	// sqr.BaseNumPairs  = append(sqr.BaseNumPairs, originalNum[0])
	n1:= 0
	totalLen := sqr.InputIntegerLen + ((maxPrecision + 1) * 2)
	// 0 1 2 3 4 5 6 7
	for i:= 0; i < totalLen; i+=2 {

		if i < sqr.InputIntegerLen {
			n1 = originalNum[i] * 10
		} else {
			n1 = 0
		}


		if i+1 < sqr.InputIntegerLen {
			n1+= originalNum[i+1]
		}

		sqr.BaseNumPairs  = append(sqr.BaseNumPairs, n1)

	}

	sqr.ResultPrecision = maxPrecision
	return nil
/*


	sqr.InputArray = originalNum[0:]
	sqr.InputPrecision = 0
	sqr.ResultMaxPrecision = maxPrecision
	sqr.CalcInputArrayStats()

	if sqr.InputArrayLen < 1 {
		return errors.New("SquareRootOp.Initialize() - ERROR: originalNum Array Length is ZERO!")
	}


	n1 := 0

	if sqr.InputIntegerLen == 1 {
		n1 = 1 + maxPrecision
		sqr.BaseNumPairs = make([]int, n1, n1 + 50)
		sqr.BaseNumPairs[0] =  originalNum[0]
		sqr.ResultPrecision = maxPrecision
		return nil
	}


	arryStart := 0

	sqr.BaseNumPairs = make([]int, 0, n1 + 50)

	if (sqr.InputIntegerLen / 2) * 2 != sqr.InputIntegerLen {
		arryStart = 1
		sqr.BaseNumPairs = append(sqr.BaseNumPairs, sqr.InputArray[0])
	}

	sqr.ResultAry = make([]int, 0, (n1 * 2) + 50)

	// 0 1 2 3 4 5 6
	for i:= arryStart; i < sqr.InputArrayLen; i+=2 {

		n1 = originalNum[i] * 10

		if i+1 != sqr.InputArrayLen {
			n1+= originalNum[i+1]
		}

		sqr.BaseNumPairs = append(sqr.BaseNumPairs, sqr.InputArray[0])

	}

	for j:=0; j< maxPrecision; j++ {

		sqr.BaseNumPairs = append(sqr.BaseNumPairs, 0)
	}

	sqr.ResultPrecision = maxPrecision

	return nil

}


func (sqr *SquareRootOp) ComputeSquareRoot() error {

	err := sqr.FindNearestSquareLessThan()

	if err != nil {
		return err
	}

	sqr.ResultIdx = 0
	n1 := (sqr.ResultAry[0] * sqr.ResultAry[0])
	sqr.BasePairIdx = 0
	sqr.TempSubtractionResult = sqr.BaseNumPairs[sqr.BasePairIdx] - n1
	sqr.TempSubtractionResult = sqr.TempSubtractionResult * 100
	sqr.BasePairIdx++
	sqr.TempSubtractionResult += sqr.BaseNumPairs[sqr.BasePairIdx]

	// This completes the first Subtraction Result
	// saved to sqr.TempSubtractionResult
	sqr.TempDivisionResult = (sqr.ResultAry[0] * 2) * 10

	n2 := 0
	n3 := 0
	lenBasePairs := len(sqr.BaseNumPairs)
	for sqr.BasePairIdx < lenBasePairs {

		for i := 1; i < 1000000; i++ {
			n1 = sqr.TempDivisionResult + i
			n2 = n1 * i
			if n2 > sqr.TempSubtractionResult {
				n2 = (sqr.TempDivisionResult+(i-1)) * (i - 1)
				n3 = (sqr.TempSubtractionResult - n2) * 100
				sqr.BasePairIdx++
				if sqr.BasePairIdx >= lenBasePairs {
					return nil
				}
				n3 += sqr.BaseNumPairs[sqr.BasePairIdx]
				sqr.TempSubtractionResult = n3
				sqr.ResultAry = append(sqr.ResultAry, (i-1))
				sqr.ResultIdx = len(sqr.ResultAry) - 1
				sqr.TempDivisionResult = 0
				n1 = sqr.ResultAry[0]
				for j := 1; j <= sqr.ResultIdx; j++ {
					n1 = (n1 * 10) + sqr.ResultAry[j]
				}
				sqr.TempDivisionResult = (n1 * 2) * 10
				break
			}
		}
	}


	return nil

}


func (sqr *SquareRootOp) FindNearestSquareLessThan() (error) {

	twoDigit := sqr.BaseNumPairs[0]

	if twoDigit < 0 {
		return errors.New("FindNearestSquareLessThan() - ERROR: twoDigit is less than Zero!")
	}


	if twoDigit == 0 {
		sqr.ResultAry = append(sqr.ResultAry, 0)
		return nil
	}

	for i:=1 ; i < 10; i++ {

		if i * i > twoDigit {
			sqr.ResultAry = append(sqr.ResultAry, i-1)
			return nil
		}
	}

	return errors.New("FindNearestSquareLessThan() - ERROR: Could NOT locate nearest square less than first base pair!")
}


func (sqr *SquareRootOp) CalcInputArrayStats() (error) {

	sqr.InputArrayLen = len(sqr.InputArray)

	for i:=0; i< sqr.InputArrayLen ; i++ {

		if sqr.InputArray[i] > 0 {
			if sqr.InputFirstIdx == -1 {
				sqr.InputFirstIdx = i
			}

			sqr.InputLastIdx = i
		}

	}

	sqr.InputIntegerLen = sqr.InputArrayLen - sqr.InputPrecision	- sqr.InputFirstIdx

	return nil
}
*/