package sun

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SunTimesResponse struct {
	Status  string       `json:"status"`
	Results SunTimesData `json:"results"`
}

type SunTimesData struct {
	Date       string `json:"date"`
	Sunrise    string `json:"sunrise"`
	Sunset     string `json:"sunset"`
	FirstLight string `json:"first_light"`
	LastLight  string `json:"last_light"`
	Dawn       string `json:"dawn"`
	Dusk       string `json:"dusk"`
	SolarNoon  string `json:"solar_noon"`
	GoldenHour string `json:"golden_hour"`
	DayLength  string `json:"day_length"`
	Timezone   string `json:"timezone"`
	UTCOffset  int    `json:"utc_offset"`
}

func FetchSunTimes(date string, latitude string, longitude string) (*SunTimesData, error) {
	url := "https://api.sunrisesunset.io/json?date=" + date + "&lat=" + latitude + "&lng=" + longitude
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var sunTimes SunTimesResponse
	if err := json.Unmarshal(body, &sunTimes); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &sunTimes.Results, nil
}
