package cache

import (
	"context"
	helper3p "covid19-india/internal/third_party_helpers"
	"fmt"
	"time"
)

var geoCache *RedisCache

func init() {
	geoCache = CreateRedisCache(30*time.Minute, "geo")
}

type geoState map[string]string

// GetStateFromLatLong Fetches user's state (from geo coordinates) from cache
func GetStateFromLatLong(lat float64, lng float64) (string, error) {
	key := fmt.Sprintf("%f,%f", lat, lng)

	res, err := geoCache.Get(context.TODO(), key, &geoState{}, func() (interface{}, error) {
		// cache miss
		if state, err := helper3p.GetStateFromLatLong(lat, lng); err != nil {
			return nil, err
		} else {
			return &geoState{"State": state}, nil
		}
	})

	if err != nil {
		return "", err
	}

	var data = *res.(*geoState)

	return data["State"], nil
}
