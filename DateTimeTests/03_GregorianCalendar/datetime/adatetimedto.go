package datetime

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
)

// ADateTimeDto - This is a light data transfer object designed expressly
// for transmission of raw date time information.
//
type ADateTimeDto struct {
	date DateTransferDto
	time TimeTransferDto
	julianDayNumber JulianDayNoDto     // Encapsulates Julian Day Number/Time
	dateTimeFmt        string          // Date Time Format String. Empty string or default is
	//                                 //   "2006-01-02 15:04:05.000000000 -0700 MST"
	tag string
	lock *sync.Mutex
}

// Compare - Receives a pointer to an incoming ADateTimeDto instance and
// compares the internal date/time data fields to those of the current ADateTimeDto
// instance. If successful, an integer value is returned describing the result
// of this comparison.
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
//
// IMPORTANT
//
// Complete accuracy can only be guaranteed if the 'isLeapYear' data fields are 
// correctly initialized in both the current ADateTimeDto instance and in
// the second instance passed by input parameter, 'timeTransDto2'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dateTimeTransDto2   *ADateTimeDto
//     - A pointer to another incoming ADateTimeDto instance. The date/time
//       values of this second instance will be compared to those of the current
//       instance.
//
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  comparisonResult    int
//     - If successful this method will return will return an integer value
//       describing the result of comparing 'timeTransDto2' to the current 
//       ADateTimeDto instance.
//
//       Comparison Result
//
//       -1 - The current ADateTimeDto instance is LESS THAN 'timeTransDto2'.
//
//        0 - The current ADateTimeDto instance is EQUAL to 'timeTransDto2'.
//
//       +1 - The current ADateTimeDto instance is GREATER THAN 'timeTransDto2'.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) Compare(
	aDateTimeDto2 *ADateTimeDto,
	ePrefix string) (
	comparisonResult int,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.Compare() "

	comparisonResult = math.MinInt32
	err = nil

	if aDateTimeDto2 == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'aDateTimeDto2' is a nil pointer!\n")

		return comparisonResult, err
	}

	if aDateTimeDto2.lock == nil {
		aDateTimeDto2.lock = new(sync.Mutex)
	}

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_,
	err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto,
		ePrefix + "Testing validity of 'aDateTimeDto'. ")

	if err != nil {
		return comparisonResult, err
	}

	_,
	err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto2,
		ePrefix + "Testing validity of 'aDateTimeDto2'. ")

	if err != nil {
		return comparisonResult, err
	}

	comparisonResult,
	err = aDateTimeDto.date.Compare(
		&aDateTimeDto2.date,
		ePrefix + "Comparing 'aDateTimeDto2.date'. " )

	if err != nil {
		return comparisonResult, err
	}

	if comparisonResult != 0 {
		return comparisonResult, err
	}

	comparisonResult,
		err = aDateTimeDto.time.Compare(
		&aDateTimeDto2.time,
		ePrefix + "Comparing 'aDateTimeDto2.time'. " )

	return comparisonResult, err
}

// CompareAreYearsAdjacent - Compares the year value of the current
// 'ADateTimeDto' instance, 'aDateTimeDto1.year', to the
// year value encapsulated by a second, input parameter instance,
// 'aDateTimeDto2.year', to determine if the two year values
// differ a value of plus or minus and are therefore 'adjacent'
// years. 'Adjacent' year values differ by 1 such that year1 - year2
// is always equal to plus or minus 1 (+1 or -1).
//
// For the purposes of this description, the current instance , aDateTimeDto.year,
// is labeled 'aDateTimeDto1.year' and the input parameter instance,
// aDateTimeDto2, is labeled as 'aDateTimeDto2.year'.
//
// There are three possible outcomes from this comparison:
//
//  1. If the current aDateTimeDto1.year value IS adjacent to
//     aDateTimeDto2.year, and (aDateTimeDto1.year - aDateTimeDto2.year)
//     is equal to +1, this method will return an integer value of plus 1
//     ('+1'). In this case, aDateTimeDto1.year is GREATER than
//     aDateTimeDto2.year by a value of '+1'.
//
//  2. If the current aDateTimeDto1.year value is NOT adjacent
//     to aDateTimeDto2.year this method will return an integer
//     value of zero ('0'). In this case (aDateTimeDto1.year - aDateTimeDto2.year)
//     is Greater than plus one (+1) or Less than minus one (-1).
//
//  3. If the current aDateTimeDto1.year value, IS adjacent to
//     aDateTimeDto2.year and  (aDateTimeDto1.year - aDateTimeDto2.year)
//     is equal to minus 1 ('-1'), this method will return an integer
//     value of minus 1 ('-1'). In this case, aDateTimeDto1.year is LESS
//     THAN aDateTimeDto2.year by a value of minus 1 ('-1').
//
//
//  For all cases, this method will return one of these three
//  integer values:
//
//   +1 = Years ARE ADJACENT aDateTimeDto1.year > aDateTimeDto2.year
//    0 = Years ARE NOT ADJACENT
//   -1 = Years ARE ADJACENT aDateTimeDto1.year < aDateTimeDto2.year
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  aDateTimeDto2            *ADateTimeDto
//     - A pointer to another instance of ADateTimeDto. This instance will
//       be compared to the current ADateTimeDto instance.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
// yearAdjacentComparison    int
//     - This method will return a integer value containing the comparison
//       result indicating whether the year values from the current ADateTimeDto
//       instance and input parameter 'aDateTimeDto2' have adjacent year values.
//
//       For the purposes of this description, the current instance , aDateTimeDto,
//       is labeled 'aDateTimeDto1.year' and the input parameter instance,
//       aDateTimeDto2, is labeled as 'aDateTimeDto2.year'.
//
//       For all all comparison outcomes, this method will return one of these
//       three integer values:
//
//        +1 = Years ARE ADJACENT aDateTimeDto1.year > aDateTimeDto2.year
//         0 = Years ARE NOT ADJACENT
//        -1 = Years ARE ADJACENT aDateTimeDto1.year < aDateTimeDto2.year
//
//       If an error is encountered, the return parameter, 'yearAdjacentComparison'
//       will be set to a very large negative value, math.MinInt32.
//
func (aDateTimeDto *ADateTimeDto) CompareAreYearsAdjacent(
	aDateTimeDto2 *ADateTimeDto) (
	yearAdjacentComparison int ) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	yearAdjacentComparison = math.MinInt32

	if aDateTimeDto2 == nil {
		return yearAdjacentComparison
	}

	if aDateTimeDto2.lock == nil {
		aDateTimeDto2.lock = &sync.Mutex{}
	}

	var err error

	_,
	yearAdjacentComparison,
	err = aDateTimeDto.date.CompareAreYearsAdjacent(
		&aDateTimeDto2.date,
		"")

	if err != nil {
		yearAdjacentComparison = 0
	}

	return yearAdjacentComparison
}


