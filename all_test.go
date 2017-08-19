package round

import (
	"testing"
	"math"
)

var vf = []float64{
	4.9790119248836735e+00,
	7.7388724745781045e+00,
	-2.7688005719200159e-01,
	-5.0106036182710749e+00,
	9.6362937071984173e+00,
	2.9263772392439646e+00,
	5.2290834314593066e+00,
	2.7279399104360102e+00,
	1.8253080916808550e+00,
	-8.6859247685756013e+00,
}

var round = []float64{
	5.0000000000000000e+00,
	8.0000000000000000e+00,
	-0.0000000000000000e+00,
	-5.0000000000000000e+00,
	1.0000000000000000e+01,
	3.0000000000000000e+00,
	5.0000000000000000e+00,
	3.0000000000000000e+00,
	2.0000000000000000e+00,
	-9.0000000000000000e+00,
}

var vfroundSC = [][2]float64{
	{0, 0},
	{1.390671161567e-309, 0}, // denormal
	{0.49999999999999994, 0}, // 0.5-epsilon
	{0.5, 1},
	{0.5000000000000001, 1}, // 0.5+epsilon
	{NaN(), NaN()},
	{Inf(1), Inf(1)},
	{2.2517998136852485e+15, 2.251799813685249e+15}, // 1 bit fraction
	{4.503599627370497e+15, 4.503599627370497e+15},  // large integer
}

func TestRound(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Round(vf[i]); round[i] != f {
			t.Errorf("Round(%g) = %g, want %g", vf[i], f, round[i])
		}
	}
	for i := 0; i < len(vfroundSC); i++ {
		if f := Round(vfroundSC[i][0]); !alike(vfroundSC[i][1], f) {
			t.Errorf("Round(%g) = %g, want %g", vfroundSC[i][0], f, vfroundSC[i][1])
		}
	}
}

func alike(a, b float64) bool {
	switch {
	case a != a && b != b: // math.IsNaN(a) && math.IsNaN(b):
		return true
	case a == b:
		return math.Signbit(a) == math.Signbit(b)
	}
	return false
}

// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN() float64 { return math.Float64frombits(uvnan) }

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf(sign int) float64 {
	var v uint64
	if sign >= 0 {
		v = uvinf
	} else {
		v = uvneginf
	}
	return math.Float64frombits(v)
}