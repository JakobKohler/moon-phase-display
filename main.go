package main

import (
	"fmt"
	"log"

	"github.com/JakobKohler/moon-phase-display/api"
	"github.com/JakobKohler/moon-phase-display/util"
)

func main() {
	client := api.NewClient()
	moonPhase, err := client.FetchMoonPhase()

	if err != nil {
		log.Fatalln(err)
	}

	outString := util.BuildPrettyMoonPhaseString(moonPhase)
	fmt.Printf(outString)
}
