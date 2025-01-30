package datetime

import (
	"fmt"
	"strings"
	"sync"
)

var mCalendarYearTypeStringToCode = map[string]CalendarYearNumType{
	"None"                : CalendarYearNumType(0),
	"Astronomical"        : CalendarYearNumType(1),
	"Before Common Era"   : CalendarYearNumType(2),
	"BeforeCommonEra"     : CalendarYearNumType(2),
	"Common Era"          : CalendarYearNumType(3),
	"CommonEra"           : CalendarYearNumType(3),
}

var mCalendarYearNumTypeLwrCaseStringToCode = map[string]CalendarYearNumType{
	"none"              : CalendarYearNumType(0),
	"astronomical"      : CalendarYearNumType(1),
	"before common era" : CalendarYearNumType(2),
	"beforecommonera"   : CalendarYearNumType(2),
	"common era"        : CalendarYearNumType(3),
	"commonera"         : CalendarYearNumType(3),
}

var mCalendarYearNumTypeCodeToString = map[CalendarYearNumType]string{
	CalendarYearNumType(0)  : "None",
	CalendarYearNumType(1)  : "Astronomical",
	CalendarYearNumType(2)  : "Before Common Era",
	CalendarYearNumType(3)  : "Common Era",
}


// CalendarYearNumType - An enumeration of calendar year types.
//
// The datetime software library supports two types of year numbering
// modes, 'Astronomical Year Numbering' and 'Common Era Year Numbering'.
// For more information on year numbering modes, reference:
//  Type: CalendarYearNumMode
//  Source File: datetime\calendaryearnumbermodeenum.go
//
// The first of these two modes is known as 'Astronomical Year Numbering'.
// This year numbering system includes a year zero (0). In other words,
// the date January 1, year 1 is immediately preceded by the date December
// 31, year zero (0). Likewise, the date January 1, year 0 is immediately
// preceded by the date December 31, year -1.
//
// As shown, years occurring before year zero are preceded with a minus sign
// (January 1, -1). This year numbering system is the default numbering
// system used by the datetime software library. In addition, all year values
// for the type 'CalendarDateTime' are stored internally using the
// 'Astronomical Year Numbering' system. For more information on the
// 'Astronomical Year Numbering System', reference:
//      https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
// The second type of calendar year numbering mode supported by the datetime
// Software Library is the 'Common Era' numbering system. In the 'Common Era'
// year numbering system the date January 1, year 1 is labeled as 'CE' or
// 'C.E.' designating year 1 as the first year of the Common Era. January 1,
// year 1 is immediately preceded by the date December 31, year 1 BCE.
// In other words, the date December 31, year 0 in according to 'Astronomical
// Year Numbering' is formatted as December 31, year 1 BCE in the Common Era
// Year Numbering system. Note that 'BCE' or 'B.C.E' means 'Before Common Era'
//
// For more information on the 'Common Era Year Numbering System', reference:
//      https://en.wikipedia.org/wiki/Common_Era
//
// This year number type, 'CalendarYearNumType', allows the user to designate
// one of three types of years:
//   1. Astronomical Year Numbering
//        All years less than zero, equal to zero or greater than zero prefixed
//        with the appropriate numerical sign (minus sign for negative years).
//        For example, years prior to year 0 are numbered -1, -2, -3, -4 etc.
//
//   2. Before Common Era Numbering
//        All years preceding the date January 1, year 1. These years typically
//        have the suffix BCE, for Before Common Era. For example, the date
//        immediately preceding January 1, 0001 CE is December 31, 0001 BCE.
//        year preceding year 0001 BCE are numbered as 2 BCE, 3 BCE, 4 BCE etc.
//
//   3. Common Era Numbering
//        All years following, or greater than, the date December 31, 1 BCE. These
//        years often have the suffix CE signaling they are years in the Common Era.
//        For example years following the date December 31, 1 BCE are numbered 1 CE,
//        2 CE, 3 CE, 4 CE etc.
//
// The type 'CalendarYearNumType' is structured as enumeration designating a year
// value as belonging to one of three types listed above. Since Go does not directly
// support enumerations, the 'CalendarYearNumType' structure has been adapted to
// function in a manner similar to a classic enumeration.
//
// 'CalendarYearNumType' is declared as a type 'int'. The method names are effectively
// represent an enumeration of calendar year number types.
// These methods are listed as follows:
//
//
// None              (0) - None - Signals that the Calendar Year Number Type
//                         (CalendarYearNumType) Type is not initialized. This is
//                         an error condition.
//
// Astronomical      (1) - Signals that the year number value includes a year zero.
//                         In other words, the date January 1, year 1 is immediately
//                         preceded by the date December 31, year 0. Likewise date,
//                         January 1, year 0 is immediately preceded by the date,
//                         December 31, year -1.
//                         Reference:
//                           https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
// Before Common Era (2) - Signals that the year number is less than the date
//                         January 1, 0001 CE. The date immediately preceding
//                         January 1, 0001 CE is December 31, 0001 BCE. All
//                         years before year 1 CE are numbered as 2 BCE, 3 BCE
//                         4 BCE etc. Reference:
//                            https://en.wikipedia.org/wiki/Common_Era
//
// CommonEra         (3) - Signals that the year number is greater than the date
//                         December 31, 0001 BCE. Years following the date December 31,
//                         0001 BCE are numbered as 1 CE, 2 CE, 3 CE, 4 CE etc.
//                         Reference:
//                            https://en.wikipedia.org/wiki/Common_Era
//
// For easy access to these enumeration values, use the global variable 'CalYearMode'.
// Example: CalYearType.CommonEra()
//
// Otherwise you will need to use the formal syntax.
// Example: CalendarYearNumType(0).CommonEra()
//
// Depending on your editor, intellisense (a.k.a. intelligent code completion) may not
// list the CalendarYearNumType methods in alphabetical order. Be advised that all
// 'CalendarYearNumType' methods beginning with 'X', as well as the method 'String()',
// are utility methods and not part of the enumeration values.
//
type CalendarYearNumType int

