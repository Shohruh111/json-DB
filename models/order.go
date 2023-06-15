package models

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type Order struct {
	Id       string       `json:"id"`
	UserId   string       `json:"user_id"`
	Sum      int          `json:"sum"`
	SumCount int          `json:"sum_count"`
	Status   string       `json:"status"`
	Orders   []*OrderItem `json:"orders"`
}
type OrderItem struct {
	ProductId  string `json:"product_id"`
	OrderId    string `json:"order_id"`
	Count      int    `json:"count"`
	TotalPrice int    `json:"total_price"`
}
type OrderGetList struct {
	Count  int
	Orders []*Order
}
type CreateOrder struct {
	UserId   string `json:"user_id"`
	Sum      int    `json:"sum"`
	SumCount int    `json:"sum_count"`
	Status   string `json:"status"`
}
type OrderGetListRequest struct {
	Offset int
	Limit  int
}

type UpdateOrder struct {
	Id       string       `json:"id"`
	UserId   string       `json:"user_id"`
	Sum      int          `json:"sum"`
	SumCount int          `json:"sum_count"`
	Status   string       `json:"status"`
	Orders   []*OrderItem `json:"orders"`
}
