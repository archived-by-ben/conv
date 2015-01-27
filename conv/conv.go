/*
Program conv is a simple unit conversion tool for common measurements.
*/

package main

import (
	"fmt"
	"github.com/bengarrett/conv/common"
	"github.com/bengarrett/conv/si"
	"github.com/bengarrett/conv/symbols"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	x, y float64
	p    = fmt.Printf
)

// Init checks for the existence of user arguments otherwise the help is displayed.
func init() {
	ver := "0.1"
	// no arguments provided
	if len(os.Args) < 2 {
		err := fmt.Errorf("No measurement or unit were provided")
		fmt.Println(err)
		pHelp()
		os.Exit(0)
	}
	// user requested help
	switch os.Args[1] {
	case "-h", "--help", "/?":
		pHelp()
		os.Exit(0)
	case "-v", "--version", "/v":
		p("conv %s\n", ver)
		pCopyright()
		os.Exit(0)
	case "-x":
		/* Test functions argument (delete before making public) */
		fmt.Println("experimental mode:")
		//symbols.Test()
		os.Exit(0)
	}
}

// Handle user input then calculate and print all conversions.
func main() {
	m, x := input()
	pResults(m, x)
	p("\n")
}

func pResults(m string, x float64) {
	switch m {
	case "f":
		y = si.Celsius(common.Fahrenheit(x))
		pLegacy(x, y, m, "c")
	case "c":
		y = si.Fahrenheit(common.Celsius(x))
		pLegacy(x, y, m, "f")
	case "hp":
		y = si.Watt(common.Horsepower(x))
		pMetric(x, y, m, "w")
	case "w":
		y = si.Horsepower(common.Watt(x))
		pLegacy(x, y, m, "hp")
	case "kph", "kmh":
		y = si.Mph(common.Kmh(x))
		pLegacy(x, y, m, "mph")
		pSupplement(common.Kmh(x), "mps")
		pSupplement(si.Kn(common.Kmh(x)), "kn")
	case "ct":
		y = si.Gram(common.Carat(x))
		pMetric(x, y, m, "g")
	case "g", "oz", "lb", "st":
		// TODO: Add support for kg
		switch m {
		case "g":
			y = si.Ounce(common.Gram(x))
			pLegacy(x, y, m, "oz")
		case "oz":
			y = si.Gram(common.Ounce(x))
			pMetric(x, y, m, "g")
			x = y
		case "lb":
			y = si.Gram(common.Pound(x))
			pMetric(x, y, m, "g")
			x = y
		case "st":
			y = si.Gram(common.Stone(x))
			pMetric(x, y, m, "g")
			x = y
		}
		if y = si.Ounce(common.Gram(x)); y >= 0.1 && m != "g" {
			p("\nAnd also")
		} else if y = si.Pound(common.Gram(x)); y >= 0.1 && m == "g" {
			p("\nAnd also")
		}
		if y = si.Pound(common.Gram(x)); m != "lb" && y >= 0.1 {
			pSupplement(y, "lb")
		}
		if y = si.Stone(common.Gram(x)); m != "st" && y >= 0.1 {
			pSupplement(y, "st")
		}
		if y = si.Ounce(common.Gram(x)); m != "g" && m != "oz" {
			pSupplement(y, "oz")
		}
	case "cm", "m", "km", "in", "ft", "yd", "mi", "nm":
		switch m {
		case "cm":
			y = si.Inch(common.Centimetre(x))
			pLegacy(x, y, m, "in")
			y = si.Metre(common.Centimetre(x))
		case "m":
			y = si.Yard(common.Metre(x))
			pLegacy(x, y, m, "yd")
			y = x
		case "km":
			y = si.Mile(common.Kilometre(x))
			pMetric(x, y, m, "mi")
			y = si.Metre(common.Kilometre(x))
		case "in":
			y = si.Metre(common.Inch(x))
			pMetric(x, y, m, "m")
		case "ft":
			y = si.Metre(common.Foot(x))
			pMetric(x, y, m, "m")
		case "yd":
			y = si.Metre(common.Yard(x))
			pMetric(x, y, m, "m")
		case "mi":
			y = si.Kilometre(common.Mile(x))
			pMetric(x, y, m, "km")
			y = si.Metre(common.Mile(x))
		case "nm":
			y = si.Kilometre(common.Nautical(x))
			pMetric(x, y, m, "km")
			y = si.Metre(common.Nautical(x))
		}
		if m != "m" {
			x = y
		}
		p("\nAnd also")
		if y = si.Nautical(common.Metre(x)); m != "nm" && y >= 0.1 {
			pSupplement(y, "nm")
		}
		if y = si.Mile(common.Metre(x)); m != "mi" && m != "km" && y >= 0.1 {
			pSupplement(y, "mi")
		}
		if y = si.Yard(common.Metre(x)); m != "yd" && y >= 0.1 {
			pSupplement(y, "yd")
		}
		if y = si.Foot(common.Metre(x)); m != "ft" && y >= 0.1 {
			pSupplement(y, "ft")
		}
		if m != "in" && m != "cm" {
			y = si.Inch(common.Metre(x))
			pSupplement(y, "in")
		}
	default:
		if len(m) > 1 {
			err := fmt.Errorf("The unit type '%s' does not exist", m)
			fmt.Println(err)
		} else {
			err := fmt.Errorf("No unit type was supplied")
			fmt.Println(err)
		}
		os.Exit(0)
	}
}

