package datetime

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
)

type timeTransferDtoUtility struct {

	lock *sync.Mutex

}

// compare - This methods receives two input parameters of type
// TimeTransferDto labeled as 'timeTransDtoOne' and 'timeTransDtoTwo'.
// It then proceeds to compare the total nanoseconds value for both
// parameters. The comparison result is returned as an integer value
// set for one of three possible comparison outcomes:
//
//   If 'timeTransDtoOne' is LESS THAN 'timeTransDtoTwo', this method
//   will return an integer value of minus one (-1).
//
//   If 'timeTransDtoOne' is EQUAL to 'timeTransDtoTwo', this method
//   will return an integer value of zero (0).
//
//   If 'timeTransDtoOne' is GREATER THAN 'timeTransDtoTwo', this
//   method will return an integer value of plus one (+1).
//
// Summary of Returned Comparison Results
//
//    Comparison           Comparison
//      Result               Status
//    ----------           ----------
//       -1      timeTransDtoOne < timeTransDtoTwo
//        0      timeTransDtoOne = timeTransDtoTwo
//       +1      timeTransDtoOne > timeTransDtoTwo
//
// Be advised that this method will subject both input parameters to
// validity testing. If either 'timeTransDtoOne' or 'timeTransDtoTwo'
// is judged as invalid, this method will return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  timeTransDtoOne          *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This method will
//       compare the total nanoseconds value of this TimeTransferDto
//       instance to the total nanoseconds value of the second input
//       parameter 'timeTransDtoTwo'.
//
//       If this 'timeTransDtoOne' instance is invalid, an error will be
//       returned.
//
//
//  timeTransDtoTwo          *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This method will
//       compare the total nanoseconds value of this TimeTransferDto
//       instance to the total nanoseconds value of the first input
//       parameter 'timeTransDtoOne'.
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
//       comparison of total nanosecond values associated with the input
//       parameters, 'timeTransDtoOne' and 'timeTransDtoTwo'.
//
//       If 'timeTransDtoOne' is LESS THAN 'timeTransDtoTwo', this method
//       will return an integer value of minus one (-1).
//
//       If 'timeTransDtoOne' is EQUAL to 'timeTransDtoTwo', this method
//       will return an integer value of zero (0).
//
//       If 'timeTransDtoOne' is GREATER THAN 'timeTransDtoTwo', this
//       method will return an integer value of plus one (+1).
//
//       Summary of Returned Comparison Results
//
//          Comparison           Comparison
//            Result               Status
//          ----------           ----------
//             -1      timeTransDtoOne < timeTransDtoTwo
//              0      timeTransDtoOne = timeTransDtoTwo
//             +1      timeTransDtoOne > timeTransDtoTwo
//
//
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (timeTransDtoUtil *timeTransferDtoUtility) compare(
	timeTransDtoOne *TimeTransferDto,
	timeTransDtoTwo *TimeTransferDto,
	ePrefix string) (
	compareResult int,
	err error) {


	if timeTransDtoUtil.lock == nil {
		timeTransDtoUtil.lock = new(sync.Mutex)
	}

	timeTransDtoUtil.lock.Lock()

	defer timeTransDtoUtil.lock.Unlock()

	ePrefix += "timeTransferDtoUtility.compare() "

	compareResult = math.MinInt32
	err = nil

	if timeTransDtoOne == nil {
		err = errors.New(ePrefix +
			"\nError: Input parameter 'timeTransDtoOne' is a nil pointer!")
		return compareResult, err
	}

	if timeTransDtoOne.lock == nil {
		timeTransDtoOne.lock = new(sync.Mutex)
	}

	if timeTransDtoTwo == nil {
		err = errors.New(ePrefix +
			"\nError: Input parameter 'timeTransDtoTwo' is a nil pointer!")
		return compareResult, err
	}

	if timeTransDtoTwo.lock == nil {
		timeTransDtoTwo.lock = new(sync.Mutex)
	}

	timeTransDtoNanobot := timeTransferNanobot{}

	_, err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDtoOne,
		ePrefix + "Testing 'timeTransDtoOne' Validity ")

	if err != nil {
		return compareResult, err
	}

	_, err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDtoTwo,
		ePrefix + "Testing 'timeTransDtoTwo' Validity ")

	if err != nil {
		return compareResult, err
	}


	if timeTransDtoOne.totalTimeNanoseconds >
		timeTransDtoTwo.totalTimeNanoseconds {
		compareResult = 1
	} else if timeTransDtoOne.totalTimeNanoseconds <
		timeTransDtoTwo.totalTimeNanoseconds {
		compareResult = -1
	} else {
		// Must be equal
		compareResult = 0
	}

	return compareResult, err
}

