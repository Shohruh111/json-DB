package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
)

func main() {
	cfg := config.Load()
	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)
	// for {
	// 	var (
	// 		answer string
	// 	)
	// 	fmt.Printf("1.Create Product\n2.Get By Product Id\n3.Get Product List\n4.Update Product Info\n5.Delete Product\n6.Exit\nEnter: ")
	// 	fmt.Scan(&answer)
	// 	if answer == "1" {
	// 		var (
	// 			name  string
	// 			price string
	// 		)
	// 		fmt.Printf("First Name: ")
	// 		fmt.Scan(&name)
	// 		fmt.Printf("Last Name: ")
	// 		fmt.Scan(&price)
	// 		user, err := con.ProductCreate(&models.CreateProduct{
	// 			Name:  name,
	// 			Price: price,
	// 		})

	// 		if err != nil {
	// 			log.Println("product create err:", err)
	// 			continue
	// 		}

	// 		fmt.Printf("%+v\n", user)
	// 		fmt.Println("Product Created successfully")
	// 	} else if answer == "2" {
	// 		var id string
	// 		fmt.Printf("Enter id of user: ")
	// 		fmt.Scan(&id)
	// 		byid, err := con.GetByIdPoduct(&models.ProductPrimaryKey{Id: id})
	// 		if err != nil {
	// 			log.Printf("Error while GetById: %+v\n", err)
	// 			continue
	// 		}
	// 		fmt.Println(byid)
	// 	} else if answer == "3" {
	// 		var dataLimit int
	// 		var page int
	// 		var answer string
	// 		fmt.Printf("Input Data limit: ")
	// 		fmt.Scan(&dataLimit)
	// 		for {
	// 			fmt.Printf("Press e if you want end. Press n to continue: ")
	// 			fmt.Scan(&answer)
	// 			if answer == "n" {
	// 				fmt.Println("Input page:")
	// 				fmt.Scan(&page)
	// 				respProduct, err := con.ProductGetList(&models.ProductGetListRequest{
	// 					Offset: (page - 1) * dataLimit,
	// 					Limit:  dataLimit,
	// 				})

	// 				if err != nil {
	// 					fmt.Println(err)
	// 					continue
	// 				}

	// 				for _, product := range respProduct.Products {
	// 					fmt.Println(product)
	// 				}
	// 			} else if answer == "e" {
	// 				break
	// 			}
	// 		}
	// 	} else if answer == "4" {
	// 		var (
	// 			id           string
	// 			updatedname  string
	// 			updatedprice string
	// 		)
	// 		fmt.Printf("Enter id of user: ")
	// 		fmt.Scan(&id)
	// 		fmt.Printf("Enter New Name: ")
	// 		fmt.Scan(&updatedname)
	// 		fmt.Printf("Enter new Last Name: ")
	// 		fmt.Scan(&updatedprice)
	// 		product, err := con.ProductUpdate(&models.UpdateProduct{
	// 			Id:    id,
	// 			Name:  updatedname,
	// 			Price: updatedprice,
	// 		})

	// 		if err != nil {
	// 			fmt.Println(err)
	// 			continue
	// 		}

	// 		fmt.Printf("%+v\n", product)
	// 		fmt.Println("products information successfully updated !")
	// 	} else if answer == "5" {
	// 		var (
	// 			id string
	// 		)
	// 		fmt.Printf("Enter id of Product: ")
	// 		fmt.Scan(&id)
	// 		err = con.ProductDelete(&models.ProductPrimaryKey{Id: id})
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			continue
	// 		}
	// 		fmt.Println("Product successfully deleted !")
	// 	} else {
	// 		break
	// 	}
	// }
	// resp, err := con.OrderCreate(&models.CreateOrder{
	// 	UserId:   "13526574-a20f-4a43-b9ff-9aef94bcd639",
	// 	Sum:      0,
	// 	SumCount: 0,
	// 	Status:   "In process...",
	// })
	// fmt.Println(resp)
	// con.Strg.Order().CreteOrderItem(&models.OrderItem{
	// 	ProductId: "637ad914-8407-428f-be65-71b3af2418a8",
	// 	OrderId:   "0cad2f33-81d2-42b8-8703-9339abeb252d",
	// 	Count:     4,
	// })
	err = con.AddOrderItem(&models.OrderItem{
		ProductId: "637ad914-8407-428f-be65-71b3af2418a8",
		OrderId:   "0cad2f33-81d2-42b8-8703-9339abeb252d",
		Count:     4,
	})
}
