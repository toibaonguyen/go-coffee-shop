package repositories

import "coffee-shop/dto"

type ProductRepository interface {
	GetMenuProducts(limit int, size int, search string, category string) ([]dto.MenuProduct, error)
}

type ProductRepositoryImpl struct {
}

func (ProductRepositoryImpl) GetMenuProducts(limit int, size int, search string, category string) ([]dto.MenuProduct, error) {
	var products []dto.MenuProduct
	products = append(products, dto.MenuProduct{
		Id:       "p1",
		Name:     "Espresso",
		ImageURL: "https://example.com/images/espresso.png",
		Status:   "available",
		Variants: []dto.ProductVariant{
			{
				Id:   "p1-s",
				Size: "S",
				Price: dto.Price{
					IsDiscount:      false,
					CurrentPrice:    25000,
					OriginalPrice:   25000,
					DiscountPercent: 0,
				},
				Stock: 100,
			},
		},
	})

	return products, nil
}
