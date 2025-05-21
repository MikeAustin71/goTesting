package strStuff

// D:\GoProjects\MikeAustin71\goTesting\PackageStruct02

import (
  "fmt"
  custErr "golangmikesamples/PackageStruct02/pkgStr/customErr"
)

// D:\GoProjects\MikeAustin71\goTesting\PackageStruct02

type strStuffHelper struct{}

type myStdError = custErr.StdBasicError

func (sStuffHelper *strStuffHelper) SetStr(
  mainStrStuff *StrStuff,
  str string) error {

  ePrefix := "strStuffHelper.SetStr()"

  if mainStrStuff == nil {

    return &custErr.StdBasicError{
      ErrPrefix:  ePrefix,
      ErrMessage: "Error: Input parameter 'mainStrStuff' is a 'nil' Pointer.",
    }
  }

  if len(str) == 0 {
    return &custErr.StdBasicError{
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
