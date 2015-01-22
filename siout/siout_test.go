package siout

import "testing"

func TestFahrenheit(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{100, 310.9277777777778},
		{0, 255.3722222222222},
		{-540, -44.627777777777766},
	}
	for _, c := range cases {
		got := Fahrenheit(c.in)
		if got != c.want {
			t.Errorf("Fahrenheit(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCelsius(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{300, 573.15},
		{0, 273.15},
		{-494, -220.85000000000002},
	}
	for _, c := range cases {
		got := Celsius(c.in)
		if got != c.want {
			t.Errorf("Celsius(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