var lockCalendarYearNumType sync.Mutex

// None - Signals that the CalendarYearNumType Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (calYrNumType CalendarYearNumType) None() CalendarYearNumType {

	lockCalendarYearNumType.Lock()

	defer lockCalendarYearNumType.Unlock()

	return CalendarYearNumType(0)
}

// Astronomical - Signals that the year numbering system includes
// a year zero. In other words, the date January 1, year 1 is
// immediately preceded by the date December 31, year 0.
// As its name implies, Astronomical Year numbering is frequently
// used in astronomical and mathematical calculations.
//
// Reference:
//      https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
// This method is part of the standard enumeration.
//
func (calYrNumType CalendarYearNumType) Astronomical() CalendarYearNumType {

	lockCalendarYearNumType.Lock()

	defer lockCalendarYearNumType.Unlock()

	return CalendarYearNumType(1)
}

// BCE - Before Common Era. Signals that the year numbering system does NOT
// include a year zero and that the year value is less than the date,
// January 1, year 1 CE (Common Era). Year values preceding the date,
// January 1, year 1 CE are styled as year 1 BCE, year 2 BCE, year 3 BCE
// etc.
//
// Therefore the date immediately preceding January 1, 0001 CE is,
// December 31, 0001 BCE (Before Common Era).
//
// Reference:
//      https://en.wikipedia.org/wiki/Common_Era
//
// This method is part of the standard enumeration.
//
func (calYrNumType CalendarYearNumType) BCE() CalendarYearNumType {

	lockCalendarYearNumType.Lock()

	defer lockCalendarYearNumType.Unlock()

	return CalendarYearNumType(2)
}

// CE - Common Era. Signals that the year numbering system does NOT
// include a year zero and that the year value is greater than or
// equal to the date, January 1, year 1 CE (Common Era). Year values
// following the date, January 1, year 1 CE are styled as year 2 CE,
// year 3 CE, year 4 CE etc.
//
// Therefore the date immediately following December 31, 0001 BCE is,
// January 1, 0001 CE (Common Era).
//
// Reference:
//      https://en.wikipedia.org/wiki/Common_Era
//
// This method is part of the standard enumeration.
//
func (calYrNumType CalendarYearNumType) CE() CalendarYearNumType {

	lockCalendarYearNumType.Lock()

	defer lockCalendarYearNumType.Unlock()

	return CalendarYearNumType(3)
}


