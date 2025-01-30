package datetime

import (
	"fmt"
	"sync"
)

type TimeTransferDto struct {

	hour                 int                 // Hour component of time value
	minute               int                 // Minute component of time value
	second               int                 // Second component of time value
	hasLeapSecond        bool                // 'true' = time value includes a leap second
	nanosecond           int                 // Nanosecond component of time value
	totalTimeNanoseconds int64               // Total number of nanoseconds in the time value
	timeZone             TimeZoneDefinition  // Contains a detailed definition and descriptions of the Time
	tag                  string              // Tag Description string
	lock *sync.Mutex                         // Used for implementing thread safe operations.
}


// Compare - This methods receives an input parameters of type
// TimeTransferDto ('timeTransDtoTwo') and proceeds to compare the
// that time value with the time value of the current TimeTransferDto
// instance. The actual comparison is performed on the total
// nanoseconds values for the two TimeTransferDto instances.
//
// The comparison result is returned as an integer value set for one
// of three possible comparison outcomes:
//
//   If the current TimeTransferDto instance is LESS THAN
//   'timeTransDtoTwo', this method will return an integer value of
//   minus one (-1).
//
//   If the current TimeTransferDto instance is EQUAL to
//   'timeTransDtoTwo', this method will return an integer value of
//   zero (0).
//
//   If the current TimeTransferDto instance is GREATER THAN
//   'timeTransDtoTwo', this method will return an integer value of
//   plus one (+1).
//
//   Summary of Returned Comparison Results
//
//   Comparison                     Comparison
//     Result                         Status
//   ----------                     ----------
//      -1      current TimeTransferDto instance < timeTransDtoTwo
//       0      current TimeTransferDto instance = timeTransDtoTwo
//      +1      current TimeTransferDto instance > timeTransDtoTwo
//
// Be advised that this method will subject both the current
// TimeTransferDto instance and parameter 'timeTransDtoTwo' to
// validity testing. If either of the two TimeTransferDto
// objects is judged as invalid, this method will return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  timeTransDtoTwo          *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This method will
//       compare the total nanoseconds value of this TimeTransferDto
//       instance to the total nanoseconds value of the current
//       TimeTransferDto instance.
//
//       If this 'timeTransDtoTwo' instance is invalid, an error will be
//       returned.
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
//     - If this method completes successfully, a comparison result will
//       be returned as an integer value. The result is based on a
//       comparison of total nanosecond values associated with the current
//       TimeTransferDto instance and input parameter 'timeTransDtoTwo'.
//
//       If the current TimeTransferDto instance is LESS THAN
//       'timeTransDtoTwo', this method will return an integer value of
//       minus one (-1).
//
//       If the current TimeTransferDto instance is EQUAL to
//       'timeTransDtoTwo', this method will return an integer value of
//       zero (0).
//
//       If the current TimeTransferDto instance is GREATER THAN
//       'timeTransDtoTwo', this method will return an integer value of
//       plus one (+1).
//
//       Summary of Returned Comparison Results
//
//       Comparison                     Comparison
//         Result                         Status
//       ----------                     ----------
//          -1      current TimeTransferDto instance < timeTransDtoTwo
//           0      current TimeTransferDto instance = timeTransDtoTwo
//          +1      current TimeTransferDto instance > timeTransDtoTwo
//
//
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (timeTransDto *TimeTransferDto) Compare(
	timeTransDtoTwo *TimeTransferDto,
	ePrefix string) (
	compareResult int,
	err error) {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	ePrefix += "TimeTransferDto.Compare() "

	timeTransDtoUtil := timeTransferDtoUtility{}

	compareResult,
	err = timeTransDtoUtil.compare(
		timeTransDto,
		timeTransDtoTwo,
		ePrefix)

	return compareResult, err
}


// CopyIn - Populates the current TimeTransferDto instance with a deep copy
// of member data elements extracted from the the incoming TimeTransferDto
// instance, 'incomingTimeTransDto'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  incomingTimeTransDto  *TimeTransferDto
//     - Data elements from input parameter, 'incomingTimeTransDto', will be
//       used to populate the current TimeTransferDto instance. When successfully
//       completed, all member data variables from 'incomingTimeTransDto' and
//       the current TimeTransferDto instance will be identical.
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
func (timeTransDto *TimeTransferDto) CopyIn(
	incomingTimeTransDto *TimeTransferDto,
	ePrefix string) error {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	ePrefix += "TimeTransferDto.CopyIn() "

	timeTransDtoUtil := timeTransferDtoUtility{}

	return timeTransDtoUtil.copyIn(
		timeTransDto,
		incomingTimeTransDto,
		ePrefix)
}

