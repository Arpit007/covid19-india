package helpers

import (
	"covid19-india/internal/config"
	"covid19-india/internal/models"
	"covid19-india/internal/utils"
	"errors"
	"fmt"
	"time"
)

func GetPlaceFromLatLng(lat float64, lng float64) (*models.GeoResponse, error) {
	client := utils.GetClient(time.Second * 5)

	res, err := client.Get(prepareRequestUrl(lat, lng))
	if err != nil {
		return nil, err
	}

	var data models.GeoResponse
	if err := utils.DecodeResponseBody(res, &data); err != nil {
		return nil, err
	}

	if len(data.Items) == 0 {
		return nil, errors.New("unable to decode location")
	}

	return &data, nil
}

func GetStateFromLatLong(lat float64, lng float64) (string, error) {
	place, err := GetPlaceFromLatLng(lat, lng)

	if err != nil {
		return "", err
	}

	state := place.Items[0].Address.State

	if len(state) == 0 {
		return "", errors.New("unable to identify state")
	}

	return state, nil
}

func prepareRequestUrl(lat float64, lng float64) string {
	return fmt.Sprintf("https://revgeocode.search.hereapi.com/v1/revgeocode?at=%f,%f&apikey=%s", lat, lng, config.ENV.HereMapsApiKey)
}
