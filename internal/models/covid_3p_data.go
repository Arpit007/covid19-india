package models

import "time"

type Covid3pData struct {
	Active          int64  `json:"active,string,omitempty"`
	Confirmed       int64  `json:"confirmed,string,omitempty"`
	Deaths          int64  `json:"deaths,string,omitempty"`
	DeltaConfirmed  int64  `json:"deltaconfirmed,string,omitempty"`
	DeltaDeaths     int64  `json:"deltadeaths,string,omitempty"`
	DeltaRecovered  int64  `json:"delta_recovered,string,omitempty"`
	LastUpdatedTime string `json:"lastupdatedtime"` // Last Sync time of data provider
	MigratedOther   int64  `json:"migratedother,string,omitempty"`
	Recovered       int64  `json:"recovered,string,omitempty"`
	State           string `json:"state"`
	StateCode       string `json:"state_code"`
	StateNotes      string `json:"statenotes"`
}

type Covid3pDataResponse struct {
	StateWise []Covid3pData `json:"statewise"`
}

func (data *Covid3pData) ToCovidData() (*CovidData, error) {
	syncTime, err := time.Parse("02/01/2006 15:04:05", data.LastUpdatedTime)
	currentTime := time.Now()

	if err != nil {
		return nil, err
	}

	var region string

	if data.State == "Total" {
		region = "India"
	} else {
		region = data.State
	}

	model := &CovidData{
		Region:         region,
		ActiveCases:    data.Active,
		ConfirmedCases: data.Confirmed,
		Deaths:         data.Deaths,
		Recovered:      data.Recovered,
		RemoteSyncTime: syncTime,
		CreatedAt:      currentTime,
		UpdatedAt:      currentTime,
	}

	return model, nil
}
