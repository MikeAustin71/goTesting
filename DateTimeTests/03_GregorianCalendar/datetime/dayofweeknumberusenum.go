package datetime

import (
	"fmt"
	"sync"
)

var mUsDayOfWeekNoStringToCode = map[string]UsDayOfWeekNo{
	"NONE"            : UsDayOfWeekNo(-1),
	"SUNDAY"          : UsDayOfWeekNo(0),
	"MONDAY"          : UsDayOfWeekNo(1),
	"TUESDAY"         : UsDayOfWeekNo(2),
	"WEDNESDAY"       : UsDayOfWeekNo(3),
	"THURSDAY"        : UsDayOfWeekNo(4),
	"FRIDAY"          : UsDayOfWeekNo(5),
	"SATURDAY"        : UsDayOfWeekNo(6),
	"None"            : UsDayOfWeekNo(-1),
	"Sunday"          : UsDayOfWeekNo(0),
	"Monday"          : UsDayOfWeekNo(1),
	"Tuesday"         : UsDayOfWeekNo(2),
	"Wednesday"       : UsDayOfWeekNo(3),
	"Thursday"        : UsDayOfWeekNo(4),
	"Friday"          : UsDayOfWeekNo(5),
	"Saturday"        : UsDayOfWeekNo(6),
	"none"            : UsDayOfWeekNo(-1),
	"sunday"          : UsDayOfWeekNo(0),
	"monday"          : UsDayOfWeekNo(1),
	"tuesday"         : UsDayOfWeekNo(2),
	"wednesday"       : UsDayOfWeekNo(3),
	"thursday"        : UsDayOfWeekNo(4),
	"friday"          : UsDayOfWeekNo(5),
	"saturday"        : UsDayOfWeekNo(6),
	"SUN"             : UsDayOfWeekNo(0),
	"MON"            : UsDayOfWeekNo(1),
	"TUE"            : UsDayOfWeekNo(2),
	"WED"            : UsDayOfWeekNo(3),
	"THU"            : UsDayOfWeekNo(4),
	"THURS"          : UsDayOfWeekNo(4),
	"FRI"            : UsDayOfWeekNo(5),
	"SAT"            : UsDayOfWeekNo(6) ,
	"Sun"             : UsDayOfWeekNo(0),
	"Mon"            : UsDayOfWeekNo(1),
	"Tue"            : UsDayOfWeekNo(2),
	"Wed"            : UsDayOfWeekNo(3),
	"Thu"            : UsDayOfWeekNo(4),
	"Thurs"          : UsDayOfWeekNo(4),
	"Fri"            : UsDayOfWeekNo(5),
	"Sat"            : UsDayOfWeekNo(6) ,
	"sun"             : UsDayOfWeekNo(0),
	"mon"            : UsDayOfWeekNo(1),
	"tue"            : UsDayOfWeekNo(2),
	"wed"            : UsDayOfWeekNo(3),
	"thu"            : UsDayOfWeekNo(4),
	"thurs"          : UsDayOfWeekNo(4),
	"fri"            : UsDayOfWeekNo(5),
	"sat"            : UsDayOfWeekNo(6) ,
}

var mUsDayOfWeekNoCodeToString = map[UsDayOfWeekNo]string{
	UsDayOfWeekNo(-1)  : "None",
	UsDayOfWeekNo(0)   : "Sunday",
	UsDayOfWeekNo(1)   : "Monday",
	UsDayOfWeekNo(2)   : "Tuesday",
	UsDayOfWeekNo(3)   : "Wednesday",
	UsDayOfWeekNo(4)   : "Thursday",
	UsDayOfWeekNo(5)   : "Friday",
	UsDayOfWeekNo(6)   : "Saturday",
}

var mUsDayOfWeekNoCodeTo3CharAbbrvDay = map[UsDayOfWeekNo]string{
	UsDayOfWeekNo(0)   : "Sun",
	UsDayOfWeekNo(1)   : "Mon",
	UsDayOfWeekNo(2)   : "Tue",
	UsDayOfWeekNo(3)   : "Wed",
	UsDayOfWeekNo(4)   : "Thu",
	UsDayOfWeekNo(5)   : "Fri",
	UsDayOfWeekNo(6)   : "Sat",
}

