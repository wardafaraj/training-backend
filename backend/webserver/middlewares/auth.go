package middlewares

import (
	"net/http"
	"training/backend/config"
	"training/package/crypto"
	"training/package/log"
	"training/package/util"
	"training/package/wrappers"
	"strings"

	"github.com/labstack/echo/v4"
)

const dataHash = "DATA-HASH"
const dataSignature = "DATA-SIGNATURE"
const systemName = "SYSTEM-NAME"

// KeyAuth middleware
func KeyAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			hash := c.Request().Header.Get(dataHash)
			signature := c.Request().Header.Get(dataSignature)
			system := c.Request().Header.Get(systemName)

			// if SkipperKeyAuth(c) {
			// 	return next(c)
			// }

			cfg, err := config.New()
			if util.IsError(err) {
				log.Errorf("error initializing config")
				return wrappers.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
			}
			pubKey, err := cfg.GetSystemPublicKey(system)
			if util.IsError(err) {
				return wrappers.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			}
			isValid, err := crypto.Verify(pubKey, hash, signature)
			if isValid && !util.IsError(err) {
				return next(c)
			}
			return wrappers.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")

		}
	}
}

// SkipperKeyAuth
func SkipperKeyAuth(c echo.Context) bool {
	url := c.Request().URL.Path
	if strings.HasSuffix(url, "") ||
		strings.HasPrefix(url, "") ||
		strings.Contains(url, "/") {
		return true
	}
	return false
}
