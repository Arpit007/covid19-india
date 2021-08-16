package controller

import (
	"covid19-india/internal/helpers"
	"covid19-india/internal/models"
	"covid19-india/internal/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CovidInfoController struct{}

func (self CovidInfoController) RegisterRoutes(g *echo.Group) {
	// Register routes
	g.POST("/refresh", self.refreshData)
	g.GET("/geo", self.getCovidDataByGeo)
}

// refreshData godoc
// @Summary Populate Covid19 Data
// @Description Fetches and persists India's covid 19 data in DB
// @Tags covidApi
// @Produce  json
// @Success 201 {object} models.SimpleMessageResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /v1/covid/refresh [post]
func (self CovidInfoController) refreshData(c echo.Context) error {
	if err := helpers.RefreshCovidData(); err != nil {
		return utils.HandleError(err, http.StatusInternalServerError, c)
	}

	res := models.SimpleMessageResponse{Message: "Data ingested successfully"}

	return c.JSON(http.StatusCreated, res)
}

// getCovidDataByGeo godoc
// @Summary Get Covid Data for State
// @Description Get India & State's covid data based on geo-location
// @Tags covidApi
// @Param lat query float32 true "Latitude" minimum(-90) maximum(90)
// @Param lng query float32 true "Longitude" minimum(-180) maximum(180)
// @Produce  json
// @Success 201 {object} models.GeoCovidDataResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /v1/covid/geo [get]
func (self CovidInfoController) getCovidDataByGeo(c echo.Context) error {
	lat, err := strconv.ParseFloat(c.QueryParam("lat"), 32)

	// Validate latitude
	if err != nil || !validateLatitude(lat) {
		return utils.HandleError(errors.New("invalid latitude"), http.StatusBadRequest, c)
	}

	lng, err := strconv.ParseFloat(c.QueryParam("lng"), 32)

	// Validate longitude
	if err != nil || !validateLongitude(lng) {
		return utils.HandleError(errors.New("invalid longitude"), http.StatusBadRequest, c)
	}

	if data, err := helpers.GetCovidDataForUserGeo(lat, lng); err != nil {
		return utils.HandleError(err, http.StatusInternalServerError, c)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func validateLatitude(lat float64) bool {
	return lat >= -90 && lat <= 90
}

func validateLongitude(lng float64) bool {
	return lng >= -180 && lng <= 180
}