var mUsDayOfWeekNoCodeTo2CharAbbrvDay = map[UsDayOfWeekNo]string{
	UsDayOfWeekNo(0)   : "Su",
	UsDayOfWeekNo(1)   : "Mo",
	UsDayOfWeekNo(2)   : "Tu",
	UsDayOfWeekNo(3)   : "We",
	UsDayOfWeekNo(4)   : "Th",
	UsDayOfWeekNo(5)   : "Fr",
	UsDayOfWeekNo(6)   : "Sa",
}


// UsDayOfWeekNo - An enumeration of U.S. (United States) day of the week
// numbers.
//
// The most common day of the week numbering system used internationally,
// is the ISO 8601 standard which specifies that the week begins on Monday.
// ISO stands for the International Organization for Standardization (ISO).
//    https://www.iso.org/home.html
//
// Under the ISO 8601 Standard, days of the week are numbered beginning with
// one (1) for Monday and ending with seven (7) for Sunday. For an enumeration
// of ISO 8601 days of the week, reference type, ISO8601DayOfWeekNo defined
// in source file:
//    datetime/dayofweeknumberiso8601enum.go
//
// The United States, Canada, Australia and New Zealand put Sunday as the
// first day of the week on their calendars. The first day of the week,
// 'Sunday', is numbered as zero with the last day of the week being numbered
// as 6, for 'Saturday'.
//
// The type defined here, UsDayOfWeekNo, provides an enumeration for United
// States day of the week numbers as listed below:
//
//              Week         Week Day
//              Day           Number
//            ========     ===========
//             Sunday          = 0
//             Monday          = 1
//             Tuesday         = 2
//             Wednesday       = 3
//             Thursday        = 4
//             Friday          = 5
//             Saturday        = 6
//
// Since the Go Programming Language does not directly support
// enumerations, the UsDayOfWeekNo type has been adapted to function
// in a manner similar to classic enumerations. 'UsDayOfWeekNo' is
// declared as a type 'int'. The method names effectively represent
// an enumeration of U.S. Day Of The Week numbers. These methods are
// listed as follows:
//
//
//  None           (-1) - None - Signals that the U.S. Day Of The Week Number
//                        (UsDayOfWeekNo) Type is not initialized. This is
//                        an error condition.
//
//  Sunday          (0) - In accordance with the U.S. day of the week numbering
//                        system, Sunday is the first day of the week and is
//                        represented by a week day number value of 'zero'.
//
//  Monday          (1) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Monday', with a weekday number value of '1'.
//
//
//  Tuesday         (2) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Tuesday', with a weekday number value of '2'.
//
//  Wednesday       (3) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Wednesday', with a weekday number value of '3'.
//
//  Thursday        (4) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Thursday', with a weekday number value of '4'.
//
//  Friday          (5) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Friday', with a weekday number value of '5'.
//
//  Saturday        (6) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Saturday', with a weekday number value of '6'.
//                        Under the U.S. day of the week numbering system,
//                        'Saturday' is the last day of the week.
//
// Resources:
//   https://www.timeanddate.com/date/week-numbers.html
//
//
// Usage:
//
// For easy access to these enumeration values, use the global variable 'UsWeekDayNo'.
// Example: UsWeekDayNo.Tuesday()
//
// Otherwise you will need to use the formal syntax.
// Example: UsDayOfWeekNo(0).Tuesday()
//
// Depending on your editor, intellisense (a.k.a. intelligent code completion) may not
// list the 'UsDayOfWeekNo' methods in alphabetical order.
//
// Finally, be advised that all 'UsDayOfWeekNo' methods beginning with 'X', as well as
// the method 'String()' are utility methods and not part of the enumeration.
//
type UsDayOfWeekNo int

var lockUsDayOfWeekNo sync.Mutex

// None - Signals that the UsDayOfWeekNo Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) None() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(-1)
}

