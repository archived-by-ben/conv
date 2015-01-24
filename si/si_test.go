package si

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

func TestKmh(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{10, 35.99999999999971},
	}
	for _, c := range cases {
		got := Kmh(c.in)
		if got != c.want {
			t.Errorf("Kmh(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMph(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{10, 22.369362920544024},
	}
	for _, c := range cases {
		got := Mph(c.in)
		if got != c.want {
			t.Errorf("Mph(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKn(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{10, 19.438461717893492},
	}
	for _, c := range cases {
		got := Kn(c.in)
		if got != c.want {
			t.Errorf("Kn(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMps(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 1},
	}
	for _, c := range cases {
		got := Mps(c.in)
		if got != c.want {
			t.Errorf("Mps(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCarat(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{2, 10},
	}
	for _, c := range cases {
		got := Carat(c.in)
		if got != c.want {
			t.Errorf("Carat(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestOunce(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 28},
	}
	for _, c := range cases {
		got := Ounce(c.in)
		if got != c.want {
			t.Errorf("Ounce(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPound(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 453.59237},
	}
	for _, c := range cases {
		got := Pound(c.in)
		if got != c.want {
			t.Errorf("Pound(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestStone(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 6350.29318},
	}
	for _, c := range cases {
		got := Stone(c.in)
		if got != c.want {
			t.Errorf("Stone(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