// CopyOut - Returns a deep copy of the current TimeTransferDto instance.
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
//  TimeTransferDto
//     - A deep copy of the current TimeTransferDto instance.
//
//
//  error
//     - If successful, the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (timeTransDto *TimeTransferDto) CopyOut(
	ePrefix string) (TimeTransferDto, error) {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	ePrefix += "TimeTransferDto.CopyOut() "

	timeTransDtoUtil := timeTransferDtoUtility{}

	newTimeTransDto := TimeTransferDto{}

	return timeTransDtoUtil.copyOut(
		&newTimeTransDto,
		ePrefix)
}

// Empty - Resets the internal data fields of the current TimeTransferDto
// instance to invalid values. Effectively, the current TimeTransferDto
// instance is rendered blank and invalid.
//
// If the current TimeTransferDto instance is submitted for validity
// testing after calling this method, those tests will fail.
//
// Note that this method differs from method,
// 'TimeTransferDto.SetThisInstanceToZero()'. The 'empty' operation
// will convert all member variables to invalid data values.
// The 'TimeTransferDto.SetThisInstanceToZero()' method will convert
// only some member variables to invalid values.
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
func (timeTransDto *TimeTransferDto) Empty(
	ePrefix string) error {

	if timeTransDto.lock == nil {
		timeTransDto.lock = &sync.Mutex{}
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	timeTransDtoMech := timeTransferDtoMechanics{}

	return timeTransDtoMech.empty(
		timeTransDto,
		ePrefix)

}

// Equal - Receives a pointer to a TimeTransferDto object and compares
// the data values to those of the current TimeTransferDto instance.
// If the data values contained in the two instances are equal, this
// method returns a boolean value set to 'true'.
//
// Be advised that if the either the current TimeTransferDto object or
// the input parameter 'timeTransDto2' are evaluated as invalid, this
// this method will return an error message.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  timeTransDto2            *TimeTransferDto
//     - The data values of this TimeTransferDto object will be
//       compared to those contained in the current TimeTransferDto
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
//       TimeTransferDto and those contained in the input parameter,
//       'timeTransDto2' are equal. A return value of 'true' signals
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
func (timeTransDto *TimeTransferDto) Equal(
	timeTransDto2 *TimeTransferDto,
	ePrefix string) (
	areEqual bool,
	err error) {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	ePrefix += "TimeTransferDto.Equal() "

	areEqual = false

	var compareInt int

	timeTransDtoUtil := timeTransferDtoUtility{}

	compareInt,
	err = timeTransDtoUtil.compare(
		timeTransDto,
		timeTransDto2,
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

// ExchangeValues - Receives a pointer to an incoming TimeTransferDto
// object and proceeds to exchange the data values of all internal
// member variables with those contained in the current TimeTransferDto
// instance.
//
// If the method completes successfully, the current TimeTransferDto
// instance will be populated with the original data values from
// input parameter 'timeTransDtoTwo'.  Likewise, 'timeTransDtoTwo'
// will be populated with the original values copied from the current
// TimeTransferDto instance.
//
// If either the current TimeTransferDto instance or the input parameter
// 'timeTransDtoTwo' are judged invalid, this method will return an
// error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  timeTransDtoTwo          *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This is one of the
//       TimeTransferDto objects used in the data exchange. Data values from
//       this instance will be copied to the current TimeTransferDto
//       instance while the original values from the current TimeTransferDto
//       instance will be copied to this instance, 'timeTransDtoTwo'.
//
//       If 'timeTransDtoTwo' is an invalid instance, an error will
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
func (timeTransDto *TimeTransferDto) ExchangeValues(
	timeTransDto2 *TimeTransferDto,
	ePrefix string) (err error) {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	ePrefix += "TimeTransferDto.ExchangeValues() "

	timeTransDtoMech := timeTransferDtoMechanics{}

	err = timeTransDtoMech.exchangeValues(
		timeTransDto,
		timeTransDto2,
		ePrefix)

	return err
}

// GetHasLeapSecond - Returns a boolean value signaling
// whether or not the time value encapsulated by the
// current TimeTransferDto instance contains a 'leap second'.
//
//
// Leap Second
//
// A 'leap second' is a one-second adjustment that is occasionally
// applied to Coordinated Universal Time (UTC) in order to
// accommodate the difference between precise time (as measured by
// atomic clocks) and imprecise observed solar time (known as UT1
// and which varies due to irregularities and long-term slowdown
// in the Earth's rotation). If this return parameter is set to
// 'true', it signals that the time value  encapsulated by the
// current TimeTransferDto instance contains a 'leap second'.
// For more information on the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
// IMPORTANT
//
// This method does validate the current TimeTransferDto
// instance before returning the value. To run a validity
// check on the TimeTransferDto instance first call one
// of the two following methods:
//
//  TimeTransferDto.IsValidInstance() bool
//                OR
//  TimeTransferDto.IsValidInstanceError(ePrefix string) error
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
//     - A return value of 'true' signals that a 'leap second'
//       is present in the time value encapsulated by the current
//       TimeTransferDto instance
//
func (timeTransDto *TimeTransferDto) GetHasLeapSecond() bool {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	return timeTransDto.hasLeapSecond
}

// GetHour - Returns the hours component of the time value
// encapsulated by the current TimeTransferDto instance.
//
//
// IMPORTANT
//
// This method does validate the current TimeTransferDto
// instance before returning the value. To run a validity
// check on the TimeTransferDto instance first call one
// of the two following methods:
//
//  TimeTransferDto.IsValidInstance() bool
//                OR
//  TimeTransferDto.IsValidInstanceError(ePrefix string) error
//
//
func (timeTransDto *TimeTransferDto) GetHour() int {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	return timeTransDto.hour
}

// GetMinute - Returns the minutes component of the time value
// encapsulated by the current TimeTransferDto instance.
//
//
// IMPORTANT
//
// This method does validate the current TimeTransferDto
// instance before returning the value. To run a validity
// check on the TimeTransferDto instance first call one
// of the two following methods:
//
//  TimeTransferDto.IsValidInstance() bool
//                OR
//  TimeTransferDto.IsValidInstanceError(ePrefix string) error
//
//
func (timeTransDto *TimeTransferDto) GetMinute() int {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	return timeTransDto.minute
}

// GetNanosecond - Returns the nanoseconds component of the time value
// encapsulated by the current TimeTransferDto instance.
//
//
// IMPORTANT
//
// This method does validate the current TimeTransferDto
// instance before returning the value. To run a validity
// check on the TimeTransferDto instance first call one
// of the two following methods:
//
//  TimeTransferDto.IsValidInstance() bool
//                OR
//  TimeTransferDto.IsValidInstanceError(ePrefix string) error
//
//
func (timeTransDto *TimeTransferDto) GetNanosecond() int {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	return timeTransDto.nanosecond
}

// GetSecond - Returns the seconds component of the time value
// encapsulated by the current TimeTransferDto instance.
//
//
// IMPORTANT
//
// This method does validate the current TimeTransferDto
// instance before returning the value. To run a validity
// check on the TimeTransferDto instance first call one
// of the two following methods:
//
//  TimeTransferDto.IsValidInstance() bool
//                OR
//  TimeTransferDto.IsValidInstanceError(ePrefix string) error
//
//
func (timeTransDto *TimeTransferDto) GetSecond() int {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	return timeTransDto.second
}

// GetTotalNanoseconds - Returns a deep copy of the
// time zone definition object associated with the
// time value encapsulated by the current
// TimeTransferDto instance.
//
// A TimeZoneDefinition object provides complete
// information on a time zone. For more information
// on type TimeZoneDefinition, reference the source
// file documentation:
//  datetime/timezonedefinition.go
//
//
// IMPORTANT
//
// This method does validate the current TimeTransferDto
// instance before returning the value. To run a validity
// check on the TimeTransferDto instance first call one
// of the two following methods:
//
//  TimeTransferDto.IsValidInstance() bool
//                OR
//  TimeTransferDto.IsValidInstanceError(ePrefix string) error
//
//
func (timeTransDto *TimeTransferDto) GetTimeZone() TimeZoneDefinition {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	return timeTransDto.timeZone.CopyOut()
}

// GetTotalNanoseconds - Returns the total time value
// expressed in nanoseconds for the combined hours,
// minutes, seconds and nanoseconds contained in the
// the time value encapsulated by the current
// TimeTransferDto instance.
//
// Note that the returned value is of type int64.
//
//
// IMPORTANT
//
// This method does validate the current TimeTransferDto
// instance before returning the value. To run a validity
// check on the TimeTransferDto instance first call one
// of the two following methods:
//
//  TimeTransferDto.IsValidInstance() bool
//                OR
//  TimeTransferDto.IsValidInstanceError(ePrefix string) error
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
//     - This method will return the total time in nanoseconds. As
//       such, this value represents the sum of hours, minutes, seconds
//       and nanoseconds expressed as total nanoseconds.
//
func (timeTransDto *TimeTransferDto) GetTotalNanoseconds() int64 {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	return timeTransDto.totalTimeNanoseconds
}

// GetTimeValue - Returns a series of values specifying the time
// value encapsulated by the current TimeTransferDto instance.
//
// Taken together, the returned parameters of 'hours', 'minutes',
// 'seconds', 'nanoseconds' and time zone definition provides the
// complete time contain in the current TimeTransferDto instance.
//
// Be advised that this method validates the current TimeTransferDto
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
//       current TimeTransferDto instance.
//
//
//  minute              int
//     - The minute component of the time value encapsulated by the
//       current TimeTransferDto instance.
//
//
//  second              int
//     - The second component of the time value encapsulated by the
//       current TimeTransferDto instance.
//
//
//  nanosecond          int
//     - The nanosecond component of the time value encapsulated by
//       the current TimeTransferDto instance.
//
//
//  timeZoneDef         TimeZoneDefinition
//     - The the time zone definition used to configure the time
//       value encapsulated by the current TimeTransferDto
//       instance.
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
//       If an error occurs it will most likely be due to failure to
//       properly initialize the current TimeTransferDto instance.
//
func (timeTransDto *TimeTransferDto) GetTimeValue(
	ePrefix string) (
	hour int,
	minute int,
	second int,
	nanosecond int,
	timeZone TimeZoneDefinition,
	err error) {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	ePrefix += "TimeTransferDto.GetTimeValue() "

	timeZone, err = TimeZoneDefinition{}.New()

	if err != nil {
		return hour,
			minute,
			second,
			nanosecond,
			TimeZoneDefinition{},
			fmt.Errorf(ePrefix + "\n" +
				"Error returned by  TimeZoneDefinition{}.New()\n" +
				"Error='%v'\n", err.Error())
	}

	timeTransDtoNanobot := timeTransferNanobot{}

	_, err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDto,
		ePrefix + "Testing 'timeTransDto' Validity ")

	if err != nil {

		return hour,
			minute,
			second,
			nanosecond,
			timeZone,
			err
	}

	hour = timeTransDto.hour
	minute = timeTransDto.minute
	second = timeTransDto.second
	nanosecond = timeTransDto.nanosecond
	timeZone = timeTransDto.timeZone.CopyOut()

	return hour,
		minute,
		second,
		nanosecond,
		timeZone,
		err
}

// GetTag - Returns the tag description string associated with
// the current instance of TimeTransferDto. The tag description
// is stored in internal member variable, 'TimeTransferDto.tag'.
// This member variable is used to store any descriptive information
// required by the user.
//
// NOTE: To insert descriptive information in this tag member
// variable, call method TimeTransferDto.SetTagDescription().
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//          -- NONE --
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string - The returned string contains the tag description string
//           stored in internal member variable, 'TimeTransferDto.tag'.
//
func (timeTransDto *TimeTransferDto) GetTag() string {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	return timeTransDto.tag
}

// IsCoordinatedUniversalTime - Returns a boolean value signaling whether
// the time zone is set to Universal Coordinated Time.
//
// Coordinated Universal Time (or UTC) is the primary time standard by
// which the world regulates clocks and time. It is within about 1 second
// of mean solar time at 0° longitude, and is not adjusted for daylight
// saving time. It is effectively a successor to Greenwich Mean Time (GMT).
//   Reference: https://en.wikipedia.org/wiki/Coordinated_Universal_Time
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
//  isUtcTime           bool
//     - If this message completes with no errors, this value will be
//       set to 'true' if the current TimeTransferDto instance has a
//       time value with a time zone configured for Coordinated Universal
//       Time. This time is usually, but not always, referred to as the
//       'UTC' time zone.
//
//       If the time zone is NOT configured as Coordinated Universal Time
//       this return value is set to 'false'.
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
//       If an error occurs it will most likely be due to failure to
//       properly initialize the current TimeTransferDto instance.
//
func (timeTransDto *TimeTransferDto) IsCoordinatedUniversalTime(
	ePrefix string) (
	isUtcTime bool,
	err error) {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	ePrefix += "TimeTransferDto.IsCoordinatedUniversalTime() "

	timeTransDtoUtil := timeTransferDtoUtility{}

	isUtcTime,
		err = timeTransDtoUtil.isTimeZoneCoordinatedUniversalTime(
		timeTransDto,
		ePrefix)

	return isUtcTime, err
}

// IsValidInstance - Tests the current TimeTransferDto instance
// for validity.
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
//     - If the current TimeTransferDto instance is valid and
//       properly initialized, this method will return 'true'.
//       If the current TimeTransferDto instance is invalid,
//       a value of 'false' will be returned.
//
func (timeTransDto *TimeTransferDto) IsValidInstance() bool {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	var isValid bool

	timeTransDtoNanobot := timeTransferNanobot{}

	isValid, _ = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDto,
		"")

	return isValid
}

// IsValidInstanceError - Similar to method TimeTransferDto.IsValidInstance().
// However, this method returns a error message.
//
// This method will test the current TimeTransferDto instance for validity
// and return an error if the instance is invalid.
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
//     - This method will analyze and test the current instance of TimeTransferDto
//       for validity. If the instance is invalid, a type 'error' will be returned
//       encapsulating an appropriate error message. The error message will be prefixed
//       with the error prefix string (ePrefix) passed as an input parameter.
//
//       If the current TimeTransferDto is valid an properly initialized, this
//       returned error type will be set to 'nil'.
//
func (timeTransDto *TimeTransferDto) IsValidInstanceError(
	ePrefix string) error {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	ePrefix += "TimeTransferDto.IsValidInstanceError() "

	var err error

	timeTransDtoNanobot := timeTransferNanobot{}

	_, err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDto,
		ePrefix)

	return err
}