// String - Returns a string with the name of the enumeration associated
// with this instance of 'CalendarYearNumType'. The 'String' method is not
// appropriate for year number suffixes. For year number suffixes (e.g.
// "", "BCE", "CE") see the method CalendarYearNumType.XYearSuffix().
//
// 'Strings' returned by this method are:
//    1. "Astronomical"
//    2. "Before Common Era"
//    3. "Common Era"
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t := CalendarYearNumType(0).None()
// str := t.String()
//     str is now equal to 'None'
//
// t = CalendarYearNumType(0).Astronomical()
// str := t.String()
//     str is now equal to 'Astronomical'
//
// t = CalendarYearNumType(0).BCE()
// str = t.String()
//     str is now equal to 'Before Common Era'
//
// t = CalendarYearNumType(0).CE()
// str = t.String()
//     str is now equal to 'Common Era'
//
func (calYrNumType CalendarYearNumType) String() string {

	lockCalendarYearNumType.Lock()

	defer lockCalendarYearNumType.Unlock()

	result, ok := mCalendarYearNumTypeCodeToString[calYrNumType]

	if !ok {
		return ""
	}

	return result
}

// XGetCalendarYearNumberMode - Returns the Calendar Year Numbering
// Mode associated with this Calendar Year Number Type.
//
// CalendarYearNumMode - An enumeration of calendar numbering systems.
// CalendarYearNumMode is used to designate the specific type of calendar
// year numbering applied to a date time operation.
//
// Since Go does not directly support enumerations, the 'CalendarYearNumMode'
// type has been adapted to function in a manner similar to classic enumerations.
// 'CalendarYearNumMode' is declared as a type 'int'. The method names effectively
// represent an enumeration of calendar specification types. These methods are
// listed as follows:
//
//
// None             (0) - None - Signals that the Calendar Year Numbering Mode
//                        (CalendarYearNumMode) Type is not initialized. This is
//                        an error condition.
//
// Astronomical     (1) - Signals that the year numbering system includes a year
//                        zero. In other words, the date January 1, year 1 is
//                        immediately preceded by the date December 31, year 0.
//                        Reference:
//                          https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
// CommonEra        (2) - Signals that the year numbering system does NOT include
//                        a year zero. In other words, the date January 1, year 1 CE
//                        is immediately preceded by the date December 31, year 1 BCE.
//                        Reference:
//                          https://en.wikipedia.org/wiki/Common_Era
//
func (calYrNumType CalendarYearNumType) XGetCalendarYearNumberMode() CalendarYearNumMode {

	lockCalendarYearNumType.Lock()

	defer lockCalendarYearNumType.Unlock()

	if calYrNumType == 1 {
		return CalendarYearNumMode(1)
	} else if calYrNumType == 2 {
		return CalendarYearNumMode(2)
	} else if calYrNumType == 3 {
		return CalendarYearNumMode(2)
	}

	return CalendarYearNumMode(0)
}

