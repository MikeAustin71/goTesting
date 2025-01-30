package main

import "fmt"

// AmericaTimeZones - IANA Time Zones in America
type americaTimeZones string

func (amTz americaTimeZones) Chicago() string { return "America/Chicago"}
func (amTz americaTimeZones) NewYork() string {return "America/New_York"}
func (amTz americaTimeZones) Denver() string {return "America/Denver"}
func (amTz americaTimeZones) LosAngeles() string {return "America/Los_Angeles"}
func (amTz americaTimeZones) Argentina() argentinaTimeZones {return argentinaTimeZones("")}

// America/Argentina/Buenos_Aires

type argentinaTimeZones string

func (arg argentinaTimeZones) BuenosAires() string {return "America/Argentina/Buenos_Aires"}
func (arg argentinaTimeZones) Catamarca() string {return "America/Argentina/Catamarca"}

type europeTimeZones string

func (euTz europeTimeZones) London() string {return "Europe/London"}
func (euTz europeTimeZones) Paris() string {return "Europe/Paris"}
func (euTz europeTimeZones) Madrid() string {return "Europe/Madrid"}
func (euTz europeTimeZones) Berlin() string {return "Europe/Berlin"}

type asiaTimeZones string

func (asTz asiaTimeZones) Tokyo() string {return "Asia/Tokyo"}
func (asTz asiaTimeZones) HoChiMinh() string {return "Asia/Ho_Chi_Minh"}
func (asTz asiaTimeZones) Seoul() string {return "Asia/Seoul"}
func (asTz asiaTimeZones) Shanghai() string {return "Asia/Shanghai"}


type IanaTimeZone struct {
	America americaTimeZones
	Asia	asiaTimeZones
	Europe europeTimeZones
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
