package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
)

type OrderService struct {
	orderDao *dao.OrderDao
}

func NewService(dao *dao.OrderDao) *OrderService {
	return &OrderService{orderDao: dao}
}

func (service *OrderService) DeleteOrderById(id string) error {
	return service.orderDao.DeleteById(id)
}

func (service *OrderService) QueryOrderById(id string) (order *model.DemoOrder, err error) {
	order, err = service.orderDao.QueryOrderById(id)
	return order, err
}

func (service *OrderService) UpdateByOrderNo(m map[string]interface{}, orderNo string) error {
	return service.orderDao.UpdateByNo(orderNo, m)
}

func (service *OrderService) CreateOrder(order *model.DemoOrder) error {
	_, err := service.orderDao.QueryOrderByNo(order.OrderNo)
	if err == gorm.ErrRecordNotFound {
		return service.orderDao.CreateOrder(order)
	}
	if err != nil {
		return err
	}
	return nil
}

func (service *OrderService) QueryOrders(page, pageSize int) (orders []*model.DemoOrder, err error) {
	//页查询 page为页数 pagesize为单页展示条目数量 默认page=1 pagesize=100
	//当page的页数小于等于零的时候  offset不生效
	orders, err = service.orderDao.QueryOrders(page, pageSize)
	return orders, err
}

//根据user_name做模糊查找、根据创建时间、金额排序
func (service *OrderService) QueryOrdersByName(userName, orderBy, desc string) (orders []*model.DemoOrder, err error) {
	return service.orderDao.QueryOrdersByName(userName, orderBy, desc)
}

//获取改变信息并保存
func (service *OrderService) UpdateById(m map[string]interface{}, id string) error {
	return service.orderDao.UpdateById(m,id)
}