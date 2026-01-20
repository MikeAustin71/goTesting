package strStuff

import (
	custErr "github.com/mikeaustin71/PackageStruct02/pkgStr/customErr"
	emp "github.com/mikeaustin71/PackageStruct02/pkgStr/employee"
)

type myStdErr = custErr.StdBasicError

type StrStuff struct {
	MyStr      string
	MyEmployee emp.Employee
}

func (s *StrStuff) SetTheString(str string) error {

	ePrefix := "StrStuff.SetTheEmployee"

	if len(str) == 0 {

		return &myStdErr{
			ErrPrefix:  ePrefix,
			ErrMessage: "Input parameter 'str' is empty",
		}
	}

	err := new(strStuffHelper).SetStr(s, str)

	if err != nil {

		return &myStdError{
			ErrPrefix:  ePrefix,
			ErrMessage: "Used to be err.Error()",
		}

	}

	return nil
}

func (s *StrStuff) PrintTheString() error {
	ePrefix := "StrStuff.PrintTheString"

	if len(s.MyStr) == 0 {

		return &myStdErr{
			ErrPrefix: ePrefix,
			ErrMessage: "StrStuff 'str' member variable is empty!\n" +
				"You MUST set the string first.",
		}
	}

	err := new(strStuffHelper).PrintStr(s)

	if err != nil {
		return &myStdError{
			ErrPrefix:  ePrefix,
			ErrMessage: err.Error(),
		}
	}

	return nil
}
