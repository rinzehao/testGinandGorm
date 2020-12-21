package alert

import (
	"go.uber.org/zap"
	"log"
	"testGinandGorm/common/redis_utils"
	"testGinandGorm/pkg/model"
)

type OrderDao interface {
	CreateOrder(s *model.OrderCtx) error
	Delete(id string) error
	UpdateByNo(no string, m map[string]interface{}) error
	QueryOrderById(id string) (*model.OrderMould, error)
	QueryOrderByNo(no string) (*model.OrderMould, error)
	QueryOrders(page, pageSize int) (orders []*model.OrderMould, err error)
	QueryOrdersByName(userName, orderBy, desc string) (orders []*model.OrderMould, err error)
	UpdateById(m map[string]interface{}, id string) error
}

type MyOrderDao struct {
	db    OrderDB
	cache redis_utils.Cache
}

func NewMyOrderDao(db OrderDB, cache redis_utils.Cache) *MyOrderDao {
	return &MyOrderDao{db: db, cache: cache}
}

func (dao *MyOrderDao) CreateOrder(s *model.OrderCtx) error {
	return nil
}

func (dao *MyOrderDao) Delete(id string) error {
	return nil
}

func (dao *MyOrderDao) UpdateByNo(no string, m map[string]interface{}) error {
	return nil
}

func (dao *MyOrderDao) QueryOrderById(id string) (order *model.OrderMould, err error) {

	//step1 cache
	var flag bool
	if flag, err = dao.cache.Exist(id); err != nil {
		return nil, err
	}
	if flag == false {
		log.Println("queryTarget form cache fail")
	}
	if flag == true {
		if err = dao.cache.GetStruct(id, order); err != nil {
			return nil, err
		}
	}
	if order != nil {
		return order, nil
	}

	//step2 get db
	order, err = dao.db.QueryOrderById(id)
	if err != nil {
		return nil, err
	}
	//step3 set cache
	if order != nil {
		dao.cache.SetStruct(id, order)
	}

	return order, nil
}

func (dao *MyOrderDao) QueryOrderByNo(no string) (*model.OrderMould, error) {
	return nil, nil
}

func (dao *MyOrderDao) QueryOrders(page, pageSize int) (orders []*model.OrderMould, err error) {
	return nil, nil
}

func (dao *MyOrderDao) QueryOrdersByName(userName, orderBy, desc string) (orders []*model.OrderMould, err error) {
	return nil, nil
}

func (dao *MyOrderDao) UpdateById(m map[string]interface{}, id string) error {
	return nil
}
