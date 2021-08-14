package helpers

import (
	. "covid19-india/configs"
	. "covid19-india/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type GeoResponse struct {
	Items []GeoPlace `json:"items"`
}

func GetPlaceFromLatLng(lat float64, lng float64) (*GeoResponse, error) {
	client := &http.Client{Timeout: time.Second * 5}

	res, err := client.Get(prepareRequestUrl(lat, lng))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var geoResponse GeoResponse
	if err := json.NewDecoder(res.Body).Decode(&geoResponse); err != nil {
		return nil, err
	}

	if len(geoResponse.Items) == 0 {
		return nil, errors.New("unable to decode location")
	}

	return &geoResponse, nil
}

func prepareRequestUrl(lat float64, lng float64) string {
	return fmt.Sprintf("https://revgeocode.search.hereapi.com/v1/revgeocode?at=%f,%f&apikey=%s", lat, lng, ENV.HereMapsApiKey)
}
