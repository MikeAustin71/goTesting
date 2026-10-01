package main

import (
	"fmt"
	"math/big"

	"github.com/mikeaustin71/Math/bFloatNatLog/02NatLog"
)

func main() {

	m := big.NewFloat(1.5)
	prec := uint(2048)

	ctx := naturalLogCalcs.NewLnAGMContext(m, prec)
	_, err := ctx.LnAGMMantissa()
	if err != nil {
		fmt.Printf("LnAGMMantissa failed: %v", err)
	}

	fmt.Printf("main.go executed without error!")

}
