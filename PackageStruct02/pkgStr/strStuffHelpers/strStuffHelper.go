package pkgStr

// D:\GoProjects\MikeAustin71\goTesting\PackageStruct02

import (
  mainRepo "golangmikesamples/PackageStruct02/pkgStr"
  custErr "golangmikesamples/PackageStruct02/pkgStr/customErr"
)

// D:\GoProjects\MikeAustin71\goTesting\PackageStruct02

type strStuffHelper struct{}

func (sStuffHelper *strStuffHelper) SetStr(
  mainStrStuff *mainRepo.StrStuff,
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
