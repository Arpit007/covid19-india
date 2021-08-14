package models

import "time"

type CovidRegionResponse struct {
	Region         string    `json:"region"`
	ActiveCases    int64     `json:"activeCases,string,omitempty"`
	ConfirmedCases int64     `json:"confirmedCases,string,omitempty"`
	Deaths         int64     `json:"deaths,string,omitempty"`
	Recovered      int64     `json:"recovered,string,omitempty"`
	RemoteSyncTime time.Time `json:"remoteSyncTime"`
}

type UserFeedResponse struct {
	India         CovidRegionResponse `json:"india"`
	State         CovidRegionResponse `json:"state"`
	LastUpdatedAt time.Time           `json:"last_updated_at"`
}

type DataIngestResponse struct {
	Message   string    `json:"message"`
	UpdatedAt time.Time `json:"updatedAt"`
}
