package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"slices"

	"github.com/JakobKohler/moon-phase-display/model"
)

func buildUrl(apiBase string, lat string, lon string) string {
	endpoint, err := url.Parse(apiBase)

	if err != nil {
		log.Fatal(err)
	}

	queryParams := url.Values{}
	queryParams.Set("latitude", lat)
	queryParams.Set("longitude", lon)
	queryParams.Set("aboveHorizon", "false")

	endpoint.RawQuery = queryParams.Encode()

	return endpoint.String()
}

func fetchBodyPositions(apiBase string, lat string, lon string) (model.BodyPositionResponse, error) {
	url := buildUrl(apiBase, lat, lon)

	data, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res model.BodyPositionResponse

	if err := json.Unmarshal(body, &res); err != nil {
		log.Fatal(err)
	}

	return res, nil
}

func (c *Client) FetchMoonPhase() (float64, error) {
	data, err := fetchBodyPositions(c.BaseURL, "0", "0")

	idx := slices.IndexFunc(data.Data, func(b model.BodyInfo) bool { return b.Name == "Moon" })

	if err != nil || idx < 0 {
		return -1, err
	}

	moonPhase := data.Data[idx].Phase
	moonPhase /= 100

	return moonPhase, nil
}
