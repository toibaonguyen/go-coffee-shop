package errorhandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	var appHttpErr *AppHttpError
	if errors.As(err, &appHttpErr) {
		c.JSON(appHttpErr.Code, map[string]string{
			"status":  appHttpErr.Status,
			"message": appHttpErr.Message,
		})
		return
	}

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		if httpErr.Code == http.StatusBadRequest {
			c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Validation failed",
			})
			return
		}

		c.JSON(httpErr.Code, map[string]string{
			"message": httpErr.Message.(string),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, map[string]string{
		"message": "Internal Server Error",
	})
}
