// Convert package combines common and si funcs to
// return the unit measurement conversions.
package convert

import (
	"fmt"
	"github.com/bengarrett/conv/common"
	"github.com/bengarrett/conv/si"
	"math"
	"strconv"
	"strings"
)

// Set to true will force Round() to display many decimals.
// If set to true, convert_test.go results will break.
var float bool = false

// Rounds a float for display.
// Decimal values less than 0.1 are dropped.
func Round(x float64) string {
	if float == true {
		return fmt.Sprintf("%.49f", x)
	}
	y, _ := strconv.ParseFloat(fmt.Sprintf("%.0f", x), 64)
	if z := y - x; math.Abs(z) < 0.099 {
		return fmt.Sprintf("%.0f", x)
	}
	return fmt.Sprintf("%.1f", x)
}

// calcPrefix calculates and prints a metric prefix, using a
// supplied number, exponent, factor symbol and unit symbol.
func Prefix(x float64, e int, fs string, us string) string {
	if y := x / math.Pow10(e); y > 1 {
		return fmt.Sprintf("%s%s%s, ", Round(y), fs, strings.TrimSpace(us))
	}
	return ""
}

func Cf(x float64) string {
	y := common.Celsius(x)
	z := si.Fahrenheit(y)
	return Round(z)
}

func Cmin(x float64) string {
	y := common.Centimetre(x)
	z := si.Inch(y)
	return Round(z)
}

func Cmft(x float64) string {
	y := common.Centimetre(x)
	z := si.Foot(y)
	return Round(z)
}

func Cmmi(x float64) string {
	y := common.Centimetre(x)
	z := si.Mile(y)
	return Round(z)
}

func Cmnm(x float64) string {
	y := common.Centimetre(x)
	z := si.Nautical(y)
	return Round(z)
}

func Cmyd(x float64) string {
	y := common.Centimetre(x)
	z := si.Yard(y)
	return Round(z)
}

func Ctg(x float64) string {
	y := common.Carat(x)
	z := si.Gram(y)
	return Round(z)
}

func Fc(x float64) string {
	y := common.Fahrenheit(x)
	z := si.Celsius(y)
	return Round(z)
}

func Ftcm(x float64) string {
	y := common.Foot(x)
	z := si.Centimetre(y)
	return Round(z)
}

func Ftin(x float64) string {
	y := common.Foot(x)
	z := si.Inch(y)
	return Round(z)
}

func Ftkm(x float64) string {
	y := common.Foot(x)
	z := si.Kilometre(y)
	return Round(z)
}

func Ftm(x float64) string {
	y := common.Foot(x)
	z := si.Metre(y)
	return Round(z)
}

func Ftmi(x float64) string {
	y := common.Foot(x)
	z := si.Mile(y)
	return Round(z)
}

func Ftnm(x float64) string {
	y := common.Foot(x)
	z := si.Nautical(y)
	return Round(z)
}

func Ftyd(x float64) string {
	y := common.Foot(x)
	z := si.Yard(y)
	return Round(z)
}

func Glb(x float64) string {
	y := common.Gram(x)
	z := si.Pound(y)
	return Round(z)
}

func Goz(x float64) string {
	y := common.Gram(x)
	z := si.Ounce(y)
	return Round(z)
}

func Gst(x float64) string {
	y := common.Gram(x)
	z := si.Stone(y)
	return Round(z)
}

func Gukl(x float64) string {
	y := common.GallonUK(x)
	z := si.Litre(y)
	return Round(z)
}

func Gukbbl(x float64) string {
	y := common.GallonUK(x)
	z := si.Barrel(y)
	return Round(z)
}

func Gukcum(x float64) string {
	y := common.GallonUK(x)
	z := si.Cubicmetre(y)
	return Round(z)
}

func Gukgus(x float64) string {
	y := common.GallonUK(x)
	z := si.GallonUS(y)
	return Round(z)
}

func Gusl(x float64) string {
	y := common.GallonUS(x)
	z := si.Litre(y)
	return Round(z)
}

func Gusbbl(x float64) string {
	y := common.GallonUS(x)
	z := si.Barrel(y)
	return Round(z)
}

func Guscum(x float64) string {
	y := common.GallonUS(x)
	z := si.Cubicmetre(y)
	return Round(z)
}

func Gusguk(x float64) string {
	y := common.GallonUS(x)
	z := si.GallonUK(y)
	return Round(z)
}

func Hpw(x float64) string {
	y := common.Horsepower(x)
	z := si.Watt(y)
	return Round(z)
}

func Incm(x float64) string {
	y := common.Inch(x)
	z := si.Centimetre(y)
	return Round(z)
}

func Inft(x float64) string {
	y := common.Inch(x)
	z := si.Foot(y)
	return Round(z)
}

func Inkm(x float64) string {
	y := common.Inch(x)
	z := si.Kilometre(y)
	return Round(z)
}

func Inm(x float64) string {
	y := common.Inch(x)
	z := si.Metre(y)
	return Round(z)
}

func Inmi(x float64) string {
	y := common.Inch(x)
	z := si.Mile(y)
	return Round(z)
}

func Innm(x float64) string {
	y := common.Inch(x)
	z := si.Nautical(y)
	return Round(z)
}

func Inyd(x float64) string {
	y := common.Inch(x)
	z := si.Yard(y)
	return Round(z)
}

func Kmhkn(x float64) string {
	y := common.Kmh(x)
	z := si.Kn(y)
	return Round(z)
}

