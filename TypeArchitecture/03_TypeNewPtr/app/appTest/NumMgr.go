package appTest

import "fmt"

type NumMgr struct {
	Num  int
}

func (nMgr NumMgr) New() NumMgr {

	newNumMgr := NumMgr{}

	return newNumMgr
}

func (nMgr NumMgr) NewPtr() *NumMgr {

	return new(NumMgr)
}

func (nMgr NumMgr) This() *NumMgr {
	return &nMgr
}

func (nMgr *NumMgr) Empty() {

	nMgr.Num = -1

}

func (nMgr *NumMgr) SetNum(n int)  {

	nMgr.Num = n

	return
}

func (nMgr *NumMgr) MulNum(n int) (*NumMgr, error){
	nMgr.Num = nMgr.Num * n

	return nMgr, nil
}

func (nMgr NumMgr) AddNum(num1, num2 int) {

	nMgr.SetNum(num1+num2)

	fmt.Println("Internal num1+num2 Value: ", nMgr.Num)
}

func (nMgr *NumMgr) AddNums(num1, num2 int) int {
	return num1 + num2
}
