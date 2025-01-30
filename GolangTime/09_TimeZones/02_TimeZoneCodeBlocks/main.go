package main

// Taken from
// https://bl.ocks.org/joyrexus/a56717634a672dcdfd48
// J Voigt's Block

import (
	"fmt"
	"time"
	"encoding/binary"
)

func main() {

	const shortForm = "2006-01-02"
	loc, _ := time.LoadLocation("America/Chicago")

	t, err := time.ParseInLocation(shortForm, "2015-01-19", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	d := decode(encode(t))

	fmt.Println(d)
	fmt.Println(d.Format(shortForm))
}

// decode unmarshals a time.
func decode(b []byte) time.Time {
	i := int64(binary.BigEndian.Uint64(b))
	return time.Unix(i, 0)
}

// encode marshals a time.
func encode(t time.Time) []byte {
	buf := make([]byte, 8)
	u := uint64(t.Unix())
	binary.BigEndian.PutUint64(buf, u)
	return buf
}
// Output
// 2015-01-19 00:00:00 -0600 CST
// 2015-01-19


