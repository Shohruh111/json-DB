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
	// OrderPayment(con)
	// UserHistory(con)
	con.ActiveProduct()
}
func UserHistory(con *controller.Controller) {
	userHistory, err := con.UserHistory()
	if err != nil {
		fmt.Println(err)
	}
	for key, val := range userHistory {
		fmt.Printf("\t\t %s:\n", key)
		fmt.Println(val)
	}
}
func User(con *controller.Controller) {
	con.UserCreate(&models.CreateUser{
		FirstName: "Daminick",
		LastName:  "Toretto",
		Balance:   600_000,
	})
}

func Category(con *controller.Controller) {
	con.CategoryCreate(&models.CreateCategory{
		Name: "Ilmiy",
	})
}

func Product(con *controller.Controller) {

	var categoryId = "0486c9a7-6aaa-450c-8006-1c216ac5650b"
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
	// 	UserId: "d33c4572-81b8-4837-9aaa-252154a3f629",
	// })

	var (
		orderId   = "e75db884-2b1d-4ff6-a66f-7d41b2c56697"
		productId = "558db6d5-e43c-4840-b369-7effe0ebba14"
		count     = 2
	)

	con.AddOrderItem(&models.CreateOrderItem{
		OrderId:   orderId,
		ProductId: productId,
		Count:     count,
	})
}

func OrderPayment(con *controller.Controller) {

	var orderId = "d4daa0c2-d13d-4ab4-89f1-677ec606fc3a"

	err := con.OrderPayment(&models.OrderPayment{
		OrderId: orderId,
	})

	fmt.Println(err)
}