// NewFromComponents - This method returns a new, fully populated
// TimeTransferDto instance.
//
// TimeTransferDto objects encapsulate a complete time value including
// time zone.  Time values are maintained with nanosecond accuracy.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  hour                int
//     - The hour component of a time value.
//
//
//  minute              int
//     - The minute component of a time value.
//
//  second              int
//     - The second component of a time value.
//
//
//  hasLeapSecond       bool
//     - If this parameter is set to 'true', it signals that the time value
//       includes a "Leap Second". This parameter is rarely used and is almost
//       always set to 'false'.
//
//       A leap second is a one-second adjustment that is occasionally applied
//       to Coordinated Universal Time (UTC) in order to accommodate the
//       difference between precise time (as measured by atomic clocks) and
//       imprecise observed solar time (known as UT1 and which varies due to
//       irregularities and long-term slowdown in the Earth's rotation). If
//       this return parameter is set to 'true',      the time calculation
//       will assume the duration of the relevant 'day' is 24-hours plus one
//       second. Otherwise, the duration of a day is assumed to consist of
//       exactly 24-hours. For more information on the 'leap second',
//       reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  nanosecond          int
//       The nanosecond component of a time value.
//
//
//  timeZoneLocation    string
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
//  ePrefix           string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newTimeTransDto     TimeTransferDto
//     - If this method completes successfully, a new TimeTransferDto
//       instance encapsulating a complete time value will be returned.
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (timeTransDto TimeTransferDto) NewFromComponents(
	hour int,
	minute int,
	second int,
	hasLeapSecond bool,
	nanosecond int,
	timeZoneLocation string,
	ePrefix string) (
	newTimeTransDto TimeTransferDto,
	err error) {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	ePrefix += "TimeTransferDto.NewFromComponents() "

	err = nil

	newTimeTransDto = TimeTransferDto{}

	timeTransUtil := timeTransferDtoUtility{}

	err = timeTransUtil.setTimeTransferDto(
		&newTimeTransDto,
		hour,
		minute,
		second,
		hasLeapSecond,
		nanosecond,
		timeZoneLocation,
		ePrefix)

	return newTimeTransDto, err
}

