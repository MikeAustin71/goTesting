package datetime

import (
	"math"
	"math/big"
	"sync"
)

// DateTransferDto - This is a light data transfer object designed expressly
// for transmission of raw date information.
//
type DateTransferDto struct {
	astronomicalYear     int64                        // Year Number expressed using Astronomical Year Numbering.
	yearNumberingMode    CalendarYearNumMode          // Designates the year numbering system associated with year value.
	yearNumType          CalendarYearNumType          // Astronomical Year, BCE Year or CE Year
	month                int                          // Month Number
	day                  int                          // Day Number day of the month
	hasLeapSecond        bool                         // If 'true' it signals that the day length consists of 24-hours and 1-second
	calendarBaseData     ICalendarBaseData            // Interface - Calendar Specific Base Data
	tag                  string                       // Tag Description string
	lock                 *sync.Mutex                  // Used to coordinate thread safe operations
}

// CopyIn - Populates the current DateTransferDto instance with a deep copy
// of member data elements extracted from the the incoming DateTransferDto
// instance, 'incomingDateTransDto'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  incomingDateTransDto  *DateTransferDto
//     - Data elements from input parameter, 'incomingDateTransDto' will be
//       used to populate the current DateTransferDto instance. When successfully
//       completed, all member data variables from 'incomingDateTransDto' and
//       the current DateTransferDto instance will be identical.
//
//
//  ePrefix           string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) CopyIn(
	incomingDateTransDto *DateTransferDto,
	ePrefix string) error {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.CopyIn() "

	dateTransUtil := dateTransferDtoUtility{}

	return dateTransUtil.copyIn(dateTransDto, incomingDateTransDto, ePrefix)
}


// Compare - Compares two instances of DateTransferDto date values.
// The method will return one of three comparison values.
//
//    Comparison                Comparison
//      Result                    Status
//    ----------                ----------
//       -1          DateTransferDto < incomingDateTransDto
//        0          DateTransferDto = incomingDateTransDto
//       +1          DateTransferDto > incomingDateTransDto
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  incomingDateTransDto     *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This method will
//       compare the date value of the current DateTransferDto instance
//       to to that of this input parameter ('incomingDateTransDto').
//
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
//  compareResult            int
//     - If this method completes successfully, a comparison
//       result will be returned as an integer value. The result
//       is based on a comparison of date values between the current
//       instance DateTransferDto and input parameter
//       'incomingDateTransDto'.
//
//       If the current DateTransferDto object is LESS THAN
//       'incomingDateTransDto', this method will return an integer
//       value of minus one (-1).
//
//       If the current DateTransferDto object is EQUAL to
//       'incomingDateTransDto' this method will return an integer
//       value of zero (0).
//
//       If the current DateTransferDto object is GREATER THAN
//       'incomingDateTransDto', this method will return an integer
//       value of plus one (+1).
//
//
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) Compare(
	incomingDateTransDto *DateTransferDto,
	ePrefix string)(
	compareResult int,
	err error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.CopyIn() "

	dateTransUtil := dateTransferDtoUtility{}

	compareResult,
	err = dateTransUtil.compare(
		dateTransDto,
		incomingDateTransDto,
		ePrefix)

	return compareResult, err
}