// CompareTotalTimeNanoseconds - Compares the total time nanoseconds values
// for the current ADateTimeDto instance and another ADateTimeDto instance
// passed as an input parameter.
//
// The comparison result is returned as an integer value.
//
// The total time nanoseconds value is the total value of hours, minutes,
// seconds and nanoseconds converted to total nanoseconds.
//
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  aDateTimeDto2            *ADateTimeDto
//     - A pointer to another incoming ADateTimeDto instance. The total
//       time nanoseconds value for the current ADateTimeDto instance
//       will be compared to total time nanoseconds value of this second instance.
//
//
//  ePrefix                  string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  comparisonResult    int
//     - If successful this method will return will return an integer value
//       describing the result of comparing the total time in nanoseconds
//       for the current ADateTimeDto instance and the input parameter,
//       'aDateTimeDto2'
//
//       Comparison Result
//
//       -1 - The current ADateTimeDto instance total time nanoseconds
//            value is LESS THAN the 'aDateTimeDto2' total time nanoseconds value.
//
//        0 - The current ADateTimeDto instance total time nanoseconds
//            value is EQUAL to the 'aDateTimeDto2' total time nanoseconds value.
//
//       +1 - The current ADateTimeDto instance total time nanoseconds
//            value is GREATER THAN the 'aDateTimeDto2' total time nanoseconds value.
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) CompareTotalTimeNanoseconds(
	aDateTimeDto2 *ADateTimeDto,
	ePrefix string) (
	comparisonResult int,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.CompareTotalTimeNanoseconds() "

	comparisonResult = math.MinInt32
	err = nil

	if aDateTimeDto2 == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'aDateTimeDto2' is a 'nil' pointer!\n")

		return comparisonResult, err
	}

	if aDateTimeDto2.lock == nil {
		aDateTimeDto2.lock = &sync.Mutex{}
	}

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_,
	err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto,
		ePrefix + "Testing validity of 'aDateTimeDto'. ")

	if err != nil {
		return comparisonResult, err
	}

	_,
	err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto2,
		ePrefix + "Testing validity of 'aDateTimeDto2'. ")

	if err != nil {
		return comparisonResult, err
	}

	comparisonResult,
	err = aDateTimeDto.time.Compare(
		&aDateTimeDto2.time,
		ePrefix + "Comparing time for 'aDateTimeDto2'. ")

	return comparisonResult, err
}

// CompareYears - Compares the years values for the current ADateTimeDto
// instance and another ADateTimeDto instance passed as an input
// parameter.
//
// The comparison result is returned as an integer value.
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  aDateTimeDto2       *ADateTimeDto
//     - A pointer to another incoming ADateTimeDto instance. The year
//       value for the current ADateTimeDto instance will be compared
//       to the year value of this second instance.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  comparisonResult    int
//     - This method will return will return an integer value describing the
//       result of comparing current ADateTimeDto instance year value to the
//       input parameter 'aDateTimeDto2' year value.
//
//       There are three possible outcomes resulting from this comparison:
//
//       Comparison Result
//
//       -1 - The current ADateTimeDto instance year value is LESS THAN the
//            'aDateTimeDto2' year value.
//
//        0 - The current ADateTimeDto instance year value is EQUAL to the
//            'aDateTimeDto2' year value.
//
//       +1 - The current ADateTimeDto instance year value is GREATER THAN
//            the 'aDateTimeDto2' year value.
//
//       If an error is encountered, the return parameter, 'comparisonResult'
//       will be set to a very large negative value, math.MinInt32.
//
func (aDateTimeDto *ADateTimeDto) CompareYears(
	aDateTimeDto2 *ADateTimeDto) (
	comparisonResult int) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	comparisonResult = math.MinInt32

	if aDateTimeDto2 == nil {
		return comparisonResult
	}

	if aDateTimeDto2.lock == nil {
		aDateTimeDto2.lock = &sync.Mutex{}
	}

	var err error

	comparisonResult,
	err = aDateTimeDto.date.CompareYears(
		&aDateTimeDto2.date,
		"")

	if err != nil {
		comparisonResult = math.MinInt32
	}

	return comparisonResult
}

// CopyIn - Receives a pointer to an incoming ADateTimeDto and performs
// a deep copy of all internal data fields to the current ADateTimeDto
// instance.
//
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTimeTransDto2  *ADateTimeDto
//    - Data field values from this ADateTimeDto instance
//      will be copied into the current ADateTimeDto instance.
//      If the method completes successfully, data fields in both
//      instances will have identical values.
//
//  ePrefix            string
//    - Error Prefix. This is an error prefix which is included in all returned
//      error messages. Usually, it contains the names of the calling
//      method or methods. If no error prefix is desired, simply provide
//      an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) CopyIn(
	dateTimeTransDto2 *ADateTimeDto,
	ePrefix string) error {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = new(sync.Mutex)
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.CopyIn() "

	dTimeTransDtoUtil := aDateTimeDtoUtility{}

	return dTimeTransDtoUtil.copyIn(
		aDateTimeDto,
		dateTimeTransDto2,
		ePrefix)
}

// CopyOut - Returns a deep copy of the current ADateTimeDto
// instance.
//
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
// ADateTimeDto
//     - If this method completes successfully, a deep copy of the current
//       ADateTimeDto instance will be returned as a new instance of type
//       ADateTimeDto.
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
func (aDateTimeDto *ADateTimeDto) CopyOut(
	ePrefix string) (
	ADateTimeDto, error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.CopyOut() "

	dTimeTransDtoUtil := aDateTimeDtoUtility{}

	return dTimeTransDtoUtil.copyOut(aDateTimeDto, ePrefix)
}

// Empty - Resets the internal data fields of the current 
// ADateTimeDto instance to invalid values. Effectively,
// the current ADateTimeDto instance is rendered blank and
// invalid.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix       string
//   This is an error prefix which is included in all returned
//    error messages. Usually, it contains the names of the calling
//    method or methods.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If successful, the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) Empty(
	ePrefix string) error {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	dTimeTransDtoMech := aDateTimeDtoMechanics{}

	return dTimeTransDtoMech.empty(
		aDateTimeDto,
		ePrefix)
}

