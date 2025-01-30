package datetime

import (
	"fmt"
	"sync"
)


var mIso8601DayOfWeekNoStringToCode = map[string]ISO8601DayOfWeekNo{
	"NONE"           : ISO8601DayOfWeekNo(0),
	"MONDAY"         : ISO8601DayOfWeekNo(1),
	"TUESDAY"        : ISO8601DayOfWeekNo(2),
	"WEDNESDAY"      : ISO8601DayOfWeekNo(3),
	"THURSDAY"       : ISO8601DayOfWeekNo(4),
	"FRIDAY"         : ISO8601DayOfWeekNo(5),
	"SATURDAY"       : ISO8601DayOfWeekNo(6),
	"SUNDAY"         : ISO8601DayOfWeekNo(7),
	"None"           : ISO8601DayOfWeekNo(0),
	"Monday"         : ISO8601DayOfWeekNo(1),
	"Tuesday"        : ISO8601DayOfWeekNo(2),
	"Wednesday"      : ISO8601DayOfWeekNo(3),
	"Thursday"       : ISO8601DayOfWeekNo(4),
	"Friday"         : ISO8601DayOfWeekNo(5),
	"Saturday"       : ISO8601DayOfWeekNo(6),
	"Sunday"         : ISO8601DayOfWeekNo(7),
	"none"           : ISO8601DayOfWeekNo(0),
	"monday"         : ISO8601DayOfWeekNo(1),
	"tuesday"        : ISO8601DayOfWeekNo(2),
	"wednesday"      : ISO8601DayOfWeekNo(3),
	"thursday"       : ISO8601DayOfWeekNo(4),
	"friday"         : ISO8601DayOfWeekNo(5),
	"saturday"       : ISO8601DayOfWeekNo(6),
	"sunday"         : ISO8601DayOfWeekNo(7),
	"MON"            : ISO8601DayOfWeekNo(1),
	"TUE"            : ISO8601DayOfWeekNo(2),
	"WED"            : ISO8601DayOfWeekNo(3),
	"THU"            : ISO8601DayOfWeekNo(4),
	"THURS"          : ISO8601DayOfWeekNo(4),
	"FRI"            : ISO8601DayOfWeekNo(5),
	"SAT"            : ISO8601DayOfWeekNo(6) ,
	"SUN"            : ISO8601DayOfWeekNo(7),
	"Mon"            : ISO8601DayOfWeekNo(1),
	"Tue"            : ISO8601DayOfWeekNo(2),
	"Wed"            : ISO8601DayOfWeekNo(3),
	"Thu"            : ISO8601DayOfWeekNo(4),
	"Thurs"          : ISO8601DayOfWeekNo(4),
	"Fri"            : ISO8601DayOfWeekNo(5),
	"Sat"            : ISO8601DayOfWeekNo(6) ,
	"Sun"            : ISO8601DayOfWeekNo(7),
	"mon"            : ISO8601DayOfWeekNo(1),
	"tue"            : ISO8601DayOfWeekNo(2),
	"wed"            : ISO8601DayOfWeekNo(3),
	"thu"            : ISO8601DayOfWeekNo(4),
	"thurs"          : ISO8601DayOfWeekNo(4),
	"fri"            : ISO8601DayOfWeekNo(5),
	"sat"            : ISO8601DayOfWeekNo(6) ,
	"sun"            : ISO8601DayOfWeekNo(7),
}


var mIso8601DayOfWeekNoCodeToString = map[ISO8601DayOfWeekNo]string{
	ISO8601DayOfWeekNo(0)   : "None",
	ISO8601DayOfWeekNo(1)   : "Monday",
	ISO8601DayOfWeekNo(2)   : "Tuesday",
	ISO8601DayOfWeekNo(3)   : "Wednesday",
	ISO8601DayOfWeekNo(4)   : "Thursday",
	ISO8601DayOfWeekNo(5)   : "Friday",
	ISO8601DayOfWeekNo(7)   : "Sunday",
}

var mIso8601DayOfWeekNoCodeTo3CharAbbrvDay = map[ISO8601DayOfWeekNo]string{
	ISO8601DayOfWeekNo(1)   : "Mon",
	ISO8601DayOfWeekNo(2)   : "Tue",
	ISO8601DayOfWeekNo(3)   : "Wed",
	ISO8601DayOfWeekNo(4)   : "Thu",
	ISO8601DayOfWeekNo(5)   : "Fri",
	ISO8601DayOfWeekNo(6)   : "Sat",
	ISO8601DayOfWeekNo(7)   : "Sun",
}