// copyIn - Makes a deep copy of incoming TimeTransferDto instance
// 'incomingTimeTransDto' and stores the data in the internal member
// variables of 'oldTimeTransDto', the original TimeTransferDto
// instance. All member variable data values in 'oldTimeTransDto'
// will be overwritten.
//
// If this method completes successfully, the internal member variable
// data values for both 'oldTimeTransDto' and 'incomingTimeTransDto'
// will be identical.
//
// This version of the 'copyIn' methods performs no validity checking on
// input parameters 'oldTimeTransDto' or 'incomingTimeTransDto'. It only
// performs the low level deep copy operation.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  oldTimeTransDto        *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This method will
//       perform a deep copy of member variable data values from input
//       parameter 'incomingTimeTransDto' to 'oldTimeTransDto'. All
//       original member variable data values contained in
//       'oldTimeTransDto' will be overwritten.
//
//
//  incomingTimeTransDto   *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This method will
//       perform a deep copy and transfer member variable data values
//       from 'incomingTimeTransDto' to 'oldTimeTransDto'.
//
//
//  ePrefix                string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil' and member variable data values will be copied
//       from 'incomingTimeTransDto' to 'oldTimeTransDto'.
//
//       If errors are encountered during processing, the returned error
//       Type will encapsulate an error message. Note that this error
//       message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
//
func (timeTransDtoUtil *timeTransferDtoUtility) copyIn(
	oldTimeTransDto *TimeTransferDto,
	incomingTimeTransDto *TimeTransferDto,
	ePrefix string) (err error) {

	if timeTransDtoUtil.lock == nil {
		timeTransDtoUtil.lock = new(sync.Mutex)
	}

	timeTransDtoUtil.lock.Lock()

	defer timeTransDtoUtil.lock.Unlock()

	ePrefix += "timeTransferDtoUtility.copyIn() "

	if oldTimeTransDto == nil {
		err = errors.New(ePrefix +
			"\nError: Input parameter 'oldTimeTransDto' is a nil pointer!")
		return err
	}

	if oldTimeTransDto.lock == nil {
		oldTimeTransDto.lock = new(sync.Mutex)
	}

	if incomingTimeTransDto == nil {
		err = errors.New(ePrefix +
			"\nError: Input parameter 'incomingTimeTransDto' is a nil pointer!")
		return err
	}

	if incomingTimeTransDto.lock == nil {
		incomingTimeTransDto.lock = new(sync.Mutex)
	}

	timeTransDtoNanobot := timeTransferNanobot{}

	_, err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		incomingTimeTransDto,
		ePrefix + "Testing 'incomingTimeTransDto' Validity ")

	if err != nil {
		return err
	}

	timeTransDtoMech := timeTransferDtoMechanics{}

	err = timeTransDtoMech.empty(
		oldTimeTransDto,
		ePrefix + "Emptying 'oldTimeTransDto' ")

	if err != nil {
		return err
	}

	oldTimeTransDto.hour  = incomingTimeTransDto.hour
	oldTimeTransDto.minute = incomingTimeTransDto.minute
	oldTimeTransDto.second = incomingTimeTransDto.second

	oldTimeTransDto.hasLeapSecond =
		incomingTimeTransDto.hasLeapSecond

	oldTimeTransDto.nanosecond = incomingTimeTransDto.nanosecond
	oldTimeTransDto.totalTimeNanoseconds =
		incomingTimeTransDto.totalTimeNanoseconds

	oldTimeTransDto.timeZone =
		incomingTimeTransDto.timeZone.CopyOut()

	oldTimeTransDto.tag = incomingTimeTransDto.tag

	_, err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		oldTimeTransDto,
		ePrefix + "Testing 'oldTimeTransDto' Validity ")

	return err
}


// copyOut - Returns a deep copy of input parameter 'timeTransDto' which
// is a pointer to an instance of 'TimeTransferDto'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  timeTransDto        *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
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
//  newTimeTransDto     TimeTransferDto
//     - If successful this method returns a deep copy of the input
//       parameter, 'timeTransDto'. When the copy operation is completed
//       member variables contained in return value, 'newTimeTransDto'
//       will be identical to those contained in input parameter,
//      'timeTransDto'.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (timeTransDtoUtil *timeTransferDtoUtility) copyOut(
	timeTransDto *TimeTransferDto,
	ePrefix string) (
	newTimeTransDto TimeTransferDto,
	err error) {

	if timeTransDtoUtil.lock == nil {
		timeTransDtoUtil.lock = new(sync.Mutex)
	}

	timeTransDtoUtil.lock.Lock()

	defer timeTransDtoUtil.lock.Unlock()

	ePrefix += "timeTransferDtoUtility.copyOut() "

	err = nil

	newTimeTransDto = TimeTransferDto{}

	if timeTransDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'timeTransDto' is a nil pointer!")

		return newTimeTransDto, err
	}

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDtoMech := timeTransferDtoMechanics{}

	err = timeTransDtoMech.empty(
		&newTimeTransDto,
		ePrefix + "Emptying 'newTimeTransDto' ")

	if err != nil {
		return newTimeTransDto, err
	}

	timeTransDtoNanobot := timeTransferNanobot{}

	_, err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDto,
		ePrefix + "Testing 'timeTransDto' Validity ")

	if err != nil {
		return newTimeTransDto, err
	}

	newTimeTransDto.hour = timeTransDto.hour
	newTimeTransDto.minute = timeTransDto.minute
	newTimeTransDto.second = timeTransDto.second
	newTimeTransDto.hasLeapSecond = timeTransDto.hasLeapSecond
	newTimeTransDto.nanosecond = timeTransDto.nanosecond

	newTimeTransDto.totalTimeNanoseconds =
		timeTransDto.totalTimeNanoseconds

	newTimeTransDto.timeZone = timeTransDto.timeZone.CopyOut()

	_, err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		&newTimeTransDto,
		ePrefix + "Testing 'newTimeTransDto' Validity ")

	return newTimeTransDto, err
}

