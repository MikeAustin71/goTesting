package datetime

import (
	"fmt"
	"sync"
)

var mDayOfWeekNumberingSystemTypeStringToCode = map[string] DayOfWeekNumberingSystemType{
	"None"                      : DayOfWeekNumberingSystemType(0),
	"NONE"                      : DayOfWeekNumberingSystemType(0),
	"none"                      : DayOfWeekNumberingSystemType(0),
	"US Day Of Week"            : DayOfWeekNumberingSystemType(1),
	"US Day Of The Week"        : DayOfWeekNumberingSystemType(1),
	"USDayOfWeek"               : DayOfWeekNumberingSystemType(1),
	"USDayOfTheWeek"            : DayOfWeekNumberingSystemType(1),
	"US DAY OF WEEK"            : DayOfWeekNumberingSystemType(1),
	"USDAYOFWEEK"               : DayOfWeekNumberingSystemType(1),
	"US DAY OF THE WEEK"        : DayOfWeekNumberingSystemType(1),
	"USDAYOFTHEWEEK"            : DayOfWeekNumberingSystemType(1),
	"us day of week"            : DayOfWeekNumberingSystemType(1),
	"us day of the week"        : DayOfWeekNumberingSystemType(1),
	"usdayofweek"               : DayOfWeekNumberingSystemType(1),
	"usdayoftheweek"            : DayOfWeekNumberingSystemType(1),
	"ISO 8601 Day Of Week"      : DayOfWeekNumberingSystemType(2),
	"ISO8601DayOfWeek"          : DayOfWeekNumberingSystemType(2),
	"ISO8601DayOfTheWeek"       : DayOfWeekNumberingSystemType(2),
	"ISO 8601 DAY OF WEEK"      : DayOfWeekNumberingSystemType(2),
	"ISO 8601 DAY OF THE WEEK"  : DayOfWeekNumberingSystemType(2),
	"ISO8601DAYOFWEEK"          : DayOfWeekNumberingSystemType(2),
	"ISO8601DAYOFTHEWEEK"       : DayOfWeekNumberingSystemType(2),
	"iso 8601 day of week"      : DayOfWeekNumberingSystemType(2),
	"iso 8601 day of the week"  : DayOfWeekNumberingSystemType(2),
	"iso8601dayofweek"          : DayOfWeekNumberingSystemType(2),
	"iso8601dayoftheweek"       : DayOfWeekNumberingSystemType(2),
}

var mDayOfWeekNumberingSystemTypeCodeToString = map[DayOfWeekNumberingSystemType]string{
	DayOfWeekNumberingSystemType(0)  : "None",
	DayOfWeekNumberingSystemType(1)  : "US Day Of Week",
	DayOfWeekNumberingSystemType(2)  : "ISO 8601 Day Of Week",
}


// DayOfWeekNumberingSystemType - An enumeration of different numbering systems
// used to number the days of the week.
//
// Currently, the datetime software library supports two types of
// Day of the Week Numbering Systems:
//
//   1. US Day Of The Week Numbering System.
//
//   2. ISO 8601 Standard Day of the Week Numbering System.
//
//
// US Day Of The Week Numbering System
//
// The United States, Canada, Australia and New Zealand put Sunday as the
// first day of the week on their calendars. The first day of the week,
// 'Sunday', is numbered as zero with the last day of the week being numbered
// as 6 for 'Saturday'. US Day of the Week Numbers are listed below:
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
// For more information on the US Day of the Week Numbering System,
// reference type UsDayOfWeekNo, Source Code File:
//    datetime/dayofweeknumberusenum.go
//
//
// ISO 8601 Standard Day Of The Week Numbering System
//
// The most common day of the week numbering system used internationally,
// is the ISO 8601 standard which specifies that the week begins on Monday.
// ISO stands for the International Organization for Standardization (ISO).
//    https://www.iso.org/home.html
//
// ISO 8601 standard is used in Western Europe, Scandinavia, and most of
// Eastern Europe as well as many other nations across the globe. It is
// considered an international standard.
//
// Under the ISO 8601 Standard, days of the week are numbered beginning with
// one (1) for Monday and ending with seven (7) for Sunday. ISO 8601 week day
// numbers are listed as follows:
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
// For more information on the ISO 8601 Standard, reference type,
// ISO8601DayOfWeekNo defined in source file:
//    datetime/dayofweeknumberiso8601enum.go
//
//
// Since the Go Programming Language does not directly support enumerations,
// the DayOfWeekNumberingSystemType type has been adapted to function in a
// manner similar to classic enumerations. DayOfWeekNumberingSystemType is
// declared as a type 'int'. The method names effectively represent an
// enumeration of Day Of The Week Numbering Systems supported by the datetime
// software library. These methods are listed as follows:
//
//  None              (0) - None - Signals that the Day Of the Week Numbering System
//                          type, DayOfWeekNumberingSystemType is not initialized.
//                          This is an error condition.
//
//  UsDayOfWeek       (1) - Signals that the DayOfWeekNumberingSystemType has been
//                          correctly initialized and set to the US Day of the Week
//                          numbering system.
//
//  ISO8601DayOfWeek  (2) - Signals that the DayOfWeekNumberingSystemType has been
//                          correctly initialized and set to the ISO 8601 Standard
//                          for week day numbering.
//
//
// Resources:
//   https://en.wikipedia.org/wiki/ISO_8601#Week_dates
//   https://www.timeanddate.com/date/week-numbers.html
//
// Usage:
//
// For easy access to these enumeration values, use the global variable 'DayOfWeekNoType'.
// Examples:
//      DayOfWeekNoType.UsDayOfWeek()
//      DayOfWeekNoType.ISO8601DayOfWeek()
//
// Otherwise you will need to use the formal syntax.
//  Example: DayOfWeekNumberingSystemType(0).ISO8601DayOfWeek()
//
// Depending on your editor, intellisense (a.k.a. intelligent code completion) may not
// list the DayOfWeekNumberingSystemType methods in alphabetical order.
//
// Finally, be advised that all DayOfWeekNumberingSystemType methods beginning with 'X', as well
// as the method 'String()' are utility methods and not part of the enumeration.
//
type DayOfWeekNumberingSystemType int

