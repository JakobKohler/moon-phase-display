package util

import (
	"math"
)

type MoonPhase struct {
	Icons []string
}

func NewMoonPhase() *MoonPhase {
	return &MoonPhase{
		Icons: []string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	}
}

func (m *MoonPhase) GetIcon(phase float64) string {
	iconIdx := int(math.Round(float64(len(m.Icons))*phase)) - 1
	if iconIdx < 0 {
		iconIdx = 0
	}
	return m.Icons[iconIdx]
}

func (m *MoonPhase) isFullMoon(phase float64) bool {
	iconIdx := int(math.Round(float64(len(m.Icons)) * phase))
	return iconIdx == 14 //Full moon icon is displayed
}