// CompareAreYearsAdjacent - Compares the astronomical year values
// encapsulated by the current DateTransferDto instance, and input
// parameter 'dateTransDtoTwo'. This comparison will determine if
// the two year values are adjacent. Adjacent years are defined
// here as differing by a value of plus or minus one year. If the
// year values are 'Adjacent', the difference as calculated by
// subtracting the year value of 'dateTransDtoTwo.year' from the
// current 'DateTransferDto.year'. This calculation is styled as
// DateTransferDto.year - dateTransDtoTwo.year and will always
// equal to plus or minus 1 (+1 or -1), if the two years are
// adjacent years.
//
// The 'Adjacent Year' comparison implemented by calculating the
// difference in year values will always have one of three possible
// outcomes (DateTransferDto.year - dateTransDtoTwo.year):
//
//  1. If 'DateTransferDto.year' IS ADJACENT to 'dateTransDtoTwo.year',
//     and the difference as calculated by 'DateTransferDto.year -
//     dateTransDtoTwo.year' is plus one (+1), it signals that
//     'DateTransferDto.year' is GREATER than 'dateTransDtoTwo.year'
//     by a value of '+1'. For this case, the returned boolean value
//     of 'areYearsAdjacent' is set to 'true' and the returned integer
//     value of 'compareResult' is set to +1.
//
//  2. If 'DateTransferDto.year' IS NOT ADJACENT to 'dateTransDtoTwo.year',
//     the calculation result of 'DateTransferDto.year - dateTransDtoTwo.year'
//     will be GREATER THAN plus one (+1) or minus one (-1). In this
//     case, the returned boolean value of 'areYearsAdjacent' will be set
//     to 'false' and the returned integer value of 'compareResult' will
//     be set to zero ('0').
//
//  3. If 'DateTransferDto.year' IS ADJACENT to 'dateTransDtoTwo.year',
//     and the difference as calculated by 'DateTransferDto.year -
//     dateTransDtoTwo.year' is minus one (-1), it signals that
//     'DateTransferDto.year' is LESS THAN 'dateTransDtoTwo.year' by a
//     value of minus one (-1). For this case, the returned boolean
//     value of 'areYearsAdjacent' is set to 'true' and the returned
//     integer value of 'compareResult' is set to -1.
//
// If either the current DateTransferDto instance or the input parameter
// 'dateTransDtoTwo' is judged to be invalid, this method will return an
// error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDtoTwo          *DateTransferDto
//     - A pointer to an instance of DateTransferDto. The astronomical
//       year value encapsulated by this instance will be compared to
//       the astronomical year value contained in the current DateTransferDto
//       instance in order to determine whether the two years are adjacent years
//       as defined in the narrative above.
//
//       If 'dateTransDtoTwo' is an invalid instance, an error will
//       be returned.
//
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
//  areYearsAdjacent         bool
//     - This boolean return value will signal whether the astronomical
//       year values contained in the current 'DateTransferDto' instance
//       and the input parameter 'dateTransDtoTwo' are adjacent years.
//       Adjacent years differ by a value of plus or minus 1-year.
//
//       If the current 'DateTransferDto' instance and 'dateTransDtoTwo'
//       are adjacent years, this boolean value is set to 'true'.
//       Otherwise, it is set to 'false'.
//
//
//  compareResult            int
//     - This integer return value will specify the comparison result
//       obtained from determining whether the current 'DateTransferDto'
//       instance and input parameter 'dateTransDtoTwo' are adjacent
//       years. As result of this comparison, this return parameter,
//       'compareResult', will be populated with one of three possible values.
//
//       1. If the current DateTransferDto instance IS ADJACENT to
//          input parameter 'dateTransDtoTwo', and the difference as
//          calculated by 'DateTransferDto.year - dateTransDtoTwo.year'
//          is plus one (+1), it signals that 'DateTransferDto.year' is
//          GREATER than 'dateTransDtoTwo.year' by a value of '+1'. For
//          this case, the returned integer value of 'compareResult' is
//          set to +1.
//
//       2. If the current DateTransferDto instance IS NOT ADJACENT to
//          'dateTransDtoTwo.year', the calculation result of
//          'DateTransferDto.year - dateTransDtoTwo.year' will be GREATER
//          THAN plus one (+1) or minus one (-1). In this case, the returned
//          integer value of 'compareResult' will be set to zero ('0').
//
//       3. If 'dateTransDtoOne.year' IS ADJACENT to 'dateTransDtoTwo.year',
//          and the difference as calculated by 'dateTransDtoOne.year -
//          'dateTransDtoTwo.year' is minus one (-1), it signals that
//          'dateTransDtoOne.year' is LESS THAN 'dateTransDtoTwo.year' by a
//          value of minus one (-1). For this case, the returned integer
//          value of 'compareResult' is set to -1.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) CompareAreYearsAdjacent(
	dateTransDtoTwo *DateTransferDto,
	ePrefix string) (
	areYearsAdjacent bool,
	compareResult int,
	err error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.CompareAreYearsAdjacent() "

	dateTransUtil := dateTransferDtoUtility{}

	areYearsAdjacent,
	compareResult,
	err = dateTransUtil.compareAreYearsAdjacent(
		dateTransDto,
		dateTransDtoTwo,
		ePrefix)

	return areYearsAdjacent, compareResult, err
}

// CompareYears - Compares the astronomical year values for the current
// DateTransferDto instance and another DateTransferDto instance passed
// as an input parameter ('dateTransDto2').
//
// The comparison result is returned as an integer value signaling whether
// DateTransferDto instance astronomical year value is GREATER THAN, EQUAL TO,
// or LESS THAN that of input parameter 'dateTransDto2'. That means that one
// of three possible comparison results will be returned by this method:
//
//      Return                Comparison
//      Value                   Result
//
//       -1 - The current DateTransferDto instance astronomical year value is
//            LESS THAN the 'timeTransDto2' astronomical year value.
//
//        0 - The current DateTransferDto instance astronomical year value is
//            EQUAL to the 'timeTransDto2' astronomical year value.
//
//       +1 - The current DateTransferDto instance astronomical year value is
//            GREATER THAN the 'timeTransDto2' astronomical year value.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDto2            *DateTransferDto
//     - A pointer to another incoming DateTransferDto instance. The
//       astronomical year value for the current DateTransferDto instance
//       will be compared to the year value of this second instance.
//
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
//  comparisonResult         int
//     - This method will return will return an integer value describing
//       the result of comparing the astronomical year values for the
//       current DateTransferDto instance and input parameter, 'dateTransDto2'.
//
//       Comparison Result
//       One of three possible comparison results will be returned.
//
//      Return                Comparison
//      Value                   Result
//       -1 - The current DateTransferDto instance astronomical year value is
//            LESS THAN the 'timeTransDto2' astronomical year value.
//
//        0 - The current DateTransferDto instance astronomical year value is
//            EQUAL to the 'timeTransDto2' astronomical year value.
//
//       +1 - The current DateTransferDto instance astronomical year value is
//            GREATER THAN the 'timeTransDto2' astronomical year value.
//
//
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) CompareYears(
	dateTransDto2 *DateTransferDto,
	ePrefix string) (
	comparisonResult int,
	err error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = &sync.Mutex{}
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.CompareYears() "

	dateTransUtil := dateTransferDtoUtility{}

	comparisonResult,
	err = dateTransUtil.compareYears(
		dateTransDto,
		dateTransDto2,
		ePrefix)

	return comparisonResult, err
}

// CopyOut - Returns a deep copy of the current DateTransferDto instance.
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
//  DateTransferDto
//     - A deep copy of the current DateTransferDto instance.
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) CopyOut(
	ePrefix string) (DateTransferDto, error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.CopyOut() "

	dateTransUtil := dateTransferDtoUtility{}

	newDateTransDto := DateTransferDto{}

	return dateTransUtil.copyOut(
		&newDateTransDto,
		ePrefix)
}

