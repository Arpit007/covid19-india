package cache

import (
	"context"
	"covid19-india/internal/helpers"
	"covid19-india/internal/models"
	"errors"
	"fmt"
	"time"
)

var rCache *RedisCache

func init() {
	rCache = CreateRedisCache(time.Minute*30, "geo")
}

func GetPlaceFromLatLng(lat float64, lng float64, out *models.GeoResponse) (*models.GeoResponse, error) {
	key := fmt.Sprintf("%f,%f", lat, lng)

	err := rCache.Get(context.TODO(), key, out, func(out interface{}) error {
		return helpers.GetPlaceFromLatLng(lat, lng, out.(*models.GeoResponse))
	})

	if err != nil {
		return nil, err
	}

	if len(out.Items) == 0 {
		return nil, errors.New("unable to decode location")
	} else {
		return out, nil
	}
}
