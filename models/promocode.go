package models

type Promocode struct {
	Name         string             `json:"name"`
	Price        int                `json:"price"`
	Discount     int                `json:"discount"`
	DiscountType string             `json:"discount_type"`
	PriceLimit   int                `json:"price_limit"`
	Promocodes   []*CreatePromocode `json:"promocode"`
}
type UpdatePromocode struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	PriceLimit   int    `json:"price_limit"`
}
type CreatePromocode struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	PriceLimit   int    `json:"price_limit"`
}
type RemovePromocode struct{
	Name string `json:"name"`
}