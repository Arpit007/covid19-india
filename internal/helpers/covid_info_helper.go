package helpers

import (
	"context"
	"covid19-india/internal/cache"
	"covid19-india/internal/dao"
	"covid19-india/internal/models"
	"covid19-india/internal/models/transformers"
	helper3p "covid19-india/internal/third_party_helpers"
	"errors"
	"github.com/sirupsen/logrus"
)

// RefreshCovidData Fetch the Covid data and persist in DB
func RefreshCovidData(ctx context.Context) error {
	// Fetch covid data from third party
	covid3pData, err := helper3p.FetchCovid3pData()

	if err != nil {
		return err
	}

	if len(covid3pData) == 0 {
		return errors.New("no covid data found from remote")
	}

	var covidData []models.CovidData
	for _, data3p := range covid3pData {
		if data, err := data3p.ToCovidData(); err != nil {
			return err
		} else {
			covidData = append(covidData, *data)
		}
	}

	// Persist data
	if err := dao.PersistCovidData(ctx, covidData); err != nil {
		return err
	}

	// Reset the cache
	go func() {
		if err := cache.ResetCovidDataCache(context.TODO(), covidData); err != nil {
			logrus.Error("Error resetting cache ", err)
		}
	}()

	return nil
}

// GetCovidDataForUserGeo Get Covid data for user's state & India based on their geo-coordinates
func GetCovidDataForUserGeo(ctx context.Context, lat float64, lng float64) (*models.GeoCovidDataResponse, error) {
	// Get user's state from coordinates
	state, err := helper3p.GetStateFromLatLong(lat, lng)

	if err != nil {
		return nil, err
	}

	// Get covid data for state & India
	data, err := cache.GetCovidDataForRegions(ctx, []string{state, "India"})

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("no data found")
	}

	return transformers.ToGeoCovidDataResponse(data), nil
}
