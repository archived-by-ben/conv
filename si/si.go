// si package takes units of SI (International System)
// and converts them into common metric and imperial units.
package si

func Barrel(x float64) float64 {
	return x * 6.2898105697751
}

func Carat(x float64) float64 {
	return x / 0.2
}

func Celsius(x float64) float64 {
	return x - 273.15
}

func Centimetre(x float64) float64 {
	return x * 100
}

func Cubicmetre(x float64) float64 {
	return x
}

func Fahrenheit(x float64) float64 {
	return x*9/5 - 459.67
}

func Foot(x float64) float64 {
	return x / 0.3048
}

func Gram(x float64) float64 {
	return x
}

func GallonUK(x float64) float64 {
	return x * 219.97
}

func GallonUS(x float64) float64 {
	return x * 264.17
}

func Horsepower(x float64) float64 {
	return x / 746
}

func Inch(x float64) float64 {
	return x / 0.0254
}

func Kilometre(x float64) float64 {
	return x / 1000
}

func Kmh(x float64) float64 {
	return x / 0.27777777777778
}

func Kn(x float64) float64 {
	return x / 0.514444
}

func Litre(x float64) float64 {
	return x / 0.001
}

func Metre(x float64) float64 {
	return x
}

func Mile(x float64) float64 {
	return x / 1609.344
}

func Mph(x float64) float64 {
	return x / 0.44704
}

func Mps(x float64) float64 {
	return x
}

func Nautical(x float64) float64 {
	return x / 1852
}

func Ounce(x float64) float64 {
	return x / 28.3495231
}

func Pound(x float64) float64 {
	return x / 453.59237
}

func Stone(x float64) float64 {
	return x / 6350.29318
}

func Watt(x float64) float64 {
	return x
}

func Yard(x float64) float64 {
	return x / 0.9144
}
