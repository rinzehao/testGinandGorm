package alert

import (
	"github.com/jinzhu/gorm"
	"testGinandGorm/pkg/model"
)

type orderService interface {
	Delete( *model.OrderMade) error
	QueryID( *model.OrderMade) error
	UpdateNo( *model.OrderMade) error
	Create( *model.OrderMade) error
	QueryMulti( *model.OrderMade) error
	QueryName( *model.OrderMade) error
	UpdateID( *model.OrderMade) error
}

type MyOrderService struct{
	orderDao MyOrderDao
}

func NewOrderService(dao MyOrderDao) *MyOrderService {
	return &MyOrderService{orderDao: dao }
}

func (service *MyOrderService) Delete(ctx *model.OrderMade) error{
	return service.orderDao.Delete(ctx.OrderID)
}

func (service *MyOrderService) QueryID(ctx *model.OrderMade) error{
	order, err := service.orderDao.QueryOrderById(ctx.OrderID)
	if err != nil {
		return err
	}
	ctx.Order =order
	return nil
}

func (service *MyOrderService) UpdateNo(ctx *model.OrderMade) error{
	return service.orderDao.UpdateByNo(ctx.GetOrderNo(), ctx.GetUpdateMap())
}

func (service *MyOrderService) Create(ctx *model.OrderMade) error{
	_, err := service.orderDao.QueryOrderByNo(ctx.OrderNo)
	if err == gorm.ErrRecordNotFound {
		val := ctx.GetOrder().(*model.DemoOrder)
		return service.orderDao.CreateOrder(val)
	}
	if err != nil {
		return err
	}
	return nil
}

func (service *MyOrderService) QueryMulti(ctx *model.OrderMade) error{
	orders, err := service.orderDao.QueryOrders(ctx.Page, ctx.PageSize)
	if err != nil {
		return err
	}
	ctx.Group=orders
	return nil
}

func (service *MyOrderService) QueryName(ctx *model.OrderMade) error{
	orders, err := service.orderDao.QueryOrdersByName(ctx.UserName,ctx.OrderBy,ctx.Desc)
	if err != nil {
		return err
	}
	ctx.Group=orders
	return nil
}

func (service *MyOrderService) UpdateID(ctx *model.OrderMade) error{
	return service.orderDao.UpdateById(ctx.OrderID, ctx.UpdateMap)
}