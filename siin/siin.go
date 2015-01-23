// Siin package converts units of SI (International System of Units)
// into common metric and imperial units.
package siin

func Fahrenheit(si float64) float64 {
	return si*9/5 - 459.67
}

func Celsius(si float64) float64 {
	return si - 273.15
}

func Horsepower(si float64) float64 {
	return si / 746
}

func Watt(si float64) float64 {
	return si
}
