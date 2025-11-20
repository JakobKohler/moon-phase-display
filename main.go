package main

import (
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

	log.Println(moonPhase)
	outString := util.BuildPrettyMoonPhaseString(moonPhase)
	log.Println(outString)
}
