package main

import (
	"fmt"
	"github.com/bengarrett/conv/siin"
	"github.com/bengarrett/conv/siout"
	"github.com/bengarrett/conv/units"
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

// init checks for the existence of user arguments otherwise the help is displayed.
func init() {
	if len(os.Args) < 2 {
		help()
		os.Exit(0)
	}
}

func main() {
	/* Handle user input */
	uunit, fin := input()

	/* Calculate and display conversions */
	switch {
	case uunit == "f":
		in, out = "fah", "cel"
		// takes the float input (fin) and converts it to a International System (SI) unit
		// then takes the SI unit and converts it into Celsius
		fout = siin.Celsius(siout.Fahrenheit(fin))
		// displays the input and the results to the user
		temp(fin, fout, in, out)
	case uunit == "c":
		in, out := "cel", "fah"
		fout = siin.Fahrenheit(siout.Celsius(fin))
		temp(fin, fout, in, out)
	default:
		help()
	}
	os.Exit(0)
}

// temp displays temperature conversions.
func temp(fin float64, fout float64, in string, out string) {
	fmt.Printf("%.1f %s converts to ", fin, units.Info(in, "sym"))
	fmt.Printf("%.1f %s (%s)\n", fout, units.Info(out, "sym"), units.Info(out, "nam"))
}

// help displays the end user help.
func help() {
	fmt.Printf("conv is a tool that converts common usage metric and imperial units of \nmeasurements.\n\n")
	fmt.Printf("Usage:\n\tconv number unit\n\n")
	fmt.Printf("The %s units are:\n\n", units.Info("fah", "qty"))
	fmt.Printf("\tf\t%s (%s)\n", units.Info("fah", "sym"), units.Info("fah", "nam"))
	fmt.Printf("\tc\t%s (%s)\n", units.Info("cel", "sym"), units.Info("cel", "nam"))

	/* Please keep this as it is a requirement of the CC licence. */
	p("\nCreated by Ben Garrett as a learning project for Go (lang).\n")
	p("Personal page: <http://bens.zone> Source code: <https://github.com/bengarrett>\n")
}

// input parses user arguments to determine the unit type and value.
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
	// convert value into an float64 type so we can use it with calculations.
	fi, _ := strconv.ParseFloat(uv, 64)
	fi = float64(fi)
	return uu, fi
}
