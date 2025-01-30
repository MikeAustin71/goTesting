package datetime

import (
	"errors"
	"sync"
)

type timeTransferDtoMechanics struct {

	lock *sync.Mutex

}

// empty - Receives a pointer to a TimeTransferDto instance
// and proceeds to set the internal data elements to invalid
// values.
//
// If, after calling this method, the instance 'timeTransDto'
// is submitted for validity testing, it will fail.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  timeTransDto        *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This method WILL
//       change the values of internal member variables to achieve
//       the method's objectives.
//
//       The 'empty' operation will set all internal member variables
//       to invalid data values.
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
func (timeTransDtoMech *timeTransferDtoMechanics) empty(
	timeTransDto *TimeTransferDto,
	ePrefix string) (err error) {

	if timeTransDtoMech.lock == nil {
		timeTransDtoMech.lock = new(sync.Mutex)
	}

	timeTransDtoMech.lock.Lock()

	defer timeTransDtoMech.lock.Unlock()

	ePrefix += "timeTransferDtoMechanics.empty() "

	err = nil

	if timeTransDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'timeTransDto' is a 'nil' pointer!\n")

		return err
	}

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.hour = -1
	timeTransDto.minute = -1
	timeTransDto.second = -1
	timeTransDto.hasLeapSecond = false
	timeTransDto.nanosecond = -1
	timeTransDto.totalTimeNanoseconds = -1
	timeTransDto.tag = ""

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setZeroTimeZoneDef(
		&timeTransDto.timeZone,
		ePrefix + "Zeroing Time Zone ")

	return err
}

// exchangeValues - Receives pointers to two TimeTransferDto objects and
// proceeds to exchange the data values of all internal member variables.
// If the method completes successfully, 'timeTransDtoOne' will be populated
// with the original data values of 'timeTransDtoTwo'. Likewise, 'timeTransDtoTwo'
// will be populated with the original values of 'timeTransDtoOne'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  timeTransDtoOne          *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This is one of the
//       two TimeTransferDto objects used in the data exchange. Data values
//       from this instance will be copied to input parameter 'timeTransDtoTwo'
//       while the original values from 'timeTransDtoTwo' will be copied
//       to this instance, 'timeTransDtoOne'.
//
//       If 'timeTransDtoOne' is an invalid instance, an error will
//       be returned.
//
//
//  timeTransDtoTwo          *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This is one of the
//       TimeTransferDto objects used in the data exchange. Data values from
//       this instance will be copied to input parameter 'timeTransDtoOne'
//       while the original values from 'timeTransDtoOne' will be copied
//       to this instance, 'timeTransDtoTwo'.
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
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func(timeTransDtoMech *timeTransferDtoMechanics) exchangeValues(
	timeTransDto1    *TimeTransferDto,
	timeTransDto2    *TimeTransferDto,
	ePrefix string) (err error) {

	if timeTransDtoMech.lock == nil {
		timeTransDtoMech.lock = new(sync.Mutex)
	}

	timeTransDtoMech.lock.Lock()

	defer timeTransDtoMech.lock.Unlock()

	ePrefix += "timeTransferDtoUtility.setTimeTransferDto() "

	if timeTransDto1 == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'timeTransDto1' is a 'nil' pointer!\n")

		return err
	}

	if timeTransDto2 == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'timeTransDto2' is a 'nil' pointer!\n")

		return err
	}

	timeTransDtoNanobot := timeTransferNanobot{}

	_,
		err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDto1,
		ePrefix + "Testing Validity of 'timeTransDto1'. ")

	if err != nil {
		return err
	}

	_,
		err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDto2,
		ePrefix + "Testing Validity of 'timeTransDto2'. ")

	if err != nil {
		return err
	}

	tempTimeTransDto := TimeTransferDto{}
	tempTimeTransDto.hour = timeTransDto1.hour
	tempTimeTransDto.minute = timeTransDto1.minute
	tempTimeTransDto.second = timeTransDto1.second
	tempTimeTransDto.hasLeapSecond = timeTransDto1.hasLeapSecond
	tempTimeTransDto.nanosecond = timeTransDto1.nanosecond
	tempTimeTransDto.totalTimeNanoseconds = timeTransDto1.totalTimeNanoseconds
	tempTimeTransDto.timeZone = timeTransDto1.timeZone.CopyOut()
	tempTimeTransDto.tag = timeTransDto1.tag

	timeTransDto1.hour = timeTransDto2.hour
	timeTransDto1.minute = timeTransDto2.minute
	timeTransDto1.second = timeTransDto2.second
	timeTransDto1.hasLeapSecond = timeTransDto2.hasLeapSecond
	timeTransDto1.nanosecond = timeTransDto2.nanosecond
	timeTransDto1.totalTimeNanoseconds = timeTransDto2.totalTimeNanoseconds
	timeTransDto1.timeZone = timeTransDto2.timeZone.CopyOut()
	timeTransDto1.tag = timeTransDto2.tag

	timeTransDto2.hour = tempTimeTransDto.hour
	timeTransDto2.minute = tempTimeTransDto.minute
	timeTransDto2.second = tempTimeTransDto.second
	timeTransDto2.hasLeapSecond = tempTimeTransDto.hasLeapSecond
	timeTransDto2.nanosecond = tempTimeTransDto.nanosecond
	timeTransDto2.totalTimeNanoseconds = tempTimeTransDto.totalTimeNanoseconds
	timeTransDto2.timeZone = tempTimeTransDto.timeZone.CopyOut()
	timeTransDto2.tag = tempTimeTransDto.tag

	_,
		err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDto1,
		ePrefix + "Testing Validity of output 'timeTransDto1'. ")

	if err != nil {
		return err
	}

	_,
		err = timeTransDtoNanobot.testTimeTransferDtoValidity(
		timeTransDto2,
		ePrefix + "Testing Validity of output 'timeTransDto2'. ")

	return err
}

// setInstanceToZero - This method receives a pointer to a TimeTransferDto
// instance and proceeds to set all of the member variable data values to
// their native 'zero' state. When this operation is completed some, but not
// all, of the member variables will contain 'invalid' values. If after
// calling this method, the 'timeTransDto' instance is submitted for validity
// testing, it will fail those tests.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  timeTransDto        *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This method will
//       set all of the internal member variables to their native
//       'zero' states. Some, but not all. of the member variables will
//       be set to invalid values.
//
//       After calling this method, submitting this TimeTransferDto instance
//       for validity testing will result in failure.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//       --- NONE ---
//
func (timeTransDtoMech *timeTransferDtoMechanics) setInstanceToZero(
	timeTransDto *TimeTransferDto) {

	if timeTransDtoMech.lock == nil {
		timeTransDtoMech.lock = new(sync.Mutex)
	}

	timeTransDtoMech.lock.Lock()

	defer timeTransDtoMech.lock.Unlock()

	if timeTransDto == nil {
		panic("timeTransferDtoMechanics.setInstanceToZero()\n" +
			"Input parameter 'timeTransDto' is a nil pointer.")
	}

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	timeTransDto.hour = 0
	timeTransDto.minute = 0
	timeTransDto.second = 0
	timeTransDto.hasLeapSecond = false
	timeTransDto.nanosecond = 0
	timeTransDto.totalTimeNanoseconds = 0
	timeTransDto.timeZone, _ = TimeZoneDefinition{}.New()
	timeTransDto.tag = ""

	return
}
