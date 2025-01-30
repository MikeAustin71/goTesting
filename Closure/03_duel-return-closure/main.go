/*	Closure Example
	In the example below, g1 and g2 get separate copies
	of 'cntr'.
*/

package main

import "fmt"

type addItUp func() int

func genFunc(f int) addItUp {

	cntr:= 0

	if f == 1 {
		return func() int {
			cntr++
			return cntr
		}
	}

	return func() int {
		cntr+=4
		return cntr
	}
}

func main() {
	g1 := genFunc(1)
	g2 := genFunc(2)

	fmt.Println(g1())	// 1
	fmt.Println(g1())	// 2
	fmt.Println(g2())	// 4
	fmt.Println(g2())	// 8
	fmt.Println(g1())	// 3
	fmt.Println(g2())   // 12
}

/*	Output
	$ go run main.go
	1
	2
	4
	8
	3
	12
 */