// ExchangeValues - Performs a data exchange. The data fields
// from incoming ADateTimeDto instance 'dateTimeTransDto2'
// are copied to the current ADateTimeDto instance and the
// data fields from the current instance are in turn copied to
// 'dateTimeTransDto2'.
//
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  aDateTimeDto2       *ADateTimeDto
//     - A pointer to another incoming ADateTimeDto instance. The
//       internal data values of this second instance will be populated 
//       with data fields from the current instance. The current
//       ADateTimeDto will then be populated with the original values
//       from this input parameter, 'aDateTimeDto2'.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note: This error
//       message will incorporate the text passed by the Error Prefix
//       input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) ExchangeValues(
	dateTimeTransDto2 *ADateTimeDto,
	ePrefix string) error {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.ExchangeValues() "

	aDateTimeDtoMech := aDateTimeDtoMechanics{}

	return aDateTimeDtoMech.exchangeValues(
		aDateTimeDto,
		dateTimeTransDto2,
		ePrefix)
}

// GetAbsoluteYearValue - Returns the 'year' value as an absolute
// or positive value.
//
// For the numerical sign of this 'year' value, see instance method,
// 'ADateTimeDto.GetYearNumberSign()'.
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int64
//     - This method returns the absolute or positive 'year' value.
//
func (aDateTimeDto *ADateTimeDto) GetAbsoluteYearValue() int64 {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()


	return aDateTimeDto.date.GetAbsoluteYearValue()
}

// GetAbsoluteYearBigIntValue - Returns the 'year' value as an
// absolute or positive value of type *big.Int.
//
// For the numerical sign of this 'year' value, see instance method,
// 'ADateTimeDto.GetYearNumberSign()'.
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *big.Int
//     - This method returns the absolute or positive 'year' value.
//
func (aDateTimeDto *ADateTimeDto) GetAbsoluteYearBigIntValue() *big.Int {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.GetAbsoluteYearBigIntValue()
}

// GetDateHasLeapSecond - Returns a boolean value signaling whether
// the day specified by this date includes a 'Leap Second'.
//
// The standard 'day' has a duration of 24-hours. If this method returns
// a boolean value of  'true' is signals that the day identified by this
// date consists of 24-hours + 1-second.
//
// A leap second is a one-second adjustment that is occasionally applied
// to Coordinated Universal Time (UTC) in order to accommodate the
// difference between precise time (as measured by atomic clocks) and
// imprecise observed solar time (known as UT1 and which varies due
// to irregularities and long-term slowdown in the Earth's rotation).
//
// Leap seconds are extremely rare occurrences.
//
// For more information on the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
// This method will only signal whether the date includes a leap second.
// To discover whether the time value includes a leap second, call method
// 'ADateTimeDto.GetTimeHasLeapSecond()'.
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - This method returns a boolean value signaling whether the date
//       includes a 'leap second'. If the return value is set to 'true',
//       the day specified by this date has a duration of 24-hours and
//       one-second. If this return value is set to 'false', it signals
//       that the day specified by this date is a standard day consisting
//       of exactly 24-hours.
//
func (aDateTimeDto *ADateTimeDto) GetDateHasLeapSecond() bool {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.date.GetHasLeapSecond()
}

// GetDateTag - Returns the Text Tag or text description assigned to the
// to the date member variable of the current ADateTimeDto instance.
//
// This value is set using the method ADateTimeDto.SetDateTag().
//
// The date tag is used to associate descriptive information with the
// date value encapsulated in the current ADateTimeDto instance.
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - The text description assigned by the user to the date value
//       of the current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetDateTag() string {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.date.tag
}

// GetDateTime - Returns the individual date/time components of 
// the current ADateTimeDto instance.
//
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isLeapYear               bool
//     - A boolean value signaling whether the the return parameter 'year'
//       is a leap year. If the 'year' value is a leap year, this return
//       parameter is set to 'true'. Otherwise, it is set to 'false'.
//
//
//  year                     int64
//     - An int64 value containing the year number. Note that this year
//       value should be interpreted in the context the year type return
//       parameter discussed below. This year value could be formatted
//       according to the Astronomical Year Numbering System or the
//       Common Era Year Numbering System.
//
//
//  yearType                 CalendarYearNumType
//     - The year number type associated with the return value 'yearValue'
//       described above. 'yearType' classifies return parameter 'yearValue'
//       as one of three year types:
//
//         1. Astronomical Year
//         2. BCE - Before Common Era
//         3. CE  - Common Era
//
//       For more information on Astronomical and Common Era Year
//       Numbering, reference:
//           Source File: datetime\calendaryearnumbertypeenum.go
//           https://en.wikipedia.org/wiki/Astronomical_year_numbering
//           https://en.wikipedia.org/wiki/Common_Era
//
//
//  month                    int
//     - The month number component of this date value.
//
//
//  day                      int
//     - The day number component of this date value.
//
//
//  hour                     int
//     - The hour component of the time value encapsulated by the
//       current TimeTransferDto instance.
//
//
//  minute                   int
//     - The minute component of the time value encapsulated by the
//       current TimeTransferDto instance.
//
//
//  second                   int
//     - The second component of the time value encapsulated by the
//       current TimeTransferDto instance.
//
//
//  nanosecond               int
//     - The nanosecond component of the time value encapsulated by
//       the current TimeTransferDto instance.
//
//
//  timeZoneDef              TimeZoneDefinition
//     - The the time zone definition used to configure the time
//       value encapsulated by the current TimeTransferDto
//       instance.
//
//
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'.
//
//       If errors are encountered during processing, the returned error
//       Type will encapsulate an error message. Note that this error
//       message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) GetDateTime(
	ePrefix string) (
	isLeapYear bool,
	year int64,
	yearType CalendarYearNumType,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	timeZone TimeZoneDefinition,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.GetDateTime() "

	isLeapYear = aDateTimeDto.date.GetIsLeapYear()

	year,
	yearType,
	month,
	day,
	err = aDateTimeDto.date.GetDate(ePrefix)

	if err != nil {
		return isLeapYear,
		year,
		yearType,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		timeZone,
		err
	}

	hour,
	minute,
	second,
	nanosecond,
	timeZone,
	err = aDateTimeDto.time.GetTimeValue(ePrefix)

	return isLeapYear,
		year,
		yearType,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		timeZone,
		err
}

