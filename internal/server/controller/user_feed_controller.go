package controller

import (
	"covid19-india/internal/dao"
	"covid19-india/internal/helpers"
	"covid19-india/internal/models"
	"covid19-india/internal/server/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserFeedController struct{}

const indiaRegion = "India"

func (self UserFeedController) RegisterRoutes(g *echo.Group) {
	// Register routes
	g.GET("/geo", self.getCovidDataByGeo)
}

func (self UserFeedController) getCovidDataByGeo(c echo.Context) error {
	lat, err := strconv.ParseFloat(c.QueryParam("lat"), 64)
	if err != nil {
		return utils.HandleError(errors.New("invalid latitude"), http.StatusBadRequest, c)
	}

	lng, err := strconv.ParseFloat(c.QueryParam("lng"), 64)
	if err != nil {
		return utils.HandleError(errors.New("invalid longitude"), http.StatusBadRequest, c)
	}

	location, err := helpers.GetPlaceFromLatLng(lat, lng)

	if err != nil {
		return utils.HandleError(err, http.StatusInternalServerError, c)
	}

	state := location.Items[0].Address.State

	data, err := dao.GetCovidDataForStates([]string{state, indiaRegion})

	if err != nil {
		return utils.HandleError(err, http.StatusInternalServerError, c)
	}

	if len(data) == 0 {
		return utils.HandleError(errors.New("no data found"), http.StatusInternalServerError, c)
	}

	return c.JSON(http.StatusOK, models.ToUserFeedResponse(data))
}
