package main

import (
	"fmt"
)

type ErrorReturnBasic struct {
	ErrPrefix  string
	ReturnFunc string
	ErrContext string
	ErrMessage string
}

func (e *ErrorReturnBasic) Error() string {

	var errStr string
	foundCnt := 0

	if e.ErrPrefix != "" {
		errStr = e.ErrPrefix + "\n"
		foundCnt++
	}

	if e.ReturnFunc != "" {
		errStr += e.ReturnFunc + "\n"
		foundCnt++
	}

	if e.ErrContext != "" {
		errStr += e.ErrContext + "\n"
		foundCnt++
	}

	if e.ErrMessage != "" {
		errStr += "Error:" + e.ErrMessage + "\n"
		foundCnt++
	}

	if foundCnt == 0 {
		errStr = "No Error parameters provided!\n"
	}

	return errStr
}

func main() {

	var err error

	_, err = test01("main")

	if err != nil {

		fmt.Printf("%s", err.Error())
	}

	_, err = test02("main")

	if err != nil {

		fmt.Printf("%s", err.Error())

	}

	_, err = test03("main")

	if err != nil {

		fmt.Printf("%s", err.Error())
	}

	_, err = test04("main")

	if err != nil {

		fmt.Printf("%s", err.Error())
	}

	_, err = test05("main")

	if err != nil {

		fmt.Printf("%s", err.Error())
	}

	return

}

func test01(errPrefix string) (int, error) {

	var err error

	err = fmt.Errorf("test error")

	if err != nil {

		return 0,
			&ErrorReturnBasic{
				ErrPrefix:  errPrefix,
				ReturnFunc: "test01",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	} else {
		fmt.Println("Test passed! - No Error Detected.")
	}

	return 0, nil
}

func test02(errPrefix string) (int, error) {

	var err error

	err = fmt.Errorf("unusual test error #2")

	if err != nil {

		return 0,
			&ErrorReturnBasic{
				ErrPrefix:  errPrefix,
				ReturnFunc: "test02",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	} else {
		fmt.Println("Test passed! - No Error Detected.")
	}

	return 0, nil
}

func test03(errPrefix string) (int, error) {

	var err error

	err = fmt.Errorf("wierd test error #3")

	var errBasic = new(ErrorReturnBasic)

	if err != nil {
		errBasic.ErrPrefix = errPrefix
		errBasic.ReturnFunc = "test03"
		errBasic.ErrContext = ""
		errBasic.ErrMessage = err.Error()
		return 0, errBasic

	} else {
		fmt.Println("Test passed! - No Error Detected.")
	}

	return 0, nil
}

func test04(errPrefix string) (int, error) {

	var err error

	err = fmt.Errorf("wierd test error #4")

	if err != nil {

		var errBasic = new(ErrorReturnBasic)

		errBasic.ErrPrefix = errPrefix
		errBasic.ReturnFunc = "test04"
		errBasic.ErrContext = "Context Message test04"
		errBasic.ErrMessage = err.Error()

		return 0, errBasic

	} else {
		fmt.Println("Test # 4 passed! - No Error Detected.")

	}

	return 0, err
}

func test05(errPrefix string) (int, error) {

	var err error

	err = fmt.Errorf("fatal test error #5")

	if err != nil {

		return 0,
			&ErrorReturnBasic{
				ErrPrefix:  errPrefix,
				ReturnFunc: "test05",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	} else {
		fmt.Println("Test passed! - No Error Detected.")
	}

	return 0, nil
}
