package controller

import (
	"covid19-india/internal/cache"
	"covid19-india/internal/dao"
	"covid19-india/internal/helpers"
	"covid19-india/internal/models"
	"covid19-india/internal/models/transformers"
	"covid19-india/internal/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type CovidDataController struct{}

const indiaRegion = "India"

func (self CovidDataController) RegisterRoutes(g *echo.Group) {
	// Register routes
	g.POST("/refresh", self.refreshData)
	g.GET("/geo", self.getCovidDataByGeo)
}

// refreshData godoc
// @Summary Populate Covid19 Data
// @Description Fetches and persists India's covid 19 data in DB
// @Tags covidApi
// @Produce  json
// @Success 201 {object} models.DataIngestResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /v1/data/refresh [post]
func (self CovidDataController) refreshData(c echo.Context) error {
	data, err := helpers.FetchCovid3pData()

	if err != nil {
		return utils.HandleError(err, http.StatusInternalServerError, c)
	}

	if len(data) == 0 {
		return utils.HandleError(errors.New("no covid data found from remote"), http.StatusInternalServerError, c)
	}

	covidData, err := dao.PersistCovidData(data)

	if err != nil {
		return utils.HandleError(err, http.StatusInternalServerError, c)
	}

	go func() {
		if err := cache.ResetCovidDataCache(covidData); err != nil {
			logrus.Error("Error resetting cache ", err)
		}
	}()

	updatedAt := time.Now().In(transformers.IstTimeZone).Format(time.RFC1123)
	res := models.DataIngestResponse{Message: "Data ingested successfully", UpdatedAt: updatedAt}

	return c.JSON(http.StatusCreated, res)
}

// getCovidDataByGeo godoc
// @Summary Get Covid Data for State
// @Description Get India & State's covid data based on geo-location
// @Tags covidApi
// @Param lat query float32 true "Latitude"
// @Param lng query float32 true "Longitude"
// @Produce  json
// @Success 201 {object} models.GeoCovidDataResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /v1/data/geo [get]
func (self CovidDataController) getCovidDataByGeo(c echo.Context) error {
	lat, err := strconv.ParseFloat(c.QueryParam("lat"), 64)
	if err != nil || lat < -90 || lat > 90 { // Validation
		return utils.HandleError(errors.New("invalid latitude"), http.StatusBadRequest, c)
	}

	lng, err := strconv.ParseFloat(c.QueryParam("lng"), 64)
	if err != nil || lng < -180 || lng > 180 { // Validation
		return utils.HandleError(errors.New("invalid longitude"), http.StatusBadRequest, c)
	}

	// Fetch user's state from geo-coordinates
	state, err := cache.GetStateFromLatLong(lat, lng)

	if err != nil {
		return utils.HandleError(err, http.StatusInternalServerError, c)
	}

	// Get covid data for state & India
	data, err := cache.GetCovidDataForRegions([]string{state, indiaRegion})

	if err != nil {
		return utils.HandleError(err, http.StatusInternalServerError, c)
	}

	if len(data) == 0 {
		return utils.HandleError(errors.New("no data found"), http.StatusInternalServerError, c)
	}

	return c.JSON(http.StatusOK, transformers.ToGeoCovidDataResponse(data))
}
