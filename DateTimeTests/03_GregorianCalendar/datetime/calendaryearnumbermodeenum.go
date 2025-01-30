package datetime

import (
	"fmt"
	"strings"
	"sync"
)

var mCalendarYearNumModeStringToCode = map[string]CalendarYearNumMode{
	"None"           : CalendarYearNumMode(0),
	"Astronomical"   : CalendarYearNumMode(1),
	"CommonEra"      : CalendarYearNumMode(2),
}

var mCalendarYearNumModeLwrCaseStringToCode = map[string]CalendarYearNumMode{
	"none"           : CalendarYearNumMode(0),
	"astronomical"   : CalendarYearNumMode(1),
	"commonera"      : CalendarYearNumMode(2),
}

var mCalendarYearNumModeCodeToString = map[CalendarYearNumMode]string{
	CalendarYearNumMode(0)  : "None",
	CalendarYearNumMode(1)  : "Astronomical",
	CalendarYearNumMode(2)  : "CommonEra",
}

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
// For easy access to these enumeration values, use the global variable 'CalYearMode'.
// Example: CalYearMode.CommonEra()
//
// Otherwise you will need to use the formal syntax.
// Example: CalendarYearNumMode(0).CommonEra()
//
// Depending on your editor, intellisense (a.k.a. intelligent code completion) may not
// list the CalendarYearNumMode methods in alphabetical order. Be advised that all
// 'CalendarYearNumMode' methods beginning with 'X', as well as the method 'String()',
// are utility methods and not part of the enumeration values.
//
type CalendarYearNumMode int

var lockCalendarYearNumMode sync.Mutex

// None - Signals that the CalendarYearNumMode Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (calYrNumMode CalendarYearNumMode) None() CalendarYearNumMode {

	lockCalendarYearNumMode.Lock()

	defer lockCalendarYearNumMode.Unlock()

	return CalendarYearNumMode(0)
}

// Astronomical - Signals that the year numbering system includes
// a year zero. In other words, the date January 1, year 1 is
// immediately preceded by the date December 31, year 0.
// As its name implies, Astronomical Year numbering is frequently
// used in astronomical calculations.
//
// Reference:
//      https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
// This method is part of the standard enumeration.
//
func (calYrNumMode CalendarYearNumMode) Astronomical() CalendarYearNumMode {

	lockCalendarYearNumMode.Lock()

	defer lockCalendarYearNumMode.Unlock()

	return CalendarYearNumMode(1)
}

// CommonEra - Signals that the year numbering system does NOT
// include a year zero. In other words, the date January 1, year 1 CE
// is immediately preceded by the date December 31, year 1 BCE.
//
// Reference:
//      https://en.wikipedia.org/wiki/Common_Era
//
// This method is part of the standard enumeration.
//
func (calYrNumMode CalendarYearNumMode) CommonEra() CalendarYearNumMode {

	lockCalendarYearNumMode.Lock()

	defer lockCalendarYearNumMode.Unlock()

	return CalendarYearNumMode(2)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'CalendarYearNumMode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= CalendarYearNumMode(0).CommonEra()
// str := t.String()
//     str is now equal to 'CommonEra'
//
func (calYrNumMode CalendarYearNumMode) String() string {

	lockCalendarYearNumMode.Lock()

	defer lockCalendarYearNumMode.Unlock()

	result, ok := mCalendarYearNumModeCodeToString[calYrNumMode]

	if !ok {
		return ""
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether
// the current Calendar Year Numbering Mode value is
// valid.
//
// In order to be valid, Calendar Year Numbering Mode
// MUST be equal to one of the two following values:
//
// 1. CalendarYearNumMode(0).Astronomical() - Value = 1
//               OR
// 2. CalendarYearNumMode(0).CommonEra() - Value = 2
//
// Specifically the enumeration CalendarYearNumMode(0).None()
// is considered, "INVALID".
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  YearNumberingSystem := CalendarYearNumMode(0).Astronomical()
//
//  isValid := YearNumberingSystem.XIsValid()
//
func (calYrNumMode CalendarYearNumMode) XIsValid() bool {

	lockCalendarYearNumMode.Lock()

	defer lockCalendarYearNumMode.Unlock()

	if calYrNumMode > 2 ||
		calYrNumMode < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of CalendarYearNumMode is returned set to the value
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
//                        exact match. Therefore, 'gregorian' will NOT
//                        match the enumeration name, 'Gregorian'.
//
//                        If 'false' a case insensitive search is conducted
//                        for the enumeration name. In this case, 'gregorian'
//                        will match match enumeration name 'Gregorian'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// CalendarYearNumMode - Upon successful completion, this method will return
//                       a new instance of CalendarYearNumMode set to the value
//                       of the enumeration matched by the string search performed
//                       on input parameter, 'valueString'.
//
// error        - If this method completes successfully, the returned error
//                Type is set equal to 'nil'. If an error condition is encountered,
//                this method will return an error type which encapsulates an
//                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t, err := CalendarYearNumMode(0).XParseString("CommonEra", true)
//
//     t is now equal to CalendarYearNumMode(0).CommonEra()
//
func (calYrNumMode CalendarYearNumMode) XParseString(
	valueString string,
	caseSensitive bool) (CalendarYearNumMode, error) {

	lockCalendarYearNumMode.Lock()

	defer lockCalendarYearNumMode.Unlock()

	ePrefix := "CalendarYearNumMode.XParseString() "

	if len(valueString) < 4 {
		return CalendarYearNumMode(0).None(),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n" +
				"String length is less than '4'.\n" +
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var calendarYearNumberingMode CalendarYearNumMode

	if caseSensitive {

		calendarYearNumberingMode, ok = mCalendarYearNumModeStringToCode[valueString]

		if !ok {
			return CalendarYearNumMode(0).None(),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CalendarYearNumMode Value.\n" +
					"valueString='%v'\n", valueString)
		}

	} else {

		calendarYearNumberingMode, ok = mCalendarYearNumModeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return CalendarYearNumMode(0).None(),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CalendarYearNumMode Value.\n" +
					"valueString='%v'\n", valueString)
		}
	}

	return calendarYearNumberingMode, nil
}

// XValue - This method returns the enumeration value of the current
// CalendarYearNumMode instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (calYrNumMode CalendarYearNumMode) XValue() CalendarYearNumMode {

	lockCalendarYearNumMode.Lock()

	defer lockCalendarYearNumMode.Unlock()

	return calYrNumMode
}

// XValueInt - This method returns the integer value of the current
// CalendarYearNumMode instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (calYrNumMode CalendarYearNumMode) XValueInt() int {

	lockCalendarYearNumMode.Lock()

	defer lockCalendarYearNumMode.Unlock()

	return int(calYrNumMode)
}

// CalYearMode - public global variable of
// type CalendarYearNumMode.
//
// This variable serves as an easier, short hand
// technique for accessing CalendarYearNumMode
// values.
//
// Usage:
// CalYearMode.None(),
// CalYearMode.Astronomical(),
// CalYearMode.CommonEra(),
//
var CalYearMode CalendarYearNumMode

