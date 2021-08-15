package transformers

import (
	"covid19-india/internal/models"
	"time"
)

var IstTimeZone *time.Location

func init() {
	secondsDelta := int((5*time.Hour + 30*time.Minute).Seconds()) // IST: +05:30
	IstTimeZone = time.FixedZone("IST", secondsDelta)
}

// ToCovidRegionResponse Transform CovidData to API response CovidRegionResponse
func ToCovidRegionResponse(data *models.CovidData) *models.CovidRegionResponse {
	return &models.CovidRegionResponse{
		Region:         data.Region,
		ActiveCases:    data.ActiveCases,
		ConfirmedCases: data.ConfirmedCases,
		Deaths:         data.Deaths,
		Recovered:      data.Recovered,
	}
}

// ToGeoCovidDataResponse Transform data to API response GeoCovidDataResponse
func ToGeoCovidDataResponse(covidData []models.CovidData) *models.GeoCovidDataResponse {
	var indiaData models.CovidData
	var stateData models.CovidData

	for _, data := range covidData {
		switch data.Region {
		case "India":
			indiaData = data
		default:
			stateData = data
		}
	}

	return &models.GeoCovidDataResponse{
		India:         *ToCovidRegionResponse(&indiaData),
		State:         *ToCovidRegionResponse(&stateData),
		LastUpdatedAt: indiaData.UpdatedAt.In(IstTimeZone).Format(time.RFC1123),
	}
}