// XIsValid - Returns a boolean value signaling whether
// the current Calendar Year Numbering Type value is
// valid.
//
// In order to be valid, Calendar Year Numbering Type
// MUST be equal to one of the three following values:
//
// 1. CalendarYearNumType(0).Astronomical() - Value = 1
//               OR
// 2. CalendarYearNumType(0).CE() - Value = 2
//               OR
// 3. CalendarYearNumType(0).BCE() - Value = 3
//
// Specifically the enumeration CalendarYearNumType(0).None()
// is considered, "INVALID".
//
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  YearNumberingSystem := CalendarYearNumType(0).Astronomical()
//
//  isValid := YearNumberingSystem.XIsValid()
//
func (calYrNumType CalendarYearNumType) XIsValid() bool {

	lockCalendarYearNumType.Lock()

	defer lockCalendarYearNumType.Unlock()

	if calYrNumType > 3 ||
		calYrNumType < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of CalendarYearNumType is returned set to the value
// of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// valueString   string - A string which will be matched against the
//                        enumeration string values. If 'valueString'
//                        is equal to one of the enumeration names, this
//                        method will proceed to successful completion
//                        and return the correct enumeration value.
//
// caseSensitive   bool - If 'true' the search for enumeration names
//                        will be case sensitive and will require an
//                        exact match. Therefore, 'astronomical' will NOT
//                        match the enumeration name, 'Astronomical'.
//
//                        If 'false' a case insensitive search is conducted
//                        for the enumeration name. In this case, 'astronomical'
//                        will match match enumeration name 'Astronomical'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// CalendarYearNumType - Upon successful completion, this method will return
//                       a new instance of CalendarYearNumType set to the value
//                       of the enumeration matched by the string search performed
//                       on input parameter, 'valueString'.
//
// error               - If this method completes successfully, the returned error
//                       Type is set equal to 'nil'. If an error condition is encountered,
//                       this method will return an error type which encapsulates an
//                       appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t, err := CalendarYearNumType(0).XParseString("Common Era", true)
//
//     t is now equal to CalendarYearNumType(0).CommonEra()
//
// t, err = CalendarYearNumType(0).XParseString("CommonEra", true)
//
//     t is now equal to CalendarYearNumType(0).CommonEra()
//
// Note that the method will successfully match the strings "Common Era"
// and "CommonEra". The same is true for "Before Common Era" and
// "BeforeCommonEra".
//
func (calYrNumType CalendarYearNumType) XParseString(
	valueString string,
	caseSensitive bool) (CalendarYearNumType, error) {

	lockCalendarYearNumType.Lock()

	defer lockCalendarYearNumType.Unlock()

	ePrefix := "CalendarYearNumType.XParseString() "

	valueString = strings.Trim(valueString, " ")

	if len(valueString) < 4 {
		return CalendarYearNumType(0).None(),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n" +
				"String length is less than '4'.\n" +
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var calendarYearNumberType CalendarYearNumType

	if caseSensitive {

		calendarYearNumberType, ok = mCalendarYearTypeStringToCode[valueString]

		if !ok {
			return CalendarYearNumType(0).None(),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid Calendar Year Numnber Type Value.\n" +
					"valueString='%v'\n", valueString)
		}

	} else {

		calendarYearNumberType, ok = mCalendarYearNumTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return CalendarYearNumType(0).None(),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid Calendar Year Number Type Value.\n" +
					"valueString='%v'\n", valueString)
		}
	}

	return calendarYearNumberType, nil
}


// XValue - This method returns the enumeration value of the current
// CalendarYearNumType instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (calYrNumType CalendarYearNumType) XValue() CalendarYearNumType {

	lockCalendarYearNumType.Lock()

	defer lockCalendarYearNumType.Unlock()

	return calYrNumType
}


// XValueInt - This method returns the integer value of the current
// CalendarYearNumType instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (calYrNumType CalendarYearNumType) XValueInt() int {

	lockCalendarYearNumType.Lock()

	defer lockCalendarYearNumType.Unlock()

	return int(calYrNumType)
}

// XYearSuffix - Returns a string suitable of use as a year suffix or
// year label.
//
// Years employing the Astronomical Year Numbering system have no year
// suffix and this method therefore returns an empty string ("") for
// 'Astronomical' years.
//
// Years associated with the Type CalendarYearNumType(0).BCE will receive
// a year suffix of 'BCE'.
//
// Years associated with the Type CalendarYearNumType(0).CE will receive
// a year suffix of 'CE'.
//
// All other 'CalendarYearNumType' values will return the string, "ERROR".
//
func (calYrNumType CalendarYearNumType) XYearSuffix() string {

	lockCalendarYearNumType.Lock()

	defer lockCalendarYearNumType.Unlock()

	if calYrNumType == 1 {
		// Astronomical
		return ""
	} else if calYrNumType == 2 {
		return "BCE"
	} else if calYrNumType == 3 {
		return "CE"
	}

	return "ERROR"
}

// CalYearType - public global variable of type,
// CalendarYearNumType.
//
// This variable serves as an easier, short hand
// technique for accessing CalendarYearNumType
// values.
//
// Usage:
// yearType := CalYearType.None()
// yearType = CalYearType.Astronomical()
// yearType = CalYearType.BCE()
// yearType = CalYearType.CE()
//
var CalYearType CalendarYearNumType

