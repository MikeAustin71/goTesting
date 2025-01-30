package main

import (
	"fmt"
)

func main() {

	d := make(map[int] map[string]string)

	d[0] = make(map[string]string,0)

	d[0]["0 PM"] =  "0 pm"
	d[1] = make(map[string]string,0)
	d[1]["1 PM"] = "1 pm"

	fmt.Println("d[0]", d[0])


	/*
		d[1]["1 PM"] = "1 pm"
		d[2]["2 PM"] = "2 pm"
		d[3]["3 PM"] = "3 pm"
		d[4]["4 PM"] = "4 pm"
		d[5]["5 PM"] = "5 pm"
		d[6]["6 PM"] = "6 pm"
		d[7]["7 PM"] = "7 pm"
		d[8]["8 PM"] = "8 pm"
		d[9]["9 PM"] = "9 pm"
		d[10]["0 P.M."] = "0 pm"
		d[11]["1 P.M."] = "1 pm"
		d[12]["2 P.M."] = "2 pm"
		d[13]["3 P.M."] = "3 pm"
		d[14]["4 P.M."] = "4 pm"
		d[15]["5 P.M."] = "5 pm"
		d[16]["6 P.M."] = "6 pm"
		d[17]["7 P.M."] = "7 pm"
		d[18]["8 P.M."] = "8 pm"
		d[19]["9 P.M."] = "9 pm"
		d[20]["0PM"] = "0 pm"

		fmt.Println("---------------------")
		fmt.Println("d[0]= ", d[0])
		/*
		   type path []byte

		   func (p *path) TruncateAtFinalSlash() {
		       i := bytes.LastIndex(*p, []byte("/"))
		       if i >= 0 {
		           *p = (*p)[0:i]
		       }
		   }

		   func main() {
		       pathName := path("/usr/bin/tso") // Conversion from string to path.
		       pathName.TruncateAtFinalSlash()
		       fmt.Printf("%s\n", pathName)
		   }

	*/

}
