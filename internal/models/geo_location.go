package models

type Address struct {
	Label       string `json:"label"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
	StateCode   string `json:"state_code"`
	State       string `json:"state"`
	County      string `json:"county"`
	City        string `json:"city"`
	District    string `json:"district"`
	PostalCode  string `json:"postalCode"`
}

type GeoPlace struct {
	Title   string  `json:"title"`
	ID      string  `json:"id"`
	Address Address `json:"address"`
}

type GeoResponse struct {
	Items []GeoPlace `json:"items"`
}
