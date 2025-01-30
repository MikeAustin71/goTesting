package common

type NumTree struct {
	Num  int
}

func (nTree NumTree) Add(n1, n2 int) int {

	return n1 + n2
}

func (nTree NumTree) New() NumTree {

	return NumTree {}
}

func (nTree *NumTree) AddToThis(n int) {
	nTree.Num += n
}

