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
	in, out   string
	fin, fout float64
	p         = fmt.Printf
)

// Init checks for the existence of user arguments otherwise the help is displayed.
func init() {
	// no arguments provided
	if len(os.Args) < 2 {
		err := fmt.Errorf("No measurement or unit were provided")
		fmt.Println(err)
		help()
		os.Exit(0)
	}
	// user requested help
	switch os.Args[1] {
	case "-h", "--help", "/?":
		help()
		os.Exit(0)
	}
}

// Handle user input then calculate and display conversions.
func main() {
	uunit, fin := input()
	switch uunit {
	case "f":
		// takes the float input (fin) and converts it to a International System (SI) unit
		// then takes the SI unit and converts it into Celsius
		fout = siin.Celsius(siout.Fahrenheit(fin))
		// displays the input and the results to the user
		legacy(fin, fout, uunit, "c")
	case "c":
		fout = siin.Fahrenheit(siout.Celsius(fin))
		legacy(fin, fout, uunit, "f")
	case "hp":
		fout = siin.Watt(siout.Horsepower(fin))
		metric(fin, fout, uunit, "w")
	case "w":
		fout = siin.Horsepower(siout.Watt(fin))
		legacy(fin, fout, uunit, "hp")
	default:
		err := fmt.Errorf("The unit type '%s' does not exist", uunit)
		fmt.Println(err)
	}
}

// Help displays the end user help.
func help() {
	p("conv is a tool that converts common use units of measurements.\n\n")
	p("Usage:\n\tconv measurement unit\n\n")
	p("Example:\n\tconv 100f\n\n")
	p("\t100.0 °F converts to 37.8 °C (Celsius)\n\n")
	// dynamically generate this help?
	//p("The %s units are:\n\n", units.Info("f", "cat"))
	p("The permitted units are:\n\n")
	measurements := []string{"c", "f", "w", "hp"}
	helpTemplate(measurements)
	/* Please keep this as it is a requirement of the CC licence. */
	p("\nCreated by Ben Garrett as a learning project for Go (lang).\n")
	p("Personal: <http://bens.zone> Source: <https://github.com/bengarrett/conv>\n")
}

// HelpTemplate is used by the Help function to print permitted units.
func helpTemplate(unitCol []string) {
	for _, u := range unitCol {
		p("\t%s\t%s (%s)\n", u, units.Info(u, "sym"), units.Info(u, "nam"))
	}
}

// Input parses user arguments to determine the measurement unit and value.
func input() (uunit string, fin float64) {
	// create a slice to append strings
	s := []string{}
	s = append(s, os.Args[1]) // append 1st user argument
	if len(os.Args[1:]) > 1 {
		s = append(s, os.Args[2]) // append 2nd user argument if provided
	}
	// ignore all other user arguments but combine 1st & 2nd into a single string
	ui := strings.Join(s, "")
	ui = strings.ToLower(ui)
	// create a reg expression to replace all non-alpha characters
	uu := regexp.MustCompile("[^a-z]").ReplaceAllString(ui, "")
	// create reg expression to replace all non float numbers and characters
	uv := regexp.MustCompile("[^-?0-9+.?[^0-9]*").ReplaceAllString(ui, "")
	if len(uv) < 1 {
		err := fmt.Errorf("No measurement was provided")
		fmt.Println(err)
	}
	// convert value into an float64 type so we can use it with calculations.
	fi, _ := strconv.ParseFloat(uv, 64)
	fi = float64(fi)
	return uu, fi
}

// Legacy is a print function for imperial and other measurements.
func legacy(fin float64, fout float64, in string, out string) {
	symIn := units.Info(in, "sym")
	name := units.Info(out, "nam")
	symOut := units.Info(out, "sym")
	if ceil, floor := math.Ceil(fin), math.Floor(fin); ceil == floor {
		p("%.0f%s converts to ", fin, symIn)
	} else {
		p("%.1f%s converts to ", fin, symIn)
	}
	if ceil, floor := math.Ceil(fout), math.Floor(fout); ceil == floor {
		p("%.0f%s (%s)\n", fout, symOut, name)
	} else {
		p("%.1f%s (%s)\n", fout, symOut, name)
	}
}

// Metric is a print function for metric measurements.
func metric(fin float64, fout float64, in string, out string) {
	symIn := units.Info(in, "sym")
	name := units.Info(out, "nam")
	symOut := units.Info(out, "sym")
	if ceil, floor := math.Ceil(fin), math.Floor(fin); ceil == floor {
		p("%.0f%s equals in %ss\n", fin, symIn, name)
	} else {
		p("%.1f%s equals in %ss\n", fin, symIn, name)
	}
	prefix(fout, symOut)
}

// Prefix is a print function for metric prefixes that are applied
// to measurements.
func prefix(unit float64, measure string) {
	if calc := unit / math.Pow10(12); calc > 1 && calc < 10 {
		p("%.1fT%s, ", calc, measure)
	} else if calc > 1 {
		p("%.fT%s, ", calc, measure)
	}
	if calc := unit / math.Pow10(9); calc > 1 && calc < 10 {
		p("%.1fG%s, ", calc, measure)
	} else if calc > 1 {
		p("%.fG%s, ", calc, measure)
	}
	if calc := unit / math.Pow10(6); calc > 1 && calc < 10 {
		p("%.1fM%s, ", calc, measure)
	} else if calc > 1 {
		p("%.fM%s, ", calc, measure)
	}
	if calc := unit / math.Pow10(3); calc > 1 && calc < 10 {
		p("%.1fk%s, ", calc, measure)
	} else if calc > 1 {
		p("%.fk%s, ", calc, measure)
	}
	p("%.0f%s\n", unit, measure)
}
