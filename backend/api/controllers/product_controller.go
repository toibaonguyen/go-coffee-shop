package controllers

import (
	"coffee-shop/api/models"
	"coffee-shop/services/customerservices"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	CustomerProductService customerservices.ProductService
}

func (pc ProductController) GetProducts(ctx echo.Context) error {
	limitStr := ctx.QueryParam("limit")
	sizeStr := ctx.QueryParam("size")
	limit, _ := strconv.Atoi(limitStr)
	size, _ := strconv.Atoi(sizeStr)
	search := ctx.QueryParam("search")
	category := ctx.QueryParam("category")

	products, err := pc.CustomerProductService.GetProductsForCustomerMenu(limit, size, search, category)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, models.ResponseData{
		Status: "OK",
		Data:   products,
	})
}
