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
	hp  = unit{"horsepower", " hp"}
	w   = unit{"watts", " W"}
	kph = unit{"kilometres per hour", " km/h"}
	mph = unit{"miles per hour", " mph"}
	kn  = unit{"knots", "kn"}
	mps = unit{"metres per second", "m/s"}
	ct  = unit{"carats", " ct"}
	g   = unit{"grams", " g"}
	oz  = unit{"ounces", " oz"}
	lb  = unit{"pounds", " ℔"}
	st  = unit{"stones", " st"}
	cm  = unit{"centimetres", "cm"}
	in  = unit{"inches", "″"}
	yd  = unit{"yards", " yd"}
	ft  = unit{"feet", "′"}
	m   = unit{"metres", "m"}
	km  = unit{"kilometres", "km"}
	mi  = unit{"miles", " mi"}
	nm  = unit{"nautical miles", " NM"}
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
	case "cm":
		return extr(cm, request)
	case "in":
		return extr(in, request)
	case "yd":
		return extr(yd, request)
	case "ft":
		return extr(ft, request)
	case "m":
		return extr(m, request)
	case "km":
		return extr(km, request)
	case "mi":
		return extr(mi, request)
	case "nm":
		return extr(nm, request)
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
