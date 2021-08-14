package models

import "time"

type CovidRegionResponse struct {
	Region         string    `json:"region"`                          // Data Region
	ActiveCases    int64     `json:"activeCases,string,omitempty"`    // Total Active Cases
	ConfirmedCases int64     `json:"confirmedCases,string,omitempty"` // Total Confirmed Cases
	Deaths         int64     `json:"deaths,string,omitempty"`         // Total Deaths
	Recovered      int64     `json:"recovered,string,omitempty"`      // Total Recovered
	RemoteSyncTime time.Time `json:"remoteSyncTime"`                  // Instant at which data was updated by 3rd Party Data Provider
}

type UserFeedResponse struct {
	India         CovidRegionResponse `json:"india"`           // India's covid data
	State         CovidRegionResponse `json:"state"`           // State's covid data
	LastUpdatedAt time.Time           `json:"last_updated_at"` // Data last updated at
}

type DataIngestResponse struct {
	Message   string    `json:"message"`   // Success Message
	UpdatedAt time.Time `json:"updatedAt"` // Data Updated At
}

type HealthCheckResponse struct {
	Status string `json:"status"` // Status of server
}

type ErrorResponse struct {
	Status  string `json:"status"`  // Error
	Message string `json:"message"` // Error Message
}
