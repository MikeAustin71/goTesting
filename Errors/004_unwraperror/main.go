package main

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
)

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

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"main()",
		"")

	if err != nil {

		fmt.Printf("main()\nError returned by ePref.ErrPrefixDto{}.NewIEmpty()/n%v/n", err.Error())

		return
	}

	err = testControl(ePrefix)

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

func testControl(errPrefDto *ePref.ErrPrefixDto) error {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"testControl",
		"")

	if err != nil {
		return err
	}

	err = test021(ePrefix)

	if err != nil {

		return err

	}

	return nil
}

func test021(errPrefDto *ePref.ErrPrefixDto) error {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"test021",
		"")

	if err != nil {
		return err
	}

	err = test022(ePrefix)

	if err != nil {

		return err

	}

	return nil
}

func test022(errPrefDto *ePref.ErrPrefixDto) error {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"test022",
		"")

	if err != nil {
		return err
	}

	err = test023(ePrefix)

	if err != nil {

		return err

	}

	return nil
}

func test023(errPrefDto *ePref.ErrPrefixDto) error {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"test023",
		"")

	if err != nil {
		return err
	}

	err = &ReturnBasicError{
		ErrPrefix:  ePrefix.String(),
		ErrMessage: "Error thrown by 'test023'",
	}

	if err != nil {

		return err

	}

	return nil
}
