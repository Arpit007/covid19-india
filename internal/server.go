package internal

import (
	_ "covid19-india/docs"
	"covid19-india/internal/config"
	"covid19-india/internal/controller"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/swag"
	"net/http"
)

// StartServer godoc
// @title Covid 19 API Server
// @version 1.0
// @description Get covid data based on your geo-location in India
// @contact.name Arpit Bhatnagar
// @contact.email arpitbhatnagar10@gmail.com
// @BasePath /
func StartServer() {
	app := echo.New()

	// Middleware
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${remote_ip}\t${time_rfc3339}\t${method}\t${uri}\t${status}\t${latency_human}\n",
	}))
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	// Register Routes
	registerRoutes(app)

	// Start server
	port := fmt.Sprintf(":%s", config.ENV.Port)
	app.Logger.Fatal(app.Start(port))
}

func registerRoutes(e *echo.Echo) {
	// Register Swagger routes
	e.GET("/", index)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Register Index Controller
	new(controller.IndexController).RegisterRoutes(e)
}

// Index /
// Redirect to swagger page
func index(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
}
