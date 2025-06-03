package routes

import (
	"training/backend/webserver/middlewares"
	"training/package/validator"

	"github.com/labstack/echo/v4"
)

// Routers function
func Routers(app *echo.Echo) {

	//Common middleware
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger(true))
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())
	// app.Use(middlewares.KeyAuth())
	app.Use(middlewares.Session())

	//initialize custom validator
	app.Validator = validator.GetValidator()

	//web routers
	LookUpApiRouters(app)
}
