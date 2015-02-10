/*
Program conv is a simple unit conversion tool for common measurements.
*/

package main

import (
	"bytes"
	"fmt"
	"github.com/bengarrett/conv/calculate"
	"github.com/bengarrett/conv/symbols"
	"math"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

const (
	app = "conv" // application name
	ver = "1.00" // version
)

var (
	// initialise the buffer which will be used to save text intended for the user's console
	buf bytes.Buffer
	// bs returns the content of the buffer as a string
	bs = buf.String
	// bt formats strings and saves them to the buffer
	bt = func(format string, a ...interface{}) {
		buf.WriteString(fmt.Sprintf(format, a...))
	}
	// p is a print to console short-cut
	p = fmt.Printf
)

// Init function handles arguments that force the application to exit back to the operating system.
func init() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		if runtime.GOOS == "windows" {
			if arg == "/?" {
				arg = "-h"
			} else {
				arg = regexp.MustCompile("[/]").ReplaceAllString(arg, "-")
			}
		}
		switch arg {
		case "-h", "--help":
			help()
			p("%s\n", bs())
			os.Exit(0)
		case "-v", "--version":
			bt("%s %s (%s-%s/%s)\n", app, ver, runtime.Compiler, runtime.GOARCH, runtime.GOOS)
			copyright()
			p("%s\n", bs())
			os.Exit(0)
		case "--sample":
			// run a loop to display sample results of all units
			unitDump := symbols.UnitData()
			// sort units by key name
			mapofUnits := make([]string, 0, len(unitDump))
			for unitItem := range unitDump {
				mapofUnits = append(mapofUnits, unitItem)
			}
			sort.Strings(mapofUnits)
			for i, _ := range mapofUnits {
				u := mapofUnits[i]
				eg := fmt.Sprint("100", u)
				bt("\n[%s] results from \"%v\"\n", symbols.Proper(u), eg)
				m, x, _ := input([]string{"place-holder", eg})
				process(m, x)
				bt("\n")
			}
			p("%s\n", bs())
			os.Exit(0)
		}
	}
}

// Main function handles the calculation and display of measurement conversions.
func main() {
	// parse user input
	m, x, err := input(os.Args)
	// if there were no problems with user input, use process() to calculate and save the results to the buffer.
	// if the user values created errors, return those errors instead.
	if err == nil {
		err = process(m, x)
	}
	// print any errors and abort
	if err != nil {
		if runtime.GOOS == "windows" {
			err = fmt.Errorf("%s: %s\nTry 'conv /?' for more information.", app, err)
		} else {
			err = fmt.Errorf("%s: %s\nTry 'conv --help' for more information.", app, err)
		}
		fmt.Println(err)
		os.Exit(0)
	}
	// if no errors print the buffer to the user's console to display conversion results.
	p("%s\n", bs())
	// EOP
}

// Input function parses the user's arguments to determine the measurement unit and value.
func input(i []string) (m string, x float64, err error) {
	var args string = ""
	if len(i) < 2 {
		err = fmt.Errorf("no measurement or unit were provided")
	} else {
		// create a slice to append strings
		s := []string{}
		s = append(s, i[1]) // append 1st user argument
		if len(i[1:]) > 1 {
			s = append(s, i[2]) // append 2nd user argument if provided
		}
		// ignore all other user arguments but combine 1st & 2nd into a single string
		args = strings.Join(s, "")
		args = strings.ToLower(args)
	}
	// create a regular expression to replace all non-alphabetic characters
	if m = regexp.MustCompile("[^a-z]").ReplaceAllString(args, ""); len(m) < 1 {
		err = fmt.Errorf("no unit type was supplied")
	} else {
		// create regex to removal anything that isn't a number, decimal or sign
		if val := regexp.MustCompile("[^-?0-9+.?[^0-9]*").ReplaceAllString(args, ""); len(val) < 1 {
			err = fmt.Errorf("no measurement was provided")
		} else {
			// convert val from a string type into a float64 so we can use it in calculations.
			if x, err = strconv.ParseFloat(val, 64); err != nil {
				err = fmt.Errorf("could not parse '%v'", val)
			} else {
				x = float64(x)
			}

		}
	}
	// return the parsed measurement and value
	return m, x, err
}

