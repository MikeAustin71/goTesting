package main

import "fmt"

// ReturnBasicError
// Custom error message used to specify the
// function returning the error message
type ReturnBasicError struct {
	ErrPrefix  string
	ReturnFunc string
	ErrContext string
	ErrMessage string
}

// Error
// The following format is used to initialize
// this type of error:
//
//	if err != nil {
//
//		return Decimal{},
//			&ReturnBasicError{
//				ErrPrefix:  ePrefix.String(),
//				ReturnFunc: "dec, err := new(Decimal).NewBigInt(\n" +
//				"  big.NewInt(0).Set(bNum.bigInt), bNum.precision)\n",
//				ErrContext: "",
//				ErrMessage:   err.Error(),
//			}
//	}
//
//	NOTE:
//
//	Element 'ErrContext' is optional
func (e *ReturnBasicError) Error() string {

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

func (e *ReturnBasicError) Unwrap() error {

	return fmt.Errorf("%w", e)

}

func main() {

	err := testControl()

	if err != nil {

		fmt.Printf("\n")
		fmt.Printf("Error Return Example\n")
		fmt.Printf("--------------------\n")
		fmt.Printf("%v", err)
		fmt.Printf("--------------------\n")

	}

}

//func test01() error {
//
//	ePrefix := "test01"
//
//	return fmt.Errorf("%w",
//		&ReturnBasicError{
//			ErrPrefix:  ePrefix,
//			ErrMessage: "Error thrown by test01()",
//		})
//
//}

func testControl() error {

	ePrefix := "testControl"

	err := test021()

	if err != nil {

		return fmt.Errorf("%v\n%w", ePrefix, err)

	}

	return nil
}

func test021() error {

	ePrefix := "test021"

	err := test022()

	if err != nil {

		return fmt.Errorf("%v\n%w", ePrefix, err)

	}

	return nil
}

func test022() error {

	ePrefix := "test022"

	err := test023()

	if err != nil {

		return fmt.Errorf("%v\n%w", ePrefix, err)

	}

	return nil
}

func test023() error {

	ePrefix := "test023"

	err := &ReturnBasicError{
		ErrPrefix:  ePrefix,
		ErrMessage: "Error thrown by 'test023'",
	}

	if err != nil {

		return fmt.Errorf("%v\n%w", ePrefix, err)

	}

	return nil
}