// Empty - Resets the internal data fields of the current DateTransferDto
// instance to invalid values. Effectively, the current DateTransferDto
// instance is rendered blank and invalid.
//
// If the current DateTransferDto instance is submitted for validity
// testing after calling this method, those tests will fail.
//
// Note that this method differs from method DateTransferDto.NewZeroInstance()
// where only some of the internal member variables are set to invalid values.
// In this case all of the member variables are set to invalid values.
//
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
func (dateTransDto *DateTransferDto) Empty(
	ePrefix string) error {

	if dateTransDto.lock == nil {
		dateTransDto.lock = &sync.Mutex{}
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	dateTransDtoMech := dateTransferDtoMechanics{}

	return dateTransDtoMech.empty(
		dateTransDto,
		ePrefix)

}

// Equal - Receives a pointer to a DateTransferDto object and compares
// the data values to those of the current DateTransferDto instance.
// If the data values contained in the two instances are equal, this
// method returns a boolean value set to 'true'.
//
// Be advised that if the either the current DateTransferDto object or
// the input parameter 'dateTransDto2' are evaluated as invalid, this
// this method will return an error message.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDto2            *DateTransferDto
//     - The data values of this DateTransferDto object will be
//       compared to those contained in the current DateTransferDto
//       instance. If all the data values are equal, this method will
//       return a boolean value set to 'true'.
//
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
//  areEqual                 bool
//     - If this method completes successfully, this boolean flag will
//       signal whether the data value contained in the current
//       DateTransferDto and those contained in the input parameter,
//       'dateTransDto2' are equal. A return value of 'true' signals
//       equality while a value of 'false' shows the two compared
//       instances are NOT equal.
//
//
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) Equal(
	dateTransDto2 *DateTransferDto,
	ePrefix string) (
	areEqual bool,
	err error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.Equal() "

	areEqual = false

	var compareInt int

	dateTransUtil := dateTransferDtoUtility{}

	compareInt,
		err = dateTransUtil.compare(
		dateTransDto,
		dateTransDto2,
		ePrefix)

	if err != nil {
		return areEqual, err
	}

	if compareInt == 0 {
		areEqual = true
	} else {
		areEqual = false
	}

	return areEqual, err
}

// ExchangeValues - Receives a pointer to an incoming DateTransferDto
// object and proceeds to exchange the data values of all internal
// member variables with those contained in the current DateTransferDto
// instance.
//
// If the method completes successfully, the current DateTransferDto
// instance will be populated with the original data values from
// input parameter 'dateTransDtoTwo'. Likewise, 'dateTransDtoTwo'
// will be populated with the original values copied from the current
// DateTransferDto instance.
//
// If either the current DateTransferDto instance or the input parameter
// 'dateTransDtoTwo' are judged invalid, this method will return an
// error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDtoTwo          *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This is one of the
//       DateTransferDto objects used in the data exchange. Data values from
//       this instance will be copied to the current DateTransferDto
//       instance while the original values from the current DateTransferDto
//       instance will be copied to this instance, 'dateTransDtoTwo'.
//
//       If 'dateTransDtoTwo' is an invalid instance, an error will
//       be returned.
//
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
//  error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) ExchangeValues(
	dateTransDtoTwo *DateTransferDto,
	ePrefix string) error {

	if dateTransDto.lock == nil {
		dateTransDto.lock = &sync.Mutex{}
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.ExchangeValues() "

	dateTransDtoMech := dateTransferDtoMechanics{}

	return dateTransDtoMech.exchangeValues(
		dateTransDto,
		dateTransDtoTwo,
		ePrefix)
}

// GetAbsoluteYearBigIntValue - Returns the absolute value of the
// Astronomical Year component of the date value encapsulated
// in the current DateTransferDto instance. The returned value
// is formatted as a big.Int number.
//
// Returning the absolute value for the year ensures that the
// returned value will always be a positive number.
//
//
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto instance
// before returning the value. To run a validity check on the
// DateTransferDto instance first call one of the two following
// methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
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
//  *big.Int
//     - The returned big.Int value will always be a positive number.
//
//
func (dateTransDto *DateTransferDto) GetAbsoluteYearBigIntValue(
	) *big.Int {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	var absYearValue int64

	absYearValue = dateTransDto.astronomicalYear

	if absYearValue < 0 {
		absYearValue *= -1
	}

	return big.NewInt(absYearValue)
}

// GetAbsoluteYearValue - Returns the absolute value of the
// Astronomical Year component of the date value encapsulated
// in the current DateTransferDto instance. Returning the absolute
// value for the year ensures that the returned value will always
// be a positive number.
//
//
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto instance
// before returning the value. To run a validity check on the
// DateTransferDto instance first call one of the two following
// methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
//
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
//     - The returned int64 value will always be a positive number.
//
//
func (dateTransDto *DateTransferDto) GetAbsoluteYearValue() int64 {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	if dateTransDto.astronomicalYear < 0 {
		return dateTransDto.astronomicalYear * -1
	}

	return dateTransDto.astronomicalYear
}

// GetDate - Returns the components that constitute the date
// value encapsulated by the current DateTransferDto instance.
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
//  yearValue                int64
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
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) GetDate(
	ePrefix string) (
	yearValue int64,
	yearType CalendarYearNumType,
	month int,
	day int,
	err error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.CopyOut() "

	yearValue = math.MinInt64
	yearType = CalendarYearNumType(0).None()
	month = -1
	day = -1

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_,
		err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		ePrefix)

	if err != nil {
		return yearValue, yearType, month, day, err
	}

	dateTransDtoMech := dateTransferDtoMechanics{}

	yearValue,
		yearType,
		err = dateTransDtoMech.getConvertedYearValueByType(
		dateTransDto.astronomicalYear,
		CalendarYearNumType(0).Astronomical(),
		dateTransDto.yearNumberingMode,
		ePrefix)

	if err != nil {
		return yearValue, yearType, month, day, err
	}

	month = dateTransDto.month
	day = dateTransDto.day

	return yearValue, yearType, month, day, err
}

