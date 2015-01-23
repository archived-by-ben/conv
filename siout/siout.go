// Siout package converts common units
// into units of SI (International System of Units).
package siout

func Fahrenheit(f float64) float64 {
	return (f + 459.67) * 5 / 9
}

func Celsius(c float64) float64 {
	return c + 273.15
}

func Horsepower(hp float64) float64 {
	return hp * 746
}

func Watt(w float64) float64 {
	return w
}
