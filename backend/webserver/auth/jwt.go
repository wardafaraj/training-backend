package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthJWT() echo.MiddlewareFunc {
	if accessTokenMiddleware != nil {
		return accessTokenMiddleware
	}
	accessTokenMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &JWTCustomClaims{},
		SigningKey:              []byte(GetJWTSecret()),
		TokenLookup:             "cookie:" + accessTokenCookieName,
		ErrorHandlerWithContext: JWTErrorChecker,
		Skipper:                 SkipperLoginCheck,
	})
	return accessTokenMiddleware
}

func SkipperLoginCheck(c echo.Context) bool {

	if strings.HasSuffix(c.Path(), "/auth/api/v1/auth/login") ||
		strings.HasSuffix(c.Path(), "/auth/api/v1/auth/register") {
		return true
	}
	return false
}

func JWTErrorChecker(err error, c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
