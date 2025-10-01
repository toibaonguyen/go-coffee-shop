package api

import (
	errorhandler "coffee-shop/api/error-handler"
	"coffee-shop/api/router"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.HTTPErrorHandler = errorhandler.HTTPErrorHandler
	router.SetProductRoute(e)
}
