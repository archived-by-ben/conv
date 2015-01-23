// Package units returns the name, symbols and categories of
// the units used in conversion.
package units

import (
	"fmt"
)

type unit struct {
	name, symbol string
}

// Unit data containing the name, symbol & category.
var (
	c  = unit{"Celsius", "°C"}
	f  = unit{"Fahrenheit", "°F"}
	hp = unit{"Horsepower", "hp"}
	w  = unit{"Watt", "W"}
	kw = unit{"kilowatt", "kW"}
	mw = unit{"megawatt", "MW"}
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
