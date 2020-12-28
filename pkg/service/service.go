package service

import (
	"github.com/jinzhu/gorm"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
)

//type Service interface {
//	Schema () string
//	Delete(model.DeleteContext) error
//	Query(model.QueryContext) error
//	UpdateByNo(model.UpdateContext) error
//	Create(model.CreateContext) error
//	QueryOrders(model.QueryObjectsContext) error
//	QueryOrdersByName(model.QueryByNameContext) error
//	UpdateById(model.UpdateContext) error
//}

type OrderService struct {
	orderDao dao.OrderDao
}

func NewOrderService(dao dao.OrderDao) *OrderService {
	return &OrderService{orderDao: dao}
}

func (g *OrderService) Schema() string {
	return "order"
}

func (service *OrderService) Delete(ctx model.DeleteContext) error {
	return service.orderDao.DeleteOrderById(ctx.Param().(string))
}

func (service *OrderService) Query(ctx model.QueryContext) error {
	order, err := service.orderDao.QueryOrderById(ctx.Param().(string))
	if err != nil {
		return err
	}
	ctx.SetResult(order)
	return nil
}

func (service *OrderService) UpdateByNo(ctx model.UpdateContext) error {
	updateMap := ctx.Param().(map[string]interface{})
	return service.orderDao.UpdateByNo(ctx.GetIdentify(), updateMap)
}

func (service *OrderService) Create(ctx model.CreateContext) error {
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

func (service *OrderService) QueryOrders(ctx model.QueryObjectsContext) error {
	orders, err := service.orderDao.QueryOrders(ctx.Page(), ctx.PageSize())
	if err != nil {
		return err
	}
	var array []interface{}
	for _, v := range orders {
		array = append(array, v)
	}
	ctx.SetResult(array)
	return nil
}

func (service *OrderService) QueryOrdersByName(ctx model.QueryByNameContext) error {
	if ctx.Desc() {
		orders, err := service.orderDao.QueryOrdersByName(ctx.Param().(string), ctx.Order(), "DESC")
		if err != nil {
			return err
		}
		var array []interface{}
		for _, v := range orders {
			array = append(array, v)
		}
		ctx.SetResult(array)
	}
	if !ctx.Desc() {
		orders, err := service.orderDao.QueryOrdersByName(ctx.Param().(string), ctx.Order(), "ASC")
		if err != nil {
			return err
		}
		var array []interface{}
		for _, v := range orders {
			array = append(array, v)
		}
		ctx.SetResult(array)
	}
	return nil
}

func (service *OrderService) UpdateById(ctx model.UpdateContext) error {
	ctx.SetResult(ctx.Param().(map[string]interface{}))
	return service.orderDao.UpdateById(ctx.GetIdentify(), ctx.Param().(map[string]interface{}))
}