// GetDate - Returns the date components of the current ADateTimeDto
// instance.
//
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isLeapYear               bool
//     - A boolean value signaling whether the the return parameter 'year'
//       is a leap year. If the 'year' value is a leap year, this return
//       parameter is set to 'true'. Otherwise, it is set to 'false'.
//
//
//  year                     int64
//     - An int64 value containing the year number. Note that this year
//       value should be interpreted in the context the year type return
//       parameter discussed below. This year value could be formatted
//       according to the Astronomical Year Numbering System or the
//       Common Era Year Numbering System.
//
//
//  yearType                 CalendarYearNumType
//     - The year number type associated with the return value 'yearValue'
//       described above. 'yearType' classifies return parameter 'yearValue'
//       as one of three year types:
//
//         1. Astronomical Year
//         2. BCE - Before Common Era
//         3. CE  - Common Era
//
//       For more information on Astronomical and Common Era Year
//       Numbering, reference:
//           Source File: datetime\calendaryearnumbertypeenum.go
//           https://en.wikipedia.org/wiki/Astronomical_year_numbering
//           https://en.wikipedia.org/wiki/Common_Era
//
//
//  month                    int
//     - The month number component of this date value.
//
//
//  day                      int
//     - The day number component of this date value.
//
//
func (aDateTimeDto *ADateTimeDto) GetDate(
	ePrefix string) (
	isLeapYear bool,
	year int64,
	yearType CalendarYearNumType,
	month int,
	day int,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.GetDate() "

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_,
	err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto,
		ePrefix + "Testing validity of 'aDateTimeDto'. ")

	if err != nil {
		return isLeapYear,
			year,
			yearType,
			month,
			day,
			err
	}

	isLeapYear = aDateTimeDto.GetIsLeapYear()

	year,
	yearType,
	month,
	day,
	err = aDateTimeDto.date.GetDate(ePrefix)

	return isLeapYear,
		year,
		yearType,
		month,
		day,
		err
}

// GetDateDto - Returns an instance of DateTransferDto. This
// type encapsulates the date value contain in the current
// ADateTimeDto instance.
//
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix                  string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  dateDto                  DateTransferDto
//     - An instance of DateTransferDto which contains the date value
//       encapsulated by the current ADateTimeDto instance.
//
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) GetDateDto(
	ePrefix string) (
	dateDto DateTransferDto,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.GetDate() "

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_,
		err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto,
		ePrefix + "Testing validity of 'aDateTimeDto'. ")

	if err != nil {
		return dateDto,	err
	}

	dateDto,
	err = aDateTimeDto.date.CopyOut(ePrefix)

	return dateDto, err
}

// GetDay - Returns the day number associated with the date value
// contained in the current ADateTimeDto instance.
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method returns the day number component of the
//       date encapsulated by the current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetDay() int {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.date.GetDay()
}

// GetDaysInYear - Returns the the number of days in the year
// specified by the current ADateTimeDto instance.
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method will return the number of days in the year component
//       of the date value for the current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetDaysInYear(
	ePrefix string) (
	daysInYear int,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.GetDaysInYear() "

	daysInYear = math.MinInt32
	err = nil

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_,
		err =
		aDateTimeDtoNanobot.testDateTransferDtoValidity(
			aDateTimeDto,
			ePrefix + "Testing validity of 'aDateTimeDto'. ")

	if err != nil {
		return daysInYear, err
	}

	daysInYear,
		err = aDateTimeDto.date.GetDaysInYear(ePrefix)

	return daysInYear, err
}

// GetDateTimeStr - Returns the equivalent date time string
// reflecting the date time value of the current ADateTimeDto
// instance. The Date Time Format was previously specified and
// is extracted from internal member variable, 'calDTime.dateTimeFmt'.
//
// To manage the Date Time Format for the current ADateTimeDto
// instance, reference method ADateTimeDto.SetDateTimeFormat().
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix       string
//    This is an error prefix which is included in all returned
//    error messages. Usually, it contains the names of the calling
//    method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  dateTimeStr              string
//     - If successful, this method will return a string containing
//       the date time formatted according to the Date Time Format
//       associated with this ADateTimeDto instance. Reference
//       method ADateTimeDto.SetDateTimeFormat().
//
//
//  err                      error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) GetDateTimeStr(
	ePrefix string) (
	dateTimeStr string,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.GetDateTimeStr() "

	dateTimeStr = ""

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_, err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto,
		ePrefix + " Checking Current Instance Validity: 'aDateTimeDto' ")

	if err != nil {
		return dateTimeStr, err
	}

	var usDayOfWeekNumber UsDayOfWeekNo

	usDayOfWeekNumber,
	err = aDateTimeDto.date.GetUsDayOfWeekNo(
		aDateTimeDto.julianDayNumber,
		ePrefix)

	if err != nil {
		return dateTimeStr, err
	}

	aDateTimeDtoUtil := aDateTimeDtoUtility{}

	dateTimeStr,
	err = aDateTimeDtoUtil.generateDateTimeStr(
		aDateTimeDto.date.astronomicalYear,
		aDateTimeDto.date.month,
		aDateTimeDto.date.day,
		usDayOfWeekNumber,
		aDateTimeDto.time.hour,
		aDateTimeDto.time.minute,
		aDateTimeDto.time.second,
		aDateTimeDto.time.nanosecond,
		aDateTimeDto.dateTimeFmt,
		ePrefix)

	return dateTimeStr, err
}

// GetHour - Returns the hour component of the time value
// encapsulated by the current  ADateTimeDto instance.
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method returns the hour  component of the time
//       value encapsulated by the current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetHour() int {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.time.GetHour()
}

// GetIsLeapYear - Returns a boolean value signaling whether the year
// value contained in the date encapsulated within the current ADateTimeDto
// instance is a leap year.
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If this return value is set to 'true', it signals that the
//       year value encapsulated by the current ADateTimeDto instance
//       is a leap year. Otherwise, this method returns 'false' signaling
//       that the year value is NOT a leap year.
//
func (aDateTimeDto *ADateTimeDto) GetIsLeapYear() bool {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.date.GetIsLeapYear()

}

// GetMinute - Returns the minute component of the time value
// encapsulated in the current ADateTimeDto instance.
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method will return the minute component of the time value
//       encapsulated in the current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetMinute() int {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.time.GetMinute()
}

// GetMonth - Returns the month number incorporated in the date
// value encapsulated in the current ADateTimeDto instance.
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method will return the month component of the date value
//       encapsulated in the current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetMonth() int {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.date.GetMonth()
}

// GetNanosecond - Returns the nanosecond component of the time value
// encapsulated in the current ADateTimeDto instance.
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method will return the nanosecond component of the time value
//       encapsulated in the current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetNanosecond() int {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.time.GetNanosecond()
}

// GetOrdinalDayNoInYear - Returns the Ordinal Day Number in the year
// value encapsulated in the current ADateTimeDto instance. This is
// otherwise referred to as the Ordinal Date.
//
// The ordinal day number is a calendar date typically consisting of
// a year and a day number ranging between 1 and 366 (starting on
// January 1st). For more information on the Ordinal Day Number or the
// Ordinal Date, reference:
//    https://en.wikipedia.org/wiki/Ordinal_date
//
// Be advised that if the current ADateTimeDto contains invalid member
// variable data values, an error message will be returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ordinalDayNoInYear  int
//     - If the method completes successfully, this value will represent
//       ordinal or sequential day number in the year encapsulated within
//       the current ADateTimeDto instance.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) GetOrdinalDayNoInYear(
	ePrefix string) (
	ordinalDayNo int,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.GetOrdinalDayNoInYear() "

	ordinalDayNo = math.MinInt32

	err = nil

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_,
	err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto,
		ePrefix +
			"Testing validity of current ADateTimeDto instance. ")

	if err != nil {
		return ordinalDayNo, err
	}

	ordinalDayNo,
	err = aDateTimeDto.date.GetOrdinalDayNoInYear(ePrefix)


	return ordinalDayNo, err
}

