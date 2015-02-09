// calculate package does the mathematics of the measurements.
package calculate

import (
	"fmt"
	"math"
	"strconv"
)

const (
	float = false
)

// Converter takes a value and converts it from in unit to out unit.
func Converter(x float64, in string, out string) (result string) {
	// note: common_si() resets err to nil
	if y, err := common_si(x, in); err != nil {
		// check for any common_si errors
		return fmt.Sprint(err)
	} else if z, err := si_common(y, out); err != nil {
		// if previous check okay, then check for any si_common errors
		return fmt.Sprint(err)
	} else {
		// if both previous checks out, then return formatted result
		return Round(z)
	}
}

// Rounds a float for display.
// Decimal values less than 0.1 are dropped.
func Round(x float64) (rounded string) {
	if float == true {
		rounded = fmt.Sprintf("%.49f", x)
	} else {
		y, _ := strconv.ParseFloat(fmt.Sprintf("%.0f", x), 64)
		if z := y - x; math.Abs(z) < 0.099 {
			rounded = fmt.Sprintf("%.0f", x)
		} else {
			rounded = fmt.Sprintf("%.1f", x)
		}
	}
	return rounded
}

// si_common takes a SI measurement and converts it into a common measurement.
func si_common(x float64, out string) (common float64, err error) {
	switch out {
	default:
		err := fmt.Errorf("si_common could not find unit '%v'", out)
		return common, err
	case "Barrel", "bbl":
		common = x * 6.2898105697751
	case "Carat", "ct":
		common = x / 0.2
	case "Celsius", "c":
		common = x - 273.15
	case "Centimetre", "cm":
		common = x * 100
	case "Cubicmetre", "cum":
		common = x
	case "Fahrenheit", "f":
		common = x*9/5 - 459.67
	case "Foot", "ft":
		common = x / 0.3048
	case "Gram", "g":
		common = x
	case "GallonUK", "guk":
		common = x * 219.97
	case "GallonUS", "gus":
		common = x * 264.17
	case "Horsepower", "hp":
		common = x / 746
	case "Inch", "in":
		common = x / 0.0254
	case "kg":
		common = x / 1000
	case "Kilometre", "km":
		common = x / 1000
	case "Kmh", "kmh":
		common = x / 0.27777777777778
	case "Kn", "kn":
		common = x / 0.514444
	case "Litre", "l":
		common = x / 0.001
	case "Metre", "m":
		common = x
	case "Mile", "mi":
		common = x / 1609.344
	case "Mph", "mph":
		common = x / 0.44704
	case "Mps", "mps":
		common = x
	case "Nautical", "nm":
		common = x / 1852
	case "Ounce", "oz":
		common = x / 28.3495231
	case "Pound", "lb":
		common = x / 453.59237
	case "Stone", "st":
		common = x / 6350.29318
	case "Watt", "w":
		common = x
	case "Yard", "yd":
		common = x / 0.9144
	case "bps":
		common = x
	case "kbps":
		common = x * 1000
	case "mbps":
		common = x * 1000000
	case "bs":
		common = x * 0.125
	case "kbs":
		common = x * 1000 * 0.125
	case "mbs":
		common = x * 1000000 * 0.125
	}
	return common, nil
}

// common_si takes a common measurement and converts it into a si measurement.
func common_si(x float64, in string) (si float64, err error) {
	switch in {
	default:
		err := fmt.Errorf("common_si could not find unit '%v'", in)
		return si, err
	case "Barrel", "bbl":
		si = x * 0.158987294928
	case "Carat", "ct":
		si = x * 0.2
	case "Celsius", "c":
		si = x + 273.15
	case "Centimetre", "cm":
		si = x * 0.01
	case "Cubicmetre", "cum":
		si = x
	case "Fahrenheit", "f":
		si = (x + 459.67) * 5 / 9
	case "Foot", "ft":
		si = x * 0.3048
	case "Gram", "g":
		si = x
	case "GallonUK", "guk":
		si = x / 219.97
	case "GallonUS", "gus":
		si = x / 264.17
	case "Horsepower", "hp":
		si = x * 746
	case "Inch", "in":
		si = x * 0.0254
	case "kg":
		si = x * 1000
	case "Kilometre", "km":
		si = x * 1000
	case "Kmh", "kmh":
		si = x * 0.27777777777778
	case "Kn", "kn":
		si = x * 0.514444
	case "Litre", "l":
		si = x * 0.001
	case "Metre", "m":
		si = x
	case "Mile", "mi":
		si = x * 1609.344
	case "Mph", "mph":
		si = x * 0.44704
	case "Mps", "mps":
		si = x
	case "Nautical", "nm":
		si = x * 1852
	case "Ounce", "oz":
		si = x * 28.3495231
	case "Pound", "lb":
		si = x * 453.59237
	case "Stone", "st":
		si = x * 6350.29318
	case "Watt", "w":
		si = x
	case "Yard", "yd":
		si = x * 0.9144
	case "bps":
		si = x
	case "kbps":
		si = x * 1000
	case "mbps":
		si = x * 1000000
	case "bs":
		si = x / 0.125
	case "kbs":
		si = x * 1000 / 0.125
	case "mbs":
		si = x * 1000000 / 0.125
	}
	return si, nil
}
