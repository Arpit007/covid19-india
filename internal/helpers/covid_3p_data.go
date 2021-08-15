package helpers

import (
	"covid19-india/internal/models"
	"covid19-india/internal/utils"
	"errors"
	"time"
)

// FetchCovid3pData Fetches Covid Data from 3rd Party API
func FetchCovid3pData() ([]models.Covid3pData, error) {
	client := utils.GetClient(time.Second * 10)

	res, err := client.Get("https://data.covid19india.org/data.json")
	if err != nil {
		return nil, err
	}

	var data models.Covid3pDataResponse
	if err := utils.DecodeResponseBody(res, &data); err != nil {
		return nil, err
	}

	if len(data.StateWise) == 0 {
		return nil, errors.New("unable to fetch covid 19 data")
	}

	return data.StateWise, nil
}
