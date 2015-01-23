/*
Program conv is a simple unit conversion tool for common measurements.
*/

package main

import (
	"fmt"
	"github.com/bengarrett/conv/siin"
	"github.com/bengarrett/conv/siout"
	"github.com/bengarrett/conv/units"
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
	}
}

// Handle user input then calculate and display any conversions.
func main() {
	uin, x := input()
	switch uin {
	case "f":
		y = siin.Celsius(siout.Fahrenheit(x))
		pLegacy(x, y, uin, "c")
	case "c":
		y = siin.Fahrenheit(siout.Celsius(x))
		pLegacy(x, y, uin, "f")
	case "hp":
		y = siin.Watt(siout.Horsepower(x))
		pMetric(x, y, uin, "w")
	case "w":
		y = siin.Horsepower(siout.Watt(x))
		pLegacy(x, y, uin, "hp")
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

// Legacy formats and prints the results of imperial conversions and other measurements.
func pLegacy(x float64, y float64, uin string, uout string) {
	symin, symout, nameout := units.Info(uin, "sym"), units.Info(uout, "sym"), units.Info(uout, "nam")
	calcWhole(x)
	p("%s converts to ", symin)
	calcWhole(y)
	p("%s (%s)\n", symout, nameout)
}

// Metric formats and prints the results of metric conversions.
func pMetric(x float64, y float64, uin string, uout string) {
	symin, symout, nameout := units.Info(uin, "sym"), units.Info(uout, "sym"), units.Info(uout, "nam")
	calcWhole(x)
	p("%s equals to ", symin)
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

// calcWhole determines if a float value should display as a decimal value
// or as a more readable whole number.
func calcWhole(x float64) {
	if ceil, floor := math.Ceil(x), math.Floor(x); ceil == floor {
		p("%.0f", x)
	} else {
		p("%.1f", x)
	}
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
	slice := []string{"c", "f", "w", "hp"}
	pHelpT(slice)
	/* Please keep this as it is a requirement of the CC licence. */
	p("\nCreated by Ben Garrett as a learning project for Go (lang).\n")
	p("Personal: <http://bens.zone> Source: <https://github.com/bengarrett/conv>\n")
}

// pHelpT is a template used by the phelp function to print permitted units.
func pHelpT(s []string) {
	for _, u := range s {
		p("\t%s\t%s (%s)\n", u, units.Info(u, "sym"), units.Info(u, "nam"))
	}
}
