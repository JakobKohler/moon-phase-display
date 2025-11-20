package util

import (
	"math"
)

type MoonPhase struct {
	Icons map[float64]string
}

func NewMoonPhase() *MoonPhase {
	return &MoonPhase{
		Icons: map[float64]string{
			0.000: " ",
			0.037: " ",
			0.074: " ",
			0.111: " ",
			0.148: " ",
			0.185: " ",
			0.222: " ",
			0.259: " ",
			0.296: " ",
			0.333: " ",
			0.370: " ",
			0.407: " ",
			0.444: " ",
			0.481: " ",
			0.518: " ",
			0.555: " ",
			0.592: " ",
			0.629: " ",
			0.666: " ",
			0.703: " ",
			0.740: " ",
			0.777: " ",
			0.814: " ",
			0.851: " ",
			0.888: " ",
			0.925: " ",
			0.962: " ",
			1.000: " ",
		},
	}
}

func (m *MoonPhase) GetIcon(phase float64) string {
	if phase < 0 {
		phase = 0
	} else if phase > 1 {
		phase = 1
	}

	var closest float64
	var minDiff float64 = 1.0

	for k := range m.Icons {
		diff := math.Abs(k - phase)
		if diff < minDiff {
			minDiff = diff
			closest = k
		}
	}
	return m.Icons[closest]
}
