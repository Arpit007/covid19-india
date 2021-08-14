package utils

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HandleError(err error, code int, c echo.Context) error {
	log.Error(err)

	return c.JSON(code, ErrorResponse{Status: "Error", Message: err.Error()})
}
