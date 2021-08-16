package models

// CovidRegionResponse Covid data for a region
type CovidRegionResponse struct {
	Region         string `json:"region"`                          // Data Region
	ActiveCases    int64  `json:"activeCases,string,omitempty"`    // Total Active Cases
	ConfirmedCases int64  `json:"confirmedCases,string,omitempty"` // Total Confirmed Cases
	Deaths         int64  `json:"deaths,string,omitempty"`         // Total Deaths
	Recovered      int64  `json:"recovered,string,omitempty"`      // Total Recovered
}

// GeoCovidDataResponse India & a State's covid data
type GeoCovidDataResponse struct {
	India         CovidRegionResponse `json:"india"`           // India's covid data
	State         CovidRegionResponse `json:"state"`           // State's covid data
	LastUpdatedAt string              `json:"last_updated_at"` // Data last updated at
}

// SimpleMessageResponse Response with a string message
type SimpleMessageResponse struct {
	Message string `json:"message"`
}

// ErrorResponse Error message response
type ErrorResponse struct {
	Status  string `json:"status"`  // Status
	Message string `json:"message"` // Error Message
}
