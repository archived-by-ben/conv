/*
Le Système International d'Unités
International System of Units
SI to metric or imperial to units
*/

package siin

func Fahrenheit(si float64) float64 {
	return si*9/5 - 459.67
}

func Celsius(si float64) float64 {
	return si - 273.15
}
