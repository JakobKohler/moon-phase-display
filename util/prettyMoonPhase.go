package util

const starsColor = "e8c55c"

func BuildPrettyMoonPhaseString(moonPhase float64) string {
	moonPhaseUtil := NewMoonPhase()
	currentPhaseIcon := moonPhaseUtil.GetIcon(moonPhase)

	if moonPhaseUtil.isFullMoon(moonPhase) {
		return "%%{F#" + starsColor + "}\uE370%%{F-} %%{F}" + currentPhaseIcon + "%%{F-}%%{F#" + starsColor + "}\uE370"
	}

	return currentPhaseIcon
}