// GetRemainingDaysInYear - Returns the number of days remaining
// in the year.
//
// Be advised that if the current ADateTimeDto contains invalid member
// variable data values, an error message will be returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix                  string
//    This is an error prefix which is included in all returned
//    error messages. Usually, it contains the names of the calling
//    method or methods. This string will be prefixed to any returned
//    error messages.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  remainingDaysInYear      int
//     - The number of days remaining in the year. For example, January 1, 2021,
//       is the first day of a non-leap year or standard year containing 365 days.
//       Therefore there are 364-days remaining in this year.
//
//
//  err                      error
//     - If the method completes successfully, the returned error Type is set equal
//       to 'nil'. If errors are encountered during processing, the returned 'error'
//       instance will encapsulate an error message. Note that this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) GetRemainingDaysInYear(
	ePrefix string) (
	remainingDaysInYear int,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.GetRemainingDaysInYear() "

	remainingDaysInYear = math.MinInt32

	err = nil

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_,
		err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto,
		ePrefix +
			"Testing validity of current ADateTimeDto instance. ")

	if err != nil {
		return remainingDaysInYear, err
	}

	remainingDaysInYear,
	err = aDateTimeDto.date.GetRemainingDaysInYear(ePrefix)

	return remainingDaysInYear, err
}

// GetSecond - Returns the second component of the time value
// encapsulated in the current ADateTimeDto instance.
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method will return the second component of the time value
//       encapsulated in the current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetSecond() int {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.time.GetSecond()
}

// GetTag - Returns the Text Tag or text description assigned to the
// current instance of ADateTimeDto.
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - The text description assigned by the user to the current
//       ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetTag() string {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.tag
}

// GetTimeHasLeapSecond - Returns a boolean value signaling whether
// the seconds specified by this time value includes a 'Leap Second'.
//
// The standard 'day' has a duration of 24-hours. This means that the
// maximum time value for a day is 23-hours, 59-minutes and 59-seconds.
// If this method returns a boolean value of 'true' is signals that the
// time value includes a leap second and is set to 23-hours, 59-minutes
// and 60-seconds.
//
// A leap second is a one-second adjustment that is occasionally applied
// to Coordinated Universal Time (UTC) in order to accommodate the
// difference between precise time (as measured by atomic clocks) and
// imprecise observed solar time (known as UT1 and which varies due
// to irregularities and long-term slowdown in the Earth's rotation).
//
// Leap seconds are extremely rare occurrences.
//
// For more information on the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
// This method will only signal whether the time value includes a leap
// second. To discover whether the date value includes a leap second,
// call method 'ADateTimeDto.GetDateHasLeapSecond()'.
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - This method returns a boolean value signaling whether the time
//       value includes a 'leap second'. If the return value is set to
//       'true', it signals that the time value is set to 23-hours,
//       59-minutes and 60-seconds. If this return value is set to 'false',
//       it signals that the time value does NOT include a leap second.
//
func (aDateTimeDto *ADateTimeDto) GetTimeHasLeapSecond() bool {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.time.GetHasLeapSecond()
}

// GetTimeTag - Returns the Text Tag or text description assigned to the
// to the time member variable of the current ADateTimeDto instance.
//
// This value is set using the method ADateTimeDto.SetTimeTag().
//
// The time tag is used to associate descriptive information with the
// time value encapsulated in the current ADateTimeDto instance.
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - The text description assigned by the user to the time value
//       of the current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetTimeTag() string {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.time.tag
}

// GetTime - Returns the time components of the current ADateTimeDto
// instance.
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  hour                int
//     - The hour component of the time value encapsulated by the
//       current ADateTimeDto instance.
//
//
//  minute              int
//     - The minute component of the time value encapsulated by the
//       current ADateTimeDto instance.
//
//
//  second              int
//     - The second component of the time value encapsulated by the
//       current ADateTimeDto instance.
//
//
//  nanosecond          int
//     - The nanosecond component of the time value encapsulated by
//       the current ADateTimeDto instance.
//
//
//  timeZoneDef         TimeZoneDefinition
//     - The the time zone definition used to configure the time
//       value encapsulated by the current ADateTimeDto instance.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'.
//
//       If errors are encountered during processing, the returned error
//       Type will encapsulate an error message. Note that this error
//       message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) GetTime(
	ePrefix string) (
	hour int, 
	minute int, 
	second int,
	nanosecond int,
	timeZone TimeZoneDefinition,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.GetTime() "

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_,
	err =
		aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto,
		ePrefix + "Testing validity of 'aDateTimeDto'. ")

	if err != nil {
		return hour,
		minute,
		second,
		nanosecond,
		timeZone,
		err
	}

	hour,
	minute,
	second,
	nanosecond,
	timeZone,
	err = aDateTimeDto.time.GetTimeValue(ePrefix)

	return hour,
		minute,
		second,
		nanosecond,
		timeZone,
		err
}

// GetTotalTimeInNanoseconds - For the current ADateTimeDto instance
// this method returns the sum of Hours, Minutes, Seconds and Nanoseconds
// expressed as total nanoseconds.
//
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  totalTimeInNanoseconds     int64
//     - If successful this method will return a the total of
//       Hours, Minutes, Seconds and Nanoseconds for the current
//       instance of ADateTimeDto expressed as total
//       nanoseconds.
//
//
//  err                         error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) GetTotalTimeInNanoseconds(
	ePrefix string) (
	totalTimeInNanoseconds int64, err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.GetTotalTimeInNanoseconds() "

	totalTimeInNanoseconds = math.MinInt64

	ePrefix += "ADateTimeDto.GetTime() "

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_,
		err =
		aDateTimeDtoNanobot.testDateTransferDtoValidity(
			aDateTimeDto,
			ePrefix + "Testing validity of 'aDateTimeDto'. ")

	if err != nil {
		return totalTimeInNanoseconds, err
	}

	totalTimeInNanoseconds = aDateTimeDto.time.GetTotalNanoseconds()

	return totalTimeInNanoseconds, err
}