var mIso8601DayOfWeekNoCodeTo2CharAbbrvDay = map[ISO8601DayOfWeekNo]string{
	ISO8601DayOfWeekNo(1)   : "Mo",
	ISO8601DayOfWeekNo(2)   : "Tu",
	ISO8601DayOfWeekNo(3)   : "We",
	ISO8601DayOfWeekNo(4)   : "Th",
	ISO8601DayOfWeekNo(5)   : "Fr",
	ISO8601DayOfWeekNo(6)   : "Sa",
	ISO8601DayOfWeekNo(7)   : "Su",
}

// ISO8601DayOfWeekNo - An enumeration of ISO 8601 Standard, Day Of The
// Week numbers. ISO stands for the International Organization for
// Standardization (ISO). https://www.iso.org/home.html
//
// The most common day of the week numbering system used internationally,
// is the ISO 8601 standard. This standard is used in Western Europe,
// Scandinavia, and most of Eastern Europe as well as many other nations
// across the globe.
//
// The ISO 8601 standard specifies that the week begins on Monday. Days
// of the week are numbered beginning with one (1) for Monday and ending
// with seven (7) for Sunday. Western European Calendars therefore show
// the first day of the week as Monday.
//
// In contrast, calendars in the United States, Canada, Australia and
// New Zealand put Sunday as the first day the week. These countries
// number the days of the week beginning with zero (0) for Sunday and
// ending with Saturday as day number 6. This method of numbering week
// days is styled as the US Day Of The Week Numbering System. For more
// information on the US Day Of The Week  Numbering System, see type
// UsDayOfWeekNo defined in source file:
//    datetime/dayofweeknumberusenum.go
//
// The ISO8601DayOfWeekNo type defined here provides an enumeration for
// ISO 8601 Day Of The Week numbers as listed below:
//
//              Week         Week Day
//              Day           Number
//            ========     ===========
//             Monday          = 1
//             Tuesday         = 2
//             Wednesday       = 3
//             Thursday        = 4
//             Friday          = 5
//             Saturday        = 6
//             Sunday          = 7
//
//
// Since the Go Programming Language does not directly support enumerations,
// the ISO8601DayOfWeekNo type has been adapted to function in a manner
// similar to classic enumerations. ISO8601DayOfWeekNo is declared as a
// type 'int'. The method names effectively represent an enumeration of
// ISO 8601 Day Of The Week numbers. These methods are listed as follows:
//
//  None            (0) - None - Signals that the ISO 8601 Day Of The Week Number
//                        (ISO8601DayOfWeekNo) Type is not initialized. This is
//                        an error condition.
//
//  Monday          (1) - Signals that the ISO8601DayOfWeekNo instance is equal
//                        to 'Monday', with a weekday number value of '1'. Under
//                        the ISO 8601 standard, 'Monday' is the first day of the
//                        week.
//
//  Tuesday         (2) - Signals that the ISO8601DayOfWeekNo instance is equal
//                        to 'Tuesday', with a weekday number value of '2'.
//
//  Wednesday       (3) - Signals that the ISO8601DayOfWeekNo instance is equal
//                        to 'Wednesday', with a weekday number value of '3'.
//
//  Thursday        (4) - Signals that the ISO8601DayOfWeekNo instance is equal
//                        to 'Thursday', with a weekday number value of '4'.
//
//  Friday          (5) - Signals that the ISO8601DayOfWeekNo instance is equal
//                        to 'Friday', with a weekday number value of '5'.
//
//  Saturday        (6) - Signals that the ISO8601DayOfWeekNo instance is equal
//                        to 'Saturday', with a weekday number value of '6'.
//
//  Sunday          (7) - In accordance with the ISO 8601 Day Of The week numbering
//                        system, Sunday is the last day of the week and is
//                        represented by a week day number value of '7'.
//
// Resources:
//   https://en.wikipedia.org/wiki/ISO_8601#Week_dates
//   https://www.timeanddate.com/date/week-numbers.html
//
//
// Usage:
//
// For easy access to these enumeration values, use the global variable 'Iso8601'.
// Example: Iso8601WeekDayNo.Tuesday()
//
// Otherwise you will need to use the formal syntax.
//  Example: ISO8601DayOfWeekNo(0).Tuesday()
//
// Depending on your editor, intellisense (a.k.a. intelligent code completion) may not
// list the 'ISO8601DayOfWeekNo' methods in alphabetical order.
//
// Finally, be advised that all ISO8601DayOfWeekNo methods beginning with 'X', as well
// as the method 'String()' are utility methods and not part of the enumeration.
//
type ISO8601DayOfWeekNo int

