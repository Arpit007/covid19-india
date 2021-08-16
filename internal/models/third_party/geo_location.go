package third_party

// Address 3rd Party Reverse geocoding response model
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

// GeoPlace 3rd Party Reverse geocoding response model
type GeoPlace struct {
	Title   string  `json:"title"`
	ID      string  `json:"id"`
	Address Address `json:"address"`
}

// GeoResponse 3rd Party Reverse geocoding response
type GeoResponse struct {
	Items []GeoPlace `json:"items"`
}
