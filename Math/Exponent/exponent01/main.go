package main

import (
  "fmt"
  "math/big"
)

// FuncReturnError
// Custom error message used to specify the
// function returning the error message
type FuncReturnError struct {
  ErrPrefix  string
  ReturnFunc string
  ErrContext string
  ErrMessage string
}

// Error
// The following format is used to initialize
// this type of error:
//
// EXAMPLE #1:
//
//	if err != nil {
//
//		return Decimal{},
//			&FuncReturnError{
//				ErrPrefix:  ePrefix.String(),
//				ReturnFunc: "dec, err := new(Decimal).NewBigInt(\n" +
//				"  big.NewInt(0).Set(bNum.bigInt), bNum.precision)\n",
//				ErrContext: "",
//				ErrMessage:   err.Error(),
//			}
//	}
//
//
//	EXAMPLE #2:
//
//	return Decimal{},
//	&FuncReturnError{
//		ErrPrefix: ePrefix.String(),
//		ReturnFunc: "dec, err := new(Decimal).NewBigInt(\n" +
//		"    big.NewInt(0).Set(bNum.bigInt), bNum.precision)",
//		ErrContext: "",
//		ErrMessage: err.Error(),
//	}
//
//	NOTE:
//
//	Element 'ErrContext' is optional
func (e *FuncReturnError) Error() string {

  var errStr string
  foundCnt := 0

  if e.ErrPrefix != "" {
    errStr = e.ErrPrefix + "\n"
    foundCnt++
  }

  if e.ReturnFunc != "" {
    errStr += "Error returned by: \\n  " + e.ReturnFunc + "\n"
    foundCnt++
  }

  if e.ErrContext != "" {
    errStr += e.ErrContext + "\n"
    foundCnt++
  }

  if e.ErrMessage != "" {
    errStr += "Error: \n  " + e.ErrMessage + "\n"
    foundCnt++
  }

  if foundCnt == 0 {
    errStr = "No Error parameters provided!\n"
  } else {

    errStr += "\n"
  }

  return errStr
}

func (e *FuncReturnError) GetError() error {
  return fmt.Errorf("%w", e.Error())
}

func (e *FuncReturnError) Unwrap() error {
  return fmt.Errorf("%w", e.Error())
}

func main() {
  err := test03()

  if err != nil {
    fmt.Println(err)
  }

  return
}

func test01() error {

  ePrefix := "main-test01 - Scale Factor Test #1"

  expectedPrecisionUint := uint(2)

  expectedScaleFactor := big.NewInt(100)

  bigInt10 := big.NewInt(10)

  bigIntExpectedPrecision := big.NewInt(int64(expectedPrecisionUint))

  actualScaleFactor := big.NewInt(0).Exp(bigInt10, bigIntExpectedPrecision, nil)

  fmt.Printf("\n%v\n"+
    "actualScaleFactor = %v\n\n",
    ePrefix, actualScaleFactor.Text(10))

  if expectedScaleFactor.Cmp(actualScaleFactor) != 0 {

    errMsg := fmt.Sprintf("%v\n"+
      "Error: Computed Scale Factor is INCORRECT!\n"+
      "Because expectedScaleFactor.Cmp(actualScaleFactor) != 0\n"+
      "Expected actualScaleFactor = '%v'\n"+
      "  Actual actualScaleFactor = '%v'\n"+
      "expectedPrecisionUint = '%v' \n\n",
      ePrefix, expectedScaleFactor.Text(10),
      actualScaleFactor.Text(10),
      expectedPrecisionUint)

    return &FuncReturnError{
      ErrPrefix:  ePrefix,
      ReturnFunc: "",
      ErrContext: fmt.Sprintf("expectedScaleFactor.Cmp(actualScaleFactor) != 0"),
      ErrMessage: errMsg,
    }

  }

  fmt.Printf("\n\n********** Success! **********\n"+
    "Function: %v\n"+
    "Expected actualScaleFactor = '%v'\n"+
    "  Actual actualScaleFactor = '%v'\n"+
    "expectedPrecisionUint = '%v' \n\n",
    ePrefix,
    expectedScaleFactor.Text(10),
    actualScaleFactor.Text(10),
    expectedPrecisionUint)

  return nil

}

