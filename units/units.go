package units

// name, argument, symbol, quantity
type unit struct {
	n, a, s, q string
}

var (
	request = [4]string{"nam", "arg", "sym", "qty"}
	cel     = unit{"Celsius", "C", "°C", "temperature"}
	fah     = unit{"Fahrenheit", "F", "°F", "temperature"}
)

func Info(u string, req string) string {
	switch {
	case u == "cel":
		return extractdata(cel, req)
	case u == "fah":
		return extractdata(fah, req)
	}
	return ""
}

func extractdata(u unit, req string) string {
	switch {
	case req == "nam":
		return name(u)
	case req == "arg":
		return argument(u)
	case req == "sym":
		return symbol(u)
	case req == "qty":
		return quantity(u)
	}
	return ""
}

func name(u unit) string {
	return u.n
}

func argument(u unit) string {
	return u.a
}

func symbol(u unit) string {
	return u.s
}

func quantity(u unit) string {
	return u.q
}
