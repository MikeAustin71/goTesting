package main

import "fmt"

type IntPointers struct {
	Int1 *int
	Int2 *int
}

type FloatPointers struct {
	Float1 *float64
	Float2 *float64
}

type NumberPointers struct {
	IntPtrs   IntPointers
	FloatPtrs FloatPointers
}

func main() {
	BaseRun02()
}

func BaseRun02() {
	origIntA := 1
	origIntB := 20

	origFloatA := 62.14
	origFloatB := 97.21838

	nmPtr1 := &IntPointers{Int1: &origIntA, Int2: &origIntB}

	floatPtr1 := &FloatPointers{Float1: &origFloatA, Float2: &origFloatB}

	numPointers := NumberPointers{IntPtrs: *nmPtr1, FloatPtrs: *floatPtr1}

	fmt.Printf("\n\n          BaseRun02\n")
	fmt.Printf("-------------------------------\n")
	fmt.Printf("Original Before Modification:\n")

	fmt.Printf("numPointers.IntPtrs.Int1-Pointer: %v  numPointers.IntPtrs.Int1-Value %v\n",
		numPointers.IntPtrs.Int1, *numPointers.IntPtrs.Int1)

	fmt.Printf("numPointers.IntPtrs.Int2-Pointer: %v  numPointers.IntPtrs.Int2-Value %v\n",
		numPointers.IntPtrs.Int2, *numPointers.IntPtrs.Int2)

	fmt.Printf("numPointers.FloatPtrs.Float1-Pointer: %v  numPointers.FloatPtrs.Float1-Value %v\n",
		numPointers.FloatPtrs.Float1, *numPointers.FloatPtrs.Float1)

	fmt.Printf("numPointers.FloatPtrs.Float2-Pointer: %v  numPointers.FloatPtrs.Float2-Value %v\n",
		numPointers.FloatPtrs.Float2, *numPointers.FloatPtrs.Float2)

	Caller03(&numPointers)

	fmt.Printf("\n\n         BaseRun02\n")
	fmt.Printf("-------------------------------\n")
	fmt.Printf("\nBaseRun02-Recap - After Modification:\n")

	fmt.Printf("numPointers.IntPtrs.Int1-Pointer: %v  numPointers.IntPtrs.Int1-Value %v\n",
		numPointers.IntPtrs.Int1, *numPointers.IntPtrs.Int1)

	fmt.Printf("numPointers.IntPtrs.Int2-Pointer: %v  numPointers.IntPtrs.Int2-Value %v\n",
		numPointers.IntPtrs.Int2, *numPointers.IntPtrs.Int2)

	fmt.Printf("numPointers.FloatPtrs.Float1-Pointer: %v  numPointers.FloatPtrs.Float1-Value %v\n",
		numPointers.FloatPtrs.Float1, *numPointers.FloatPtrs.Float1)

	fmt.Printf("numPointers.FloatPtrs.Float2-Pointer: %v  numPointers.FloatPtrs.Float2-Value %v\n",
		numPointers.FloatPtrs.Float2, *numPointers.FloatPtrs.Float2)

}

func Caller03(nPtrs *NumberPointers) {
	fmt.Printf("\n-------------------------------\n")
	fmt.Printf("    Running Caller 03\n")
	fmt.Printf("-------------------------------\n")
	fmt.Printf("Caller03-Original Before Modification:\n")

	fmt.Printf("nPtrs.IntPtrs.Int1-Pointer: %v  nPtrs.IntPtrs.Int1-Value %v\n",
		nPtrs.IntPtrs.Int1, *nPtrs.IntPtrs.Int1)

	fmt.Printf("nPtrs.IntPtrs.Int2-Pointer: %v  nPtrs.IntPtrs.Int2-Value %v\n",
		nPtrs.IntPtrs.Int2, *nPtrs.IntPtrs.Int2)

	fmt.Printf("nPtrs.FloatPtrs.Float1-Pointer: %v  nPtrs.FloatPtrs.Float1-Value %v\n",
		nPtrs.FloatPtrs.Float1, *nPtrs.FloatPtrs.Float1)

	fmt.Printf("nPtrs.FloatPtrs.Float2-Pointer: %v  nPtrs.FloatPtrs.Float2-Value %v\n",
		nPtrs.FloatPtrs.Float2, *nPtrs.FloatPtrs.Float2)

	testPtrs10(nPtrs)

	fmt.Printf("\nCaller03-Recap - After Modification\n")

	fmt.Printf("nPtrs.IntPtrs.Int1-Pointer: %v  nPtrs.IntPtrs.Int1-Value %v\n",
		nPtrs.IntPtrs.Int1, *nPtrs.IntPtrs.Int1)

	fmt.Printf("nPtrs.IntPtrs.Int2-Pointer: %v  nPtrs.IntPtrs.Int2-Value %v\n",
		nPtrs.IntPtrs.Int2, *nPtrs.IntPtrs.Int2)

	fmt.Printf("nPtrs.FloatPtrs.Float1-Pointer: %v  nPtrs.FloatPtrs.Float1-Value %v\n",
		nPtrs.FloatPtrs.Float1, *nPtrs.FloatPtrs.Float1)

	fmt.Printf("nPtrs.FloatPtrs.Float2-Pointer: %v  nPtrs.FloatPtrs.Float2-Value %v\n",
		nPtrs.FloatPtrs.Float2, *nPtrs.FloatPtrs.Float2)

}

