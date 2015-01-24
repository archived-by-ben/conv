// Package units returns the name, symbols and categories of
// the units used in conversion.
package symbols

import (
	"fmt"
)

type unit struct {
	name, symbol string
}

// Unit data containing the name, symbol & category.
var (
	c   = unit{"Celsius", "°C"}
	f   = unit{"Fahrenheit", "°F"}
	hp  = unit{"horsepower", "hp"}
	w   = unit{"watt", "W"}
	kph = unit{"kilometres per hour", "km/h"}
	mph = unit{"miles per hour", "mph"}
	kn  = unit{"knot", "kn"}
	mps = unit{"metres per second", "m/s"}
	ct  = unit{"carat", "ct"}
	g   = unit{"gram", "g"}
	oz  = unit{"ounce", "oz"}
	lb  = unit{"pound", "℔"}
	st  = unit{"stone", "st"}
)

// Info extracts data from the collection of unit data.
// Here it fetches the row of data then uses Extr to fetch
// the piece of data.
func Info(unit string, request string) string {
	switch unit {
	case "c":
		return extr(c, request)
	case "f":
		return extr(f, request)
	case "hp":
		return extr(hp, request)
	case "w":
		return extr(w, request)
	case "kph":
		return extr(kph, request)
	case "mph":
		return extr(mph, request)
	case "kn":
		return extr(kn, request)
	case "mps":
		return extr(mps, request)
	case "ct":
		return extr(ct, request)
	case "g":
		return extr(g, request)
	case "oz":
		return extr(oz, request)
	case "lb":
		return extr(lb, request)
	case "st":
		return extr(st, request)
	}
	err := fmt.Errorf("the unit %s is cannot be found", unit)
	return fmt.Sprint(err)
}

// Extra extracts the data from the unit collection.
func extr(u unit, req string) string {
	switch req {
	case "nam":
		return u.name
	case "sym":
		return u.symbol
	}
	err := fmt.Errorf("the requested field %s does not exist", req)
	return fmt.Sprint(err)
}