// GetDay - Returns the day number component associated with the
// date value encapsulated in the current ADateTimeDto
// instance.
//
//
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto instance
// before returning the value. To run a validity check on the
// DateTransferDto instance first call one of the two following
// methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   --- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - The day component contained in the date value encapsulated
//       in the current DateTransferDto instance.
//
func (dateTransDto *DateTransferDto)GetDay() int {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	return dateTransDto.day
}

// GetDaysInYear - Returns the number of days contained in
// in the year component of the date value encapsulated in
// the current DateTransferDto instance.
//
// Be advised that this method will validate the the current
// DateTransferDto instance. If it is judged to be invalid
// this method will set the return parameter 'daysInYear'
// to minus one (-1) and return an error.
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
// daysInYear           int
//     - If this method completes successfully, this return parameter will
//       be populated with the number of days in the year for the year
//       value encapsulated in input parameter 'dateTransDto'. If an error
//       is encountered, this value will be set to minus one (-1).
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) GetDaysInYear(
	ePrefix string) (
	daysInYear int,
	err error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.GetDaysInYear() "

	dateTransUtil := dateTransferDtoUtility{}

	daysInYear,
	err = dateTransUtil.getDaysInYear(
		dateTransDto,
		ePrefix)

	return daysInYear, err
}

// GetHasLeapSecond - Returns a boolean flag signaling whether the
// date value encapsulated by the current DateTransferDto instance
// contains a 'Leap Second'. If the method returns 'true', then a
// 'Leap Second' is present on the specific date encapsulated by
// the current DateTransferDto instance. If the method returns
// 'false', no 'Leap Second' is present on the specified date.
//
//
// Leap Second
//
// A leap second is a one-second adjustment that is occasionally
// applied to Coordinated Universal Time (UTC) in order to accommodate
// the difference between precise time (as measured by atomic clocks)
// and imprecise observed solar time (known as UT1 and which varies due
// to irregularities and long-term slowdown in the Earth's rotation).
// If this return parameter is set to 'true', time and duration
// calculations will will assume the duration of the relevant 'day'
// is 24-hours plus one/ second. Otherwise, the duration of a day is
// assumed to consist of exactly 24-hours. For more information on the
// 'leap second', reference:
//    https://en.wikipedia.org/wiki/Leap_second
//
//
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto instance
// before returning the value. To run a validity check on the
// DateTransferDto instance first call one of the two following
// methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
//
//
func (dateTransDto *DateTransferDto) GetHasLeapSecond() bool {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	return dateTransDto.hasLeapSecond
}

// GetIsLeapYear - Returns a boolean value signaling whether
// the year value encapsulated in the current DateTransferDto
// is a leap year. If the return value is 'true', this year
// is classified as a leap year. Otherwise, the year is not
// a 'leap year'.
//
// For more information on 'leap year', reference:
//   https://en.wikipedia.org/wiki/Leap_year
//
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto instance
// before returning the value. To run a validity check on the
// DateTransferDto instance first call one of the two following
// methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   --- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If this method returns 'true', it signals that the year value
//       encapsulated in the current DateTransferDto instance is a
//       'leap year'. A return value of 'false' signals that the year
//       is not a 'leap year'.
//
func (dateTransDto *DateTransferDto) GetIsLeapYear() bool {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	if dateTransDto.calendarBaseData == nil {
		return false
	}

	var isLeapYear bool
	var err error

	isLeapYear,
	err =
	dateTransDto.calendarBaseData.IsLeapYear(
		dateTransDto.astronomicalYear,
		CalendarYearNumType(0).Astronomical(),
		"")

	if err != nil {
		return false
	}

	return isLeapYear
}

