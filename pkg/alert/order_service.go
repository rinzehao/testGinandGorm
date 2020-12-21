package alert

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"testGinandGorm/pkg/model"
)

type orderService interface {
	Delete(ctx model.OrderCtx) error
	QueryID(ctx model.OrderCtx) error
	UpdateNo(ctx model.OrderCtx) error
	Create(ctx model.OrderCtx) error
	QueryMulti(ctx model.OrderCtx) error
	QueryName(ctx model.OrderCtx) error
	UpdateID(ctx model.OrderCtx) error
}

type MyOrderService struct{
	orderDao MyOrderDao
}

func NewOrderService(dao MyOrderDao) *MyOrderService {
	return &MyOrderService{orderDao: dao }
}

func (service *MyOrderService) Delete(ctx model.OrderCtx) error{
	return nil
}

func (service *MyOrderService) QueryID(ctx model.OrderCtx) error{
	return nil
}

func (service *MyOrderService) UpdateNo(ctx model.OrderCtx) error{
	return nil
}

func (service *MyOrderService) Create(ctx model.OrderCtx) error{
	if _, err :=service.orderDao.cache.Do("Get","") ;err !=nil {
		fmt.Println("cache no found", err)
	}
	_, err := service.orderDao.QueryOrderByNo(ctx.OrderNo_())
	if err == gorm.ErrRecordNotFound {
		return service.orderDao.CreateOrder(&ctx)
	}
	if err != nil {
		return err
	}
	return nil
}

func (service *MyOrderService) QueryMulti(ctx model.OrderCtx) error{
	return nil
}

func (service *MyOrderService) QueryName(ctx model.OrderCtx) error{
	return nil
}

func (service *MyOrderService) UpdateID(ctx model.OrderCtx) error{
	return nil
}