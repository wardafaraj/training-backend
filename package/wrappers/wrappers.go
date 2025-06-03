package wrappers

import (
	"github.com/labstack/echo/v4"
)

//Error: Use the following
//http.StatusOK  - for successul response
//http.StatusAccepted - for create responses
//http.StatusNoContent - for all data request when no data is available
//http.StatusInternalServerError - for internal error

type ResponseData struct {
	Code    int         `json:"code"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(c echo.Context, statusCode int, data interface{}) error {

	return c.JSON(statusCode, ResponseData{
		Code: statusCode,
		Data: data,
	})
}

func MessageResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, ResponseData{
		Code:    statusCode,
		Message: message,
	})
}

func ErrorResponse(c echo.Context, statusCode int, err string) error {
	return c.JSON(statusCode, ResponseData{
		Code:  statusCode,
		Error: err,
	})
}
