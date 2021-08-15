package utils

import (
	"covid19-india/internal/models"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// HandleError Log & sGenerate error message response
func HandleError(err error, code int, c echo.Context) error {
	log.Error(err)

	return c.JSON(code, models.ErrorResponse{Status: "Error", Message: err.Error()})
}
