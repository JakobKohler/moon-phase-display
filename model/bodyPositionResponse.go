package model

import "time"

type Meta struct {
	Time          time.Time `json:"time"`
	EngineVersion string    `json:"engineVersion"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	Elevation     int       `json:"elevation"`
	AboveHorizon  bool      `json:"aboveHorizon"`
}

type BodyInfo struct {
	Name           string `json:"name"`
	Constellation  string `json:"constellation"`
	RightAscension struct {
		Negative bool    `json:"negative"`
		Hours    int     `json:"hours"`
		Minutes  int     `json:"minutes"`
		Seconds  float64 `json:"seconds"`
		Raw      float64 `json:"raw"`
	} `json:"rightAscension"`
	Declination struct {
		Negative   bool    `json:"negative"`
		Degrees    int     `json:"degrees"`
		Arcminutes int     `json:"arcminutes"`
		Arcseconds float64 `json:"arcseconds"`
		Raw        float64 `json:"raw"`
	} `json:"declination"`
	Altitude       float64 `json:"altitude"`
	Azimuth        float64 `json:"azimuth"`
	AboveHorizon   bool    `json:"aboveHorizon"`
	Magnitude      float64 `json:"magnitude"`
	NakedEyeObject bool    `json:"nakedEyeObject"`
	Phase          float64 `json:"phase,omitempty"`
}

type Data []BodyInfo

type Links struct {
	Self   string `json:"self"`
	Engine string `json:"engine"`
}

type BodyPositionResponse struct {
	Meta  `json:"meta"`
	Data  `json:"data"`
	Links `json:"links"`
}
