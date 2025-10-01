package customerservices

import (
	"coffee-shop/dto"
	"coffee-shop/repositories"
)

type ProductService interface {
	GetProductsForCustomerMenu(limit int, size int, search string, category string) ([]dto.MenuProduct, error)
}

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
}

func (ps ProductServiceImpl) GetProductsForCustomerMenu(limit int, size int, search string, category string) ([]dto.MenuProduct, error) {
	products, err := ps.ProductRepository.GetMenuProducts(limit, size, search, category)
	if err != nil {
		return nil, err
	}

	return products, nil
}
