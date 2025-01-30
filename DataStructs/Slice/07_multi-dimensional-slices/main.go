package main

import (
	"fmt"
)

func main()  {

	d:= make([][][]string,0)

	d =  append(d, [][]string{{"Z","XQZ"}})

	d =  append(d, [][]string{{"B","ZQX"}})

	d =  append(d, [][]string{{"C","CFJ"}})

	for i, a := range d {
		fmt.Println("i= ", i, "  a[0][0]= ", a[0][0], "  a[0][1]", a[0][1])
	}

	fmt.Println("Length of d: ", len(d))

	for i := 0; i < len(d); i++ {
		fmt.Println("d[i][0][0] ", d[i][0][0], " d[i][0][1]", d[i][0][1] )
	}

}

/* This works
var x [2][2]string

x[0][0] = "0 PM"
x[0][1] = "0 pm"

x[1][0] = "0 PM"
x[1][1] = "0 pm"

fmt.Println("x[0]", x[0])
*/


func Test1() {
	// searchStrs["0 PM"] = "0 pm"


	/* This fails as runtime error
	d = append(d[0][0],"0 PM" )

	d[0][0] = "0 PM"
	d[0][1] = "0 pm"
	*/

	d := make([][]string,2)

	d[0] = make([]string, 2)

	d[0][0] = "0 PM"
	d[0][1] =	"0 pm"

	d[1] =  make([]string,2)
	d[1][0] = "0 PM"
	d[1][1] =	"0 pm"

	var result string

	for i:=0; i < 2; i ++ {
		for j:=0; j < 2; j++ {
			result = d[i][j]
			fmt.Printf("d[%v][%v]= %v \n",i,j,result)
		}
	}


}