var lockDayOfWeekNumberingSystemType sync.Mutex

// None - Signals that the DayOfWeekNumberingSystemType Type is uninitialized
// and invalid. This is an error condition.
//
// This method is part of the standard enumeration.
//
func (dayOfWeekNoSys DayOfWeekNumberingSystemType) None() DayOfWeekNumberingSystemType {

	lockDayOfWeekNumberingSystemType.Lock()

	defer lockDayOfWeekNumberingSystemType.Unlock()

	return DayOfWeekNumberingSystemType(0)
}

// UsDayOfWeek - Signals that the DayOfWeekNumberingSystemType instance designates
// the Day of the Week Numbering System as, US Day of the Week. This week day
// numbering system is implemented in the United States, Canada, Australia, and
// New Zealand.
//
// The US Day of the Week Numbering puts Sunday as the first day of the week on
// its calendars. The first day of the week,'Sunday', is numbered as zero (0)
// while the last day of the week being numbered as six (6), for 'Saturday'.
//
// US Day of the Week Numbers are listed below:
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
// For more information on the US Day of the Week Numbering System,
// reference type UsDayOfWeekNo, Source Code File:
//    datetime/dayofweeknumberusenum.go
//
// This method is part of the standard enumeration.
//
func (dayOfWeekNoSys DayOfWeekNumberingSystemType) UsDayOfWeek() DayOfWeekNumberingSystemType {

	lockDayOfWeekNumberingSystemType.Lock()

	defer lockDayOfWeekNumberingSystemType.Unlock()

	return DayOfWeekNumberingSystemType(1)
}

// ISO8601DayOfWeek - Signals that the DayOfWeekNumberingSystemType instance designates
// the Day of the Week Numbering System as, ISO 8601 Standard. This day of the week
// numbering system is used in Western Europe, Scandinavia, and most of Eastern Europe
// as well as many other nations across the globe. It is considered an international
// standard.
//
//
// Under the ISO 8601 Standard, days of the week are numbered beginning with
// one (1) for Monday and ending with seven (7) for Sunday. ISO 8601 week day
// numbers are listed as follows:
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
// For more information on the ISO 8601 Standard, reference type,
// ISO8601DayOfWeekNo defined in source file:
//    datetime/dayofweeknumberiso8601enum.go
//
// This method is part of the standard enumeration.
//
func (dayOfWeekNoSys DayOfWeekNumberingSystemType) ISO8601DayOfWeek() DayOfWeekNumberingSystemType {

	lockDayOfWeekNumberingSystemType.Lock()

	defer lockDayOfWeekNumberingSystemType.Unlock()

	return DayOfWeekNumberingSystemType(2)
}

