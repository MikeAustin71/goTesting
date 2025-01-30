package main

import (
	"fmt"
	"unsafe"
)

type vertex struct {
	X int64
	Y int64
}

func main() {
	var x uint16 = 5
	fmt.Println("unsafe.Sizeof reports size of uint x =", unsafe.Sizeof(x), "\n")


	a := float64(0)
	fmt.Println("unsafe.Sizeof reports size of float64 a =", unsafe.Sizeof(a),"\n")


	y := "This a string of 30 characters"
	fmt.Println("len reports size of y =", len(y))
	fmt.Println("unsafe.Sizeof reports size of string y =", unsafe.Sizeof(y),"\n")

	v := vertex{7, 9}
	fmt.Println("unsafe.Sizeof reports size of vertex v =", unsafe.Sizeof(v))
}
