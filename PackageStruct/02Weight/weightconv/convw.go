package weightconv

// weightconv is designed to convert Kilograms to Pounds
// and Pounds to Kilograms

// KToP converts Kilograms to pounds.
func KToP(k Kilogram) Pound { return Pound(k* KilogramsPerPound) }

// PToK converts pounds to kilograms
func PToK(p Pound) Kilogram { return Kilogram(p * PoundsPerKilogram )}
