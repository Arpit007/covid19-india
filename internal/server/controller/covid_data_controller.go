package controller

import (
	. "covid19-india/internal/dao"
	"covid19-india/internal/helpers"
	. "covid19-india/internal/models"
	"covid19-india/internal/server/utils"
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

func (self CovidDataController) refreshData(c echo.Context) error {
	data, err := helpers.FetchCovid3pData()

	if err != nil {
		return utils.HandleError(err, http.StatusInternalServerError, c)
	}

	if len(data) == 0 {
		return utils.HandleError(errors.New("no covid data found from remote"), http.StatusInternalServerError, c)
	}

	if err := PersistCovidData(data); err != nil {
		return utils.HandleError(err, http.StatusInternalServerError, c)
	}

	res := DataIngestResponse{Message: "Success", UpdatedAt: time.Now()}
	return c.JSON(http.StatusOK, res)
}
