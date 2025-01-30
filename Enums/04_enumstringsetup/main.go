package main

import (
	"fmt"
	"reflect"
)

type MilitaryTz string

func (mtz MilitaryTz) Z() MilitaryTz  {return MilitaryTz("Zulu")}

func (mtz MilitaryTz) String() string {

	return string(mtz)
}

// AmericaTimeZones - IANA Time Zones in America
type AmericaTimeZones string

func (amTz AmericaTimeZones) Chicago() string { return "America/Chicago"}
func (amTz AmericaTimeZones) NewYork() string {return "America/New_York"}
func (amTz AmericaTimeZones) Denver() string {return "America/Denver"}
func (amTz AmericaTimeZones) LosAngeles() string {return "America/Los_Angeles"}
func (amTz AmericaTimeZones) Argentina() ArgentinaTimeZones {return ArgentinaTimeZones("")}

// America/Argentina/Buenos_Aires

type ArgentinaTimeZones string

func (arg ArgentinaTimeZones) BuenosAires() string {return "America/Argentina/Buenos_Aires"}
func (arg ArgentinaTimeZones) Catamarca() string {return "America/Argentina/Catamarca"}

type EuropeTimeZones string

func (euTz EuropeTimeZones) London() string {return "Europe/London"}
func (euTz EuropeTimeZones) Paris() string {return "Europe/Paris"}
func (euTz EuropeTimeZones) Madrid() string {return "Europe/Madrid"}
func (euTz EuropeTimeZones) Berlin() string {return "Europe/Berlin"}

type AsiaTimeZones string

func (asTz AsiaTimeZones) Tokyo() string {return "Asia/Tokyo"}
func (asTz AsiaTimeZones) HoChiMinh() string {return "Asia/Ho_Chi_Minh"}
func (asTz AsiaTimeZones) Seoul() string {return "Asia/Seoul"}
func (asTz AsiaTimeZones) Shanghai() string {return "Asia/Shanghai"}


type IanaTimeZone struct {
  America AmericaTimeZones
	Asia	AsiaTimeZones
	Europe EuropeTimeZones
}

func (iana IanaTimeZone) Cuba() string {return "Cuba"}

var IanaTz = IanaTimeZone{}



func main() {
	chicago := IanaTimeZone{}.America.Chicago()

	fmt.Println("Chicago: ", chicago)

	tokoyo := IanaTimeZone{}.Asia.Tokyo()

	fmt.Println("Tokyo: ", tokoyo)

	//buenosAires := IanaTz.America.Argentina().BuenosAires()
	buenosAires := IanaTz.America.Argentina().BuenosAires()
	fmt.Println("Buenos Aires: ", buenosAires)

	cuba := IanaTimeZone{}.Cuba()
	fmt.Println("Cuba: ", cuba)

}

func test20() {
	x := MilitaryTz("").Z()

	fmt.Println("           x: ", x)
	fmt.Println("str.String(): ", x.String())
	s := reflect.TypeOf(x)
	fmt.Println("      s.Name: ", s.Name())


}