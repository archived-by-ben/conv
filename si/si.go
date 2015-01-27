// si package takes units of SI (International System)
// and converts them into common metric and imperial units.
package si

import (
	"fmt"
	"math"
	"strconv"
)

// Rounds a float for display. Decimals less than 0.1 are dropped.
func Round(x float64) string {
	y, _ := strconv.ParseFloat(fmt.Sprintf("%.0f", x), 64)
	if z := y - x; math.Abs(z) < 0.1 {
		return fmt.Sprintf("%.0f", x)
	}
	return fmt.Sprintf("%.1f", x)
}

func Fahrenheit(x float64) float64 {
	return x*9/5 - 459.67
}

func Celsius(x float64) float64 {
	return x - 273.15
}

func Horsepower(x float64) float64 {
	return x / 746
}

func Watt(x float64) float64 {
	return x
}

func Kmh(x float64) float64 {
	return x / 0.27777777777778
}

func Mph(x float64) float64 {
	return x / 0.44704
}

func Kn(x float64) float64 {
	return x / 0.514444
}

func Mps(x float64) float64 {
	return x
}

func Carat(x float64) float64 {
	return x / 0.2
}
func Gram(x float64) float64 {
	return x
}
func Ounce(x float64) float64 {
	return x * 28
}
func Pound(x float64) float64 {
	return x * 453.59237
}
func Stone(x float64) float64 {
	return x * 6350.29318
}

func Centimetre(x float64) float64 {
	return x * 100
}
func Inch(x float64) float64 {
	return x / 0.0254
}
func Yard(x float64) float64 {
	return x / 0.9144
}
func Foot(x float64) float64 {
	return x / 0.3048
}
func Metre(x float64) float64 {
	return x
}
func Kilometre(x float64) float64 {
	return x / 1000
}
func Mile(x float64) float64 {
	return x / 1609.344
}
func Nautical(x float64) float64 {
	return x / 1852
}
