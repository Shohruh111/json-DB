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
	UserHistory(con)
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
	// 	UserId: "af299662-c6e8-4559-a0da-cebf213a061d",
	// })

	var (
		orderId   = "634957ce-b166-45da-ae73-f5b00e092f5a"
		productId = "cb158bb4-c3ce-402d-b715-223a5ba1c97a"
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
