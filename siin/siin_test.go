package siin

import "testing"

func TestFahrenheit(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{100, -279.67},
		{0, -459.67},
		{-44, -538.87},
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
		{300, 26.850000000000023},
		{0, -273.15},
		{-44, -317.15},
	}
	for _, c := range cases {
		got := Celsius(c.in)
		if got != c.want {
			t.Errorf("Celsius(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestHorsepower(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{746, 1},
		{1000.1, 1.3406166219839142},
	}
	for _, c := range cases {
		got := Horsepower(c.in)
		if got != c.want {
			t.Errorf("Horsepower(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestWatt(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 1},
		{100.1, 100.1},
	}
	for _, c := range cases {
		got := Watt(c.in)
		if got != c.want {
			t.Errorf("Watt(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
