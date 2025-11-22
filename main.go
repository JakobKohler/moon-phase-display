package main

import (
	"fmt"
	"os"

	"github.com/JakobKohler/moon-phase-display/api"
	"github.com/JakobKohler/moon-phase-display/util"
)

func getMoonPhase(client api.Client) (string, error) {
	moonPhase, err := client.FetchMoonPhase()

	if err != nil {
		return "", err
	}

	outString := util.BuildPrettyMoonPhaseString(moonPhase)
	return outString, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: moon-phase-display <mode> [latitude] [longitude]")
		return
	}

	mode := os.Args[1]

	var latitude, longitude string

	if len(os.Args) > 2 {
		latitude = os.Args[2]
	}
	if len(os.Args) > 3 {
		longitude = os.Args[3]
	}

	client := api.NewClient()

	var outString string
	var err error

	switch mode {
	case "moon":
		outString, err = getMoonPhase(*client)
	case "sun":
		fmt.Printf("Unknown mode: %s\n%s\n", latitude, longitude)
	default:
		fmt.Printf("Unknown mode: %s\n", mode)
	}

	if err != nil {
		outString = ""
	}

	fmt.Printf(outString)
}
