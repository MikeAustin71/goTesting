package strStuff

import (
	"fmt"

	"github.com/mikeaustin71/PackageStruct02/pkgStr/customErr"
)

// D:\GoProjects\MikeAustin71\goTesting\PackageStruct02

type strStuffHelper struct{}

type myStdError = customErr.StdBasicError

func (sStuffHelper *strStuffHelper) Error() string {
	//TODO implement me
	panic("implement me")
}

func (sStuffHelper *strStuffHelper) SetStr(
	mainStrStuff *StrStuff,
	str string) error {

	ePrefix := "strStuffHelper.SetStr()"

	if mainStrStuff == nil {

		return &customErr.StdBasicError{
			ErrPrefix:  ePrefix,
			ErrMessage: "Error: Input parameter 'mainStrStuff' is a 'nil' Pointer.",
		}
	}

	if len(str) == 0 {
		return &customErr.StdBasicError{
			ErrPrefix:  ePrefix,
			ErrMessage: "Error: Input parameter 'str' is a empty string.",
		}
	}

	mainStrStuff.MyStr = str

	return nil
}

func (sStuffHelper *strStuffHelper) PrintStr(mainStrStuff *StrStuff) error {
	ePrefix := "strStuffHelper.PrintStr()"

	if mainStrStuff == nil {
		return &myStdError{
			ErrPrefix:  ePrefix,
			ErrMessage: "Input Parameter 'mainStrStuff' is a 'nil' pointer!\n",
		}
	}

	if len(mainStrStuff.MyStr) == 0 {
		return &myStdError{
			ErrPrefix:  ePrefix,
			ErrMessage: "Error: Input Parameter 'mainStrStuff' is a empty string!",
		}
	}

	fmt.Printf("\n\nPrinting a StrStuff String:\nmainStrStuff.MyStr: %s\n\n",
		mainStrStuff.MyStr)

	return nil
}
