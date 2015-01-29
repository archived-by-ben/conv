// Package units returns the name, symbols and categories of
// the units used in conversion.
package symbols

import (
	"fmt"
)

/*
TODO
* ounce
* pint
* quart
* pony
*/

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
	kg  = unit{"kilograms", "kg"}
	km  = unit{"kilometres", "km"}
	kn  = unit{"knots", "kn"}
	kmh = unit{"kilometres per hour", " km/h"}
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

// Info extracts data from the collection of unit data.
// Here it fetches the row of data then uses extr to fetch
// the piece of data.
func Info(unit string, req string) string {
	switch unit {
	case "bbl":
		return extr(bbl, req)
	case "c":
		return extr(c, req)
	case "cm":
		return extr(cm, req)
	case "ct":
		return extr(ct, req)
	case "f":
		return extr(f, req)
	case "ft":
		return extr(ft, req)
	case "g":
		return extr(g, req)
	case "guk":
		return extr(guk, req)
	case "gus":
		return extr(gus, req)
	case "km":
		return extr(km, req)
	case "hp":
		return extr(hp, req)
	case "kg":
		return extr(kg, req)
	case "kmh", "kph":
		return extr(kmh, req)
	case "kn":
		return extr(kn, req)
	case "in":
		return extr(in, req)
	case "l":
		return extr(l, req)
	case "lb":
		return extr(lb, req)
	case "m":
		return extr(m, req)
	case "cum":
		return extr(cum, req)
	case "mi":
		return extr(mi, req)
	case "mph", "mih":
		return extr(mph, req)
	case "mps":
		return extr(mps, req)
	case "nm":
		return extr(nm, req)
	case "st":
		return extr(st, req)
	case "oz":
		return extr(oz, req)
	case "yd":
		return extr(yd, req)
	case "w":
		return extr(w, req)
	}
	err := fmt.Errorf("the symbol %s cannot be found", unit)
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
