package third_party_helpers

import (
	models3p "covid19-india/internal/models/third_party"
	"covid19-india/internal/utils"
	"errors"
	"time"
)

// FetchCovid3pData Fetches Covid Data from 3rd Party API
func FetchCovid3pData() ([]models3p.Covid3pData, error) {
	client := utils.GetClient(10 * time.Second)

	res, err := client.Get("https://www.mohfw.gov.in/data/datanew.json")
	if err != nil {
		return nil, err
	}

	var data []models3p.Covid3pData
	if err := utils.DecodeResponseBody(res, &data); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("unable to fetch covid 19 data")
	}

	return data, nil
}
