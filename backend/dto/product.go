package dto

type Price struct {
	IsDiscount      bool
	CurrentPrice    float64
	OriginalPrice   float64
	DiscountPercent uint
}

type ProductVariant struct {
	Id    string
	Size  string
	Price Price
	Stock int
}

type MenuProduct struct {
	Id       string
	Name     string
	ImageURL string
	Variants []ProductVariant
	Status   string
}
