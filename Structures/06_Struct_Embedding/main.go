package main

import "fmt"

// ErrorStandard
// Custom error message used to specify the
// function returning the error message
type ErrorStandard struct {
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
//			&ErrorStandard{
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
func (e *ErrorStandard) Error() string {

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

func (e *ErrorStandard) Unwrap() error {

  return fmt.Errorf("%w", e)

}

type NumSepsDto struct {
  Decimal   rune
  Thousands rune
  Currency  rune
}

func (nSeps *NumSepsDto) CopyIn(nSeps2 *NumSepsDto) error {

  ePrefix := "NumSepsDto.CopyIn"

  if nSeps2 == nil {
    return &ErrorStandard{
      ErrPrefix:  ePrefix,
      ErrMessage: "Error: Input parameter nSeps2 is a 'nil' pointer!",
    }
  }

  if nSeps2.Decimal == 0 || nSeps2.Thousands == 0 || nSeps2.Currency == 0 {
    return &ErrorStandard{
      ErrPrefix: ePrefix,
      ErrMessage: "Error: Input parameter nSeps2 is invalid!\n" +
        "One or more member elements have a zero value.\n",
    }
  }

  nSeps.Decimal = nSeps2.Decimal
  nSeps.Thousands = nSeps2.Thousands
  nSeps.Currency = nSeps2.Currency

  return nil
}

func (nSeps *NumSepsDto) CopyOut() (*NumSepsDto, error) {
  ePrefix := "NumSepsDto.CopyOut"

  if nSeps.Decimal == 0 || nSeps.Thousands == 0 || nSeps.Currency == 0 {
    return nil, &ErrorStandard{
      ErrPrefix: ePrefix,
      ErrMessage: "Error: NumSepsDto is invalid!\n" +
        "One or more member elements have a zero value.\n",
    }
  }

  nSeps2 := &NumSepsDto{}

  nSeps2.Decimal = nSeps.Decimal
  nSeps2.Thousands = nSeps.Thousands
  nSeps2.Currency = nSeps.Currency

  return nSeps2, nil

}

func (nSeps *NumSepsDto) SetDecimal(decimal rune) error {

  ePrefix := "NumSepsDto.SetDecimal"

  if decimal == 0 {

    return &ErrorStandard{
      ErrPrefix:  ePrefix,
      ErrContext: "Error: Input parameter 'decimal' is zero",
    }
  }

  nSeps.Decimal = decimal

  return nil
}

func (nSeps *NumSepsDto) SetThousands(thousandsSep rune) error {

  ePrefix := "NumSepsDto.SetThousands"

  if thousandsSep == 0 {

    return &ErrorStandard{
      ErrPrefix:  ePrefix,
      ErrContext: "Error: Input parameter 'thousandsSep' is zero",
    }
  }

  nSeps.Thousands = thousandsSep

  return nil
}

func (nSeps *NumSepsDto) SetCurrency(currencySep rune) error {

  ePrefix := "NumSepsDto.SetCurrency"

  if currencySep == 0 {

    return &ErrorStandard{
      ErrPrefix:  ePrefix,
      ErrContext: "Error: Input parameter 'currencySep' is zero",
    }
  }

  nSeps.Currency = currencySep

  return nil
}

type NumStrDto struct {
  NumStr       string
  AbsIntValue  int64
  AbsFracValue int64
  SignValue    int // -1, 0 or 1
  Decimal      rune
}

func (nStr *NumStrDto) GetNumStr() (string, error) {

  ePrefix := "NumStrDto.GetNumStr"

  if nStr.NumStr == "" {
    return "", &ErrorStandard{
      ErrPrefix:  ePrefix,
      ErrMessage: "Error: This NumStrDto is empty!",
    }
  }

  return nStr.NumStr, nil
}

func (nStr *NumStrDto) SetNumberValue(
  AbsIntValue,
  AbsFracValue int64,
  SignValue int,
  DecimalChar rune) error {

  ePrefix := "NumStrDto.SetNumberValue"

  if SignValue < -1 || SignValue > 1 {
    return &ErrorStandard{
      ErrPrefix: ePrefix,
      ErrMessage: "Error: SignValue is INVALID!\n" +
        "SignValue must be between -1 and 1!\n" +
        fmt.Sprintf("SignValue: %v", SignValue),
    }
  }

  if AbsIntValue < 0 {
    AbsIntValue *= -1
  }

  if AbsFracValue < 0 {
    AbsFracValue *= -1
  }

  if AbsIntValue == 0 && AbsFracValue == 0 {
    SignValue = 0
  }

  nStr.AbsIntValue = AbsIntValue
  nStr.AbsFracValue = AbsFracValue
  nStr.SignValue = SignValue

  numSign := ""

  if SignValue < 0 {
    numSign = "-"
  }

  if DecimalChar == '.' {
    return &ErrorStandard{
      ErrPrefix: ePrefix,
      ErrMessage: "Error: Decimal Character Rune has a ZERO value!\n" +
        "The number string is invalid!\n",
    }
  }

  nStr.NumStr = numSign + fmt.Sprintf("%d%c%d", AbsIntValue, DecimalChar, AbsFracValue)

  return nil
}

func (nStr *NumStrDto) GetNumberValue() (string, error) {
  ePrefix := "NumStrDto.GetNumberValue"

  if nStr.NumStr == "" {
    return "", &ErrorStandard{
      ErrPrefix:  ePrefix,
      ErrMessage: "Error: This NumStrDto is empty!",
    }
  }

  return nStr.NumStr, nil
}

func (nStr *NumStrDto) PrintNumStr() error {
  ePrefix := "NumStrDto.PrintNumStr"

  if nStr.NumStr == "" {
    return &ErrorStandard{
      ErrPrefix:  ePrefix,
      ErrMessage: "Error: This NumStrDto is empty!",
    }
  }

  fmt.Printf("\n\n%v\n", nStr.NumStr)

  return nil
}

func (nStr *NumStrDto) ResetNumberValue() {

  if nStr.NumStr == "0" && nStr.AbsFracValue == 0 {
    nStr.SignValue = 0
  }

  numSign := ""

  if nStr.SignValue < 0 {
    numSign = "-"
  }

  nStr.NumStr = numSign + fmt.Sprintf("%d.%d", nStr.AbsIntValue, nStr.AbsFracValue)

  return
}

type DecimalNum struct {
  NumberStr     *NumStrDto
  NumSeparators *NumSepsDto
}

func (decNum *DecimalNum) SetDecimal() error {
  ePrefix := "DecimalNum.SetDecimal"

  if decNum.NumSeparators == nil {
    return &ErrorStandard{
      ErrPrefix: ePrefix,
      ErrMessage: "Error: NumSeparators is empty!\n" +
        "This instance of DecimalNum is INVALID!",
    }
  }

  if decNum.NumberStr == nil {
    return &ErrorStandard{
      ErrPrefix: ePrefix,
      ErrMessage: "Error: NumStrDto is empty!\n" +
        "This instance of DecimalNum is INVALID!",
    }
  }

  decNum.NumberStr.Decimal = decNum.NumSeparators.Decimal

  decNum.NumberStr.ResetNumberValue()

  return nil
}

func main() {

  decNum, err := T01SetDecNum(
    35,
    752,
    -1)

  if err != nil {
    fmt.Printf("%v", err.Error())
    return
  }

  fmt.Printf("Should be -35.752\n")
	
  err = decNum.NumberStr.PrintNumStr()

  if err != nil {
    fmt.Printf("%v", err.Error())
  }

}

func T01SetDecNum(
  AbsIntValue,
  AbsFracValue int64,
  SignValue int) (*DecimalNum, error) {

  ePrefix := "T01SetDecNum"

  decNum := &DecimalNum{
    NumberStr: &NumStrDto{
      AbsIntValue:  AbsIntValue,
      AbsFracValue: AbsFracValue,
      SignValue:    SignValue,
    },
    NumSeparators: &NumSepsDto{
      Decimal:   '.',
      Thousands: ',',
      Currency:  '$',
    },
  }

  err := decNum.SetDecimal()

  if err != nil {
    return nil, &ErrorStandard{
      ErrPrefix:  ePrefix,
      ErrContext: "Error returned from decNum.SetDecimal()",
      ErrMessage: err.Error(),
    }
  }

  return decNum, nil
}
