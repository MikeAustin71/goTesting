package employee

import (
	"fmt"
	//"github.com/mikeaustin71/PackageStruct02/pkgStr/customErr"
)

type employeeHelper struct{}

func (h *employeeHelper) SetTheEmployee(employee *Employee, name string, age int) error {

	ePrefix := "employeeHelper.SetTheEmployee"

	if employee == nil {
		return &myStdErr{
			ErrPrefix:  ePrefix,
			ErrMessage: "Error: Input parameter 'employee' is a nil pointer!",
		}
	}

	if len(name) == 0 {
		return &myStdErr{
			ErrPrefix:  ePrefix,
			ErrMessage: "Error: Input parameter 'str' is empty!",
		}
	}

	if age < 1 {
		return &myStdErr{
			ErrPrefix:  ePrefix,
			ErrMessage: "Error: Input parameter 'age' must be greater than '0'!",
		}
	}

	employee.Name = name
	employee.Age = age

	return nil
}

func (h *employeeHelper) PrintEmployeeData(
	employee *Employee) error {

	if employee == nil {
		return &myStdErr{
			ErrPrefix:  "employeeHelper.PrintEmployeeData",
			ErrMessage: "Error: Input parameter 'employee' is a nil pointer!",
		}
	}

	fmt.Printf("\nEmployee Info\n%v", employee)

	return nil
}
