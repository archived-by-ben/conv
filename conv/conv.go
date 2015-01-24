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
	}
}

// Handle user input then calculate and display any conversions.
func main() {
	uin, x := input()
	switch uin {
	case "f":
		y = si.Celsius(common.Fahrenheit(x))
		pLegacy(x, y, uin, "c")
	case "c":
		y = si.Fahrenheit(common.Celsius(x))
		pLegacy(x, y, uin, "f")
	case "hp":
		y = si.Watt(common.Horsepower(x))
		pMetric(x, y, uin, "w")
	case "w":
		y = si.Horsepower(common.Watt(x))
		pLegacy(x, y, uin, "hp")
	case "kph", "kmh":
		y = si.Mph(common.Kmh(x))
		pLegacy(x, y, uin, "mph")
		pSupplement(common.Kmh(x), "mps")
		pSupplement(si.Kn(common.Kmh(x)), "kn")
		p("\n")
	case "ct":
		y = si.Gram(common.Carat(x))
		pMetric(x, y, uin, "g")
	case "g":
		y = si.Ounce(common.Gram(x))
		pLegacy(x, y, uin, "oz")
		if y = si.Pound(common.Gram(x)); y > 0.05 {
			p("And also ")
			pSupplement(y, "lb")
			if y = si.Stone(common.Gram(x)); y > 0.05 {
				pSupplement(y, "st")
			}
			p("\n")
		}
	case "oz":
		y = si.Gram(common.Ounce(x))
		pMetric(x, y, uin, "g")
		if y = si.Pound(common.Ounce(x)); y > 0.05 {
			p("And also ")
			pSupplement(y, "lb")
			if y = si.Stone(common.Ounce(x)); y > 0.05 {
				pSupplement(y, "st")
			}
			p("\n")
		}
	case "lb":
		y = si.Gram(common.Pound(x))
		pMetric(x, y, uin, "g")
		p("And also ")
		if y = si.Stone(common.Pound(x)); y > 0.05 {
			pSupplement(y, "st")
		}
		y = si.Ounce(common.Pound(x))
		pSupplement(y, "oz")
		p("\n")
	case "st":
		y = si.Gram(common.Stone(x))
		pMetric(x, y, uin, "g")
		p("And also ")
		if y = si.Pound(common.Stone(x)); y > 0.05 {
			pSupplement(y, "lb")
		}
		y = si.Ounce(common.Stone(x))
		pSupplement(y, "oz")
		p("\n")
	default:
		err := fmt.Errorf("The unit type '%s' does not exist", uin)
		fmt.Println(err)
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
	uin := regexp.MustCompile("[^a-z]").ReplaceAllString(args, "")
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
	return uin, x
}

// Supplement formats and prints the results of imperial conversions and other measurements.
func pSupplement(x float64, uout string) {
	symout, nameout := symbols.Info(uout, "sym"), symbols.Info(uout, "nam")
	p("%s%s (%s)\t", fmtWhole(x), symout, nameout)
}

// Legacy formats and prints the results of imperial conversions and other measurements.
func pLegacy(x float64, y float64, uin string, uout string) {
	symin, symout, nameout := symbols.Info(uin, "sym"), symbols.Info(uout, "sym"), symbols.Info(uout, "nam")
	p("%s%s converts to ", fmtWhole(x), symin)
	p("%s%s (%s)\n", fmtWhole(y), symout, nameout)
}

// Metric formats and prints the results of metric conversions.
func pMetric(x float64, y float64, uin string, uout string) {
	symin, symout, nameout := symbols.Info(uin, "sym"), symbols.Info(uout, "sym"), symbols.Info(uout, "nam")
	p("%s%s equals to ", fmtWhole(x), symin)
	calcPrefix(y, 24, "Y", symout)
	calcPrefix(y, 21, "Z", symout)
	calcPrefix(y, 18, "E", symout)
	calcPrefix(y, 15, "P", symout)
	calcPrefix(y, 12, "T", symout)
	calcPrefix(y, 9, "G", symout)
	calcPrefix(y, 6, "M", symout)
	calcPrefix(y, 3, "k", symout)
	p("%.0f%s ", y, symout)
	if y > 1 {
		p("(%ss)\n", nameout)
	} else {
		p("(%s)\n", nameout)
	}
}

// calcPrefix calculates and prints a metric prefix, using a
// supplied number, exponent, factor symbol and unit symbol.
func calcPrefix(x float64, e int, fs string, us string) {
	if y := x / math.Pow10(e); y > 1 && y < 10 {
		p("%.1f%s%s, ", y, fs, us)
	} else if y > 1 {
		p("%.0f%s%s, ", y, fs, us)
	}
}

// fmtWhole determines if a float value should display as a decimal value
// or as a more readable whole number.
func fmtWhole(x float64) string {
	if ceil, floor := math.Ceil(x), math.Floor(x); ceil == floor {
		return fmt.Sprintf("%.0f", x)
	}
	return fmt.Sprintf("%.1f", x)
}

// phelp prints the end user help.
func pHelp() {
	p("conv is a tool that converts common use units of measurements.\n\n")
	p("Usage:\n\tconv measurement unit\n\n")
	p("Example:\n\tconv 100f\n\n")
	p("\t100.0 °F converts to 37.8 °C (Celsius)\n\n")
	// dynamically generate this help?
	//p("The %s units are:\n\n", units.Info("f", "cat"))
	p("The permitted units are:\n\n")
	slice := []string{"c", "f", "w", "hp", "kph", "mph", "ct", "oz"}
	pHelpT(slice)
	p("\nCreated by Ben Garrett as a learning project for Go (lang).\n")
	p("Source: <https://github.com/bengarrett/conv>\n")
	pCopyright()
}

// pHelpT is a template used by the phelp function to print permitted units.
func pHelpT(s []string) {
	for _, u := range s {
		p("\t%s\t%s (%s)\n", u, symbols.Info(u, "sym"), symbols.Info(u, "nam"))
	}
}

func pCopyright() {
	/* Please keep this as it is a requirement of the MIT licence. */
	p("\n\tThe MIT License (MIT)\n")
	p("\tCopyright © 2015 Ben Garrett.\n")
	p("\t<http://choosealicense.com/licenses/mit/>\n")
}
