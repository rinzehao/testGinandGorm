package service

import (
	"github.com/jinzhu/gorm"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
)

type OrderService interface {
	DeleteOrderById(*model.OrderMade) error
	QueryOrderById(*model.OrderMade) error
	UpdateByOrderNo(*model.OrderMade) error
	CreateOrder(*model.OrderMade) error
	QueryOrders(*model.OrderMade) error
	QueryOrdersByName(*model.OrderMade) error
	UpdateById(*model.OrderMade) error
}

type MyOrderService struct {
	orderDao dao.OrderDao
}

func NewOrderService(dao dao.OrderDao) *MyOrderService {
	return &MyOrderService{orderDao: dao}
}

func (service *MyOrderService) DeleteOrderById(ctx *model.OrderMade) error {
	return service.orderDao.DeleteOrderById(ctx.OrderID)
}

func (service *MyOrderService) QueryOrderById(ctx *model.OrderMade) error {
	order, err := service.orderDao.QueryOrderById(ctx.OrderID)
	if err != nil {
		return err
	}
	ctx.Order = order
	return nil
}

func (service *MyOrderService) UpdateByOrderNo(ctx *model.OrderMade) error {
	return service.orderDao.UpdateByNo(ctx.GetOrderNo(), ctx.GetUpdateMap())
}

func (service *MyOrderService) CreateOrder(ctx *model.OrderMade) error {
	_, err := service.orderDao.QueryOrderByNo(ctx.OrderNo)
	if err == gorm.ErrRecordNotFound {
		return service.orderDao.CreateOrder(ctx.Order.(*model.DemoOrder))
	}
	if err != nil {
		return err
	}
	return nil
}

func (service *MyOrderService) QueryOrders(ctx *model.OrderMade) error {
	orders, err := service.orderDao.QueryOrders(ctx.Page, ctx.PageSize)
	if err != nil {
		return err
	}
	ctx.Group = orders
	return nil
}

func (service *MyOrderService) QueryOrdersByName(ctx *model.OrderMade) error {
	orders, err := service.orderDao.QueryOrdersByName(ctx.UserName, ctx.OrderBy, ctx.Desc)
	if err != nil {
		return err
	}
	ctx.Group = orders
	return nil
}

func (service *MyOrderService) UpdateById(ctx *model.OrderMade) error {
	return service.orderDao.UpdateById(ctx.OrderID, ctx.UpdateMap)
}