func test02() error {

  ePrefix := "main-test02 - Scale Factor Test #2"

  expectedPrecisionUint := uint(2)

  expectedScaleFactor := big.NewInt(100)

  actualScaleFactor := big.NewInt(0).Exp(
    big.NewInt(10),
    big.NewInt(int64(expectedPrecisionUint)),
    nil)

  fmt.Printf("\n%v\n"+
    "actualScaleFactor = %v\n\n",
    ePrefix, actualScaleFactor.Text(10))

  if expectedScaleFactor.Cmp(actualScaleFactor) != 0 {

    errMsg := fmt.Sprintf("%v\n"+
      "Error: Computed Scale Factor is INCORRECT!\n"+
      "Because expectedScaleFactor.Cmp(actualScaleFactor) != 0\n"+
      "Expected actualScaleFactor = '%v'\n"+
      "  Actual actualScaleFactor = '%v'\n"+
      "expectedPrecisionUint = '%v' \n\n",
      ePrefix, expectedScaleFactor.Text(10),
      actualScaleFactor.Text(10),
      expectedPrecisionUint)

    return &FuncReturnError{
      ErrPrefix:  ePrefix,
      ReturnFunc: "",
      ErrContext: fmt.Sprintf("expectedScaleFactor.Cmp(actualScaleFactor) != 0"),
      ErrMessage: errMsg,
    }

  }

  fmt.Printf("\n\n********** Success! **********\n"+
    "Function: %v\n"+
    "Expected actualScaleFactor = '%v'\n"+
    "  Actual actualScaleFactor = '%v'\n"+
    "expectedPrecisionUint = '%v' \n\n",
    ePrefix,
    expectedScaleFactor.Text(10),
    actualScaleFactor.Text(10),
    expectedPrecisionUint)

  return nil

}

func test03() error {

  ePrefix := "main-test03 - Scale Factor Test #3"

  expectedPrecisionUint := uint(0)

  expectedScaleFactor := big.NewInt(1)

  actualScaleFactor := big.NewInt(0).Exp(
    big.NewInt(10),
    big.NewInt(int64(expectedPrecisionUint)), nil)

  fmt.Printf("\n%v\n"+
    "actualScaleFactor = %v\n\n",
    ePrefix, actualScaleFactor.Text(10))

  if expectedScaleFactor.Cmp(actualScaleFactor) != 0 {

    errMsg := fmt.Sprintf("%v\n"+
      "Error: Computed Scale Factor is INCORRECT!\n"+
      "Because expectedScaleFactor.Cmp(actualScaleFactor) != 0\n"+
      "Expected actualScaleFactor = '%v'\n"+
      "  Actual actualScaleFactor = '%v'\n"+
      "expectedPrecisionUint = '%v' \n\n",
      ePrefix, expectedScaleFactor.Text(10),
      actualScaleFactor.Text(10),
      expectedPrecisionUint)

    return &FuncReturnError{
      ErrPrefix:  ePrefix,
      ReturnFunc: "",
      ErrContext: fmt.Sprintf("expectedScaleFactor.Cmp(actualScaleFactor) != 0"),
      ErrMessage: errMsg,
    }

  }

  fmt.Printf("\n\n********** Success! **********\n"+
    "Function: %v\n"+
    "Expected actualScaleFactor = '%v'\n"+
    "  Actual actualScaleFactor = '%v'\n"+
    "expectedPrecisionUint = '%v' \n\n",
    ePrefix,
    expectedScaleFactor.Text(10),
    actualScaleFactor.Text(10),
    expectedPrecisionUint)

  return nil

}
