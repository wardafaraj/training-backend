package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Cors Middleware
func Cors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4322", "*"},
		AllowMethods: []string{echo.HEAD, echo.GET, echo.PUT, echo.POST, echo.DELETE},
	})
}
