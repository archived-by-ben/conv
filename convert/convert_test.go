package convert

import "testing"

func TestRound(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{100, "100"},
		{1.1, "1.1"},
		{1.25, "1.2"},
		{64.4, "64.4"},
		{81.0, "81"},
		{81.1, "81.1"},
		{81.000000000001, "81"},
		{81.01, "81"},
		{839.87365, "839.9"},
		{-839.87365, "-839.9"},
		{1.39700, "1.4"},
	}
	for _, c := range cases {
		got := Round(c.in)
		if got != c.want {
			t.Errorf("Round(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestBblcum(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10, "1.6"},
	}
	for _, c := range cases {
		got := Bblcum(c.in)
		if got != c.want {
			t.Errorf("Bblcum(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestBbll(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "159"},
	}
	for _, c := range cases {
		got := Bbll(c.in)
		if got != c.want {
			t.Errorf("Bbll(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestBblguk(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "35"},
	}
	for _, c := range cases {
		got := Bblguk(c.in)
		if got != c.want {
			t.Errorf("Bblguk(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestBblgus(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "42"},
	}
	for _, c := range cases {
		got := Bblgus(c.in)
		if got != c.want {
			t.Errorf("Bblgus(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCf(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{0, "32"},
		{100, "212"},
		{-25, "-13"},
	}
	for _, c := range cases {
		got := Cf(c.in)
		if got != c.want {
			t.Errorf("Cf(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCmin(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0.4"},
		{100, "39.4"},
		{111.1, "43.7"},
	}
	for _, c := range cases {
		got := Cmin(c.in)
		if got != c.want {
			t.Errorf("Cmin(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCmft(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{100, "3.3"},
		{111.1, "3.6"},
	}
	for _, c := range cases {
		got := Cmft(c.in)
		if got != c.want {
			t.Errorf("Cmft(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCmmi(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{111111.1, "0.7"},
		{10000000, "62.1"},
	}
	for _, c := range cases {
		got := Cmmi(c.in)
		if got != c.want {
			t.Errorf("Cmmi(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCmnm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{111111.1, "0.6"},
		{10000000, "54"},
	}
	for _, c := range cases {
		got := Cmnm(c.in)
		if got != c.want {
			t.Errorf("Cmnm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCmyd(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{100, "1"},
		{111.1, "1.2"},
	}
	for _, c := range cases {
		got := Cmyd(c.in)
		if got != c.want {
			t.Errorf("Cmyd(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCumbbl(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "6.3"},
	}
	for _, c := range cases {
		got := Cumbbl(c.in)
		if got != c.want {
			t.Errorf("Cumbbl(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCuml(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "1000"},
	}
	for _, c := range cases {
		got := Cuml(c.in)
		if got != c.want {
			t.Errorf("Cuml(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCumguk(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "220"},
	}
	for _, c := range cases {
		got := Cumguk(c.in)
		if got != c.want {
			t.Errorf("Cumguk(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCumgus(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "264.2"},
	}
	for _, c := range cases {
		got := Cumgus(c.in)
		if got != c.want {
			t.Errorf("Cumgus(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCtg(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0.2"},
		{5.5, "1.1"},
	}
	for _, c := range cases {
		got := Ctg(c.in)
		if got != c.want {
			t.Errorf("Ctg(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFc(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{32, "0"},
		{100, "37.8"},
		{-10, "-23.3"},
	}
	for _, c := range cases {
		got := Fc(c.in)
		if got != c.want {
			t.Errorf("Fc(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

///

func TestFtcm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "30.5"},
	}
	for _, c := range cases {
		got := Ftcm(c.in)
		if got != c.want {
			t.Errorf("Ftcm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFtin(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "12"},
	}
	for _, c := range cases {
		got := Ftin(c.in)
		if got != c.want {
			t.Errorf("Ftin(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFtkm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10000, "3"},
	}
	for _, c := range cases {
		got := Ftkm(c.in)
		if got != c.want {
			t.Errorf("Ftkm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFtm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10, "3"},
	}
	for _, c := range cases {
		got := Ftm(c.in)
		if got != c.want {
			t.Errorf("Ftm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFtmi(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10000, "1.9"},
	}
	for _, c := range cases {
		got := Ftmi(c.in)
		if got != c.want {
			t.Errorf("Ftmi(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFtnm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10000, "1.6"},
	}
	for _, c := range cases {
		got := Ftnm(c.in)
		if got != c.want {
			t.Errorf("Ftnm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

///
func TestGlb(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{1000, "2.2"},
		{555.5, "1.2"},
	}
	for _, c := range cases {
		got := Glb(c.in)
		if got != c.want {
			t.Errorf("Glb(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGoz(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{10, "0.4"},
		{55.5, "2"},
	}
	for _, c := range cases {
		got := Goz(c.in)
		if got != c.want {
			t.Errorf("Goz(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGst(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{1000, "0.2"},
		{55555.5, "8.7"},
	}
	for _, c := range cases {
		got := Gst(c.in)
		if got != c.want {
			t.Errorf("Gst(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGukbbl(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{100, "2.9"},
	}
	for _, c := range cases {
		got := Gukbbl(c.in)
		if got != c.want {
			t.Errorf("Gukbbl(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGukcum(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1000, "4.5"},
	}
	for _, c := range cases {
		got := Gukcum(c.in)
		if got != c.want {
			t.Errorf("Gukcum(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGukl(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "4.5"},
	}
	for _, c := range cases {
		got := Gukl(c.in)
		if got != c.want {
			t.Errorf("Gukl(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGukgus(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "1.2"},
	}
	for _, c := range cases {
		got := Gukgus(c.in)
		if got != c.want {
			t.Errorf("Gukgus(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGusbbl(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{100, "2.4"},
	}
	for _, c := range cases {
		got := Gusbbl(c.in)
		if got != c.want {
			t.Errorf("Gusbbl(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGuscum(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1000, "3.8"},
	}
	for _, c := range cases {
		got := Guscum(c.in)
		if got != c.want {
			t.Errorf("Guscum(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGusl(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "3.8"},
	}
	for _, c := range cases {
		got := Gusl(c.in)
		if got != c.want {
			t.Errorf("Gusl(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGusguk(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0.8"},
	}
	for _, c := range cases {
		got := Gusguk(c.in)
		if got != c.want {
			t.Errorf("Gusguk(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestHpw(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "746"},
		{10, "7460"},
		{5.5, "4103"},
	}
	for _, c := range cases {
		got := Hpw(c.in)
		if got != c.want {
			t.Errorf("Hpw(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIncm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "2.5"},
		{10, "25.4"},
		{5.5, "14"},
	}
	for _, c := range cases {
		got := Incm(c.in)
		if got != c.want {
			t.Errorf("Incm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestInft(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{100, "8.3"},
	}
	for _, c := range cases {
		got := Inft(c.in)
		if got != c.want {
			t.Errorf("Inft(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestInnm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{100000, "1.4"},
	}
	for _, c := range cases {
		got := Innm(c.in)
		if got != c.want {
			t.Errorf("Innm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestInyd(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{100, "2.8"},
	}
	for _, c := range cases {
		got := Inyd(c.in)
		if got != c.want {
			t.Errorf("Inyd(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestInmi(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{100000, "1.6"},
	}
	for _, c := range cases {
		got := Inmi(c.in)
		if got != c.want {
			t.Errorf("Inmi(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKmhkn(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0.5"},
		{100, "54"},
		{50.5, "27.3"},
	}
	for _, c := range cases {
		got := Kmhkn(c.in)
		if got != c.want {
			t.Errorf("Kmhkn(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKmhmph(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0.6"},
		{100, "62.1"},
		{50.5, "31.4"},
	}
	for _, c := range cases {
		got := Kmhmph(c.in)
		if got != c.want {
			t.Errorf("Kmhmph(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKmhmps(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0.3"},
		{100, "27.8"},
		{50.5, "14"},
	}
	for _, c := range cases {
		got := Kmhmps(c.in)
		if got != c.want {
			t.Errorf("Kmhmps(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKnkmh(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "1.9"},
		{10, "18.5"},
	}
	for _, c := range cases {
		got := Knkmh(c.in)
		if got != c.want {
			t.Errorf("Knkmh(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKnmps(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0.5"},
		{10, "5.1"},
	}
	for _, c := range cases {
		got := Knmps(c.in)
		if got != c.want {
			t.Errorf("Kmhmps(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKnmph(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "1.2"},
		{10, "11.5"},
		{100, "115"},
		{50.5, "58.1"},
	}
	for _, c := range cases {
		got := Knmph(c.in)
		if got != c.want {
			t.Errorf("Kmhmph(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestLbg(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "453.6"},
		{1.1, "499"},
	}
	for _, c := range cases {
		got := Lbg(c.in)
		if got != c.want {
			t.Errorf("Lbg(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestLboz(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "16"},
		{1.1, "17.6"},
	}
	for _, c := range cases {
		got := Lboz(c.in)
		if got != c.want {
			t.Errorf("Lboz(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestLbst(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{111.5, "8"},
	}
	for _, c := range cases {
		got := Lbst(c.in)
		if got != c.want {
			t.Errorf("Lbst(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMicm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "160934.4"},
	}
	for _, c := range cases {
		got := Micm(c.in)
		if got != c.want {
			t.Errorf("Micm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMift(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "5280"},
	}
	for _, c := range cases {
		got := Mift(c.in)
		if got != c.want {
			t.Errorf("Mift(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMiin(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "63360"},
	}
	for _, c := range cases {
		got := Miin(c.in)
		if got != c.want {
			t.Errorf("Miin(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMikm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "1.6"},
	}
	for _, c := range cases {
		got := Mikm(c.in)
		if got != c.want {
			t.Errorf("Mikm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMim(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "1609.3"},
	}
	for _, c := range cases {
		got := Mim(c.in)
		if got != c.want {
			t.Errorf("Mim(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMinm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10, "8.7"},
	}
	for _, c := range cases {
		got := Minm(c.in)
		if got != c.want {
			t.Errorf("Minm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMiyd(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "1760"},
	}
	for _, c := range cases {
		got := Miyd(c.in)
		if got != c.want {
			t.Errorf("Miyd(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMphkmh(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "1.6"},
		{40, "64.4"},
		{5.5, "8.9"},
	}
	for _, c := range cases {
		got := Mphkmh(c.in)
		if got != c.want {
			t.Errorf("Mphkmh(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMphmps(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0.4"},
		{100, "44.7"},
		{55.5, "24.8"},
	}
	for _, c := range cases {
		got := Mphmps(c.in)
		if got != c.want {
			t.Errorf("Mphmps(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMphkn(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0.9"},
		{100, "86.9"},
	}
	for _, c := range cases {
		got := Mphkn(c.in)
		if got != c.want {
			t.Errorf("Mphkn(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMpskmh(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10, "36"},
		{100.1, "360.4"},
	}
	for _, c := range cases {
		got := Mpskmh(c.in)
		if got != c.want {
			t.Errorf("Mpskmh(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMpsmph(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10, "22.4"},
		{100.1, "224"},
	}
	for _, c := range cases {
		got := Mpsmph(c.in)
		if got != c.want {
			t.Errorf("Mpsmph(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMpskn(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10, "19.4"},
		{55.5, "107.9"},
	}
	for _, c := range cases {
		got := Mpskn(c.in)
		if got != c.want {
			t.Errorf("Mpskn(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNmcm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "185200"},
	}
	for _, c := range cases {
		got := Nmcm(c.in)
		if got != c.want {
			t.Errorf("Nmcm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNmft(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "6076.1"},
	}
	for _, c := range cases {
		got := Nmft(c.in)
		if got != c.want {
			t.Errorf("Nmft(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNmin(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "72913.4"},
	}
	for _, c := range cases {
		got := Nmin(c.in)
		if got != c.want {
			t.Errorf("Nmin(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNmkm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "1.9"},
	}
	for _, c := range cases {
		got := Nmkm(c.in)
		if got != c.want {
			t.Errorf("Nmkm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNmm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "1852"},
	}
	for _, c := range cases {
		got := Nmm(c.in)
		if got != c.want {
			t.Errorf("Nmm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNmmi(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "1.2"},
	}
	for _, c := range cases {
		got := Nmmi(c.in)
		if got != c.want {
			t.Errorf("Nmmi(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNmyd(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "2025.4"},
	}
	for _, c := range cases {
		got := Nmyd(c.in)
		if got != c.want {
			t.Errorf("Nmyd(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestOzg(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "28.3"},
		{50, "1417.5"},
		{55.5, "1573.4"},
	}
	for _, c := range cases {
		got := Ozg(c.in)
		if got != c.want {
			t.Errorf("Ozg(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestOzlb(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{100, "6.2"},
		{111.1, "7"}, // expected 6.9?
	}
	for _, c := range cases {
		got := Ozlb(c.in)
		if got != c.want {
			t.Errorf("Ozlb(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestOzst(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1000, "4.5"},
		{555.5, "2.5"},
	}
	for _, c := range cases {
		got := Ozst(c.in)
		if got != c.want {
			t.Errorf("Ozst(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestStg(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "6350.3"},
		{1.1, "6985.3"},
	}
	for _, c := range cases {
		got := Stg(c.in)
		if got != c.want {
			t.Errorf("Stg(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestStlb(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "14"},
		{1.1, "15.4"},
	}
	for _, c := range cases {
		got := Stlb(c.in)
		if got != c.want {
			t.Errorf("Stlb(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestStoz(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "224"},
		{1.1, "246.4"},
	}
	for _, c := range cases {
		got := Stoz(c.in)
		if got != c.want {
			t.Errorf("Stoz(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestWhp(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "0"},
		{1000, "1.3"},
	}
	for _, c := range cases {
		got := Whp(c.in)
		if got != c.want {
			t.Errorf("Whp(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestYdcm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "91.4"},
	}
	for _, c := range cases {
		got := Ydcm(c.in)
		if got != c.want {
			t.Errorf("Ydcm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestYdft(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "3"},
	}
	for _, c := range cases {
		got := Ydft(c.in)
		if got != c.want {
			t.Errorf("Ydft(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestYdin(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{1, "36"},
	}
	for _, c := range cases {
		got := Ydin(c.in)
		if got != c.want {
			t.Errorf("Ydin(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestYdkm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10000, "9.1"},
	}
	for _, c := range cases {
		got := Ydkm(c.in)
		if got != c.want {
			t.Errorf("Ydkm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestYdm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10000, "9144"},
	}
	for _, c := range cases {
		got := Ydm(c.in)
		if got != c.want {
			t.Errorf("Ydm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestYdmi(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10000, "5.7"},
	}
	for _, c := range cases {
		got := Ydmi(c.in)
		if got != c.want {
			t.Errorf("Ydmi(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestYdnm(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{10000, "5"}, // should be 4.9?
	}
	for _, c := range cases {
		got := Ydnm(c.in)
		if got != c.want {
			t.Errorf("Ydnm(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