func Kmhmph(x float64) string {
	y := common.Kmh(x)
	z := si.Mph(y)
	return Round(z)
}

func Kmhmps(x float64) string {
	y := common.Kmh(x)
	z := si.Mps(y)
	return Round(z)
}

func Lbg(x float64) string {
	y := common.Pound(x)
	z := si.Gram(y)
	return Round(z)
}

func Lboz(x float64) string {
	y := common.Pound(x)
	z := si.Ounce(y)
	return Round(z)
}

func Lbst(x float64) string {
	y := common.Pound(x)
	z := si.Stone(y)
	return Round(z)
}

func Micm(x float64) string {
	y := common.Mile(x)
	z := si.Centimetre(y)
	return Round(z)
}

func Mift(x float64) string {
	y := common.Mile(x)
	z := si.Foot(y)
	return Round(z)
}

func Miin(x float64) string {
	y := common.Mile(x)
	z := si.Inch(y)
	return Round(z)
}

func Mikm(x float64) string {
	y := common.Mile(x)
	z := si.Kilometre(y)
	return Round(z)
}

func Mim(x float64) string {
	y := common.Mile(x)
	z := si.Metre(y)
	return Round(z)
}

func Minm(x float64) string {
	y := common.Mile(x)
	z := si.Nautical(y)
	return Round(z)
}

func Miyd(x float64) string {
	y := common.Mile(x)
	z := si.Yard(y)
	return Round(z)
}

func Mphkmh(x float64) string {
	y := common.Mph(x)
	z := si.Kmh(y)
	return Round(z)
}

func Mphkn(x float64) string {
	y := common.Mph(x)
	z := si.Kn(y)
	return Round(z)
}

func Mphmps(x float64) string {
	y := common.Mph(x)
	z := si.Mps(y)
	return Round(z)
}

func Nmcm(x float64) string {
	y := common.Nautical(x)
	z := si.Centimetre(y)
	return Round(z)
}

func Nmft(x float64) string {
	y := common.Nautical(x)
	z := si.Foot(y)
	return Round(z)
}

func Nmin(x float64) string {
	y := common.Nautical(x)
	z := si.Inch(y)
	return Round(z)
}

func Nmkm(x float64) string {
	y := common.Nautical(x)
	z := si.Kilometre(y)
	return Round(z)
}

func Nmm(x float64) string {
	y := common.Nautical(x)
	z := si.Metre(y)
	return Round(z)
}

func Nmmi(x float64) string {
	y := common.Nautical(x)
	z := si.Mile(y)
	return Round(z)
}

func Nmyd(x float64) string {
	y := common.Nautical(x)
	z := si.Yard(y)
	return Round(z)
}

func Ozg(x float64) string {
	y := common.Ounce(x)
	z := si.Gram(y)
	return Round(z)
}

func Ozlb(x float64) string {
	y := common.Ounce(x)
	z := si.Pound(y)
	return Round(z)
}

func Ozst(x float64) string {
	y := common.Ounce(x)
	z := si.Stone(y)
	return Round(z)
}

func Stg(x float64) string {
	y := common.Stone(x)
	z := si.Gram(y)
	return Round(z)
}

func Stlb(x float64) string {
	y := common.Stone(x)
	z := si.Pound(y)
	return Round(z)
}

func Stoz(x float64) string {
	y := common.Stone(x)
	z := si.Ounce(y)
	return Round(z)
}

func Whp(x float64) string {
	y := common.Watt(x)
	z := si.Horsepower(y)
	return Round(z)
}

func Ydcm(x float64) string {
	y := common.Yard(x)
	z := si.Centimetre(y)
	return Round(z)
}

func Ydft(x float64) string {
	y := common.Yard(x)
	z := si.Foot(y)
	return Round(z)
}

func Ydin(x float64) string {
	y := common.Yard(x)
	z := si.Inch(y)
	return Round(z)
}

func Ydkm(x float64) string {
	y := common.Yard(x)
	z := si.Kilometre(y)
	return Round(z)
}

func Ydm(x float64) string {
	y := common.Yard(x)
	z := si.Metre(y)
	return Round(z)
}

func Ydmi(x float64) string {
	y := common.Yard(x)
	z := si.Mile(y)
	return Round(z)
}

func Ydnm(x float64) string {
	y := common.Yard(x)
	z := si.Nautical(y)
	return Round(z)
}

///\\\\////\\\\\/////\\\\\\
func Cumguk(x float64) string {
	y := common.Cubicmetre(x)
	z := si.GallonUK(y)
	return Round(z)
}

func Cumgus(x float64) string {
	y := common.Cubicmetre(x)
	z := si.GallonUS(y)
	return Round(z)
}

func Cumbbl(x float64) string {
	y := common.Cubicmetre(x)
	z := si.Barrel(y)
	return Round(z)
}

func Cuml(x float64) string {
	y := common.Cubicmetre(x)
	z := si.Litre(y)
	return Round(z)
}

///

func Bblcum(x float64) string {
	y := common.Barrel(x)
	z := si.Cubicmetre(y)
	return Round(z)
}

func Bblguk(x float64) string {
	y := common.Barrel(x)
	z := si.GallonUK(y)
	return Round(z)
}

func Bblgus(x float64) string {
	y := common.Barrel(x)
	z := si.GallonUS(y)
	return Round(z)
}

func Bbll(x float64) string {
	y := common.Barrel(x)
	z := si.Litre(y)
	return Round(z)
}

///
