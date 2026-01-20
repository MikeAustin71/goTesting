package main

/*  THIS CODE FAILS!!!!
This example is designed to demonstrate Unexported Identifiers; that is,
those identifiers which are invisible to callers outside the package
in which the unexported identifier was declared.

Exported Identifiers must have the first letter of their names capitalized.

'invisibleReverseAString' is a function declared in another Package
named 'strUtilTest'. 'invisibleReverseAString' is visible to the main
function in the main package which is attempting to call this function
because the first letter of its name is lower case. Consequently, it is
invisible to all callers outside its native package, strUtilTest, in
which it was declared.

In order to Export an identifier, the first letter of the Identifier's
name MUST be capitalized. In this case, the first letter of the Identifier's
name is lower case, and it is invisible to all callers outside its native
package.
*/

import (
	"fmt"

	"github.com/mikeaustin71/Identifiers/ExportedIdentifiers/strUtilTest"
)

func main() {
	v := "Hello"
	rv := strUtilTest.InvisibleReverseAString(v)
	fmt.Println("This is base string", v)
	fmt.Println("This is the reversed string", rv)
}

/*	Output
	$ go run main.go
	# command-line-arguments
	.\main.go:30: cannot refer to unexported name strUtilTest.invisibleReverseAString
	.\main.go:30: undefined: strUtilTest.invisibleReverseAString
*/
