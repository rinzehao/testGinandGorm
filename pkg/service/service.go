package service

import (
	"github.com/jinzhu/gorm"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
)

type Service interface {
	DeleteOrderById(model.DeleteContext) error
	QueryOrderById(model.QueryContext) error
	UpdateByOrderNo(model.UpdateContext) error
	CreateOrder(model.CreateContext) error
	QueryOrders(model.QueryOrdersContext) error
	QueryOrdersByName(*model.OrderMade) error
	UpdateById(model.UpdateContext) error
}

type OrderService struct {
	orderDao dao.OrderDao
}

func NewOrderService(dao dao.OrderDao) *OrderService {
	return &OrderService{orderDao: dao}
}

func (service *OrderService) DeleteOrderById(ctx model.DeleteContext) error {
	return service.orderDao.DeleteOrderById(ctx.Param().(string))
}

func (service *OrderService) QueryOrderById(ctx model.QueryContext) error {
	order, err := service.orderDao.QueryOrderById(ctx.Param().(string))
	if err != nil {
		return err
	}
	ctx.SetResult(order)
	return nil
}

func (service *OrderService) UpdateByOrderNo(ctx model.UpdateContext) error {
	updateMap := ctx.Param().(map[string]interface{})
	return service.orderDao.UpdateByNo(ctx.GetIdentify(), updateMap)
}

func (service *OrderService) CreateOrder(ctx model.CreateContext) error {
	order := ctx.Param().(model.DemoOrder)
	_, err := service.orderDao.QueryOrderByNo(order.OrderNo)
	if err == gorm.ErrRecordNotFound {
		ctx.SetResult(order)
		return service.orderDao.CreateOrder(&order)
	}
	if err != nil {
		return err
	}
	return nil
}

func (service *OrderService) QueryOrders(ctx model.QueryOrdersContext) error {
	orders, err := service.orderDao.QueryOrders(ctx.Page().(int), ctx.PageSize().(int))
	if err != nil {
		return err
	}
	var array []interface{}
	for i, v := range orders{
		array[i] =v
	}
	ctx.SetResult(array)
	return nil
}

func (service *OrderService) QueryOrdersByName(ctx *model.OrderMade) error {
	orders, err := service.orderDao.QueryOrdersByName(ctx.UserName, ctx.OrderBy, ctx.Desc)
	if err != nil {
		return err
	}
	ctx.Group = orders
	return nil
}

func (service *OrderService) UpdateById(ctx model.UpdateContext) error {
	ctx.SetResult(ctx.Param().(map[string]interface{}))
	return service.orderDao.UpdateById(ctx.GetIdentify(), ctx.Param().(map[string]interface{}))
}
