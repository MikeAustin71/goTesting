package weightconv
import "fmt"

type Pound float64
type Kilogram float64

const (
	PoundsPerKilogram Pound = 0.45359237
	KilogramsPerPound Kilogram = 2.2046226218487758072297380134503
)

func (k Kilogram) String() string    { return fmt.Sprintf("%g Killograms", k) }
func (p Pound) String() string { return fmt.Sprintf("%g Pounds", p) }