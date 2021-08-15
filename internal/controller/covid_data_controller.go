package controller

import (
	"covid19-india/internal/cache"
	"covid19-india/internal/dao"
	"covid19-india/internal/helpers"
	"covid19-india/internal/models"
	"covid19-india/internal/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type CovidDataController struct{}

func (self CovidDataController) RegisterRoutes(g *echo.Group) {
	// Register routes
	g.POST("/refresh", self.refreshData)
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

	if err := cache.ResetCovidDataCache(covidData); err != nil {
		return utils.HandleError(err, http.StatusInternalServerError, c)
	}

	res := models.DataIngestResponse{Message: "Success", UpdatedAt: time.Now()}

	return c.JSON(http.StatusCreated, res)
}
