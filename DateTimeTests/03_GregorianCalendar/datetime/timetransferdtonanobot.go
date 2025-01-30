package datetime

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type timeTransferNanobot struct {

	lock *sync.Mutex

}


// testTimeTransferDtoValidity - Tests and evaluates the validity of a
// TimeTransferDto instance. If the instance is valid this method sets
// return parameters 'isValid' = 'true' and 'err' = nil.
//
// If the instance is invalid, the method returns 'isValid' = 'false'
// plus an error message.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  timeTransDto        *TimeTransferDto
//     - A pointer to an instance of TimeTransferDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. Member variables will only be tested
//       for validity.
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
//  isValid             bool
//     - If the method completes successfully, this boolean value will
//       signal 'false' if any 'timeTransDto' member variables contain
//       an invalid value. Absent processing errors, if the validity
//       tests show that 'timeTransDto' member variables values are valid,
//       this 'isValid' flag will be set to 'true'.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//       An error will also be returned if the method determines that
//       one or more of the 'timeTransDto' member variables contain
//       invalid values.
//
func (timeTransDtoNanobot *timeTransferNanobot) testTimeTransferDtoValidity(
	timeTransDto *TimeTransferDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if timeTransDtoNanobot.lock == nil {
		timeTransDtoNanobot.lock = new(sync.Mutex)
	}

	timeTransDtoNanobot.lock.Lock()

	defer timeTransDtoNanobot.lock.Unlock()

	ePrefix += "timeTransferNanobot.testTimeTransferDtoValidity() "

	isValid = false
	err = nil

	if timeTransDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'timeTransDto' is a 'nil' pointer!\n")

		return isValid,err
	}

	if timeTransDto.lock == nil {
		timeTransDto.lock = new(sync.Mutex)
	}

	if timeTransDto.hour > 24 ||
		timeTransDto.hour < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'timeTransDto.hour' is INVALID!\n" +
			"The 'timeTransDto.hour' value must be <= 24 and >= 0\n" +
			"timeTransDto.hour='%v'\n", timeTransDto.hour)

		return isValid, err
	}

	if timeTransDto.minute > 59 ||
		timeTransDto.minute < 0 {

		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'timeTransDto.minute' is INVALID!\n" +
			"The 'timeTransDto.minute' value must be <= 59 and >= 0\n" +
			"timeTransDto.minute='%v'\n", timeTransDto.minute)

		return isValid, err
	}

	maxSecond := 59

	if timeTransDto.hasLeapSecond {
		maxSecond = 60
	}

	if timeTransDto.second < 0 ||
		timeTransDto.second > maxSecond {

		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'timeTransDto.second' is INVALID!\n" +
			"The 'timeTransDto.second' value must be <= %v and >= 0\n" +
			"timeTransDto.second='%v'\n",
			maxSecond,
			timeTransDto.second)

		return isValid, err
	}

	if timeTransDto.nanosecond < 0 ||
		timeTransDto.nanosecond > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'timeTransDto.nanosecond' is INVALID!\n" +
			"The 'timeTransDto.nanosecond' value must be <= 999,999,999 and >= 0\n" +
			"timeTransDto.nanosecond='%v'\n",
			timeTransDto.nanosecond)

		return isValid, err
	}

	err = timeTransDto.timeZone.IsValidInstanceError(
		ePrefix + "Testing 'timeTransDto.timeZone' Validity ")

	if err != nil {
		return isValid, err
	}

	timeTransDto.totalTimeNanoseconds = 0

	timeTransDto.totalTimeNanoseconds += int64(time.Hour) * int64(timeTransDto.hour)
	timeTransDto.totalTimeNanoseconds += int64(time.Minute) * int64(timeTransDto.minute)
	timeTransDto.totalTimeNanoseconds += int64(time.Second) * int64(timeTransDto.second)
	timeTransDto.totalTimeNanoseconds += int64(timeTransDto.nanosecond)

	isValid = true

	return isValid, err
}