// GetISODayOfWeekNo - Returns the ISO Day Of The Week Number. This
// method receives a Julian Day Number and proceeds to calculate the day
// of the week number associated with that Julian Day Number Date. It
// then returns that day of the week number to the calling function
// formatted with the ISO 8601 Standard Day Of The Week Numbering System.
//
//
// ISO 8601 Standard Day Of The Week Numbering System
//
// ISO stands for the International Organization for Standardization (ISO).
//     https://www.iso.org/home.html
//
// The ISO 8601 Standard is the most common day of the week numbering system used
// internationally. This standard is used in Western Europe, Scandinavia, and most
// of Eastern Europe as well as many other nations across the globe.
//
// The ISO 8601 standard specifies that the week begins on Monday. Days of the week
// are numbered beginning with one (1) for Monday and ending with seven (7) for
// Sunday. Western European Calendars therefore show the first day of the week as
// Monday.
//
// The ISO 8601 Standard Day of the Week Numbering System is listed as follows:
//
//             ISO 8601
//             Standard
//            Day of Week       Day Of Week
//              Number             Name
//           ===============   ==============
//                 1              Monday
//                 2              Tuesday
//                 3              Wednesday
//                 4              Thursday
//                 5              Friday
//                 6              Saturday
//                 7              Sunday
//
// Use the numbering scheme listed above when accessing map data returned by
// this method for the ISO 8601 Standard Day Of The Week Numbering System.
//
// For more information on the ISO 8601 Standard Day Of The Week Numbering System,
// reference:
//   https://en.wikipedia.org/wiki/ISO_8601#Week_dates
//   https://www.timeanddate.com/date/week-numbers.html
//   Type: ISO8601DayOfWeekNo Source Code File: datetime/dayofweeknumberiso8601enum.go
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  julianDayNoDto      JulianDayNoDto
//     - This input parameter contains the data elements of a Julian Day
//       Number and Time value. Note that key Julian Day Number and Time
//       values are stored as *big.Int and *big.Float.
//
//        type JulianDayNoDto struct {
//          julianDayNo             *big.Int    // Julian Day Number expressed as integer value
//          julianDayNoFraction     *big.Float  // The Fractional Time value of the Julian
//                                              //   Day No Time
//          julianDayNoTime         *big.Float  // Julian Day Number Plus Time Fraction accurate to
//                                              //   within nanoseconds
//          julianDayNoNumericalSign int        // Sign of the Julian Day Number/Time value. This value
//                                              //   is either '+1' or '-1'
//          totalJulianNanoSeconds   int64      // Julian Day Number Time Value expressed in nanoseconds.
//                                              //   Always represents a positive value less than 36-hours
//          netGregorianNanoSeconds  int64      // Gregorian nanoseconds. Always represents a value in
//                                              //    nanoseconds which is less than 24-hours.
//          applyLeapSecond    bool // If set to 'true' it signals that the day identified
//                                  //   by this Julian Day Number has a duration go 24-hours
//                                  //   + 1-second.
//          hours               int // Gregorian Hours
//          minutes             int // Gregorian Minutes
//          seconds             int // Gregorian Seconds
//          nanoseconds         int // Gregorian Nanoseconds
//        }
//
//        The integer portion of the Julian Day Number (digits to left of
//        the decimal) represents the Julian day number and is
//        stored in 'JulianDayNoDto.julianDayNo'. The fractional
//        digits to the right of the decimal represent elapsed time
//        since noon on the Julian day number and is stored in
//        'JulianDayNoDto.julianDayNoFraction'. The combined Julian
//        Day Number Time value is stored in 'JulianDayNoDto.julianDayNoTime'.
//        All time values are expressed as Universal Coordinated Time (UTC).
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
//  isoDayOfWeekNo      ISO8601DayOfWeekNo
//     - If the method completes successfully, this parameter will contain
//       an enumeration type specifying the day of the week number formatted
//       according to the ISO 8601 Standard Day Of The Week Numbering System.
//       This enumeration value will allow access to the correct day number,
//       day name and day name abbreviation. The day of the week number is
//       calculated from input parameter 'julianDayNoDto'.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note this error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) GetISODayOfWeekNo(
	julianDayNoDto JulianDayNoDto,
	ePrefix string) (
	isoDayOfWeekNo ISO8601DayOfWeekNo,
	err error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.GetISODayOfWeekNo() "

	err = nil

	isoDayOfWeekNo = ISO8601DayOfWeekNo(0).None()

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_, err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		ePrefix + "Testing 'dateTransDto' Validity. ")

	if err != nil {
		return isoDayOfWeekNo, err
	}

	isoDayOfWeekNo,
		err = dateTransDto.calendarBaseData.GetISODayOfWeekNo(
		julianDayNoDto,
		ePrefix)

	return isoDayOfWeekNo, err
}

// GetUsDayOfWeekNo - Returns the US Day Of The Week Number. This method
// receives a Julian Day Number and proceeds to calculate the day of the
// week number associated with that Julian Day Number Date. It then returns
// that day of the week number to the calling function formatted with the
// US Day Of The Week Numbering System.
//
//
// US Day Of The Week Numbering System
//
// The United States, Canada, Australia and New Zealand put Sunday as the first
// day of the week on their calendars. The first day of the week, 'Sunday', is
// numbered as zero (0) with the last day of the week being numbered as six (6),
// for 'Saturday'. This system is referred to as the US Day Of The Week Numbering
// System and is listed as follows:
//
//                US
//            Day of Week        Day Of Week
//               Number             Name
//           ===============   ==============
//                 0               Sunday
//                 1               Monday
//                 2               Tuesday
//                 3               Wednesday
//                 4               Thursday
//                 5               Friday
//                 6               Saturday
//
// Use the numbering scheme listed above when accessing map data returned by
// this method for the US Day Of The Week Numbering System.
//
// For more information on the US Day Of The Week Numbering System, reference:
//   https://www.timeanddate.com/date/week-numbers.html
//   Type: UsDayOfWeekNo Source Code File: datetime/dayofweeknumberusenum.go
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  julianDayNoDto      JulianDayNoDto
//     - This input parameter contains the data elements of a Julian Day
//       Number and Time value. Note that key Julian Day Number and Time
//       values are stored as *big.Int and *big.Float.
//
//        type JulianDayNoDto struct {
//          julianDayNo             *big.Int    // Julian Day Number expressed as integer value
//          julianDayNoFraction     *big.Float  // The Fractional Time value of the Julian
//                                              //   Day No Time
//          julianDayNoTime         *big.Float  // Julian Day Number Plus Time Fraction accurate to
//                                              //   within nanoseconds
//          julianDayNoNumericalSign int        // Sign of the Julian Day Number/Time value. This value
//                                              //   is either '+1' or '-1'
//          totalJulianNanoSeconds   int64      // Julian Day Number Time Value expressed in nanoseconds.
//                                              //   Always represents a positive value less than 36-hours
//          netGregorianNanoSeconds  int64      // Gregorian nanoseconds. Always represents a value in
//                                              //    nanoseconds which is less than 24-hours.
//          applyLeapSecond    bool // If set to 'true' it signals that the day identified
//                                  //   by this Julian Day Number has a duration go 24-hours
//                                  //   + 1-second.
//          hours               int // Gregorian Hours
//          minutes             int // Gregorian Minutes
//          seconds             int // Gregorian Seconds
//          nanoseconds         int // Gregorian Nanoseconds
//        }
//
//        The integer portion of the Julian Day Number (digits to left of
//        the decimal) represents the Julian day number and is
//        stored in 'JulianDayNoDto.julianDayNo'. The fractional
//        digits to the right of the decimal represent elapsed time
//        since noon on the Julian day number and is stored in
//        'JulianDayNoDto.julianDayNoFraction'. The combined Julian
//        Day Number Time value is stored in 'JulianDayNoDto.julianDayNoTime'.
//        All time values are expressed as Universal Coordinated Time (UTC).
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
//  usDayOfWeekNo       UsDayOfWeekNo
//     - If the method completes successfully, this parameter will contain
//       an enumeration type specifying the day of the week number formatted
//       according to the US Day Of The Week Numbering System. This enumeration
//       value will allow access to the correct day number, day name and day
//       name abbreviation. The day of the week number is calculated from
//       input parameter 'julianDayNoDto'.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note this error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) GetUsDayOfWeekNo(
	julianDayNoDto JulianDayNoDto,
	ePrefix string) (
	usDayOfWeekNo UsDayOfWeekNo,
	err error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.GetUsDayOfWeekNo() "

	err = nil

	usDayOfWeekNo = UsDayOfWeekNo(0).None()

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_, err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		ePrefix + "Testing 'dateTransDto' Validity. ")

	if err != nil {
		return usDayOfWeekNo, err
	}

	usDayOfWeekNo,
		err = dateTransDto.calendarBaseData.GetUsDayOfWeekNo(
		julianDayNoDto,
		ePrefix)

	return usDayOfWeekNo, err
}

