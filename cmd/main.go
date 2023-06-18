package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
	"fmt"
)

func main() {
	cfg := config.Load()
	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)

	// User(con)
	// Category(con)
	// Product(con)
	// Order(con)
	OrderPayment(con)
}

func User(con *controller.Controller) {
	con.UserCreate(&models.CreateUser{
		FirstName: "Asadbek",
		LastName:  "Ergashev",
		Balance:   200_000,
	})
}

func Category(con *controller.Controller) {
	con.CategoryCreate(&models.CreateCategory{
		Name: "Ilmiy",
	})
}

func Product(con *controller.Controller) {

	var categoryId = "8ca2dea0-eff2-4ba4-978f-763efab6bc50"
	var products = []models.CreateProduct{
		{
			Name:         "Matematika",
			Price:        15_000,
			Discount:     0,
			DiscountType: "",
			CategoryID:   categoryId,
		},
		{
			Name:         "Learning Golang",
			Price:        200_000,
			Discount:     30,
			DiscountType: config.PercentDiscountType,
			CategoryID:   categoryId,
		},
		{
			Name:         "Clean Code",
			Price:        350_000,
			Discount:     40_000,
			DiscountType: config.FixDiscountType,
			CategoryID:   categoryId,
		},
	}

	for _, product := range products {
		con.ProductCreate(&product)
	}
}

func Order(con *controller.Controller) {
	// con.OrderCreate(&models.CreateOrder{
	// 	UserId: "204ff9b0-3f4e-41b3-a436-3a1fce028fa6",
	// })

	var (
		orderId   = "ff9aa3f6-7dd2-4b2e-9376-93bc47391e82"
		productId = "55288fca-0360-475e-87a8-69a25d92b8fa"
		count     = 2
	)

	con.AddOrderItem(&models.CreateOrderItem{
		OrderId:   orderId,
		ProductId: productId,
		Count:     count,
	})
}

func OrderPayment(con *controller.Controller) {

	var orderId = "ff9aa3f6-7dd2-4b2e-9376-93bc47391e82"

	err := con.OrderPayment(&models.OrderPayment{
		OrderId: orderId,
	})

	fmt.Println(err)
}