var lockIso8601DayOfWeekNo sync.Mutex

// None - Signals that the ISO8601DayOfWeekNo Type is uninitialized
// and invalid. This is an error condition.
//
// This method is part of the standard enumeration.
//
func (iso8601DofWeek ISO8601DayOfWeekNo) None() ISO8601DayOfWeekNo {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return ISO8601DayOfWeekNo(0)
}

// Monday - Signals that the ISO8601DayOfWeekNo instance is equal
// to 'Monday', with a weekday number value of one (1).
//
// The ISO 8601 standard specifies Monday as the first day of the
// week.
//
// This method is part of the standard enumeration.
//
func (iso8601DofWeek ISO8601DayOfWeekNo) Monday() ISO8601DayOfWeekNo {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return ISO8601DayOfWeekNo(1)
}

// Tuesday - Signals that the ISO8601DayOfWeekNo instance is equal
// to 'Tuesday', with a weekday number value of '2'.
//
// This method is part of the standard enumeration.
//
func (iso8601DofWeek ISO8601DayOfWeekNo) Tuesday() ISO8601DayOfWeekNo {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return ISO8601DayOfWeekNo(2)
}

// Wednesday - Signals that the ISO8601DayOfWeekNo instance is equal
// to 'Wednesday', with a weekday number value of '3'.
//
// This method is part of the standard enumeration.
//
func (iso8601DofWeek ISO8601DayOfWeekNo) Wednesday() ISO8601DayOfWeekNo {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return ISO8601DayOfWeekNo(3)
}

// Thursday - Signals that the ISO8601DayOfWeekNo instance is equal
// to 'Thursday', with a weekday number value of '4'.
//
// This method is part of the standard enumeration.
//
func (iso8601DofWeek ISO8601DayOfWeekNo) Thursday() ISO8601DayOfWeekNo {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return ISO8601DayOfWeekNo(4)
}

// Friday - Signals that the UsDayOfWeekNo instance is equal
// to 'Friday', with a weekday number value of '5'.
//
// This method is part of the standard enumeration.
//
func (iso8601DofWeek ISO8601DayOfWeekNo) Friday() ISO8601DayOfWeekNo {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return ISO8601DayOfWeekNo(5)
}

// Saturday - Signals that the ISO8601DayOfWeekNo instance is equal
// to 'Saturday', with a weekday number value of six (6).
//
// This method is part of the standard enumeration.
//
func (iso8601DofWeek ISO8601DayOfWeekNo) Saturday() ISO8601DayOfWeekNo {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return ISO8601DayOfWeekNo(6)
}

// Sunday - Signals that the ISO8601DayOfWeekNo instance is
// equal to 'Sunday', with a day number value of seven (7).
// Under the ISO 8601 Weekday Numbering System, 'Sunday' is
// the last day of the week.
//
// This method is part of the standard enumeration.
//
func (iso8601DofWeek ISO8601DayOfWeekNo) Sunday() ISO8601DayOfWeekNo {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return ISO8601DayOfWeekNo(7)
}

// String - Returns a string with the day-of-week name associated with this
// ISO8601DayOfWeekNo enumeration value.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= ISO8601DayOfWeekNo(0).Monday()
// str := t.String()
//     str is now equal to "Monday"
//
// t= ISO8601DayOfWeekNo(0).Tuesday()
// str = t.String()
//     str is now equal to "Tuesday"
//
func (iso8601DofWeek ISO8601DayOfWeekNo) String() string {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	result, ok := mIso8601DayOfWeekNoCodeToString[iso8601DofWeek]

	if !ok {
		return fmt.Sprintf("Error: Value= %d", int(iso8601DofWeek))
	}

	return result
}

// XDayOfWeekName - Returns a string with the full day-of-week name associated
// with this ISO8601DayOfWeekNo enumeration value.
//
// This method provides an identical result with that generated by 
// method 'ISO8601DayOfWeekNo.String()'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= ISO8601DayOfWeekNo(0).Monday()
// str := t.XDayOfWeekName()
//     str is now equal to "Monday"
//
func (iso8601DofWeek ISO8601DayOfWeekNo) XDayOfWeekName() string {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	result, ok := mIso8601DayOfWeekNoCodeToString[iso8601DofWeek]

	if !ok {
		return fmt.Sprintf("Error: Value= %d", int(iso8601DofWeek))
	}

	return result
}