func testPtrs10(xNumPtrs *NumberPointers) {

	fmt.Printf("\n-------------------------------\n")
	fmt.Printf("           testPtrs10\n")
	fmt.Printf("Initial Values Before Modification:\n")

	fmt.Printf("xNumPtrs.IntPtrs.Int1-Pointer: %v  xNumPtrs.IntPtrs.Int1-Value %v\n",
		xNumPtrs.IntPtrs.Int1, *xNumPtrs.IntPtrs.Int1)

	fmt.Printf("xNumPtrs.IntPtrs.Int2-Pointer: %v  xNumPtrs.IntPtrs.Int2-Value %v\n",
		xNumPtrs.IntPtrs.Int2, *xNumPtrs.IntPtrs.Int2)

	fmt.Printf("xNumPtrs.FloatPtrs.Float1-Pointer: %v  xNumPtrs.FloatPtrs.Float1-Value %v\n",
		xNumPtrs.FloatPtrs.Float1, *xNumPtrs.FloatPtrs.Float1)

	fmt.Printf("xNumPtrs.FloatPtrs.Float2-Pointer: %v  xNumPtrs.FloatPtrs.Float2-Value %v\n",
		xNumPtrs.FloatPtrs.Float2, *xNumPtrs.FloatPtrs.Float2)

	xStr := "\n\ntestPtrs10 Modifications\n"
	xStr += "*xNumPtrs.IntPtrs.Int1 += 2\n"
	xStr += "*xNumPtrs.IntPtrs.Int2 += 10\n"
	xStr += "*xNumPtrs.FloatPtrs.Float1 += 2.0\n"
	xStr += "*xNumPtrs.FloatPtrs.Float2 += 20.0\n\n"

	*xNumPtrs.IntPtrs.Int1 += 2
	*xNumPtrs.IntPtrs.Int2 += 10
	*xNumPtrs.FloatPtrs.Float1 += 2.0
	*xNumPtrs.FloatPtrs.Float2 += 20.0

	fmt.Printf("%v", xStr)

	fmt.Printf("\ntestPtrs10 Recap - After Modification\n")

	fmt.Printf("xNumPtrs.IntPtrs.Int1-Pointer: %v  xNumPtrs.IntPtrs.Int1-Value %v\n",
		xNumPtrs.IntPtrs.Int1, *xNumPtrs.IntPtrs.Int1)

	fmt.Printf("xNumPtrs.IntPtrs.Int2-Pointer: %v  xNumPtrs.IntPtrs.Int2-Value %v\n",
		xNumPtrs.IntPtrs.Int2, *xNumPtrs.IntPtrs.Int2)

	fmt.Printf("xNumPtrs.FloatPtrs.Float1-Pointer: %v  xNumPtrs.FloatPtrs.Float1-Value %v\n",
		xNumPtrs.FloatPtrs.Float1, *xNumPtrs.FloatPtrs.Float1)

	fmt.Printf("xNumPtrs.FloatPtrs.Float2-Pointer: %v  xNumPtrs.FloatPtrs.Float2-Value %v\n",
		xNumPtrs.FloatPtrs.Float2, *xNumPtrs.FloatPtrs.Float2)
}

func BaseRun01() {
	origA := 1
	origB := 20
	fmt.Printf("\n\n          BaseRun01\n")
	fmt.Printf("-------------------------------\n")
	fmt.Printf("Original Before Modification:\n")
	fmt.Printf("origA-Pointer: %v  origA-Value %v\n", &origA, origA)
	fmt.Printf("origB-Pointer: %v  origB-Value %v\n", &origB, origB)

	Caller02(&origA, &origB)

	fmt.Printf("\n            BaseRun01\n")
	fmt.Printf("-------------------------------\n")
	fmt.Printf("Recap - After Modification\n")
	fmt.Printf("origA-Pointer: %v  origA-Value %v\n", &origA, origA)
	fmt.Printf("iPtr1-Pointer: %v  iPtr1-Value %v\n", &origB, origB)

}

func Caller02(iPtr1 *int, iPtr2 *int) {
	fmt.Printf("\n-------------------------------\n")
	fmt.Printf("    Running Caller 02\n")
	fmt.Printf("-------------------------------\n")

	var mix1 *int
	var mix2 *int

	mix1 = iPtr1
	mix2 = iPtr2

	test01(mix1, mix2)
}

func Caller01(iPtr1 *int, iPtr2 *int) {
	fmt.Printf("\n-------------------------------\n")
	fmt.Printf("    Running Caller 01\n")
	fmt.Printf("-------------------------------\n")

	test01(iPtr1, iPtr2)
}

func test01(iPtr1 *int, iPtr2 *int) {

	fmt.Printf("\n-------------------------------\n")
	fmt.Printf("              test01\n")
	fmt.Printf("Initial Values:\n")
	fmt.Printf("iPtr1-Pointer: %v  iPtr1-Value %v\n", iPtr1, *iPtr1)
	fmt.Printf("iPtr2-Pointer: %v  iPtr2-Value %v\n", iPtr2, *iPtr2)
	fmt.Printf("\nModified Plus 1 Values:\n")
	*iPtr1++
	*iPtr2++
	fmt.Printf("iPtr1-Pointer: %v  iPtr1-Value %v\n", iPtr1, *iPtr1)
	fmt.Printf("iPtr2-Pointer: %v  iPtr2-Value %v\n", iPtr2, *iPtr2)

}
