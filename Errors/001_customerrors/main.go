package main

import (
	"errors"
	"fmt"
)

type TzError struct {
	ePrefix string
	errMsg  string
	err     error
}

func (e *TzError) Error() string {
	return fmt.Sprintf(e.ePrefix+
		"\n%v\n", e.errMsg)
}

func (e *TzError) Unwrap() error {
	return e.err
}

func (e *TzError) Is(target error) bool {

	_, ok := target.(*TzError)

	if !ok {
		return false
	}

	return true
}

type DirError struct {
	ePrefix string
	errMsg  string
	err     error
}

func (e *DirError) Error() string {
	return fmt.Sprintf(e.ePrefix+
		"\n%v\n", e.errMsg)
}

func (e *DirError) Unwrap() error {
	return e.err
}

func (e *DirError) Is(target error) bool {

	_, ok := target.(*DirError)

	if !ok {
		return false
	}

	return true
}

type FileError struct {
	ePrefix string
	errMsg  string
	err     error
}

func (e *FileError) Error() string {
	return fmt.Sprintf(e.ePrefix+
		"\n%v\n", e.errMsg)
}

func (e *FileError) Unwrap() error {
	return e.err
}

func (e *FileError) Is(target error) bool {

	var fileError *FileError

	ok := errors.As(target, &fileError)

	if !ok {
		return false
	}

	return true
}

type TzAbbrvError struct {
	ePrefix string
	errMsg  string
	err     error
}

func (e *TzAbbrvError) Error() string {
	return fmt.Sprintf(e.ePrefix+
		"\n%v\n", e.errMsg)
}

func (e *TzAbbrvError) Unwrap() error {
	return e.err
}

func (e *TzAbbrvError) Is(target error) bool {

	_, ok := target.(*TzAbbrvError)

	if !ok {
		return false
	}

	return true
}

func (e *TzAbbrvError) As(iErr interface{}) bool {

	var t = iErr.(TzAbbrvError)

	t.ePrefix = e.ePrefix
	t.errMsg = e.errMsg
	t.err = fmt.Errorf("%w", e.err)

	return true
}

func main() {
	test02("main() ")
}

func test01() {

	err := tzAbbreviationLookup("main() ")

	if err != nil {
		fmt.Println(err)
	}

	if errors.Is(err, &TzAbbrvError{}) {
		fmt.Println("'Is' Result - This error type 'TzAbbrvError'\n")
	}

	currentError := err

	for currentError != nil {
		fmt.Println(currentError)
		currentError = errors.Unwrap(currentError)
	}

}

func test02(ePrefix string) {

	ePrefix += "test02() "
	err := tzAbbrvLookup02(ePrefix)

	if err == nil {
		fmt.Println("No error returned from ", ePrefix)
		return
	}

	var tzAbbrvErr = &TzAbbrvError{}

	if errors.As(err, &tzAbbrvErr) {

		fmt.Println("test02() Success!!")

		fmt.Println(tzAbbrvErr)

	} else {
		fmt.Println("test02() **Failed**")
	}

}

func tzAbbrvLookup02(ePrefix string) error {

	ePrefix += "tzAbbrvLookup02() "

	return &TzAbbrvError{
		ePrefix: ePrefix,
		errMsg:  "Error Message From 2nd Time Zone Abbrv Lookup!",
		err:     nil,
	}
}

func tzAbbreviationLookup(ePrefix string) error {

	ePrefix += "tzAbbreviationLookup() "

	err := level2FileError(ePrefix)

	return &TzAbbrvError{
		ePrefix,
		"Err Msg from Level #1",
		err}
}

func level2FileError(ePrefix string) error {

	ePrefix += "level2FileError() "

	err := level3DirError(ePrefix)

	return &FileError{
		ePrefix: ePrefix,
		errMsg:  "This is a File Error from Level #2",
		err:     err,
	}
}

func level3DirError(ePrefix string) error {

	ePrefix += "level3DirError() "

	err := level4TzError(ePrefix)

	return &DirError{
		ePrefix: ePrefix,
		errMsg:  "This is a Directory Error from Level #3",
		err:     err,
	}
}

func level4TzError(ePrefix string) error {

	ePrefix += "level4TzError() "

	return &TzError{
		ePrefix: ePrefix,
		errMsg:  "This is a time zone error from Level #4",
		err:     nil,
	}

}