// XDayOfWeekNameAbbreviation - Returns a string with the abbreviated name of
// the day of the week associated with this ISO8601DayOfWeekNo enumeration value.
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
// t:= ISO8601DayOfWeekNo(0).Monday()
//
// str, err := t.XDayOfWeekNameAbbreviation(
//             3,
//            "callingFunctionName")
//
//     str is now equal to "Mon"
//
// str, err = t.XDayOfWeekNameAbbreviation(
//        2,
//        "callingFunctionName")
//
//     str is now equal to 'Mo'
//
func (iso8601DofWeek ISO8601DayOfWeekNo) XDayOfWeekNameAbbreviation(
	numberOfCharsInAbbreviation int,
	ePrefix string) (
	dayOfWeekNameAbbrv string,
	err error) {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	ePrefix += "ISO8601DayOfWeekNo.XDayOfWeekNameAbbreviation() "
	err = nil
	dayOfWeekNameAbbrv = ""

	var ok bool

	switch numberOfCharsInAbbreviation {

	case 2:

		dayOfWeekNameAbbrv,
			ok = mIso8601DayOfWeekNoCodeTo2CharAbbrvDay[iso8601DofWeek]

	case 3:

		dayOfWeekNameAbbrv,
			ok = mIso8601DayOfWeekNoCodeTo3CharAbbrvDay[iso8601DofWeek]

	default:

		ok = true

		err = fmt.Errorf(ePrefix + "\n" +
			"ERROR: Input parameter 'numberOfCharsInAbbreviation' is INVALID!\n" +
			"Valid values for 'numberOfCharsInAbbreviation are '2' or '3'.\n" +
			"numberOfCharsInAbbreviation='%v'\n",
			numberOfCharsInAbbreviation)

	}

	dayOfWeekNameAbbrv, ok = mIso8601DayOfWeekNoCodeTo3CharAbbrvDay[iso8601DofWeek]

	if !ok {
		err =
			fmt.Errorf(ePrefix + "\n" +
				"Error: ISO8601DayOfWeekNo is INVALID!\n" +
				"usDayOfWk= %d", int(iso8601DofWeek))
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
// If the current instance of ISO8601DayOfWeekNo is valid, one of
// the integer values listed above will be returned. To verify validity,
// call method func ISO8601DayOfWeekNo.XIsValid().
//
func (iso8601DofWeek ISO8601DayOfWeekNo) XDayOfWeekNumber() int {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return int(iso8601DofWeek)
}

// XDayOfWeekNumberingSystemType - Returns the Day of the Week
// Numbering System Type code associated with ISO8601DayOfWeekNo. Therefore,
// the value returned by this method is always the enumeration value,
// DayOfWeekNumberingSystemType.ISO8601DayOfWeekNo().
//
func (iso8601DofWeek ISO8601DayOfWeekNo) XDayOfWeekNumberingSystemType() DayOfWeekNumberingSystemType {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return DayOfWeekNumberingSystemType(0).ISO8601DayOfWeek()
}

// XIsValid - Returns a boolean value signaling whether the current
// ISO 8601 Day of Week Number value is valid.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  iso8601DayOfWeek := ISO8601DayOfWeekNo(0).Sunday()
//
//  isValid := iso8601DayOfWeek.XIsValid()
//    If 'isValid' is false it signals that 'iso8601DayOfWeek' contains
//    an invalid value.
//
func (iso8601DofWeek ISO8601DayOfWeekNo) XIsValid() bool {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	if iso8601DofWeek > 7 ||
		iso8601DofWeek < 1 {
		return false
	}

	return true
}


// XParseString - Receives a string and attempts to match it with the string
// value of a supported enumeration. If successful, a new instance of
// ISO8601DayOfWeekNo is returned set to the value of the associated
// enumeration.
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
// ISO8601DayOfWeekNo - Upon successful completion, this method will return a new
//                      instance of ISO8601DayOfWeekNo set to the value of the enumeration
//                      matched by the string search performed on input parameter,
//                      'valueString'.
//
// error              - If this method completes successfully, the returned error
//                      Type is set equal to 'nil'. If an error condition is encountered,
//                      this method will return an error type which encapsulates an
//                      appropriate error message.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
// All of these examples will be successfully processed.
//
// t, err := ISO8601DayOfWeekNo(0).XParseString("Tuesday")
// t, err := ISO8601DayOfWeekNo(0).XParseString("TUESDAY")
// t, err := ISO8601DayOfWeekNo(0).XParseString("Tue")
// t, err := ISO8601DayOfWeekNo(0).XParseString("TUE")
// t, err := ISO8601DayOfWeekNo(0).XParseString("tuesday")
// t, err := ISO8601DayOfWeekNo(0).XParseString("tue")
//
// In all of the above cases t is now equal to
// ISO8601DayOfWeekNo(0).Tuesday().
//
func (iso8601DofWeek ISO8601DayOfWeekNo) XParseString(
	valueString string) (ISO8601DayOfWeekNo, error) {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	ePrefix := "ISO8601DayOfWeekNo.XParseString() "

	iso8601DayOfWeek, ok := mIso8601DayOfWeekNoStringToCode[valueString]

	if !ok {
		return ISO8601DayOfWeekNo(0),
			fmt.Errorf(ePrefix+
				"\n'valueString' did NOT MATCH a valid ISO8601DayOfWeekNo Value.\n" +
				"valueString='%v'\n", valueString)
	}

	return iso8601DayOfWeek, nil
}

// XUsDayOfWeekNumber - Returns the U.S. (United States) Day of
// the Week Number as an integer value.
//
// The ISO8601DayOfWeekNo type returns the Day of the Week Number
// used in Western Europe, Scandinavia, Eastern Europe and many
// other nations across the globe. This method returns the
// equivalent Day of the Week Number for an alternative numbering
// system referred to as the US Day of the Week Numbering System.
// The US Day of the Week Numbering System is used in the United
// States, Canada, Australia, and New Zealand.
//
// The U.S. Day of Week Numbering System specifies that the first
// day of the week is 'Sunday' and is numbered as weekday number
// zero (0).
//
// 'Saturday' is the last day of the week and has a week day
// number value of six (6).
//
// Reference:
//   https://www.timeanddate.com/date/week-numbers.html
//   https://en.wikipedia.org/wiki/Julian_day
//   Type: UsDayOfWeekNo Source File: datetime/dayofweeknumberusenum.go
//
func (iso8601DofWeek ISO8601DayOfWeekNo) XUsDayOfWeekNumber() (UsDayOfWeekNo, error) {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	ePrefix := "ISO8601DayOfWeekNo.XUsDayOfWeekNumber() "

	dayNum := int(iso8601DofWeek)

	if dayNum > 7 {
		return UsDayOfWeekNo(-1),
			fmt.Errorf(ePrefix + "\n" +
				"Error: The ISO 8601 Day Of The Week Number is Greater Than seven (7) and therefore, INVALID!\n" +
				"ISO 8601 Day Of The Week Number='%v'\n", dayNum)
	}

	if dayNum < 1 {
		return UsDayOfWeekNo(-1),
			fmt.Errorf(ePrefix + "\n" +
				"Error: The ISO 8601 Day Of The Week Number is Less Than one (1) and therefore, INVALID!\n" +
				"U.S. Day Of The Week Number='%v'\n", dayNum)
	}

	if dayNum == 7 {
		dayNum = 0
	}

	return UsDayOfWeekNo(dayNum), nil
}

// XValue - This method returns the enumeration value of the current
// ISO8601DayOfWeekNo instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (iso8601DofWeek ISO8601DayOfWeekNo) XValue() ISO8601DayOfWeekNo {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return iso8601DofWeek
}


// XValueInt - This method returns the integer value of the current UsDayOfWeekNo
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (iso8601DofWeek ISO8601DayOfWeekNo) XValueInt() int {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return int(iso8601DofWeek)
}

// XZero - Returns an uninitialized ISO8601DayOfWeekNo instance.
// The returned object is set to 'None' or '0'.
//
// This method implements the IDayOfWeekNumber.
//
// Be advised that 'None' or '0' is treated as an invalid
// enumeration and will fail subsequent validation tests.
//
func (iso8601DofWeek ISO8601DayOfWeekNo) XZero() IDayOfWeekNumber {

	lockIso8601DayOfWeekNo.Lock()

	defer lockIso8601DayOfWeekNo.Unlock()

	return ISO8601DayOfWeekNo(0)
}


// Iso8601WeekDayNo - public global variable of
// type ISO8601DayOfWeekNo.
//
// This variable serves as an easier, short hand
// technique for accessing ISO8601DayOfWeekNo values.
//
// Usage:
// Iso8601WeekDayNo.None(),
// Iso8601WeekDayNo.Sunday(),
// Iso8601WeekDayNo.Monday(),
// Iso8601WeekDayNo.Tuesday(),
// Iso8601WeekDayNo.Wednesday(),
// Iso8601WeekDayNo.Thursday(),
// Iso8601WeekDayNo.Friday(),
// Iso8601WeekDayNo.Saturday(),
//
var Iso8601WeekDayNo ISO8601DayOfWeekNo


