package employee

import (
	"fmt"
	custErr "golangmikesamples/PackageStruct02/pkgStr/customErr"
)

type myStdErr = custErr.StdBasicError

type Employee struct {
	Name string
	Age  int
}

func (e *Employee) SetEmployee(name string, age int) error {

	ePrefix := "Employee.SetEmployee"

	err := new(employeeHelper).SetTheEmployee(
		e, name, age)

	if err != nil {
		return &myStdErr{
			ErrPrefix:  ePrefix,
			ErrMessage: err.Error(),
		}
	}

	return nil
}

func (e *Employee) PrintEmployeeInfo() error {
	ePrefix := "Employee.PrintEmployeeInfo"

	err := new(employeeHelper).PrintEmployeeData(e)

	if err != nil {
		return &myStdErr{
			ErrPrefix:  ePrefix,
			ErrMessage: err.Error(),
		}
	}

	return nil
}

func (e *Employee) String() string {
	return fmt.Sprintf("Employee Name: %s:\nEmployee Age: %d\n", e.Name, e.Age)
}
