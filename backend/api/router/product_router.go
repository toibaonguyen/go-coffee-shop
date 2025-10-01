package router

import (
	"coffee-shop/api/controllers"
	"coffee-shop/repositories"
	"coffee-shop/services/customerservices"

	"github.com/labstack/echo/v4"
)

func SetProductRoute(e *echo.Echo) {

	productController := controllers.ProductController{
		CustomerProductService: customerservices.ProductServiceImpl{
			ProductRepository: repositories.ProductRepositoryImpl{},
		},
	}

	e.GET("/api/products", productController.GetProducts)
}
