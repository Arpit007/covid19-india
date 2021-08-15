package helpers

import (
	"covid19-india/internal/config"
	"covid19-india/internal/models"
	"covid19-india/internal/utils"
	"errors"
	"fmt"
	"time"
)

func GetPlaceFromLatLng(lat float64, lng float64, out *models.GeoResponse) error {
	client := utils.GetClient(time.Second * 5)

	res, err := client.Get(prepareRequestUrl(lat, lng))
	if err != nil {
		return err
	}

	if err := utils.DecodeResponseBody(res, out); err != nil {
		return err
	}

	if len(out.Items) == 0 {
		return errors.New("unable to decode location")
	}

	return nil
}

func prepareRequestUrl(lat float64, lng float64) string {
	return fmt.Sprintf("https://revgeocode.search.hereapi.com/v1/revgeocode?at=%f,%f&apikey=%s", lat, lng, config.ENV.HereMapsApiKey)
}
