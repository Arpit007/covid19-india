package server

import (
	. "covid19-india/configs"
	"covid19-india/internal/server/controller"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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
	port := fmt.Sprintf(":%s", ENV.Port)
	app.Logger.Fatal(app.Start(port))
}

func registerRoutes(e *echo.Echo) {
	new(controller.IndexController).RegisterRoutes(e)
}
