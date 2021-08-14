package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type IndexController struct {
}

func (self IndexController) RegisterRoutes(e *echo.Echo) {
	// Register health-check
	e.GET("/healthcheck", healthCheck)

	// Register Controllers
	new(CovidDataController).RegisterRoutes(e.Group("/data"))
	new(UserFeedController).RegisterRoutes(e.Group("/user"))
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "Server running",
	})
}
