// Calculate package handles the mathematics for the measurement conversions.
package calculate

import (
	"fmt"
	"math"
	"strconv"
)

const (
	float = false
)

// Converter function takes a measurement value and an 'in' unit string and converts it to the requested 'out' unit.
// There is no allowed unit conversion logic in this function, that is handled by conv/process().
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

// Round function rounds a float and converts it to a string for display.
// Decimal values lower than 0.1 are dropped.
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

// Si_common function takes a SI measurement and converts it into a common measurement.
func si_common(x float64, out string) (common float64, err error) {
	switch out {
	default:
		err := fmt.Errorf("si_common could not find unit '%v'", out)
		return common, err
	case "bbl":
		common = x * 6.2898105697751
	case "ct":
		common = x / 0.2
	case "c":
		common = x - 273.15
	case "cm":
		common = x * 100
	case "cum":
		common = x
	case "f":
		common = x*9/5 - 459.67
	case "ft":
		common = x / 0.3048
	case "g":
		common = x
	case "guk":
		common = x * 219.97
	case "gus":
		common = x * 264.17
	case "hp":
		common = x / 746
	case "in":
		common = x / 0.0254
	case "kg":
		common = x / 1000
	case "km":
		common = x / 1000
	case "kmh":
		common = x / 0.27777777777778
	case "kn":
		common = x / 0.514444
	case "l":
		common = x / 0.001
	case "m":
		common = x
	case "mi":
		common = x / 1609.344
	case "mph":
		common = x / 0.44704
	case "mps":
		common = x
	case "nm":
		common = x / 1852
	case "oz":
		common = x / 28.3495231
	case "lb":
		common = x / 453.59237
	case "st":
		common = x / 6350.29318
	case "w":
		common = x
	case "yd":
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

// Common_si function takes a common measurement and converts it into a si measurement.
func common_si(x float64, in string) (si float64, err error) {
	switch in {
	default:
		err := fmt.Errorf("common_si could not find unit '%v'", in)
		return si, err
	case "bbl":
		si = x * 0.158987294928
	case "ct":
		si = x * 0.2
	case "c":
		si = x + 273.15
	case "cm":
		si = x * 0.01
	case "cum":
		si = x
	case "f":
		si = (x + 459.67) * 5 / 9
	case "ft":
		si = x * 0.3048
	case "g":
		si = x
	case "guk":
		si = x / 219.97
	case "gus":
		si = x / 264.17
	case "hp":
		si = x * 746
	case "in":
		si = x * 0.0254
	case "kg":
		si = x * 1000
	case "km":
		si = x * 1000
	case "kmh":
		si = x * 0.27777777777778
	case "kn":
		si = x * 0.514444
	case "l":
		si = x * 0.001
	case "m":
		si = x
	case "mi":
		si = x * 1609.344
	case "mph":
		si = x * 0.44704
	case "mps":
		si = x
	case "nm":
		si = x * 1852
	case "oz":
		si = x * 28.3495231
	case "lb":
		si = x * 453.59237
	case "st":
		si = x * 6350.29318
	case "w":
		si = x
	case "yd":
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
