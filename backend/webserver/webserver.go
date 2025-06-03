package webserver

import (
	"fmt"
	"training/backend/config"
	"training/backend/webserver/routes"
	"training/package/helpers"
	"training/package/log"
	"time"

	"github.com/labstack/echo/v4"
)

// StartWebserver starts a webserver
func StartWebserver() {
	// Echo instance
	e := echo.New()

	e.Renderer = Renderer()

	//Disable echo banner
	e.HideBanner = true

	// Routes
	routes.Routers(e)

	e.Server.IdleTimeout = 5 * time.Minute

	//Init cache
	helpers.Init()

	//Remove this in production
	e.Debug = true

	cfg, err := config.New()
	if err != nil {
		log.Errorf("error getting config: %v", err)
	}

	address := fmt.Sprintf("%v:%v", cfg.WebServer.Host, cfg.WebServer.Port)
	e.Logger.Fatal(e.Start(address))
}
