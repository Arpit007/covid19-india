package cache

import (
	"context"
	"covid19-india/internal/helpers"
	"errors"
	"fmt"
	"time"
)

var geoCache *RedisCache

func init() {
	geoCache = CreateRedisCache(time.Minute*30, "geo")
}

type geoState map[string]string

func GetStateFromLatLong(lat float64, lng float64) (string, error) {
	key := fmt.Sprintf("%f,%f", lat, lng)

	res, err := geoCache.Get(context.TODO(), key, &geoState{}, func() (interface{}, error) {
		if state, err := helpers.GetStateFromLatLong(lat, lng); err != nil {
			return nil, err
		} else {
			return &geoState{"State": state}, nil
		}
	})

	if err != nil {
		return "", err
	}

	if res == nil {
		return "", errors.New("unable to decode location")
	}

	var data = *res.(*geoState)

	return data["State"], nil
}
