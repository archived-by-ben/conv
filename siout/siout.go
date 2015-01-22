/*
Le Système International d'Unités
International System of Units
metric or imperial to SI unit
*/

package siout

func Fahrenheit(f float64) float64 {
	return (f + 459.67) * 5 / 9
}

func Celsius(c float64) float64 {
	return c + 273.15
}
