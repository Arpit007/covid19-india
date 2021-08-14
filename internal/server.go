package internal

import (
	_ "covid19-india/docs"
	"covid19-india/internal/config"
	controller2 "covid19-india/internal/controller"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/swag"
)

// @title Covid 19 API Server
// @version 1.0
// @description Get covid data based on your geo-location in India

// @contact.name Arpit Bhatnagar
// @contact.email arpitbhatnagar10@gmail.com
// @host localhost:8000
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
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	new(controller2.IndexController).RegisterRoutes(e)
}
