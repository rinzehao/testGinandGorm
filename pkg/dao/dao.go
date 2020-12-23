package dao

import (
	"log"
	"strconv"
	"testGinandGorm/common/redis_utils"
	"testGinandGorm/pkg/db"
	"testGinandGorm/pkg/model"
)

type OrderDao interface {
	CreateOrder(s *model.DemoOrder) error
	DeleteOrderById(id string) error
	UpdateByNo(no string, m map[string]interface{}) error
	QueryOrderById(id string) (*model.DemoOrder, error)
	QueryOrderByNo(no string) (*model.DemoOrder, error)
	QueryOrders(page, pageSize int) (orders []*model.DemoOrder, err error)
	QueryOrdersByName(userName, orderBy, desc string) (orders []*model.DemoOrder, err error)
	UpdateById(id string, m map[string]interface{}) error
}

type MyOrderDao struct {
	db    db.OrderDB
	cache *redis_utils.Cache
}

func NewMyOrderDao(db db.OrderDB, cache *redis_utils.Cache) *MyOrderDao {
	return &MyOrderDao{db: db, cache: cache}
}

func (dao *MyOrderDao) CreateOrder(s *model.DemoOrder) error {
	// step1 cache
	// ID -> Hash Hash:orderNo->order
	if err := dao.cache.HSet("orderHash", strconv.Itoa(s.ID), s.OrderNo); err != nil {
		return err
	}
	if err := dao.cache.SetStruct(s.OrderNo, s); err != nil {
		return err
	}
	// step2 db
	if err := dao.db.CreateOrder(s); err != nil {
		return err
	}
	return nil
}

func (dao *MyOrderDao) DeleteOrderById(id string) error {
	//step1 cache
	_, err := dao.cache.Hdel("orderHash", id)
	if err != nil {
		return err
	}
	//step2 db
	if err = dao.db.DeleteById(id); err != nil {
		log.Println("delete order from db failed ")
		return err
	}
	return nil
}

func (dao *MyOrderDao) UpdateByNo(no string, m map[string]interface{}) error {
	//step1 cache淘汰
	_, err := dao.cache.Hdel("orderHash", no)
	if err != nil {
		return err
	}
	//step2 db写入
	if err := dao.db.UpdateByNo(no, m); err != nil {
		return err
	}
	//step3 cache写入
	order, err := dao.db.QueryOrderByNo(no)
	if err != nil {
		return err
	}
	if err := dao.cache.SetStruct(no, order); err != nil {
		return err
	}
	return nil
}

func (dao *MyOrderDao) QueryOrderById(id string) (order *model.DemoOrder, err error) {
	//step1 cache
	var flag bool
	if flag, err = dao.cache.HExists("orderHash", id); err != nil {
		return nil, err
	}
	if flag == true {
		if err = dao.cache.GetStruct(id, &order); err != nil {
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

func (dao *MyOrderDao) QueryOrderByNo(no string) (order *model.DemoOrder, err error) {
	//step1 cache
	var flag bool
	if flag, err = dao.cache.HExists("orderHash", no); err != nil {
		return nil, err
	}
	if flag == false {
		log.Println("queryTarget form cache fail")
	}
	if flag == true {
		if err = dao.cache.GetStruct(no, order); err != nil {
			return nil, err
		}
	}
	if order != nil {
		return order, nil
	}
	//step2 get db
	order, err = dao.db.QueryOrderByNo(no)
	if err != nil {
		return nil, err
	}
	//step3 set cache
	if order != nil {
		dao.cache.SetStruct(no, order)
	}
	return order, nil
}

func (dao *MyOrderDao) QueryOrders(page, pageSize int) (orders []*model.DemoOrder, err error) {
	return dao.db.QueryOrders(page, pageSize)
}

func (dao *MyOrderDao) QueryOrdersByName(userName, orderBy, desc string) (orders []*model.DemoOrder, err error) {
	return dao.db.QueryOrdersByName(userName, orderBy, desc)
}

func (dao *MyOrderDao) UpdateById(id string, m map[string]interface{}) error {
	//step1 db写入
	if err := dao.db.UpdateById(id, m); err != nil {
		return err
	}
	//step2 cache淘汰
	order, err := dao.db.QueryOrderById(id)
	if err != nil {
		return err
	}
	_, err = dao.cache.Delete(order)
	if err != nil {
		return err
	}
	_, err = dao.cache.Hdel("orderHash", id)
	if err != nil {
		return err
	}
	//step3 cache写入
	if err := dao.cache.SetStruct(order.OrderNo, order); err != nil {
		return err
	}
	if err := dao.cache.SetStruct(id, order.OrderNo); err != nil {
		return err
	}
	return nil
}
