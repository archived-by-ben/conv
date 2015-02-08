// symbols package returns the name and symbols for the measurements.
package symbols

import (
	"fmt"
	"runtime"
)

type Unit struct {
	Name, Symbol string
}

func UnitData() (unitDetails map[string]Unit) {
	// stored data as a structure indexed as a key value map
	unitDetails = make(map[string]Unit)
	unitDetails["bbl"] = Unit{"petroleum barrel", "bbl"}
	unitDetails["c"] = Unit{"Celsius", "°C"}
	unitDetails["cm"] = Unit{"centimetres", "cm"}
	unitDetails["ct"] = Unit{"carats", " ct"}
	unitDetails["cum"] = Unit{"cubic metre", "m³"}
	unitDetails["f"] = Unit{"Fahrenheit", "°F"}
	unitDetails["ft"] = Unit{"feet", "′"}
	unitDetails["g"] = Unit{"grams", "g"}
	unitDetails["guk"] = Unit{"imperial gallon", "gal"}
	unitDetails["gus"] = Unit{"us gallon", "gal"}
	unitDetails["hp"] = Unit{"horsepower", " hp"}
	unitDetails["in"] = Unit{"inches", "″"}
	unitDetails["kg"] = Unit{"kilograms", "kg"}
	unitDetails["km"] = Unit{"kilometres", "km"}
	unitDetails["kmh"] = Unit{"kilometres per hour", " km/h"}
	unitDetails["kn"] = Unit{"knots", "kn"}
	unitDetails["l"] = Unit{"litres", "L"}
	unitDetails["lb"] = Unit{"pounds", " ℔"}
	unitDetails["m"] = Unit{"metres", "m"}
	unitDetails["mi"] = Unit{"miles", " mi"}
	unitDetails["mph"] = Unit{"miles per hour", " mph"}
	unitDetails["mps"] = Unit{"metres per second", "m/s"}
	unitDetails["nm"] = Unit{"nautical miles", " NM"}
	unitDetails["oz"] = Unit{"ounces", " oz"}
	unitDetails["st"] = Unit{"stone", " st"}
	unitDetails["w"] = Unit{"watts", " W"}
	unitDetails["yd"] = Unit{"yards", " yd"}
	// fixes for Windows Command Prompt font limitations
	if runtime.GOOS == "windows" {
		unitDetails["ft"] = Unit{"feet", "ft"}
		unitDetails["in"] = Unit{"inches", "in"}
		unitDetails["lb"] = Unit{"pounds", " lb"}
	}
	// display all data
	return unitDetails
}

func Glyph(Unit string) string {
	detail, found := UnitData()[Unit]
	if found == false {
		return fmt.Sprint("glyph(\"%v\") does not exist", Unit)
	}
	return detail.Symbol
}

func Proper(Unit string) string {
	detail, found := UnitData()[Unit]
	if found == false {
		return fmt.Sprint("proper(\"%v\") does not exist", Unit)
	}
	return detail.Name
}