// Sunday - Signals that the UsDayOfWeekNo instance is
// equal to 'Sunday', with a day number value of zero (0).
// Under the USA weekday numbering system, 'Sunday' is
// designated as the first day of the week.
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Sunday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(0)
}

// Monday - Signals that the UsDayOfWeekNo instance is equal
// to 'Monday', with a weekday number value of one (1).
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Monday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(1)
}

// Tuesday - Signals that the UsDayOfWeekNo instance is equal
// to 'Tuesday', with a weekday number value of two (2).
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Tuesday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(2)
}

// Wednesday - Signals that the UsDayOfWeekNo instance is equal
// to 'Wednesday', with a weekday number value of three (3).
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Wednesday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(3)
}

// Thursday - Signals that the UsDayOfWeekNo instance is equal
// to 'Thursday', with a weekday number value of four (4).
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Thursday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(4)
}

// Friday - Signals that the UsDayOfWeekNo instance is equal
// to 'Friday', with a weekday number value of five (5).
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Friday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(5)
}

// Saturday - Signals that the UsDayOfWeekNo instance is equal
// to 'Saturday', with a weekday number value of six (6).
//
// Under the USA weekday numbering system, 'Saturday' is the
// last day of the week.
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Saturday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(6)
}

// String - Returns a string with the day-of-week name associated with
// this UsDayOfWeekNo enumeration instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= UsDayOfWeekNo(0).Monday()
// str := t.String()
//     str is now equal to "Monday"
//
// t= UsDayOfWeekNo(0).Tuesday()
// str = t.String()
//     str is now equal to "Tuesday"
//
func (usDayOfWk UsDayOfWeekNo) String() string {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	result, ok := mUsDayOfWeekNoCodeToString[usDayOfWk]

	if !ok {
		return fmt.Sprintf("Error: Value= %d", int(usDayOfWk))
	}

	return result
}

// XDayOfWeekName - Returns a string with the day-of-week name associated with
// this UsDayOfWeekNo enumeration value.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= UsDayOfWeekNo(0).Monday()
// str := t.String()
//     str is now equal to "Monday"
//
func (usDayOfWk UsDayOfWeekNo) XDayOfWeekName() string {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	result, ok := mUsDayOfWeekNoCodeToString[usDayOfWk]

	if !ok {
		return fmt.Sprintf("Error: Value= %d", int(usDayOfWk))
	}

	return result
}

// XDayOfWeekNameAbbreviation - Returns a string with the abbreviated name of
// the day of the week associated with this UsDayOfWeekNo enumeration value.
// The returned abbreviated name of the day will either contain two (2) or
// three (3) characters depending on the specification contained in input
// parameter, 'numberOfCharsInAbbreviation'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numberOfCharsInAbbreviation   int
//     - This integer value determines the number of characters in the day of
//       week name abbreviations returned by this method. Currently, only
//       two or three character day of the week name abbreviations are supported.
//       If the input parameter, 'numberOfCharsInAbbreviation', contains any value
//       other than '2' or '3', an error will be returned.
//
//
//  ePrefix                       string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  dayOfWeekNameAbbrv            string
//     - An abbreviated day of the week name. This abbreviated day of
//       the week name will either contain 1 or 2 characters depending
//       on the specification provided by input parameter,
//       'numberOfCharsInAbbreviation'.
//
//
//  err                           error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= UsDayOfWeekNo(0).Monday()
//
// str, err := t.XDayOfWeekNameAbbreviation(
//        3,
//        "callingFunctionName")
//
//     str is now equal to 'Mon'
//
// str, err = t.XDayOfWeekNameAbbreviation(
//        2,
//        "callingFunctionName")
//
//     str is now equal to 'Mo'
//
func (usDayOfWk UsDayOfWeekNo) XDayOfWeekNameAbbreviation(
	numberOfCharsInAbbreviation int,
	ePrefix string) (
	dayOfWeekNameAbbrv string,
	err error) {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()


	ePrefix += "UsDayOfWeekNo.XDayOfWeekNameAbbreviation() "
	err = nil
	dayOfWeekNameAbbrv = ""

	var ok bool

	switch numberOfCharsInAbbreviation {

	case 2:

		dayOfWeekNameAbbrv,
			ok = mUsDayOfWeekNoCodeTo2CharAbbrvDay[usDayOfWk]

	case 3:

		dayOfWeekNameAbbrv,
			ok = mUsDayOfWeekNoCodeTo3CharAbbrvDay[usDayOfWk]

	default:

		ok = true

		err = fmt.Errorf(ePrefix + "\n" +
			"ERROR: Input parameter 'numberOfCharsInAbbreviation' is INVALID!\n" +
			"Valid values for 'numberOfCharsInAbbreviation are '2' or '3'.\n" +
			"numberOfCharsInAbbreviation='%v'\n",
			numberOfCharsInAbbreviation)

	}

	if !ok {
		err =
			fmt.Errorf(ePrefix + "\n" +
				"Error: UsDayOfWeekNo is INVALID!\n" +
				"usDayOfWk= %d", int(usDayOfWk))
	}

	return dayOfWeekNameAbbrv, err
}

