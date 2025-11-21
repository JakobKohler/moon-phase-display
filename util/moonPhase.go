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
	iconIdx := int(math.Round(float64(len(m.Icons)) * phase))
	return m.Icons[iconIdx]
}
