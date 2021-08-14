package models

import (
	"github.com/kamva/mgm/v3"
	"time"
)

type CovidData struct {
	mgm.DefaultModel `json:"-" bson:",inline"`
	Region           string    `json:"region" bson:"region"`
	ActiveCases      int64     `json:"activeCases,string,omitempty"`
	ConfirmedCases   int64     `json:"confirmedCases,string,omitempty"`
	Deaths           int64     `json:"deaths,string,omitempty"`
	Recovered        int64     `json:"recovered,string,omitempty"`
	RemoteSyncTime   time.Time `json:"remoteSyncTime"`
	CreatedAt        time.Time `bson:"created_at" json:"-"`
	UpdatedAt        time.Time `bson:"updated_at" json:"updatedAt"`
}

func (data *CovidData) ToCovidRegionResponse() *CovidRegionResponse {
	return &CovidRegionResponse{
		Region:         data.Region,
		ActiveCases:    data.ActiveCases,
		ConfirmedCases: data.ConfirmedCases,
		Deaths:         data.Deaths,
		Recovered:      data.Recovered,
		RemoteSyncTime: data.RemoteSyncTime,
	}
}

func ToUserFeedResponse(covidData []CovidData) *UserFeedResponse {
	var indiaData CovidData
	var stateData CovidData

	for _, data := range covidData {
		switch data.Region {
		case "India":
			indiaData = data
		default:
			stateData = data
		}
	}

	return &UserFeedResponse{
		India:         *indiaData.ToCovidRegionResponse(),
		State:         *stateData.ToCovidRegionResponse(),
		LastUpdatedAt: indiaData.UpdatedAt,
	}
}