// XDayOfWeekNumber - Returns an integer number equal to the Day of
// the Week Number for this instance of XDayOfWeekNumber.
//
// ISO 8601 Day Of The Week numbers as listed below:
//
//                        Monday    = 1
//                        Tuesday   = 2
//                        Wednesday = 3
//                        Thursday  = 4
//                        Friday    = 5
//                        Saturday  = 6
//                        Sunday    = 7
//
// If the current instance of UsDayOfWeekNo is valid, one of
// the integer values listed above will be returned. To verify
// validity, call method func UsDayOfWeekNo.XIsValid().
//
func (usDayOfWk UsDayOfWeekNo) XDayOfWeekNumber() int {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return int(usDayOfWk)
}

// XDayOfWeekNumberingSystemType - Returns the Day of the Week
// Numbering System Type code associated with UsDayOfWeekNo. Therefore,
// the value returned by this method is always the enumeration value,
// DayOfWeekNumberingSystemType.UsDayOfWeek().
//
//
func (usDayOfWk UsDayOfWeekNo) XDayOfWeekNumberingSystemType() DayOfWeekNumberingSystemType {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return DayOfWeekNumberingSystemType(0).UsDayOfWeek()
}

// XISO8601DayOfWeekNumber - Returns the ISO 8601 Day of the Week
// Number as an integer value.
//
// The UsDayOfWeekNo type returns the Day of the Week Number used
// in the United States, Canada, Australia and New Zealand.
// This method returns the equivalent Day of the Week Number for
// alternative numbering system using the ISO 8601 standard.
//
// The ISO 8601 standard specifies that the first day of the
// week is 'Monday' and is numbered as weekday number one (1).
// 'Sunday' is the last day of the week and has a week day
// number value of seven (7).
//
// Reference:
//   https://www.timeanddate.com/date/week-numbers.html
//   https://en.wikipedia.org/wiki/Julian_day
//   https://webspace.science.uu.nl/~gent0113/calendar/isocalendar.htm
//   Type: ISO8601DayOfWeekNo Source File: datetime/dayofweeknumberiso8601enum.go
//
func (usDayOfWk UsDayOfWeekNo) XISO8601DayOfWeekNumber(
	ePrefix string) (ISO8601DayOfWeekNo, error) {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	ePrefix += "UsDayOfWeekNo.XISO8601DayOfWeekNumber() "

	dayNum := int(usDayOfWk)

	if dayNum > 6 {
		return ISO8601DayOfWeekNo(0),
			fmt.Errorf(ePrefix + "\n" +
				"Error: The U.S. Day Of The Week Number is Greater Than '6' and therefore, INVALID!\n" +
				"U.S. Day Of The Week Number='%v'\n", dayNum)
	}

	if dayNum < 0 {
		return ISO8601DayOfWeekNo(0),
			fmt.Errorf(ePrefix + "\n" +
				"Error: The U.S. Day Of The Week Number is Less Than 'zero' and therefore, INVALID!\n" +
				"U.S. Day Of The Week Number='%v'\n", dayNum)
	}

	if dayNum == 0 {
		dayNum = 7
	}

	return ISO8601DayOfWeekNo(dayNum), nil
}

