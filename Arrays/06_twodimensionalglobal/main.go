package main


import "fmt"

var y0 = [2]int{0,10}
var y1 = [2]int{1,11}
var y2 = [2]int{2,12}
var y3 = [2]int{3,13}
var y4 = [2]int{4,14}
var y5 = [2]int{5,15}
var y6 = [2]int{6,16}

var a = [7][2]int{y0,y1,y2,y3,y4,y5,y6}

func main() {

	for k:=0; k < len(a); k++ {
		fmt.Println("k= ", a[k][0], "  2nd-Dimension= ", a[k][1])
	}

}
