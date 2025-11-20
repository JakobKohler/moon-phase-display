package util

func BuildPrettyMoonPhaseString(moonPhase float64) string {
	moonPhaseUtil := NewMoonPhase()

	currentPhaseIcon := moonPhaseUtil.GetIcon(moonPhase)

	return currentPhaseIcon
}
