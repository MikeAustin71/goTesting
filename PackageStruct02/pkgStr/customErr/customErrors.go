package customErr

import "fmt"

// StdBasicError
// Custom error message used to specify the
// function returning the error message
type StdBasicError struct {
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
//			&StdBasicError{
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
func (e *StdBasicError) Error() string {

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

func (e *StdBasicError) Unwrap() error {

	return fmt.Errorf("%w", e)

}