// NewZeroInstance - Returns a new instance of DateTransferDto with
// all of the internal member variables set to their native 'zero'
// states. Effectively, this method will return a new instance of
// TimeTransferDto which is blank and invalid.
//
// Note that this method differs from method DateTransferDto.Empty()
// where all of internal member variables are set to invalid values.
// In this case only some of the member variables are set to invalid
// values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//      --- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newTimeTransDto     TimeTransferDto
//     - This method will return a new TimeTransferDto instance
//       with all its internal member variables set to their native
//       'zero' values.
//
func (timeTransDto TimeTransferDto) NewZeroInstance() (
	newTimeTransDto TimeTransferDto) {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	timeTransDtoMech := timeTransferDtoMechanics{}

	newTimeTransDto = TimeTransferDto{}

	timeTransDtoMech.setInstanceToZero(&newTimeTransDto)

	return newTimeTransDto
}

// SetHasLeapSecond - Sets the value of internal member variable,
// 'TimeTransferDto.hasLeapSecond'. This member variable is used to specify
// whether the hour, minute, second and nanosecond time value encapsulated
// by the current TimeTransferDto instance contains a 'leap second'.
//
// A 'leap second' is a one-second adjustment that is occasionally applied
// to Coordinated Universal Time (UTC) in order to accommodate the
// difference between precise time (as measured by atomic clocks) and
// imprecise observed solar time (known as UT1 and which varies due to
// irregularities and long-term slowdown in the Earth's rotation). If
// this parameter is set to 'true', time calculations will assume the
// duration of the relevant 'day' is 24-hours plus one second. Otherwise,
// the duration of a day is assumed to consist of exactly 24-hours.
//
// If the input parameter, 'timeHasLeapSecond', is set to 'true', it signals
// that the time value identified by hour, minute, second and nanosecond
// encapsulated by this TimeTransferDto instance has a duration of 24-hours
// and one-second.
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
//  timeHasLeapSecond      bool
//     - Set this parameter to 'true' if you wish to configure the current
//       time value encapsulated by this TimeTransferDto instance contains
//       a 'leap second'. Time values configured with a 'leap second' specify
//       that the duration of the day is 24-hours plus one second. For more
//       information on the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//               --- NONE ---
//
func (timeTransDto *TimeTransferDto) SetHasLeapSecond(
	timeHasLeapSecond bool) {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	timeTransDto.hasLeapSecond = timeHasLeapSecond

	return
}

