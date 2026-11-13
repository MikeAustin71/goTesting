package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
	Wrapped error
}

// FIX 1: Change %w to %v here. Sprintf does not recognize %w.
func (e CustomError) Error() string {
	if e.Wrapped != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Wrapped)
	}
	return e.Message
}

func (e CustomError) Unwrap() error {
	return e.Wrapped
}

func main() {
	errTest02()
}

// errTest02
//
// Output:
//
//	Full Chain string:
//
//	funcLevel01: Error returned by funcLevel02:
//	funcLevel02: Error returned by funcLevel03:
//	funcLevel03: Error returned by funcLevel04:
//	funcLevel04: Serious Catastrophe
//	Call Fire Department!!
//	: error occurred in funcLevel04
//	----------------------------------------
//	Innermost Error: error occurred in funcLevel04
//
//	Process finished with the exit code 0
func errTest02() {
	err := funcLevel01()

	if err != nil {
		// Simply printing 'err' will now traverse the whole chain automatically
		fmt.Println("Full Chain string:")
		fmt.Println(err)
		fmt.Println("----------------------------------------")

		// Loop down to find the absolute innermost error safely
		currentError := err
		for {
			unwrapped := errors.Unwrap(currentError)
			if unwrapped == nil {
				break
			}
			currentError = unwrapped
		}

		fmt.Println("Innermost Error:", currentError)
		return
	}

	fmt.Println("No Errors Detected!")
}

func funcLevel01() error {
	err := funcLevel02()

	if err != nil {
		// FIX 2: Pass 'err' into Wrapped so the chain remains linked
		return CustomError{
			Message: "\nfuncLevel01: Error returned by funcLevel02",
			Wrapped: err,
		}
	}

	return nil
}

func funcLevel02() error {
	err := funcLevel03()

	if err != nil {
		return CustomError{
			Message: "\nfuncLevel02: Error returned by funcLevel03",
			Wrapped: err,
		}
	}

	return nil
}

func funcLevel03() error {
	err := funcLevel04()

	if err != nil {
		return CustomError{
			Message: "\nfuncLevel03: Error returned by funcLevel04",
			Wrapped: err,
		}
	}

	return nil
}

func funcLevel04() error {
	// FIX 3: Since Level 4 is the bottom of this execution chain,
	// we wrap a clean standard error as the absolute root cause.
	return CustomError{
		Message: "\nfuncLevel04: Serious Catastrophe\nCall Fire Department!!\n",
		Wrapped: errors.New("error occurred in funcLevel04"),
	}
}
