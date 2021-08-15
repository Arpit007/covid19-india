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
	covidDataCache = CreateRedisCache(time.Minute*30, "covIn")
}

func GetCovidDataForRegion(region string) (*models.CovidData, error) {
	if len(region) == 0 {
		return nil, nil
	}

	res, err := covidDataCache.Get(context.TODO(), region, &models.CovidData{}, func() (interface{}, error) {
		return dao.GetCovidDataForRegion(region)
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

func GetCovidDataForRegions(id []string) ([]models.CovidData, error) {
	var covidData []models.CovidData

	for _, region := range id {
		if data, err := GetCovidDataForRegion(region); err != nil {
			return nil, err
		} else {
			covidData = append(covidData, *data)
		}
	}

	return covidData, nil
}

func ResetCovidDataCache(covidData []models.CovidData) error {
	var keys []string

	for _, data := range covidData {
		keys = append(keys, data.Region)
	}

	return covidDataCache.RemoveKeys(context.TODO(), keys)
}
