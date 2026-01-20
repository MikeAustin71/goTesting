package main

// D:\GoProjects\MikeAustin71\goTesting\Enums\05_enumstringtimezones\timezones\ianatimezones.go
import (
	"fmt"

	"github.com/mikeaustin71/Enums/05_enumstringtimezones/timezones"
)

func main() {

	fmt.Println("Testing IANA Time Zones")

	str := timezones.IanaTz.America.Chicago()

	fmt.Println("US-Central: ", str)

	str = timezones.IanaTz.US.Mountain()

	fmt.Println("US-Mountain: ", str)

	fmt.Println("Cuba: ", timezones.IanaTz.Cuba())

}