// GetYearWithType - Returns the year value for the current ADateTimeDto
// instance with the associated year type.
//
// The return value year number type classifies return parameter 'yearValue'
// as one of three year types:
//
//  Astronomical      (1) - Signals that the year number value includes a year zero.
//                          In other words, the date January 1, year 1 is immediately
//                          preceded by the date December 31, year 0. Likewise date,
//                          January 1, year 0 is immediately preceded by the date,
//                          December 31, year -1.
//                          Reference:
//                            https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//  Before Common Era (2) - Signals that the year number is less than the date
//                          January 1, 0001 CE. The date immediately preceding
//                          January 1, 0001 CE is December 31, 0001 BCE. All
//                          years before year 1 CE are numbered as 2 BCE, 3 BCE
//                          4 BCE etc. Reference:
//                             https://en.wikipedia.org/wiki/Common_Era
//
//  CommonEra         (3) - Signals that the year number is greater than the date
//                          December 31, 0001 BCE. Years following the date December 31,
//                          0001 BCE are numbered as 1 CE, 2 CE, 3 CE, 4 CE etc.
//                          Reference:
//                             https://en.wikipedia.org/wiki/Common_Era
//
// For more information on Astronomical and Common Era Year
// Numbering, reference:
//     Source File: datetime\calendaryearnumbertypeenum.go
//     https://en.wikipedia.org/wiki/Astronomical_year_numbering
//     https://en.wikipedia.org/wiki/Common_Era
//
//
// Be advised that this method validates the current ADateTimeDto
// instance. If this instance is judged to be invalid, an error is
// returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  yearValue           int64
//     - The year component of the date value encapsulated in the current
//       ADateTimeDto instance. This year
//
//
//  yearType                 CalendarYearNumType
//     - The year number type associated with the 'yearValue' described
//       above. 'yearType' classifies return parameter 'yearValue' as one
//       of three year types:
//
//         1. Astronomical Year
//         2. BCE - Before Common Era
//         3. CE  - Common Era
//
//       For more information on Astronomical and Common Era Year
//       Numbering, reference:
//           Source File: datetime\calendaryearnumbertypeenum.go
//           https://en.wikipedia.org/wiki/Astronomical_year_numbering
//           https://en.wikipedia.org/wiki/Common_Era
//
//
//  err                         error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) GetYearWithType(
	ePrefix string) (
	yearValue int64,
	yearType CalendarYearNumType,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_,
		err =
		aDateTimeDtoNanobot.testDateTransferDtoValidity(
			aDateTimeDto,
			ePrefix + "Testing validity of 'aDateTimeDto'. ")

	if err != nil {
		return yearValue, yearType, err
	}

	yearValue,
	yearType,
	err = aDateTimeDto.date.GetYearWithType(ePrefix)

	return yearValue, yearType, err
}

// GetYear - Returns the year value formatted as an
// Astronomical Year Value. This method is identical to
// method ADateTimeDto.GetYearAstronomical() and is
// provided for naming convenience. To acquire the
// year value which may be formatted as a Common Era
// or Before Common Era value, see method
// ADateTimeDto.GetYearWithType().
//
// This method returns the astronomical year value
// extracted from the date value encapsulated by the
// current DateTransferDto instance.
// year value which may be formatted as a Common Era
// or
//
// The Astronomical Year Numbering System includes a year
// zero. In other words, the date January 1, year 1 is
// immediately preceded by the date December 31, year 0.
// Years prior to year zero (0) are numbered as negative
// values with a leading minus sign (-). For example,
// the year 1 is preceded by the year sequence: 0, -1, -2,
// -3, -4 etc.
//
// As its name implies, Astronomical Year numbering is
// frequently used in astronomical calculations.
//
// Reference:
//  https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//     --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int64
//     - The Astronomical Year value is returned as a type int64.
//       This is the Astronomical Year value encapsulated in the
//       current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetYear() int64 {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.date.GetYear()
}

// GetYearAstronomical - Returns the year value formatted
// as an Astronomical Year Value.
//
// The Astronomical Year Numbering System includes a year
// zero. In other words, the date January 1, year 1 is
// immediately preceded by the date December 31, year 0.
// Years prior to year zero (0) are numbered as negative
// values with a leading minus sign (-). For example,
// the year 1 is preceded by the year sequence: 0, -1, -2,
// -3, -4 etc.
//
// As its name implies, Astronomical Year numbering is
// frequently used in astronomical calculations.
//
// Reference:
//  https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//     --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int64
//     - The Astronomical Year value is returned as a type int64.
//       This is the Astronomical Year value encapsulated in the
//       current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetYearAstronomical() int64 {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.date.GetYearAstronomical()
}

// GetYearBigInt - Returns the astronomical year value for the
// current ADateTimeDto instance. This year value is returned
// as a type *big.Int.
//
// The Astronomical Year Numbering System includes a year
// zero. In other words, the date January 1, year 1 is
// immediately preceded by the date December 31, year 0.
// Years prior to year zero (0) are numbered as negative
// values with a leading minus sign (-). For example,
// the year 1 is preceded by the year sequence: 0, -1, -2,
// -3, -4 etc.
//
// As its name implies, Astronomical Year numbering is
// frequently used in astronomical calculations.
//
// Reference:
//  https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//     --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  *big.Int
//     - This method returns the astronomical year component of the date
//       value encapsulated in the current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) GetYearBigInt() *big.Int {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.date.GetYearAstronomicalBigInt()
}

// GetYearNumberSign - Returns a code for the numeric sign of the
// astronomical year value encapsulated in the current ADateTimeDto
// instance.
//
// Possible Return Values:
//  +1 = Year Value is Positive
//   0 = Year Value is Zero
//  -1 = Year Value is Negative
//
//
// Astronomical Year
//
// The Astronomical Year numbering system includes a year/ zero. In
// other words, the date January 1, year 1 is immediately preceded by
// the date December 31, year 0. Years prior to year zero (0) are
// numbered as negative values with a leading minus sign (-). For
// example, the year 1 is preceded by the year sequence: 0, -1, -2,
// -3, -4 etc.
//
// As its name implies, Astronomical Year numbering is frequently
// used in astronomical calculations.
//
// For additional information on the Astronomical Year Numbering
// System, reference:
//  https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//
// IMPORTANT
//
// This method does NOT validate the current ADateTimeDto instance
// before returning the value. To run a validity check on the
// ADateTimeDto instance first call one of the two following
// methods:
//
//  ADateTimeDto.IsValidInstance() bool
//                OR
//  ADateTimeDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//    --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method returns the numerical sign of the astronomical year
//       value encapsulated by the current ADateTimeDto instance.
//
//       Possible Return Values:
//        +1 = Year Value is Positive
//         0 = Year Value is Zero
//        -1 = Year Value is Negative
//
func (aDateTimeDto *ADateTimeDto) GetYearNumberSign() int {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	return aDateTimeDto.date.GetYearNumberSign()
}