// Process function handles logic to determine which measurement conversions should be returned.
func process(m string, x float64) (err error) {
	s, o, also := "", "", "\nAnd also " // string, output, "and also" string
	// allow for common use alternative names and spelling
	switch m {
	case "knot", "knots":
		m = "kn"
	case "kph", "kmph":
		m = "kmh"
	case "mih":
		m = "mph"
	case "metre", "meter", "metres", "meters":
		m = "m"
	case "mile", "miles":
		m = "mi"
	case "yard", "yards":
		m = "yd"
	case "litres", "liters", "litre", "liter":
		m = "l"
	case "ounce", "ounces":
		m = "oz"
	case "pound", "pounds":
		m = "lb"
	case "stone":
		m = "st"
	case "watt", "watts":
		m = "w"
	case "inch", "inches":
		m = "in"
	case "gram", "grams":
		m = "g"
	case "foot", "feet":
		m = "ft"
	case "carat", "carats":
		m = "ct"
	}
	switch m {
	default:
		err = fmt.Errorf("the unit type '%s' does not exist", m)
	// temperature
	case "c", "f":
		o = "f"
		if m == "f" {
			o = "c"
		}
		s = calculate.Converter(x, m, o)
		dspPrimary(x, s, m, o)
	// power
	case "hp", "w":
		o = "hp"
		if m == "hp" {
			o = "w"
		}
		s = calculate.Converter(x, m, o)
		dspPrimary(x, s, m, o)
	// speed
	case "kmh", "mph", "mps", "kn":
		o = "kmh"
		if m == "kmh" {
			o = "mph"
		}
		s = calculate.Converter(x, m, o)
		dspPrimary(x, s, m, o)
		bt(also)
		dspExtras(x, m, o, "kmh", "mph", "mps", "kn")
	// weight
	case "ct":
		o = "g"
		s = calculate.Converter(x, m, o)
		dspPrimary(x, s, m, o)
	case "g", "kg", "oz", "lb", "st":
		o = "g"
		if m == "g" {
			o = "oz"
		} else if m == "kg" {
			o = "st"
		}
		s = calculate.Converter(x, m, o)
		dspPrimary(x, s, m, o)
		bt(also)
		dspExtras(x, m, o, "ct", "g", "kg", "oz", "lb", "st")
	// distance - length
	case "cm", "m", "km", "in", "ft", "yd", "mi", "nm":
		o = "m"
		if m == "cm" {
			o = "in"
		} else if m == "m" {
			o = "yd"
		} else if m == "km" {
			o = "mi"
		} else if m == "in" {
			o = "cm"
		} else if m == "mi" || m == "nm" {
			o = "km"
		}
		s = calculate.Converter(x, m, o)
		dspPrimary(x, s, m, o)
		bt(also)
		dspExtras(x, m, o, "cm", "m", "km")
		bt("\n")
		dspExtras(x, m, o, "in", "ft", "yd", "mi", "nm")
	// liquid volume
	case "bbl", "cum", "guk", "gus", "l":
		o = "l"
		if m == "cum" || m == "l" {
			o = "gus"
		}
		s = calculate.Converter(x, m, o)
		dspPrimary(x, s, m, o)
		bt(also)
		dspExtras(x, m, o, "bbl", "cum", "guk", "gus", "l")
	case "bps", "kbps", "mbps":
		o = "bs"
		s = calculate.Converter(x, m, o)
		dspPrimary(x, s, m, o)
	case "bs", "kbs", "mbs":
		o = "bps"
		s = calculate.Converter(x, m, o)
		dspPrimary(x, s, m, o)
	}
	return err
}

// DspPrimary function prints the primary conversion result.
func dspPrimary(x float64, s string, uin string, uout string) {
	symin, symout, nameout := symbols.Glyph(uin), symbols.Glyph(uout), symbols.Proper(uout)
	bt("%s%s is about ", calculate.Round(x), symin)
	// apply metric prefixes
	switch uout {
	case "w", "g", "l", "bs", "bps":
		prefixes := []string{"k", "M", "G", "T", "P", "E", "Z", "Y"}
		for e, i := 24, 7; e > 0; e, i = e-3, i-1 {
			f, _ := strconv.ParseFloat(s, 64)
			b := f / math.Pow10(e)
			if b >= 1 {
				bs := strconv.FormatFloat(b, 'f', 0, 64)
				if b < 10 {
					bs = strconv.FormatFloat(b, 'f', 2, 64)
				} else if b < 100 {
					bs = strconv.FormatFloat(b, 'f', 1, 64)
				}
				if strings.TrimSpace(symout) != symout {
					bt("%v ", bs)
				} else {
					bt("%v", bs)
				}
				bt("%v%v, ", prefixes[i], strings.TrimSpace(symout))
			}
		}
	}
	bt("%s%s (%s)", s, symout, nameout)
}

// DspExtras function appends the results of all additional conversions.
func dspExtras(x float64, m string, skip string, uout ...string) {
	c := 0
	for _, u := range uout {
		if u != m && u != skip {
			c++
			if c > 1 {
				bt(", ")
			}
			bt("%v%v (%v)", calculate.Converter(x, m, u), symbols.Glyph(u), symbols.Proper(u))
		}
	}
}

// Help function prints the end user help.
func help() {
	// simulate user provided arguments for example
	example := "100f"
	m, x, _ := input([]string{"place-holder", example})
	// fetch all units data for display
	unitDump := symbols.UnitData()
	// sort units by key name
	mapofUnits := make([]string, 0, len(unitDump))
	for unitItem := range unitDump {
		mapofUnits = append(mapofUnits, unitItem)
	}
	sort.Strings(mapofUnits)
	bt("%s is a tool to convert everyday units of measurements.\n\nUsage:\t\t%s (measurement)(unit)\n", app, app)
	if runtime.GOOS == "windows" {
		bt("\t\t%s /?\tHelp\n\t\t%s /v\tVersion\n", app, app)
	} else {
		bt("\t\t%s -h | --help\n\t\t%s -v | --version\n", app, app)
	}
	bt("\nExample:\t%s %s\n\t\t", app, example)
	process(m, x)
	bt("\n\nThe permitted units are:\n\n")
	for i, data := range mapofUnits {
		bt("\t%s\t» %s\t%s\n", mapofUnits[i], strings.TrimSpace(unitDump[data].Symbol), unitDump[data].Name)
	}
	copyright()
	bt("\nSource: <https://github.com/bengarrett/conv>\n")
	symbols.UnitData()
}

// Copyright function prints the copyright notice.
func copyright() {
	/* Please keep this as it is a requirement of the MIT licence. */
	bt("\nCopyright © 2015 Ben Garrett.\nThe MIT License (MIT)\t<http://choosealicense.com/licenses/mit/>")
}
