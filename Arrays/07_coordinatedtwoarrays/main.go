package main

import "fmt"

var x = [25]int{-12,-11,-10,-9,-8,-7,-6,-5,-4,-3,-2,-1,0,1,2,3,4,5,6,7,8,9,10,11,12}
/*
Military Time Code Letter Reference:

UTC -12: Y- (e.g. Fiji)

UTC-11: X (American Samoa)

UTC-10: W (Honolulu, HI)

UTC-9: V (Juneau, AK)

UTC-8: U (PST, Los Angeles, CA)

UTC-7: T (MST, Denver, CO)

UTC-6: S (CST, Dallas, TX)

UTC-5: R (EST, New York, NY)

UTC-4: Q (Halifax, Nova Scotia

UTC-3: P (Buenos Aires, Argentina)

UTC-2: O (Godthab, Greenland)

UTC-1: N (Azores)

UTC+-0: Z (Zulu time)

UTC+1: A (France)

UTC+2: B (Athens, Greece)

UTC+3: C (Arab Standard Time, Iraq, Bahrain, Kuwait, Saudi Arabia, Yemen, Qatar)

UTC+4: D (Used for Moscow, Russia and Afghanistan, however, Afghanistan is technically +4:30 from UTC)

UTC+5: E (Pakistan, Kazakhstan, Tajikistan, Uzbekistan and Turkmenistan)

UTC+6: F (Bangladesh)

UTC+7: G (Thailand)

UTC+8: H (Beijing, China)

UTC+9: I (Tokyo, Australia)

UTC+10: K (Brisbane, Australia)

UTC+11: L (Sydney, Australia)

UTC+12: M (Wellington, New Zealand)
 */

var y = [25]string{
					"Y",
					"X",
					"W",
					"V",
					"U",
					"T",
					"S",
					"R",
					"Q",
					"P",
					"O",
					"N",
					"Z",
					"A",
					"B",
					"C",
					"D",
					"E",
					"F",
					"G",
					"H",
					"I",
					"K",
					"L",
					"M"}


func main() {

	for k:=0; k < 25; k++ {
		fmt.Println("x= ",x[k], "  2nd-Dimension= ", y[k])
	}

}