// IsValidInstance - If this instance returns 'true' it signals that
// the current ADateTimeDto instance has been correctly initialized
// and contains valid data field values.
//
// This method is identical in operation to method
// ADateTimeDto.IsValidInstanceError with the sole difference
// being that this method returns a boolean value while the other
// method returns an 'error' value.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  ePrefix         string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isValid         boolean
//     - If successful, and the date/time values specified by the current
//       ADateTimeDto instance are valid, this returned boolean value
//       is set to 'true'. A return of 'false' signals that the ADateTimeDto
//       instance contains invalid data value.
//
func (aDateTimeDto *ADateTimeDto) IsValidInstance() bool {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	isValid, _ := aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto,
		"")

	return isValid
}

// IsValidInstanceError - Analyzes the internal date fields of the
// current ADateTimeDto instance to determine validity.
//
// This method is identical in operation to method ADateTimeDto.IsValidInstance
// with the sole difference being that this method returns an 'error'
// value while the other method returns a 'boolean' value.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  ePrefix                 string
//     - Error Prefix. A string which will be prefixed to any error returned
//       by this method. Usually, the string consists of the method chain
//       used to call this method. In case of error, this text string is
//       included in the returned error message. Note: Be sure to leave a
//       space at the end of 'ePrefix'. If no error prefix is needed, simply
//       supply an empty string for this parameter.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                error
//     - If successful, and the date/time values specified by the current
//       ADateTimeDto instance are valid, the returned error Type
//       is set equal to 'nil'.
//
//       If errors are encountered during processing, or if the date/time
//       values specified by the current ADateTimeDto instance are
//       invalid, the returned error Type will encapsulate an appropriate
//       error message. Note this error message will be prefixed with the
//       method chain and text passed by input parameter, 'ePrefix'.
//
func (aDateTimeDto *ADateTimeDto) IsValidInstanceError(
	ePrefix string) ( err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	err = nil

	ePrefix += "ADateTimeDto.IsValidInstanceError() "

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_, err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto,
		ePrefix)

	return err
}

// New - Creates and returns an new instance of ADateTimeDto. This is a
// light data transfer object designed expressly for transmission of raw 
// date time information.
//
// This method defaults the Day Of The Week Numbering System Type to:
//           DayOfWeekNumberingSystemType(0).UsDayOfWeek()
//
// To configure this numbering system for alternative Day Of The Week
// Numbering Systems see method ADateTimeDto.SetDayOfWeekNoSys().
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  calendarSystem           CalendarSpec
//     - An enumeration type designating the Calendar System
//       associated with the generated date. Reference:
//       Source File: datetime\calendarspecenum.go
//
//       Possible Calendar System values include:
//       Gregorian, Julian, Revised Julian, or Revised Goucher-Parker.
//
//       Possible Enumeration Values:
//         CalendarSpec(0).Gregorian()
//         CalendarSpec(0).Julian()
//         CalendarSpec(0).RevisedJulian()
//         CalendarSpec(0).RevisedGoucherParker()
//
//
//  year                     int64
//     - An int64 value containing the year number. Note that this year
//       value should be interpreted in the context the year type return
//       parameter discussed below. This year value could be formatted
//       according to the Astronomical Year Numbering System or the
//       Common Era Year Numbering System.
//
//
//  yearType                 CalendarYearNumType
//     - The year number type associated with the return value 'yearValue'
//       described above. 'yearType' classifies return parameter 'yearValue'
//       as one of three year types:
//
//         1. Astronomical Year
//         2. BCE - Before Common Era
//         3. CE  - Common Era
//
//       For more information on Astronomical and Common Era Year
//       Numbering, reference:
//           Source File: datetime\calendaryearnumbertypeenum.go
//           https://en.wikipedia.org/wiki/Astronomical_year_numbering
//           https://en.wikipedia.org/wiki/Common_Era
//
//
//  month                    int
//     - The month number for this date/time specification.
//
//  day                      int
//     - The day number for this date/time specification. The day
//
//
//  hasLeapSecond            bool
//     - If this parameter is set to 'true', it signals that the day identified
//       by year, month and day input parameters contains a leap second.
//       This parameter is rarely used and is almost always set to 'false'.
//
//       A leap second is a one-second adjustment that is occasionally applied
//       to Coordinated Universal Time (UTC) in order to accommodate the
//       difference between precise time (as measured by atomic clocks) and
//       imprecise observed solar time (known as UT1 and which varies due to
//       irregularities and long-term slowdown in the Earth's rotation). If
//       this return parameter is set to 'true', time and duration
//       calculations will assume the duration of the relevant 'day' is
//       24-hours plus one second. Otherwise, the duration of the day is
//       assumed to consist of exactly 24-hours. For more information on
//       the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  hour                     int
//     - The hour time component for this date/time specification.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//
//  minute                   int
//     - The minute time component for this date/time specification.
//       The valid range is 0 - 59 inclusive
//
//  second                   int
//     - The second time component for this date/time specification.
//       The valid range is 0 - 60 inclusive. The value 60 is only
//       used in the case of leap seconds.
//
//  nanosecond               int
//     - The nanosecond time component for this date/time specification.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  timeZoneLocation         string
//     - A string containing the name of a valid time zone. Usually, this is
//       an IANA Time Zone as shown in the examples below:
//
//        Time Zone Strings           Time Zone
//        -------------------------------------
//
//       "America/New_York"         USA Eastern Time Zone
//       "America/Chicago"          USA Central Time Zone
//       "America/Denver"           USA Mountain Time Zone
//       "America/Los_Angeles"      USA Pacific Time Zone
//       "Local"                    The time zone on the host computer
//                                    where this code is executed.
//       "UTC"                      Coordinated Universal Time
//                                    https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//       Reference:
//           https://golang.org/pkg/time/
//           https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//           https://en.wikipedia.org/wiki/Tz_database
//           https://en.wikipedia.org/wiki/List_of_military_time_zones
//           https://www.iana.org/time-zones
//
//
//  dateTimeFmt              string
//    - This string contains the date/time format which will be used to
//      to format date/time output values. Example:
//          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  tag                string
//     - A string description to be associated with the newly created ADateTimeDto
//       instance generated by this method.
//
//
//  ePrefix                  string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the 
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newDateTimeDto     ADateTimeDto
//     - If successful this method will return a new, fully populated instance
//       of type ADateTimeDto.
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (aDateTimeDto ADateTimeDto) New (
	calendarSystem     CalendarSpec,
	year int64,
	yearType CalendarYearNumType,
	month int,
	day int,
	hasLeapSecond bool,
	hour,
	minute,
	second,
	nanosecond int,
	timeZoneLocation string,
	dateTimeFmt string,
	tag string,
	ePrefix string) (
	newDateTimeDto ADateTimeDto,
	err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = new(sync.Mutex)
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.New () "

	newDateTimeDto = ADateTimeDto{}

	newDateTimeDto.lock = new(sync.Mutex)

	err = nil

	newDateTimeDto.date, err =
		DateTransferDto{}.NewFromComponents(
			calendarSystem,
			year,
			yearType,
			month,
			day,
			hasLeapSecond,
			"",
			ePrefix)

	if err != nil {
		return ADateTimeDto{}, err
	}

	 if second != 60 {
		hasLeapSecond = false
	 }

	 newDateTimeDto.time, err =
	 	TimeTransferDto{}.NewFromComponents(
	 		hour,
	 		minute,
	 		second,
	 		hasLeapSecond,
	 		nanosecond,
	 		timeZoneLocation,
	 		ePrefix)

	if err != nil {
		return ADateTimeDto{}, err
	}

	dtMech := DTimeNanobot{}

	newDateTimeDto.dateTimeFmt =
		dtMech.PreProcessDateFormatStr(dateTimeFmt)

	newDateTimeDto.tag = tag

	return newDateTimeDto, err
}

