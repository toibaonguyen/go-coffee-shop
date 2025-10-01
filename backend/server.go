package main

import (
	"coffee-shop/api"

	"github.com/labstack/echo/v4"
)

func init() {
}

func main() {
	e := echo.New()
	api.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