// XIsValid - Returns a boolean value signaling whether the current U.S. Day
// of Week Number value is valid.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  usDayOfWeek := UsDayOfWeekNo(0).Sunday()
//
//  isValid := usDayOfWeek.XIsValid()
//    If 'isValid' is false it signals that 'usDayOfWeek' contains
//    an invalid value.
//
func (usDayOfWk UsDayOfWeekNo) XIsValid() bool {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	if usDayOfWk > 6 ||
		usDayOfWk < 0 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of UsDayOfWeekNo is returned set to the value of
// the associated enumeration.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// valueString   string - A string which will be matched against the
//                        enumeration string values. If 'valueString'
//                        is equal to a day of the week name or abbreviated
//                        name, this method will proceed to successful
//                        completion and return the correct enumeration
//                        value. The 'valueString' processing is case
//                        neutral meaning that both upper case and lower
//                        case 'valueString' values will be successfully
//                        processed.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// UsDayOfWeekNo - Upon successful completion, this method will return a new
//                 instance of UsDayOfWeekNo set to the value of the enumeration
//                 matched by the string search performed on input parameter,
//                'valueString'.
//
// error        - If this method completes successfully, the returned error
//                Type is set equal to 'nil'. If an error condition is encountered,
//                this method will return an error type which encapsulates an
//                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
// All of these examples will be successfully processed.
//
// t, err := UsDayOfWeekNo(0).XParseString("Tuesday")
// t, err := UsDayOfWeekNo(0).XParseString("TUESDAY")
// t, err := UsDayOfWeekNo(0).XParseString("Tue")
// t, err := UsDayOfWeekNo(0).XParseString("TUE")
// t, err := UsDayOfWeekNo(0).XParseString("tuesday")
// t, err := UsDayOfWeekNo(0).XParseString("tue")
//
// In all of the above cases t is now equal to UsDayOfWeekNo(0).Tuesday()
//
func (usDayOfWk UsDayOfWeekNo) XParseString(
	valueString string) (UsDayOfWeekNo, error) {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	ePrefix := "UsDayOfWeekNo.XParseString() "

	usDayOfWeek, ok := mUsDayOfWeekNoStringToCode[valueString]

	if !ok {
		return UsDayOfWeekNo(-1),
			fmt.Errorf(ePrefix+
				"\n'valueString' did NOT MATCH a valid UsDayOfWeekNo Value.\n" +
				"valueString='%v'\n", valueString)
	}

	return usDayOfWeek, nil
}

// XValue - This method returns the enumeration value of the current UsDayOfWeekNo
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (usDayOfWk UsDayOfWeekNo) XValue() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	if usDayOfWk < -1 ||
		usDayOfWk > 6 {
		return UsDayOfWeekNo(-1)
	}

	return usDayOfWk
}

// XValueInt - This method returns the integer value of the current UsDayOfWeekNo
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (usDayOfWk UsDayOfWeekNo) XValueInt() int {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return int(usDayOfWk)
}

// XZero - Returns an uninitialized UsDayOfWeekNo instance.
// The returned object is set to 'None' or '-1'.
//
// This method implements the IDayOfWeekNumber.
//
// Be advised that 'None' or '-1' is treated as an invalid
// enumeration and will fail subsequent validation tests.
//
func (usDayOfWk UsDayOfWeekNo) XZero() IDayOfWeekNumber {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(-1)
}


// UsWeekDayNo - public global variable of
// type UsDayOfWeekNo.
//
// This variable serves as an easier, short hand
// technique for accessing UsDayOfWeekNo values.
//
// Usage:
// UsWeekDayNo.None(),
// UsWeekDayNo.Sunday(),
// UsWeekDayNo.Monday(),
// UsWeekDayNo.Tuesday(),
// UsWeekDayNo.Wednesday(),
// UsWeekDayNo.Thursday(),
// UsWeekDayNo.Friday(),
// UsWeekDayNo.Saturday(),
//
var UsWeekDayNo UsDayOfWeekNo