// String - Returns a string with the day-of-week name associated with
// this DayOfWeekNumberingSystemType enumeration instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= DayOfWeekNumberingSystemType(0).UsDayOfWeek()
// str := t.String()
//     str is now equal to "US Day Of Week"
//
// t = DayOfWeekNumberingSystemType(0).ISO8601DayOfWeek()
// str = t.String()
//     str is now equal to "ISO 8601 Day Of Week"
func (dayOfWeekNoSys DayOfWeekNumberingSystemType) String() string {

	lockDayOfWeekNumberingSystemType.Lock()

	defer lockDayOfWeekNumberingSystemType.Unlock()

	result, ok := mDayOfWeekNumberingSystemTypeCodeToString[dayOfWeekNoSys]

	if !ok {
		return fmt.Sprintf("Error: Value= %d", int(dayOfWeekNoSys))
	}

	return result
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
func (dayOfWeekNoSys DayOfWeekNumberingSystemType) XIsValid() bool {

	lockDayOfWeekNumberingSystemType.Lock()

	defer lockDayOfWeekNumberingSystemType.Unlock()

	if dayOfWeekNoSys > 2 ||
		dayOfWeekNoSys < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of DayOfWeekNumberingSystemType is returned set
// to the value of the associated enumeration.
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
//                        is equal to a day of the week numbering system
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
// DayOfWeekNumberingSystemType
//     - Upon successful completion, this method will return a new instance
//       of DayOfWeekNumberingSystemType set to the value of the enumeration
//       matched by the string search performed on input parameter,
//       'valueString'.
//
// error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error type which encapsulates an
//       appropriate error message.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
// All of these examples will be successfully processed.
//
// t, err := DayOfWeekNumberingSystemType(0).XParseString("US Day Of Week")
// t, err = DayOfWeekNumberingSystemType(0).XParseString("US Day Of The Week")
// t, err = DayOfWeekNumberingSystemType(0).XParseString("USDayOfWeek")
// t, err = DayOfWeekNumberingSystemType(0).XParseString("TUE")
// t, err = DayOfWeekNumberingSystemType(0).XParseString("USDayOfTheWeek")
// t, err = DayOfWeekNumberingSystemType(0).XParseString("US DAY OF WEEK")
// t, err = DayOfWeekNumberingSystemType(0).XParseString("US DAY OF THE WEEK")
// t, err = DayOfWeekNumberingSystemType(0).XParseString("USDAYOFTHEWEEK")
// t, err = DayOfWeekNumberingSystemType(0).XParseString("us day of week")
// t, err = DayOfWeekNumberingSystemType(0).XParseString("us day of the week")
// t, err = DayOfWeekNumberingSystemType(0).XParseString("usdayofweek")
// t, err = DayOfWeekNumberingSystemType(0).XParseString("usdayoftheweek")
//
// In all of the above cases t is now equal to DayOfWeekNumberingSystemType(0).UsDayOfWeek()
//
func (dayOfWeekNoSys DayOfWeekNumberingSystemType) XParseString(
	valueString string) (DayOfWeekNumberingSystemType, error) {

	lockDayOfWeekNumberingSystemType.Lock()

	defer lockDayOfWeekNumberingSystemType.Unlock()

	ePrefix := "DayOfWeekNumberingSystemType.XParseString() "

	dayOfWeekNumSysType, ok := mDayOfWeekNumberingSystemTypeStringToCode[valueString]

	if !ok {
		return DayOfWeekNumberingSystemType(0),
			fmt.Errorf(ePrefix+
				"\n'valueString' did NOT MATCH a valid DayOfWeekNumberingSystemType Value.\n" +
				"valueString='%v'\n", valueString)
	}

	return dayOfWeekNumSysType, nil
}

// XValue - This method returns the enumeration value of the current DayOfWeekNumberingSystemType
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (dayOfWeekNoSys DayOfWeekNumberingSystemType) XValue() DayOfWeekNumberingSystemType {

	lockDayOfWeekNumberingSystemType.Lock()

	defer lockDayOfWeekNumberingSystemType.Unlock()

	return dayOfWeekNoSys
}

// XValueInt - This method returns the integer value of the current DayOfWeekNumberingSystemType
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (dayOfWeekNoSys DayOfWeekNumberingSystemType) XValueInt() int {

	lockDayOfWeekNumberingSystemType.Lock()

	defer lockDayOfWeekNumberingSystemType.Unlock()

	return int(dayOfWeekNoSys)
}


// DayOfWeekNoType - public global variable of
// type DayOfWeekNumberingSystemType.
//
// This variable serves as an easier, short hand
// technique for accessing DayOfWeekNumberingSystemType values.
//
// Usage:
// DayOfWeekNoType.None(),
// DayOfWeekNoType.UsDayOfWeek(),
// DayOfWeekNoType.ISO8601DayOfWeek(),
//
var DayOfWeekNoType DayOfWeekNumberingSystemType

