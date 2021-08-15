package models

import "time"

// Covid3pData 3rd party API Model
// Provides Covid data for a given state/region
type Covid3pData struct {
	State     string `json:"state_name"`
	Active    int64  `json:"active,string,omitempty"`
	Confirmed int64  `json:"positive,string,omitempty"`
	Recovered int64  `json:"cured,string,omitempty"`
	Deaths    int64  `json:"death,string,omitempty"`
}

// ToCovidData Transform 3rd party API model to app's covid data model
func (data *Covid3pData) ToCovidData() (*CovidData, error) {
	currentTime := time.Now()

	var region string

	// 3rd Party treats India level data same as state level data, with the State name as blank
	if len(data.State) == 0 {
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
		CreatedAt:      currentTime,
		UpdatedAt:      currentTime,
	}

	return model, nil
}
