package controller

import (
	"errors"
	"fmt"
	"log"

	"app/config"
	"app/models"
	"app/pkg/convert"
)

func (c *Controller) OrderCreate(req *models.CreateOrder) (*models.Order, error) {

	log.Printf("User create req: %+v\n", req)

	req.Status = config.OrderStatus["0"]
	resp, err := c.Strg.Order().Create(req)
	if err != nil {
		log.Printf("error while order Create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) GetByIdOrder(req *models.OrderPrimaryKey) ([]*models.CreateOrderItem, error) {

	resp, err := c.Strg.Order().GetById(req)
	if err != nil {
		log.Printf("error while order GetById: %+v\n", err)
		return nil, err
	}

	return resp.OrderItems, nil
}

func (c *Controller) OrderGetList(req *models.OrderGetListRequest) (*models.OrderGetList, error) {

	resp, err := c.Strg.Order().GetList(req)
	if err != nil {
		log.Printf("error while order GetList: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) OrderUpdate(req *models.UpdateOrder) (*models.Order, error) {

	resp, err := c.Strg.Order().Update(req)
	if err != nil {
		log.Printf("error while order Update: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) OrderDelete(req *models.OrderPrimaryKey) error {

	err := c.Strg.Order().Delete(req)
	if err != nil {
		log.Printf("error while order Delete: %+v\n", err)
		return err
	}

	return nil
}

func (c *Controller) AddOrderItem(req *models.CreateOrderItem) error {

	product, err := c.Strg.Product().GetById(&models.ProductPrimaryKey{Id: req.ProductId})
	if err != nil {
		log.Printf("error while Product => GetById: %+v\n", err)
		return err
	}

	productPrice := product.Price
	if product.DiscountType == config.PercentDiscountType {
		productPrice = int(float64(product.Price) * ((100 - float64(product.Discount)) / 100))
	} else if product.DiscountType == config.FixDiscountType {
		productPrice = product.Price - product.Discount
	}

	req.TotalPrice = req.Count * productPrice
	err = c.Strg.Order().AddOrderItem(req)
	if err != nil {
		log.Printf("error while order => AddOrderItem: %+v\n", err)
		return err
	}

	order, err := c.Strg.Order().GetById(&models.OrderPrimaryKey{Id: req.OrderId})
	if err != nil {
		log.Printf("error while Order => GetById: %+v\n", err)
		return err
	}

	order.Sum += req.TotalPrice
	order.SumCount += req.Count

	var updateOrder models.UpdateOrder
	err = convert.ConvertStructToStruct(order, &updateOrder)
	if err != nil {
		log.Printf("error while convertStructToStruct: %+v\n", err)
		return err
	}

	_, err = c.Strg.Order().Update(&updateOrder)
	if err != nil {
		log.Printf("error while order => Update: %+v\n", err)
		return err
	}

	return nil
}

func (c *Controller) RemoveOrderItem(req *models.RemoveOrderItemPrimaryKey) error {

	err := c.Strg.Order().RemoveOrderItem(req)
	if err != nil {
		log.Printf("error while order => RemoveOrderItem: %+v\n", err)
		return err
	}

	return nil
}

func (c *Controller) OrderPayment(req *models.OrderPayment) error {

	order, err := c.Strg.Order().GetById(&models.OrderPrimaryKey{Id: req.OrderId})
	if err != nil {
		log.Printf("error while Order => GetById: %+v\n", err)
		return err
	}

	user, err := c.Strg.User().GetById(&models.UserPrimaryKey{Id: order.UserId})
	if err != nil {
		log.Printf("error while User => GetById: %+v\n", err)
		return err
	}

	if order.Sum > user.Balance {
		return errors.New("Not enough balance " + user.FirstName + " " + user.LastName)
	}

	order.Status = config.OrderStatus["1"]
	fmt.Println(order.Status)
	user.Balance -= order.Sum

	var updateOrder models.UpdateOrder
	err = convert.ConvertStructToStruct(order, &updateOrder)
	if err != nil {
		log.Printf("error while convertStructToStruct: %+v\n", err)
		return err
	}

	_, err = c.Strg.Order().Update(&updateOrder)
	if err != nil {
		log.Printf("error while order => Update: %+v\n", err)
		return err
	}

	var updateUser models.UpdateUser
	err = convert.ConvertStructToStruct(user, &updateUser)
	if err != nil {
		log.Printf("error while convertStructToStruct: %+v\n", err)
		return err
	}

	_, err = c.Strg.User().Update(&updateUser)
	if err != nil {
		log.Printf("error while User => Update: %+v\n", err)
		return err
	}

	return nil
}

func (c *Controller) UserHistory() (map[string]interface{}, error) {
	fmt.Println("User history")
	var (
		userHistoryList = make(map[string]interface{})
	)
	users, err := c.Strg.User().GetList(&models.UserGetListRequest{0, 0})
	if err != nil {
		return nil, err
	}
	orders, err := c.Strg.Order().GetList(&models.OrderGetListRequest{0, 0})
	if err != nil {
		return nil, err
	}
	for _, user := range users.Users {
		orderId, err := c.findUsers(orders, &models.UserPrimaryKey{user.Id})
		if err != nil {
			continue
		}
		orders, _ := c.Strg.Order().GetById(&models.OrderPrimaryKey{orderId})
		userHistoryList[user.FirstName] = c.findOrders(orders)
	}
	return userHistoryList, err
}
func (c *Controller) findOrders(req *models.Order) []models.UserHistory {
	fmt.Println("find orders")
	ords := []models.UserHistory{}
	for _, ord := range req.OrderItems {
		prod, _ := c.Strg.Product().GetById(&models.ProductPrimaryKey{ord.ProductId})
		mode := models.UserHistory{
			ProductName: prod.Name,
			Count:       ord.Count,
			TotalCount:  ord.TotalPrice,
		}
		ords = append(ords, mode)
	}
	return ords
}

func (c *Controller) findUsers(req *models.OrderGetList, user *models.UserPrimaryKey) (string, error) {
	for _, order := range req.Orders {
		if user.Id == order.UserId {
			return order.Id, nil
		}
	}
	return "", errors.New("id not found")
}