// GetMonth - Returns the month number component of the date
// value encapsulated oby the current ADateTimeDto
// instance.
//
//
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto instance
// before returning the value. To run a validity check on the
// DateTransferDto instance first call one of the two following
// methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   --- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - The month component of the date value encapsulated in the current
//       DateTransferDto instance.
//
func (dateTransDto *DateTransferDto) GetMonth() int{

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	return dateTransDto.day
}

// GetOrdinalDayNoInYear - Returns the Ordinal Day Number in the year
// value encapsulated in the current DateTransferDto instance. This is
// otherwise referred to as the Ordinal Date.
//
// The ordinal day number is a calendar date typically consisting of
// a year and a day number ranging between 1 and 366 (starting on
// January 1st). For more information on the Ordinal Day Number or the
// Ordinal Date, reference:
//    https://en.wikipedia.org/wiki/Ordinal_date
//
// Be advised that if the current DateTransferDto contains invalid member
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
//       the current DateTransferDto instance.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) GetOrdinalDayNoInYear(
	ePrefix string)(
	ordinalDayNoInYear int,
	err error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.CompareAreYearsAdjacent() "
	ordinalDayNoInYear = math.MinInt32
	err = nil

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_,
	err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		ePrefix + "Testing validity of current 'DateTransferDto' object. ")

	if err != nil {
		return ordinalDayNoInYear, err
	}

	ordinalDayNoInYear,
	err =
		dateTransDto.calendarBaseData.GetOrdinalDayNoFromDate(
			dateTransDto.astronomicalYear,
			CalendarYearNumType(0).Astronomical(),
			dateTransDto.month,
			dateTransDto.day,
			ePrefix)

	return ordinalDayNoInYear, err
}


// GetRemainingDaysInYear - Returns the number of days remaining
// in the year.
//
// Be advised that if the current DateTransferDto contains invalid member
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
func (dateTransDto *DateTransferDto) GetRemainingDaysInYear(
	ePrefix string) (
	remainingDaysInYear int,
	err error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.GetRemainingDaysInYear() "

	remainingDaysInYear = math.MinInt32

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_,
	err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		ePrefix +
			"Testing validity of current DateTransferDto instance. ")

	if err != nil {
		return remainingDaysInYear, err
	}

	remainingDaysInYear,
	err = dateTransDto.calendarBaseData.GetRemainingDaysInYear(
		dateTransDto.astronomicalYear,
		CalendarYearNumType(0).Astronomical(),
		dateTransDto.month,
		dateTransDto.day,
		ePrefix )

	return remainingDaysInYear, err
}

// GetYear - Returns the year value formatted as an
// Astronomical Year Value. This method is identical to
// method DateTransferDto.GetYearAstronomical() and is
// provided for naming convenience. To acquire the
// year value which may be formatted as a Common Era
// or Before Common Era value, see method
// DateTransferDto.GetYearWithType().
//
// This method returns the astronomical year value
// extracted from the date value encapsulated by
// the current DateTransferDto instance.
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
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto
// instance before returning the value. To run a validity
// check on the DateTransferDto instance first call one
// of the two following methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
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
//       current DateTransferDto instance.
//
func (dateTransDto *DateTransferDto) GetYear() int64 {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	return dateTransDto.astronomicalYear
}

// GetYearAstronomical - Returns the astronomical year
// value extracted from the date value encapsulated by
// the current DateTransferDto instance.
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
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto
// instance before returning the value. To run a validity
// check on the DateTransferDto instance first call one
// of the two following methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
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
//       current DateTransferDto instance.
//
func (dateTransDto *DateTransferDto) GetYearAstronomical() int64 {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	return dateTransDto.astronomicalYear
}

// GetYearAstronomicalBigInt - Returns the astronomical year
// value extracted from the date value encapsulated by
// the current DateTransferDto instance. The year value is
// returned as a type, *big.Int.
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
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto
// instance before returning the value. To run a validity
// check on the DateTransferDto instance first call one
// of the two following methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
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
//  *big.Int
//     - The Astronomical Year value is returned as a type *big.Int.
//       This is the Astronomical Year value encapsulated in the
//       current DateTransferDto instance.
//
func (dateTransDto *DateTransferDto) GetYearAstronomicalBigInt() *big.Int {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	return big.NewInt(dateTransDto.astronomicalYear)
}

