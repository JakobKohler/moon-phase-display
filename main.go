package main

import (
	"log"

	"github.com/JakobKohler/moon-phase-display/api"
)

func main() {
	client := api.NewClient()
	moonPhase, err := client.FetchMoonPhase()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(moonPhase)
}
