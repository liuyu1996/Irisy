package services

import (
	"product/datamodels"
	"product/repositories"
)

type IOrderService interface {
	GetOrderByID(int64) (*datamodels.Order,error)
	DeleteOrderByID(int64) bool
	UpdateOrder(*datamodels.Order) error
	InsertOrder(*datamodels.Order) (int64 ,error)
	GetAllOrder() ([]*datamodels.Order, error)
	GetAllOrderInfo() (map[int]map[string]string, error)
}

func NewOrderService(orderRepository repositories.IOrderRepositories) IOrderService {
	return &OrderService{OrderRepository:orderRepository}
}

type OrderService struct {
	OrderRepository repositories.IOrderRepositories
}

func (o *OrderService) GetOrderByID(orderID int64) (order *datamodels.Order,err error) {
	return o.OrderRepository.SelectByKey(orderID)
}

func (o *OrderService) DeleteOrderByID(orderID int64) bool {
	return o.OrderRepository.Delete(orderID)
}

func (o *OrderService) UpdateOrder(order *datamodels.Order) (err error) {
	return o.OrderRepository.Update(order)
}

func (o *OrderService) InsertOrder(order *datamodels.Order) (ID int64, err error) {
	return o.OrderRepository.Insert(order)
}

func (o *OrderService) GetAllOrder() (orderArray []*datamodels.Order,err error) {
	return o.OrderRepository.SelectAll()
}

func (o *OrderService) GetAllOrderInfo() (orderMap map[int]map[string]string,err error) {
	return o.OrderRepository.SelectAllWithInfo()
}
