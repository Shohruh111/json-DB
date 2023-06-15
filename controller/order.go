package controller

import (
	"errors"
	"log"

	"app/models"
)

func (c *Controller) OrderCreate(req *models.CreateOrder) (*models.Order, error) {

	log.Printf("User create req: %+v\n", req)

	resp, err := c.Strg.Order().Create(req)
	if err != nil {
		log.Printf("error while order Create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) GetByIdOrder(req *models.OrderPrimaryKey) (*models.Order, error) {

	resp, err := c.Strg.Order().GetById(req)
	if err != nil {
		log.Printf("error while order GetById: %+v\n", err)
		return nil, err
	}

	return resp, nil
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

func (c *Controller) AddOrderItem(req *models.OrderItem) error {
	// var (
	// 	orders = []models.Order{}
	// )
	prodId, err := c.Strg.Product().GetById(&models.ProductPrimaryKey{
		Id: req.ProductId,
	})
	if err != nil {
		return err
	}
	order, err := c.Strg.Order().GetById(&models.OrderPrimaryKey{
		Id: req.OrderId,
	})
	if err != nil {
		return err
	}

	req.TotalPrice = prodId.Price * req.Count

	TotalOrderPrice := req.TotalPrice + order.Sum

	TotalCount := req.Count + order.SumCount
	_, err = c.Strg.Order().Update(&models.UpdateOrder{
		Id:       req.OrderId,
		UserId:   order.UserId,
		Sum:      TotalOrderPrice,
		SumCount: TotalCount,
	})
	if err != nil {
		return err
	}

	c.Strg.Order().CreteOrderItem(req)
	// data, err := ioutil.ReadFile(c.Cfg.Path + c.Cfg.OrderFileName)
	// if err != nil {
	// 	return err
	// }
	// err = json.Unmarshal(data, &orders)
	// if err != nil {
	// 	return err
	// }

	// for _, order := range orders {
	// 	if order.Id == ord.OrderId {
	// 		order.Sum = order.Sum + ord.TotalPrice
	// 		order.SumCount += 1
	// 		order.Status = "In process..."
	// 		order.Orders = append(order.Orders, ord)
	// 	}
	// }
	// body, err := json.MarshalIndent(orders, "", "	")
	// if err != nil {
	// 	return err
	// }
	// err = ioutil.WriteFile(c.Cfg.Path+c.Cfg.OrderFileName, body, os.ModePerm)
	// if err != nil {
	// 	return err
	// }

	return nil
}