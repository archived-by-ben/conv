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
	bbl = unit{"petroleum barrel", "bbl"}
	c   = unit{"Celsius", "°C"}
	cm  = unit{"centimetres", "cm"}
	ct  = unit{"carats", " ct"}
	cum = unit{"cubic metre", "m³"}
	f   = unit{"Fahrenheit", "°F"}
	ft  = unit{"feet", "′"}
	g   = unit{"grams", "g"}
	guk = unit{"imperial gallon", "gal"}
	gus = unit{"us gallon", "gal"}
	hp  = unit{"horsepower", " hp"}
	in  = unit{"inches", "″"}
	km  = unit{"kilometres", "km"}
	kn  = unit{"knots", "kn"}
	kph = unit{"kilometres per hour", " km/h"}
	l   = unit{"litres", "L"}
	lb  = unit{"pounds", " ℔"}
	m   = unit{"metres", "m"}
	mi  = unit{"miles", " mi"}
	mph = unit{"miles per hour", " mph"}
	mps = unit{"metres per second", "m/s"}
	nm  = unit{"nautical miles", " NM"}
	oz  = unit{"ounces", " oz"}
	st  = unit{"stone", " st"}
	w   = unit{"watts", " W"}
	yd  = unit{"yards", " yd"}
)

/*
* litre
* ounce
* pint
* quart
* pony
 */

// Info extracts data from the collection of unit data.
// Here it fetches the row of data then uses extr to fetch
// the piece of data.
func Info(unit string, request string) string {
	switch unit {
	case "bbl":
		return extr(bbl, request)
	case "c":
		return extr(c, request)
	case "cm":
		return extr(cm, request)
	case "ct":
		return extr(ct, request)
	case "f":
		return extr(f, request)
	case "ft":
		return extr(ft, request)
	case "g":
		return extr(g, request)
	case "guk":
		return extr(guk, request)
	case "gus":
		return extr(gus, request)
	case "km":
		return extr(km, request)
	case "hp":
		return extr(hp, request)
	case "kph":
		return extr(kph, request)
	case "kn":
		return extr(kn, request)
	case "in":
		return extr(in, request)
	case "l":
		return extr(l, request)
	case "lb":
		return extr(lb, request)
	case "m":
		return extr(m, request)
	case "cum":
		return extr(cum, request)
	case "mi":
		return extr(mi, request)
	case "mph":
		return extr(mph, request)
	case "mps":
		return extr(mps, request)
	case "nm":
		return extr(nm, request)
	case "st":
		return extr(st, request)
	case "oz":
		return extr(oz, request)
	case "yd":
		return extr(yd, request)
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