// GetYearWithType - This method returns the year value and the corresponding
// year numbering type from the date value encapsulated in the current
// instance of DateTransferDto.
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
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto instance
// before returning the value. To run a validity check on the
// DateTransferDto instance first call one of the two following
// methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
//
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
//  yearValue                int64
//     - The year value contained in the date encapsulated by the current
//       DateTransferDto instance.
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
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto *DateTransferDto) GetYearWithType(
	ePrefix string) (
	yearValue int64,
	yearType CalendarYearNumType,
	err error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.GetYearWithType() "

	calMech := calendarMechanics{}

	yearValue,
		yearType,
		err = calMech.getCalendarYearByType(
		dateTransDto.astronomicalYear,
		dateTransDto.yearNumberingMode,
		ePrefix)

	return yearValue, yearType, err
}

// GetYearNumberSign - Returns a code for the numeric sign of the
// astronomical year value encapsulated in the current DateTransferDto
// instance.
//
// Possible Return Values:
//  +1 = Astronomical Year Value is Positive
//   0 = Astronomical Year Value is Zero
//  -1 = Astronomical Year Value is Negative
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
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto
// instance before returning the value. To run a validity check
// on the DateTransferDto instance first call one of the two
// following methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
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
//  int
//     - This returned integer value is a code for the numeric sign of the
//       astronomical year value encapsulated in the current DateTransferDto
//       instance.
//
//       Possible return values are listed as follows:
//         +1 = Astronomical Year Value is Positive
//          0 = Astronomical Year Value is Zero
//         -1 = Astronomical Year Value is Negative
//
func (dateTransDto *DateTransferDto) GetYearNumberSign() int {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	if dateTransDto.astronomicalYear == 0 {
		return 0
	} else if dateTransDto.astronomicalYear < 0 {
		return -1
	}

	// Must be Greater Than Zero
	return 1
}

// IsDateValid - This method tests the date value encapsulated
// in the current DateTransferDto according to the designated
// Calendar System associated with this date value. The date
// must be valid within the context of the designated calendar
// system. For example, the date 2019-02-29 would be declared
// 'INVALID' under the Gregorian Calendar because under that
// calendar system, the year 2019 is not a leap year and
// therefore the date February 29th did not exist.
//
// If the date encapsulated by the current DateTransferDto is
// judged valid under the designated calendar system, this
// method will return 'true'. Otherwise a value of 'false' is
// returned.
//
//
// IMPORTANT
//
// This method also validates the current DateTransferDto
// instance before returning the value. If any instance
// member variables are judged to be invalid, this method
// will return 'false'. Operationally, this method is identical
// to DateTransferDto.IsValidInstance().
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
//  bool
//     - If the date value encapsulated by the current DateTransferDto
//       instance qualifies as a valid date under the calendar system
//       associated with this date, a value of 'true' will be returned.
//       If the date value is judged as invalid, this boolean return
//       parameter will be set to 'false'.
//
func (dateTransDto *DateTransferDto) IsDateValid() bool {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	var isValid bool

	isValid,
	_ = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		"")

	return isValid
}

// IsDateValidError - This method tests the date value
// encapsulated in the current DateTransferDto according to
// the designated Calendar System associated with this date
// value. The date must be valid within the context of the
// designated calendar system. For example, the date
// '2019-02-29' would be declared 'INVALID' under the
// Gregorian Calendar because under that calendar system,
// the year 2019 is not a leap year and therefore the date
// February 29th did not exist.
//
// If the date encapsulated by the current DateTransferDto is
// judged valid under the designated calendar system, this
// method will return an error value of 'nil'. Otherwise
// an error type containing an appropriate error message
// will be returned.
//
//
// IMPORTANT
//
// This method also validates the current DateTransferDto
// instance before returning the value. If any instance
// member variables are judged to be invalid, this method
// will return an error. Operationally, this method is
// identical to method DateTransferDto.IsValidInstanceError().
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
//  error
//     - If the date value encapsulated by the current DateTransferDto
//       instance qualifies as a valid date under the calendar system
//       associated with this date, this returned error value will be
//       set to 'nil'. If the date value is judged as invalid, this
//       error type will contain an appropriate error message.
//
//       Be advised that if the current DateTransferDto contains
//       invalid member variable data values, an error message will
//       likewise be returned.
//
func (dateTransDto *DateTransferDto) IsDateValidError(
	ePrefix string) error {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.IsDateValidError() "

	var err error

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_, err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		ePrefix)

	return err
}

// IsValidInstance - Tests the current DateTransferDto instance
// for validity. In addition, the date value will also be validated
// within the context of the associated calendar system.
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
//  bool
//     - If the current DateTransferDto instance is valid and
//       properly initialized, this method will return 'true'.
//       If the current DateTransferDto instance is invalid,
//       a value of 'false' will be returned.
//
func (dateTransDto *DateTransferDto) IsValidInstance() bool {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	var isValid bool

	isValid, _ = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		"")

	return isValid
}

// IsValidInstanceError - Similar to method DateTransferDto.IsValidInstance().
// However, this method returns a error message.
//
// This method will test the current DateTransferDto instance for validity
// and return an error if the instance is invalid. In addition, the date value
// will also be validated within the context of the calendar system associated
// with this date.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix       string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - This method will analyze and test the current instance of
//       DateTransferDto for validity. If the instance is invalid, a type
//       'error' will be returned encapsulating an appropriate error message.
//       The error message will be prefixed with the error prefix string
//       (ePrefix) passed as an input parameter.
//
//       If the current DateTransferDto is valid an properly initialized, this
//       returned error type will be set to 'nil'.
//
func (dateTransDto *DateTransferDto) IsValidInstanceError(
	ePrefix string) error {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.IsValidInstanceError() "

	var err error

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_, err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		ePrefix)

	return err
}

