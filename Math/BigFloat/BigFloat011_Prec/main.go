package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

// Log2Of10String stores the constant as a string literal.
// We use compile-time string concatenation (+) so the code remains readable
// without introducing unwanted newlines or whitespace.
const Log2Of10String = "3." +
	"32192809488736234787031942948939017586483139302458061205475639581593477660862521" +
	"58501397433593701550996573717102502518268240969842635268882753027729986553938519" +
	"51352657505568643017609190024891666941433374011903124187375109715866467540179189" +
	"65580673583077968843272588327499252244890238355997641739413792800977275668635547" +
	"79014867450578458847802710422545609722346579569554153701915764117177924716513500" +
	"23921127147339361440723397211574851007094987891658880831322194806793298232325931" +
	"19506713995078370033673424807066352750084069176263862535468801536862161841886085" +
	"89948353813214998930270441792078659226018229653715753672396606951164868368466238" +
	"58508486062990542699469279116273206134006446704847634070437352336742212830896703" +
	"64579092167721909021421962142457444658524535948448815483459251429540937353906549" +
	"44863277929842429159118113116329812576945019815750379218553848782035516019737827" +
	"72888817598743328660727123938252022133328052551248827434448842453165465061241489"

// NewLog2Of10 creates and returns a *big.Float containing log2(10) configured
// with the exact bit precision necessary to represent every digit without loss.
func NewLog2Of10() (*big.Float, error) {
	// 1. Determine number of significant decimal digits
	//    (Total characters minus the decimal point)
	decDigits := len(strings.ReplaceAll(Log2Of10String, ".", ""))

	// 2. Compute minimum required bits: ceil(decDigits * log2(10)) + guard bits
	//    log2(10) ≈ 3.321928094887362
	bitsNeeded := math.Ceil(float64(decDigits) * 3.321928094887362)

	// Round up to the next multiple of 64 bits (word boundary) + 64 guard bits
	// For 1,000 digits: ceil(3321.9) = 3322 bits -> sets precision to 3392 bits.
	prec := uint((math.Ceil(bitsNeeded/64.0) * 64) + 64)

	// 3. Initialize and parse into big.Float
	target := new(big.Float).
		SetPrec(prec).
		SetMode(big.ToNearestEven)

	_, ok := target.SetString(Log2Of10String)
	if !ok {
		return nil, fmt.Errorf("failed to parse Log2Of10String into big.Float")
	}

	return target, nil
}

func main() {
	log2Val, err := NewLog2Of10()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Allocated Precision: %d bits\n", log2Val.Prec())

	// Format back out to check that all digits are preserved intact
	// Formatting with 'f' format to the exact length of the fractional digits
	numFractionalDigits := len(Log2Of10String) - 2
	outputStr := log2Val.Text('f', numFractionalDigits)

	// Verify equality
	if outputStr == Log2Of10String {
		fmt.Printf("Success! All %d digits match the constant exactly.\n", numFractionalDigits)
	} else {
		fmt.Println("Precision discrepancy detected.")
	}
}
