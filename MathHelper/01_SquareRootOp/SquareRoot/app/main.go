package main

import (
	"fmt"
	"golangmikesamples/MathHelper/01_SquareRootOp/SquareRoot/common"
)

func main() {
	// 390626 == good
	// 2685   == good 51.8169856321264600826942986293190493782400
	//        400 good result
	//     40,000 good result
	//  4,000,000 good result
	//400,000,000 good result
	numAry := common.IntAry{}
	_ = numAry.SetIntAryWithNumStr("3906267")

	sqRoot := common.SquareRootOp{}
	_ = sqRoot.Initialize(&numAry, 32)
	_ = sqRoot.ComputeSquareRoot()

	fmt.Println("Square Root Result")
	fmt.Println("Input String:")
	fmt.Println(numAry.NumStr)
	fmt.Println()
	fmt.Println("Input Array:")
	fmt.Println(numAry.IntAry)
	fmt.Println()
	fmt.Println("Base Pairs Array")
	fmt.Println(sqRoot.BaseNumPairs)

	fmt.Println()
	fmt.Println("Precision: ", sqRoot.ResultPrecision)
	fmt.Println()
	fmt.Println("ResultAry:")
	fmt.Println(sqRoot.ResultAry.IntAry)
	fmt.Println()
	fmt.Println("ResultAry NumStr:")
	fmt.Println(sqRoot.ResultAry.NumStr)
	fmt.Println()
	fmt.Println("ResultAry Precision:", sqRoot.ResultAry.Precision)

}
