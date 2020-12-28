package profile

import (
	"github.com/jinzhu/gorm"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
	"testGinandGorm/pkg/service"
)


type OrderService struct {
	orderDao dao.OrderDao
}

func NewOrderService(dao dao.OrderDao) *OrderService {
	return &OrderService{orderDao: dao}
}

func (g *OrderService) Schema() string {
	return "order"
}

func (service *OrderService) Delete(ctx service.DeleteContext) error {
	return service.orderDao.DeleteOrderById(ctx.Param().(string))
}

func (service *OrderService) Query(ctx service.QueryContext) error {
	order, err := service.orderDao.QueryOrderById(ctx.Param().(string))
	if err != nil {
		return err
	}
	ctx.SetResult(order)
	return nil
}

func (service *OrderService) UpdateByNo(ctx service.UpdateContext) error {
	updateMap := ctx.Param().(map[string]interface{})
	return service.orderDao.UpdateByNo(ctx.GetIdentify(), updateMap)
}

func (service *OrderService) Create(ctx service.CreateContext) error {
	order := ctx.Param().(model.Order)
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

func (service *OrderService) QueryOrders(ctx service.QueryObjectsContext) error {
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

func (service *OrderService) QueryOrdersByName(ctx service.QueryByNameContext) error {
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

func (service *OrderService) UpdateById(ctx service.UpdateContext) error {
	ctx.SetResult(ctx.Param().(map[string]interface{}))
	return service.orderDao.UpdateById(ctx.GetIdentify(), ctx.Param().(map[string]interface{}))
}
