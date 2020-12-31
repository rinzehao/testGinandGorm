package profile_item

import (
	"github.com/jinzhu/gorm"
	"testGinandGorm/pkg/model"
	"testGinandGorm/pkg/service/profile"
)

type OrderDao interface {
	CreateOrder(s *model.Order) error
	DeleteOrderById(id string) error
	UpdateByNo(no string, m map[string]interface{}) error
	QueryOrderById(id string) (*model.Order, error)
	QueryOrderByNo(no string) (*model.Order, error)
	QueryOrders(page, pageSize int) (orders []*model.Order, err error)
	QueryOrdersByName(userName, orderBy, desc string) (orders []*model.Order, err error)
	UpdateById(id string, m map[string]interface{}) error
}

type OrderService struct {
	orderDao OrderDao
}

func NewOrderService(dao OrderDao) *OrderService {
	return &OrderService{orderDao: dao}
}

func (service *OrderService) Schema() string {
	return "order"
}

func (service *OrderService) Delete(ctx profile.DeleteContext) error {
	return service.orderDao.DeleteOrderById(ctx.Param().(string))
}

func (service *OrderService) Query(ctx profile.QueryContext) error {
	order, err := service.orderDao.QueryOrderById(ctx.Param().(string))
	if err != nil {
		return err
	}
	ctx.SetResult(order)
	return nil
}

func (service *OrderService) UpdateByNo(ctx profile.UpdateContext) error {
	updateMap := ctx.Param().(map[string]interface{})
	return service.orderDao.UpdateByNo(ctx.GetIdentify(), updateMap)
}

func (service *OrderService) Create(ctx profile.CreateContext) error {
	order := ctx.Param().(model.Order)
	_, err := service.orderDao.QueryOrderByNo(order.OrderNo)
	if err == gorm.ErrRecordNotFound {
		if err := service.orderDao.CreateOrder(&order); err != nil {
			return err
		}
		ctx.SetResult(order)
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func (service *OrderService) QueryOrders(ctx profile.QueryObjectsContext) error {
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

func (service *OrderService) QueryOrdersByName(ctx profile.QueryByNameContext) error {
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

func (service *OrderService) UpdateById(ctx profile.UpdateContext) error {
	ctx.SetResult(ctx.Param().(map[string]interface{}))
	return service.orderDao.UpdateById(ctx.GetIdentify(), ctx.Param().(map[string]interface{}))
}