// NewFromComponents - Returns a new DateTransferDto instance initialized from
// input parameters.
//
// This method defaults the Day Of The Week Numbering System Type to:
//           DayOfWeekNumberingSystemType(0).UsDayOfWeek()
//
// To configure this numbering system for alternative Day Of The Week
// Numbering Systems see method DateTransferDto.SetDayOfWeekNoSys().
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  calendarSystem  CalendarSpec
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
//  month               int
//     - The month number
//
//
//  day                 int
//    - The day number. This is the day number within the
//      the 'month' identified in the 'month' input parameter,
//      above.
//
//
//  hasLeapSecond       bool
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
//  tag                string
//     - A string description to be associated with the newly created DateTransferDto instance
//       generated by this method.
//
//
//  ePrefix            string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTransferDto
//     - If successful, this method will return a newly created instance of
//       DateTransferDto based on the input parameters identified above.
//
//
//  error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDto DateTransferDto) NewFromComponents(
	calendarSystem     CalendarSpec,
	year               int64,
	yearType           CalendarYearNumType,
	month              int,
	day                int,
	hasLeapSecond      bool,
	tag                string,
	ePrefix            string) (DateTransferDto, error) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.New() "

	newDateTransDto := DateTransferDto{}

	dateTransUtil := dateTransferDtoUtility{}

	err := dateTransUtil.setDateTransferDto(
		&newDateTransDto,
		year,
		yearType,
		month,
		day,
		hasLeapSecond,
		calendarSystem,
		tag,
		ePrefix)

	return newDateTransDto, err
}

// NewZeroInstance - Returns a new instance of DateTransferDto with
// all of the internal member variables set to their native 'zero'
// states. Effectively, this method will return a new instance of
// DateTransferDto which is blank and invalid.
//
// Note that this method differs from method DateTransferDto.Empty()
// where all of internal member variables are set to invalid values.
// In this case only some of the member variables are set to invalid
// values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newDateTransDto     DateTransferDto
//     - This method will return a new DateTransferDto instance
//       with all its internal member variables set to their native
//       'zero' values.
//
//
func (dateTransDto DateTransferDto) NewZeroInstance() (
newDateTransDto DateTransferDto) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	newDateTransDto = DateTransferDto{}

	dateTransDtoMech := dateTransferDtoMechanics{}

	dateTransDtoMech.setInstanceToZero(&newDateTransDto)

	return newDateTransDto
}

// SetHasLeapSecond - Sets the value of internal member variable,
// 'DateTransferDto.hasLeapSecond'. This member variable is used to specify
// whether the year, month and day encapsulated by the current DateTransferDto
// instance includes a leap second.
//
// A leap second is a one-second adjustment that is occasionally applied
// to Coordinated Universal Time (UTC) in order to accommodate the
// difference between precise time (as measured by atomic clocks) and
// imprecise observed solar time (known as UT1 and which varies due to
// irregularities and long-term slowdown in the Earth's rotation). If
// this parameter is set to 'true', time and duration calculations will
// assume the duration of the relevant 'day' is 24-hours plus one second.
// Otherwise, the duration of a day is assumed to consist of exactly
// 24-hours.
//
// If the input parameter, 'dateHasLeapSecond', is set to 'true', it signals
// that the day identified by year, month and day encapsulated by this
// DateTransferDto instance has a duration of 24-hours and one-second.
//
// Be advised that 'leap second' is very rarely used. For more information
// on the 'leap second', reference:
//    https://en.wikipedia.org/wiki/Leap_second
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateHasLeapSecond      bool
//     - Set this parameter to 'true' if you wish to configure the current
//       instance of DateTransferDto for a 'leap second'. Dates configured
//       with a leap second specify that the duration of the day is 24-hours
//       and one second. For more information on the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//               --- NONE ---
//
func (dateTransDto *DateTransferDto) SetHasLeapSecond(
	dateHasLeapSecond bool) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	dateTransDto.hasLeapSecond = dateHasLeapSecond

	return
}


// SetTagDescription - Sets the tag description associated with
// the current instance of DateTransferDto.
//
// This method will set the internal member variable DateTransferDto.tag
// with a description string as designated by the calling function.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  tagDescription  string
//     - A tag description used to set internal member variable
//       DateTransferDto.tag.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//          -- NONE --
//
//
func (dateTransDto *DateTransferDto) SetTagDescription(
	tagDescription string) {

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	dateTransDto.tag = tagDescription
}

// SetThisInstanceToZero - This method will proceed to set all of
// the current TimeTransferDto instance member variable data values
// to their native 'zero' state. Accordingly, some, but not all, of
// the member variables will be set to invalid values.
//
// If, after calling this method, the current 'timeTransDto' instance
// is submitted for validity testing, it will fail those tests.
//
// Note that this method differs from method, 'TimeTransferDto.Empty().
// The 'empty' operation will convert all member variables to invalid
// data values. This method will convert only some member variables to
// invalid values.
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
//       --- NONE ---
//
func (dateTransDto *DateTransferDto) SetThisInstanceToZero() {


	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.lock.Lock()

	defer dateTransDto.lock.Unlock()

	dateTransDtoMech := dateTransferDtoMechanics{}

	dateTransDtoMech.setInstanceToZero(dateTransDto)

	return
}

