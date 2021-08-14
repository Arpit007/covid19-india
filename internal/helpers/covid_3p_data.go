package helpers

import (
	"covid19-india/internal/models"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func FetchCovid3pData() ([]models.Covid3pData, error) {
	client := &http.Client{Timeout: time.Second * 10}

	res, err := client.Get("https://data.covid19india.org/data.json")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var covidDataResponse models.Covid3pDataResponse

	if err := json.NewDecoder(res.Body).Decode(&covidDataResponse); err != nil {
		return nil, err
	}

	if len(covidDataResponse.StateWise) == 0 {
		return nil, errors.New("unable to fetch covid 19 data")
	}

	return covidDataResponse.StateWise, nil
}
