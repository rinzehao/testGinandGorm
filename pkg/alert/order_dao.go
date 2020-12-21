package alert

import (
	"github.com/garyburd/redigo/redis"
	redis__utils "testGinandGorm/common/redis _utils"
	"testGinandGorm/pkg/model"
)

type OrderDao interface {
	CreateOrder(s *model.OrderCtx) error
	Delete(id string) error
	UpdateByNo(no string,m map[string]interface{}) error
	QueryOrderById(id string) (*model.OrderMould, error)
	QueryOrderByNo(no string) (*model.OrderMould, error)
	QueryOrders(page, pageSize int) (orders []*model.OrderMould, err error)
	QueryOrdersByName(userName, orderBy, desc string) (orders []*model.OrderMould, err error)
	UpdateById(m map[string]interface{}, id string) error
}

type MyOrderDao struct {
	db OrderDB
	cache redis__utils.Cache
}

func NewMyOrderDao(db OrderDB ,cache redis.Conn) *MyOrderDao{
	return &MyOrderDao{db: db ,cache: redis__utils.NewRedisCache(1e10*6*2)}
}

func (dao *MyOrderDao) CreateOrder(s *model.OrderCtx) error  {


	return nil
}

func (dao *MyOrderDao) Delete(id string) error  {
	return nil
}

func (dao *MyOrderDao) UpdateByNo(no string,m map[string]interface{}) error  {
	return nil
}

func (dao *MyOrderDao) QueryOrderById(id string) (*model.OrderMould, error)  {
	return nil, nil
}

func (dao *MyOrderDao) QueryOrderByNo(no string) (*model.OrderMould, error)  {
	return nil, nil
}

func (dao *MyOrderDao) QueryOrders(page, pageSize int) (orders []*model.OrderMould, err error)  {
	return nil,nil
}

func (dao *MyOrderDao) QueryOrdersByName(userName, orderBy, desc string) (orders []*model.OrderMould, err error) {
	return nil,nil
}

func (dao *MyOrderDao) UpdateById(m map[string]interface{}, id string) error  {
	return nil
}