// setTimeTransferDto - populates all the values of a TimeTransferDto
// instance passed as an input parameter.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  timeTransDto    *TimeTransferDto
//     - A pointer to a TimeTransferDto instance. If this method completes
//       all of the internal member data elements will be zeroed and populated
//       with new value supplied by the remaining input parameters.
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
//
//  error
//     - If successful the returned error Type is set equal to 'nil' and
//       the input parameter 'timeTransDto' will be repopulated with new
//       data values.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func(timeTransDtoUtil *timeTransferDtoUtility) setTimeTransferDto(
	timeTransDto *TimeTransferDto,
	hour int,
	minute int,
	second int,
	hasLeapSecond bool,
	nanosecond int,
	timeZoneLocation string,
	ePrefix string) (err error) {

	if timeTransDtoUtil.lock == nil {
		timeTransDtoUtil.lock = new(sync.Mutex)
	}

	timeTransDtoUtil.lock.Lock()

	defer timeTransDtoUtil.lock.Unlock()

	ePrefix += "timeTransferDtoUtility.setTimeTransferDto() "

	if timeTransDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'timeTransDto' is a 'nil' pointer!\n")

		return err
	}

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}
	
	if hour > 24 ||
		hour < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hour' is INVALID!\n" +
			"The 'hour' value must be <= 24 and >= 0\n" +
			"hour='%v'\n", hour)

		return err
	}

	if minute > 59 ||
		minute < 0 {

		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'minute' is INVALID!\n" +
			"The 'minute' value must be <= 59 and >= 0\n" +
			"minute='%v'\n", minute)

		return err
	}

	maxSecond := 59

	if hasLeapSecond {
		maxSecond = 60
	}

	if second < 0 ||
		second > maxSecond {

		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'second' is INVALID!\n" +
			"The 'second' value must be <= %v and >= 0\n" +
			"second='%v'\n",
			maxSecond,
			second)

		return err
	}

	if nanosecond < 0 ||
		nanosecond > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecond' is INVALID!\n" +
			"The 'nanosecond' value must be <= 999,999,999 and >= 0\n" +
			"nanosecond='%v'\n",
			nanosecond)

		return err

	}

	if timeZoneLocation == "" {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'timeZoneLocation' is INVALID!\n" +
			"timeZoneLocation = EMPTY STRING!\n")
		return err
	}

	tzDefUtil := timeZoneDefUtility{}

	baseDateTime := time.Now().UTC()

	err = tzDefUtil.setFromTimeZoneName(
		&timeTransDto.timeZone,
		baseDateTime,
		TimeZoneConversionType(0).Relative(),
		timeZoneLocation,
		ePrefix)

	if err != nil {
		return err
	}

	timeTransDto.hour = hour
	timeTransDto.minute = minute
	timeTransDto.second = second
	timeTransDto.hasLeapSecond = hasLeapSecond
	timeTransDto.nanosecond = nanosecond

	timeTransDto.totalTimeNanoseconds = 0

	timeTransDto.totalTimeNanoseconds += int64(time.Hour) * int64(hour)
	timeTransDto.totalTimeNanoseconds += int64(time.Minute) * int64(minute)
	timeTransDto.totalTimeNanoseconds += int64(time.Second) * int64(second)
	timeTransDto.totalTimeNanoseconds += int64(nanosecond)

	timeTransDtoNanobot := timeTransferNanobot{}

	_, err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDto,
		ePrefix + "Final Validity Check on 'timeTransDto' ")

	return err
}

// isTimeZoneCoordinatedUniversalTime - Returns a boolean value signaling whether
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
func(timeTransDtoUtil *timeTransferDtoUtility) isTimeZoneCoordinatedUniversalTime(
	timeTransDto *TimeTransferDto,
	ePrefix string) (
	isUtcTime bool,
	err error) {

	if timeTransDtoUtil.lock == nil {
		timeTransDtoUtil.lock = new(sync.Mutex)
	}

	timeTransDtoUtil.lock.Lock()

	defer timeTransDtoUtil.lock.Unlock()

	ePrefix += "timeTransferDtoUtility.isTimeZoneCoordinatedUniversalTime() "

	isUtcTime = false
	err = nil

	if timeTransDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'timeTransDto' is a 'nil' pointer!\n")

		return isUtcTime, err
	}

	timeTransDtoNanobot := timeTransferNanobot{}

	_, err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDto,
		ePrefix + "Testing 'timeTransDto' Validity ")

	if err != nil {
		return isUtcTime, err
	}

	if timeTransDto.timeZone.originalTimeZone.zoneOffsetTotalSeconds == 0 {
		isUtcTime = true
		return
	} else {
		isUtcTime = false
	}

	return isUtcTime, err
}