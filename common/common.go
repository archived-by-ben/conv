// Common package converts common units
// into units of SI (International System).
package common

func Barrel(x float64) float64 {
	return x * 0.158987294928
}

func Carat(x float64) float64 {
	return x * 0.2
}

func Celsius(x float64) float64 {
	return x + 273.15
}

func Centimetre(x float64) float64 {
	return x * 0.01
}

func Cubicmetre(x float64) float64 {
	return x
}

func Fahrenheit(x float64) float64 {
	return (x + 459.67) * 5 / 9
}

func Foot(x float64) float64 {
	return x * 0.3048
}

func Gram(x float64) float64 {
	return x
}

func GallonUK(x float64) float64 {
	return x / 219.97
}

func GallonUS(x float64) float64 {
	return x / 264.17
}

func Horsepower(x float64) float64 {
	return x * 746
}

func Inch(x float64) float64 {
	return x * 0.0254
}

func Kilometre(x float64) float64 {
	return x * 1000
}

func Kmh(x float64) float64 {
	return x * 0.27777777777778
}

func Kn(x float64) float64 {
	return x * 0.514444
}

func Litre(x float64) float64 {
	return x * 0.001
}

func Metre(x float64) float64 {
	return x
}

func Mile(x float64) float64 {
	return x * 1609.344
}

func Mph(x float64) float64 {
	return x * 0.44704
}

func Mps(x float64) float64 {
	return x
}

func Nautical(x float64) float64 {
	return x * 1852
}

func Ounce(x float64) float64 {
	return x * 28.3495231
}

func Pound(x float64) float64 {
	return x * 453.59237
}

func Stone(x float64) float64 {
	return x * 6350.29318
}

func Watt(x float64) float64 {
	return x
}

func Yard(x float64) float64 {
	return x * 0.9144
}
