package cache

import (
	"context"
	"covid19-india/internal/dao"
	"covid19-india/internal/models"
	"errors"
	"time"
)

var covidDataCache *RedisCache

func init() {
	covidDataCache = CreateRedisCache(30*time.Minute, "covIn")
}

// GetCovidDataForRegion Get covid data for a region from cache
func GetCovidDataForRegion(ctx context.Context, region string) (*models.CovidData, error) {
	if len(region) == 0 {
		return nil, nil
	}

	res, err := covidDataCache.Get(ctx, region, &models.CovidData{}, func(ctx context.Context) (interface{}, error) {
		// cache miss
		return dao.GetCovidDataForRegion(ctx, region)
	})

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.New("unable to fetch covid data")
	}

	data := res.(*models.CovidData)

	return data, nil
}

// GetCovidDataForRegions Get covid data for regions from cache
func GetCovidDataForRegions(ctx context.Context, id []string) ([]models.CovidData, error) {
	var covidData []models.CovidData

	for _, region := range id {
		if data, err := GetCovidDataForRegion(ctx, region); err != nil {
			return nil, err
		} else {
			covidData = append(covidData, *data)
		}
	}

	return covidData, nil
}

// ResetCovidDataCache Reset covid data cache
func ResetCovidDataCache(ctx context.Context, covidData []models.CovidData) error {
	var keys []string

	for _, data := range covidData {
		keys = append(keys, data.Region)
	}

	return covidDataCache.RemoveKeys(ctx, keys)
}
