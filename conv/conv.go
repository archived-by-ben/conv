/*
Program conv is a simple unit conversion tool for common measurements.
*/

package main

import (
	"fmt"
	"github.com/bengarrett/conv/convert"
	"github.com/bengarrett/conv/symbols"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	p = fmt.Printf
)

// Init checks for the existence of user arguments otherwise the help is displayed.
func init() {
	ver := "0.4"
	// no arguments provided
	if len(os.Args) < 2 {
		err := fmt.Errorf("No measurement or unit were provided")
		fmt.Println(err)
		pHelp()
		os.Exit(1)
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

// Handle user input then calculate and print all conversions.
func main() {
	m, x := input()
	pResults(m, x)
	p("\n")
}

func pResults(m string, x float64) {
	s := ""
	switch m {
	// temperature
	case "f":
		s = convert.Fc(x)
		pL1(x, s, m, "c")
	case "c":
		s = convert.Cf(x)
		pL1(x, s, m, "f")
	// power
	case "hp":
		s = convert.Hpw(x)
		pL1Metrics(x, s, m, "w")
	case "w":
		s = convert.Whp(x)
		pL1(x, s, m, "hp")
	// speed
	case "kmh", "kph":
		s = convert.Kmhmph(x)
		pL1(x, s, m, "mph")
		s = convert.Kmhmps(x)
		pL1x(s, "mps")
		s = convert.Kmhkn(x)
		pL1x(s, "kn")
	case "mph", "mih":
		s = convert.Mphkmh(x)
		pL1(x, s, m, "kmh")
		s = convert.Mphmps(x)
		pL1x(s, "mps")
		s = convert.Mphkn(x)
		pL1x(s, "kn")
	case "mps":
		s = convert.Mpskmh(x)
		pL1(x, s, m, "kmh")
		s = convert.Mpsmph(x)
		pL1x(s, "mph")
		s = convert.Mpskn(x)
		pL1x(s, "kn")
	case "kn":
		s = convert.Knkmh(x)
		pL1(x, s, m, "kmh")
		s = convert.Knmps(x)
		pL1x(s, "mps")
		s = convert.Knmph(x)
		pL1x(s, "mph")
	// weight
	case "ct":
		s = convert.Ctg(x)
		pL1Metrics(x, s, m, "g")
	case "g", "kg", "oz", "lb", "st":
		sOz, sLb, sSt := "0", "0", "0"
		switch m {
		case "g":
			s = convert.Goz(x)
			sLb = convert.Glb(x)
			sSt = convert.Gst(x)
			pL1(x, s, m, "oz")
		case "kg":
			x = x * 1000 // convert kg into g for calculations
			s = convert.Gst(x)
			sOz = convert.Goz(x)
			sLb = convert.Glb(x)
			x = x / 1000 // revert g back to k for display
			pL1(x, s, m, "st")
		case "oz":
			s = convert.Ozg(x)
			sLb = convert.Ozlb(x)
			sSt = convert.Ozst(x)
			pL1Metrics(x, s, m, "g")
		case "lb":
			s = convert.Lbg(x)
			sOz = convert.Lboz(x)
			sSt = convert.Lbst(x)
			pL1Metrics(x, s, m, "g")
		case "st":
			s = convert.Stg(x)
			sLb = convert.Stlb(x)
			sOz = convert.Stoz(x)
			pL1Metrics(x, s, m, "g")
		}
		if sSt != "0" || sLb != "0" || sOz != "0" {
			p("\nAnd also")
			if sSt != "0" {
				pL2(sSt, "st")
			}
			if sLb != "0" {
				pL2(sLb, "lb")
			}
			if sOz != "0" {
				pL2(sOz, "oz")
			}
		}
	// distance - length
	case "cm", "m", "km", "in", "ft", "yd", "mi", "nm":
		scm, sm, skm, sin := "0", "0", "0", "0"
		sft, syd, smi, snm := "0", "0", "0", "0"
		switch m {
		case "cm":
			s = convert.Cmin(x)
			sm = convert.Round(x / 100)
			skm = convert.Round(x / 100 / 1000)
			sft = convert.Cmft(x)
			syd = convert.Cmyd(x)
			smi = convert.Cmmi(x)
			snm = convert.Cmnm(x)
			pL1(x, s, m, "in")
		case "m":
			x = x * 100 // convert m into cm for calculations
			s = convert.Cmyd(x)
			scm = convert.Round(x / 100)
			skm = convert.Round(x / 100 / 1000)
			sin = convert.Cmin(x)
			sft = convert.Cmft(x)
			smi = convert.Cmmi(x)
			snm = convert.Cmnm(x)
			x = x / 100 // convert cm back to m for display
			pL1(x, s, m, "yd")
		case "km":
			x = x * 100 * 1000 // convert km into cm for calculations
			s = convert.Cmmi(x)
			scm = convert.Round(x / 100 / 1000)
			sin = convert.Cmin(x)
			sft = convert.Cmft(x)
			syd = convert.Cmyd(x)
			sm = convert.Round(x / 100)
			snm = convert.Cmnm(x)
			x = x / 100 / 1000 // convert cm back to km for display
			pL1(x, s, m, "mi")
		case "in":
			s = convert.Incm(x)
			sft = convert.Inft(x)
			syd = convert.Inyd(x)
			sm = convert.Inm(x)
			skm = convert.Inkm(x)
			smi = convert.Inmi(x)
			snm = convert.Innm(x)
			pL1(x, s, m, "cm")
		case "ft":
			s = convert.Ftm(x)
			scm = convert.Ftcm(x)
			sin = convert.Ftin(x)
			syd = convert.Ftyd(x)
			skm = convert.Ftkm(x)
			smi = convert.Ftmi(x)
			snm = convert.Ftnm(x)
			pL1(x, s, m, "m")
		case "yd":
			s = convert.Ydm(x)
			scm = convert.Ydcm(x)
			sin = convert.Ydin(x)
			sft = convert.Ydft(x)
			skm = convert.Ydkm(x)
			smi = convert.Ydmi(x)
			snm = convert.Ydnm(x)
			pL1(x, s, m, "m")
		case "mi":
			s = convert.Mikm(x)
			scm = convert.Micm(x)
			sin = convert.Miin(x)
			syd = convert.Miyd(x)
			sm = convert.Mim(x)
			sft = convert.Mift(x)
			snm = convert.Minm(x)
			pL1(x, s, m, "km")
		case "nm":
			s = convert.Nmkm(x)
			scm = convert.Nmcm(x)
			sin = convert.Nmin(x)
			syd = convert.Nmyd(x)
			sm = convert.Nmm(x)
			sft = convert.Nmft(x)
			smi = convert.Nmmi(x)
			pL1(x, s, m, "km")
		}
		if scm != "0" || sm != "0" || skm != "0" || sin != "0" ||
			sft != "0" || syd != "0" || smi != "0" || snm != "0" {
			p("\nAnd also")
			if snm != "0" {
				pL2(snm, "nm")
			}
			if smi != "0" {
				pL2(smi, "mi")
			}
			if skm != "0" {
				pL2(skm, "km")
			}
			if sm != "0" {
				pL2(sm, "m")
			}
			if syd != "0" {
				pL2(syd, "yd")
			}
			if sft != "0" {
				pL2(sft, "ft")
			}
			if sin != "0" {
				pL2(sin, "in")
			}
			if scm != "0" {
				pL2(scm, "cm")
			}
		}
	// liquid volume
	case "bbl", "cum", "guk", "gus", "l":
		sbbl, scum, sguk, sgus, sl := "0", "0", "0", "0", "0"
		switch m {
		case "bbl":
			s = convert.Bbll(x)
			scum = convert.Bblcum(x)
			sguk = convert.Bblguk(x)
			sgus = convert.Bblgus(x)
			pL1(x, s, m, "l")
		case "cum":
			sbbl = convert.Cumbbl(x)
			sl = convert.Cuml(x)
			s = convert.Cumguk(x)
			pL1(x, s, m, "guk")
			s = convert.Cumgus(x)
			pL1x(s, "gus")
		case "guk":
			s = convert.Gukl(x)
			sbbl = convert.Gukbbl(x)
			scum = convert.Gukcum(x)
			sgus = convert.Gukgus(x)
			pL1(x, s, m, "l")
		case "gus":
			s = convert.Gusl(x)
			sbbl = convert.Gusbbl(x)
			scum = convert.Guscum(x)
			sguk = convert.Gusguk(x)
			pL1(x, s, m, "l")
		case "l":
			x = x / 1000 // convert L into m3 for calculations
			sbbl = convert.Cumbbl(x)
			scum = convert.Round(x)
			s = convert.Cumguk(x)
			pL1(x*1000, s, m, "guk") // convert m3 back to L for display
			s = convert.Cumgus(x)
			pL1x(s, "gus")
		}
		if sbbl != "0" || scum != "0" || sguk != "0" || sgus != "0" || sl != "0" {
			p("\nAnd also")
			if sl != "0" {
				pL2(sl, "l")
			}
			if sbbl != "0" {
				pL2(sbbl, "bbl")
			}
			if sgus != "0" {
				pL2(sgus, "gus")
			}
			if sguk != "0" {
				pL2(sguk, "guk")
			}
			if scum != "0" {
				pL2(scum, "cum")
			}
		}
	// errors
	default:
		if len(m) > 1 {
			err := fmt.Errorf("The unit type '%s' does not exist", m)
			fmt.Println(err)
		} else {
			err := fmt.Errorf("No unit type was supplied")
			fmt.Println(err)
		}
		os.Exit(1)
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

// pL1 prints the conversion result on line 1.
func pL1(x float64, s string, uin string, uout string) {
	symin, symout, nameout := symbols.Info(uin, "sym"), symbols.Info(uout, "sym"), symbols.Info(uout, "nam")
	p("%s%s is about %s%s (%s)", convert.Round(x), symin, s, symout, nameout)
}

// pL1x appends additional conversion results on line 1.
func pL1x(x string, uout string) {
	symout, nameout := symbols.Info(uout, "sym"), symbols.Info(uout, "nam")
	p(", %s%s (%s)", x, symout, nameout)
}

// pL1Metrics prints the conversion result on line 1.
// It is an alternative to pL1() as is designed for metric output results.
func pL1Metrics(x float64, s string, uin string, uout string) {
	symin, symout, nameout := symbols.Info(uin, "sym"), symbols.Info(uout, "sym"), symbols.Info(uout, "nam")
	p("%s%s is around ", convert.Round(x), symin)
	y, _ := strconv.ParseFloat(s, 64)
	p("%s", convert.Prefix(y, 24, "Y", symout))
	p("%s", convert.Prefix(y, 21, "Z", symout))
	p("%s", convert.Prefix(y, 18, "E", symout))
	p("%s", convert.Prefix(y, 15, "P", symout))
	p("%s", convert.Prefix(y, 12, "T", symout))
	p("%s", convert.Prefix(y, 9, "G", symout))
	p("%s", convert.Prefix(y, 6, "M", symout))
	p("%s", convert.Prefix(y, 3, "k", symout))
	p("%s%s (%s)", s, symout, nameout)
}

// pL2 prints supplement conversions onto line 2.
func pL2(x string, uout string) {
	symout, nameout := symbols.Info(uout, "sym"), symbols.Info(uout, "nam")
	p(", %s%s (%s)", x, symout, nameout)
}

// phelp prints the end user help.
func pHelp() {
	slice := []string{"bbl", "c", "cm", "ct", "cum", "f", "ft", "g", "guk", "gus", "hp", "in", "km", "kmh", "kn", "l", "lb", "m", "mph", "mps", "mi", "nm", "oz", "st", "w", "yd"}
	p("conv is a tool that converts common use units of measurements.\n\n")
	p("Usage:\n\tconv measurement unit\n\n")
	p("Example:\n\tconv 100f\n\n")
	p("\t100 °F is about 37.8 °C (Celsius)\n\n")
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

// pCopyright prints a copyright notice.
func pCopyright() {
	/* Please keep this as it is a requirement of the MIT licence. */
	p("\n\tThe MIT License (MIT)\n")
	p("\tCopyright © 2015 Ben Garrett.\n")
	p("\t<http://choosealicense.com/licenses/mit/>\n")
}