// SetHasLeapSecond - The standard 'day' has a duration of exactly 24-hours.
// If this method's input parameter is set to 'true' is signals that the day
// identified by this ADateTimeDto instance consists of 24-hours + 1-second.
//
// A leap second is a one-second adjustment that is occasionally applied
// to Coordinated Universal Time (UTC) in order to accommodate the
// difference between precise time (as measured by atomic clocks) and
// imprecise observed solar time (known as UT1 and which varies due
// to irregularities and long-term slowdown in the Earth's rotation).
//
// For more information on the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
// This method can be used to set the value of internal member variables
// tracking whether or not the date value includes a leap second.
//
func (aDateTimeDto *ADateTimeDto) SetHasLeapSecond(
	hasLeapSecond bool) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	aDateTimeDto.date.hasLeapSecond = hasLeapSecond

	if hasLeapSecond == true &&
		aDateTimeDto.time.second == 60 {

		aDateTimeDto.time.hasLeapSecond = true

	}

	if hasLeapSecond == false {
		aDateTimeDto.time.hasLeapSecond = false
	}

}

// SetDateTag - Sets the tag description string attached to the
// internal member variable 'ADateTimeDto.date'. This text
// string is used to identify or describe the current
// ADateTimeDto date value.
//
func (aDateTimeDto *ADateTimeDto) SetDateTag(tag string) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	aDateTimeDto.date.SetTagDescription(tag)
}

// SetDateTimeFormat - Sets the value of the date/time
// format string to the value passed as an input parameter.
//
// If the input parameter 'dateTimeFmt' is an empty string,
// this method will default the date time format to:
//    "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (aDateTimeDto *ADateTimeDto) SetDateTimeFormat(
	dateTimeFmt string) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	dtMech := DTimeNanobot{}

	aDateTimeDto.dateTimeFmt =
		dtMech.PreProcessDateFormatStr(dateTimeFmt)

}

// SetJulianDayNumberDto - Sets the Julian Day Number for
// the current ADateTimeDto instance. The Julian Day Number
// is passed as input parameter 'julianDayNo' of type,
// JulianDayNoDto. If this input parameter is invalid, an
// error will be returned.
//
// julianDayNo JulianDayNoDto
//
func (aDateTimeDto *ADateTimeDto) SetJulianDayNumberDto(
	julianDayNo JulianDayNoDto,
	ePrefix string) (err error) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	ePrefix += "ADateTimeDto.SetJulianDayNumberDto() "

	err = julianDayNo.IsValidInstanceError(ePrefix)

	if err != nil {
		return err
	}

	aDateTimeDto.julianDayNumber,
	err =
		julianDayNo.CopyOut(ePrefix +
			"Copying input parameter 'julianDayNo'. ")

	return err
}

// SetTag - Sets the internal member variable 'tag'
// for the current ADateTimeDto instance. This text
// string is used to identify or describe the
// current ADateTimeDto instance.
//
func (aDateTimeDto *ADateTimeDto) SetTag(tag string) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	aDateTimeDto.tag = tag
}

// SetTimeTag - Sets the tag description string attached to the
// internal member variable 'ADateTimeDto.time'. This text
// string is used to identify or describe the current
// ADateTimeDto time value.
//
func (aDateTimeDto *ADateTimeDto) SetTimeTag(tag string) {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	aDateTimeDto.tag = tag
}

func (aDateTimeDto *ADateTimeDto) String() string {

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = &sync.Mutex{}
	}

	aDateTimeDto.lock.Lock()

	defer aDateTimeDto.lock.Unlock()

	outStr := "\n"
	separator := strings.Repeat("-", 65)
	outStr+= separator + "\n"

	if len(aDateTimeDto.tag) > 0 {
		outStr += "Date Time Tag: " + aDateTimeDto.tag + "\n"
	}
	outStr += "ADateTimeDto Instance\n"
	outStr += separator + "\n"
	outStr +=  fmt.Sprintf("   Date Has Leap Second: %v\n", aDateTimeDto.date.hasLeapSecond)
	outStr +=  fmt.Sprintf("Astronomical Year Value: %v\n", aDateTimeDto.date.astronomicalYear)
	outStr +=  fmt.Sprintf("    Year Numbering Mode: %v\n", aDateTimeDto.date.yearNumberingMode.String())
	outStr +=  fmt.Sprintf("               Date Tag: %v\n", aDateTimeDto.date.tag)
	outStr +=  fmt.Sprintf("             isLeapYear: %v\n", aDateTimeDto.GetIsLeapYear())
	outStr +=  fmt.Sprintf("                  month: %v\n", aDateTimeDto.date.month)
	outStr +=  fmt.Sprintf("                    day: %v\n", aDateTimeDto.date.day)
	outStr +=  fmt.Sprintf("   Time Has Leap Second: %v\n", aDateTimeDto.time.hasLeapSecond)
	outStr +=  fmt.Sprintf("                   hour: %v\n", aDateTimeDto.time.hour)
	outStr +=  fmt.Sprintf("                 minute: %v\n", aDateTimeDto.time.minute)
	outStr +=  fmt.Sprintf("                 second: %v\n", aDateTimeDto.time.second)
	outStr +=  fmt.Sprintf("             nanosecond: %v\n", aDateTimeDto.time.nanosecond)
	outStr +=  fmt.Sprintf("              Time Zone: %v\n", aDateTimeDto.time.timeZone.originalTimeZone.locationName)
	outStr +=  fmt.Sprintf("               Time Tag: %v\n", aDateTimeDto.time.tag)
	outStr +=  fmt.Sprintf(" This Instance Is Valid: %v\n", aDateTimeDto.IsValidInstance())
	outStr +=  separator + "\n\n"

	return outStr
}