// Input parses user arguments to determine the measurement unit and value.
func input() (string, float64) {
	// create a slice to append strings
	s := []string{}
	s = append(s, os.Args[1]) // append 1st user argument
	if len(os.Args[1:]) > 1 {
		s = append(s, os.Args[2]) // append 2nd user argument if provided
	}
	// ignore all other user arguments but combine 1st & 2nd into a single string
	args := strings.Join(s, "")
	args = strings.ToLower(args)
	// create a reg expression to replace all non-alphabetic characters
	m := regexp.MustCompile("[^a-z]").ReplaceAllString(args, "")
	// create reg expression to removal anything that isn't a number, decimal or sign
	val := regexp.MustCompile("[^-?0-9+.?[^0-9]*").ReplaceAllString(args, "")
	if len(val) < 1 {
		err := fmt.Errorf("No measurement was provided")
		fmt.Println(err)
	}
	// convert value into an float64 type so we can use it with calculations.
	x, _ := strconv.ParseFloat(val, 64)
	x = float64(x)
	// return the parsed measurement and value
	return m, x
}

// Supplement formats and prints the results of imperial and other measurement conversions.
func pSupplement(x float64, uout string) {
	symout, nameout := symbols.Info(uout, "sym"), symbols.Info(uout, "nam")
	p(", %s%s (%s)", si.Round(x), symout, nameout)
}

// Legacy formats and prints the results of imperial conversions and other measurements.
func pLegacy(x float64, y float64, uin string, uout string) {
	symin, symout, nameout := symbols.Info(uin, "sym"), symbols.Info(uout, "sym"), symbols.Info(uout, "nam")
	p("%s%s is about %s%s (%s)", si.Round(x), symin, si.Round(y), symout, nameout)
}

// Metric formats and prints the results of metric conversions.
func pMetric(x float64, y float64, uin string, uout string) {
	symin, symout, nameout := symbols.Info(uin, "sym"), symbols.Info(uout, "sym"), symbols.Info(uout, "nam")
	p("%s%s is around ", si.Round(x), symin)
	calcPrefix(y, 24, "Y", symout)
	calcPrefix(y, 21, "Z", symout)
	calcPrefix(y, 18, "E", symout)
	calcPrefix(y, 15, "P", symout)
	calcPrefix(y, 12, "T", symout)
	calcPrefix(y, 9, "G", symout)
	calcPrefix(y, 6, "M", symout)
	calcPrefix(y, 3, "k", symout)
	p("%s%s (%s)", si.Round(y), symout, nameout)
}

// calcPrefix calculates and prints a metric prefix, using a
// supplied number, exponent, factor symbol and unit symbol.
func calcPrefix(x float64, e int, fs string, us string) {
	if y := x / math.Pow10(e); y > 1 && y < 10 {
		p("%s%s%s, ", si.Round(y), fs, us)
	}
}

// phelp prints the end user help.
func pHelp() {
	slice := []string{"c", "cm", "ct", "f", "ft", "g", "hp", "in", "km", "kn", "kph", "lb", "m", "mph", "mps", "mi", "nm", "oz", "st", "w", "yd"}
	p("conv is a tool that converts common use units of measurements.\n\n")
	p("Usage:\n\tconv measurement unit\n\n")
	p("Example:\n\tconv 100f\n\n")
	p("\t100 °F converts to 37.8 °C (Celsius)\n\n")
	p("The permitted units are:\n\n")
	pHelpT(slice)
	p("\nCreated by Ben Garrett as a learning project for Go (lang).\n")
	p("Source: <https://github.com/bengarrett/conv>\n")
	pCopyright()
}

// pHelpT is a template used by the phelp function to print permitted units.
func pHelpT(s []string) {
	for _, u := range s {
		p("\t%s\t%s \t%s\n", u, strings.TrimSpace(symbols.Info(u, "sym")), symbols.Info(u, "nam"))
	}
}

func pCopyright() {
	/* Please keep this as it is a requirement of the MIT licence. */
	p("\n\tThe MIT License (MIT)\n")
	p("\tCopyright © 2015 Ben Garrett.\n")
	p("\t<http://choosealicense.com/licenses/mit/>\n")
}
