package main

import (
	"errors"
	"fmt"
)

// This example reuses the same variable names. Because
// the names are associated with different scopes, there
// is no conflict.

func main() {

	if foo, err := doSomething1(); err != nil {
		// scope of foo and err limited to
		// if block
		fmt.Println("doError1 - has an error")
		fmt.Println("Value of foo is ", foo)
		return
	}

	// scope of foo and err is now
	// at function block level
	foo, err := doSomething2()

	if err != nil {
		fmt.Println("doSomething2 error=", err)
		fmt.Println("foo = ", foo)
	}

}

func doSomething1() (interface{}, error) {
	return "some string", nil
}

func doSomething2() (interface{}, error) {
	return nil, errors.New("Something2 error")
}

/*	Output
	$ go run main.go
	doSomething2 error= Something2 error
	foo =  <nil>

*/
