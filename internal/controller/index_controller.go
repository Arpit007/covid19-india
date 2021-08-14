package controller

import (
	"covid19-india/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type IndexController struct {
}

func (self IndexController) RegisterRoutes(e *echo.Echo) {
	// Register health-check
	e.GET("/healthcheck", healthCheck)

	// Register Controllers
	new(CovidDataController).RegisterRoutes(e.Group("/v1/data"))
	new(UserFeedController).RegisterRoutes(e.Group("/v1/user"))
}

// healthCheck godoc
// @Summary Server's health-check
// @Description Check server's health
// @Tags healthcheck
// @Produce  json
// @Success 200 {object} models.HealthCheckResponse
// @Router /healthcheck [get]
func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, models.HealthCheckResponse{Status: "Server Running"})
}