// SetTagDescription - Sets the tag description associated with
// the current instance of TimeTransferDto.
//
// This method will set the internal member variable 'TimeTransferDto.tag'
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
func (timeTransDto *TimeTransferDto) SetTagDescription(
	tagDescription string) {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	timeTransDto.tag = tagDescription
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
func (timeTransDto *TimeTransferDto) SetThisInstanceToZero() {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	timeTransDtoMech := timeTransferDtoMechanics{}

	timeTransDtoMech.setInstanceToZero(timeTransDto)

	return
}

// SetTimeValue - Sets the time value for the current TimeTransferDto
// instance. The previous time value will be overwritten.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  hour                int
//     - The hour component of a time value.
//
//
//  minute              int
//     - The minute component of a time value.
//
//  second              int
//     - The second component of a time value.
//
//
//  hasLeapSecond       bool
//     - If this parameter is set to 'true', it signals that the time value
//       includes a "Leap Second". This parameter is rarely used and is almost
//       always set to 'false'.
//
//       A leap second is a one-second adjustment that is occasionally applied
//       to Coordinated Universal Time (UTC) in order to accommodate the
//       difference between precise time (as measured by atomic clocks) and
//       imprecise observed solar time (known as UT1 and which varies due to
//       irregularities and long-term slowdown in the Earth's rotation). If
//       this return parameter is set to 'true',      the time calculation
//       will assume the duration of the relevant 'day' is 24-hours plus one
//       second. Otherwise, the duration of a day is assumed to consist of
//       exactly 24-hours. For more information on the 'leap second',
//       reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  nanosecond          int
//       The nanosecond component of a time value.
//
//
//  timeZoneLocation    string
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
func (timeTransDto *TimeTransferDto) SetTimeValue(
	hour int,
	minute int,
	second int,
	hasLeapSecond bool,
	nanosecond int,
	timeZoneLocation string,
	ePrefix string) (
	err error) {

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.lock.Lock()

	defer timeTransDto.lock.Unlock()

	ePrefix += "TimeTransferDto.SetTimeValue() "

	err = nil

	timeTransUtil := timeTransferDtoUtility{}

	err = timeTransUtil.setTimeTransferDto(
		timeTransDto,
		hour,
		minute,
		second,
		hasLeapSecond,
		nanosecond,
		timeZoneLocation,
		ePrefix)

	return err
}


