package dao

import (
	"log"
	"strconv"
	"testGinandGorm/common/redis"
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
	cache *redis.Cache
}

const (
	id_prefix = "OrderID:"
	no_prefix = "OrderNo:"
)

func NewMyOrderDao(db db.OrderDB, cache *redis.Cache) *MyOrderDao {
	return &MyOrderDao{db: db, cache: cache}
}

func (dao *MyOrderDao) CreateOrder(s *model.DemoOrder) error {
	// step1 写入db
	if err := dao.db.CreateOrder(s); err != nil {
		return err
	}
	// step2 写入cache
	// ID -> OrderNo OrderNo->OrderEntity
	if err := dao.cache.SetString(id_prefix+strconv.Itoa(s.ID), no_prefix+s.OrderNo); err != nil {
		return err
	}
	if err := dao.cache.SetString(no_prefix+s.OrderNo, s); err != nil {
		return err
	}
	return nil
}

func (dao *MyOrderDao) DeleteOrderById(id string) error {
	//step1 cache
	var cacheNoKey string
	dao.cache.GetString(id_prefix+id, &cacheNoKey)
	if _, err := dao.cache.Delete(cacheNoKey); err != nil {
		return err
	}
	if _, err := dao.cache.Delete(id_prefix + id); err != nil {
		return err
	}

	//step2 mySQL_db
	if err := dao.db.DeleteById(id); err != nil {
		log.Println("delete order from mySQL_db failed ")
		return err
	}
	return nil
}

func (dao *MyOrderDao) UpdateByNo(orderNo string, m map[string]interface{}) error {
	//step1 cache淘汰
	if _, err := dao.cache.Delete(no_prefix + orderNo); err != nil {
		return err
	}
	//step2 db写入
	if err := dao.db.UpdateByNo(orderNo, m); err != nil {
		return err
	}
	//step3 cache写入
	order, err := dao.db.QueryOrderByNo(orderNo)
	if err != nil {
		return err
	}
	if err := dao.cache.SetString(no_prefix+orderNo, order); err != nil {
		return err
	}
	return nil
}

func (dao *MyOrderDao) QueryOrderById(id string) (order *model.DemoOrder, err error) {
	//step1 get from cache
	flag, err := dao.cache.Exist(id_prefix + id)
	if err != nil {
		return nil, err
	}
	var cacheNoKey string
	if flag == true {
		if err := dao.cache.GetString(id_prefix+id, &cacheNoKey); err != nil {
			return nil, err
		}
		if flag, err = dao.cache.Exist(cacheNoKey); err != nil {
			return nil, err
		}
		if flag == true {
			if err = dao.cache.GetString(cacheNoKey, &order); err != nil {
				return nil, err
			}
			if order != nil {
				return order, nil
			}
		}
	}
	//step2 get from mySQL_db
	order, err = dao.db.QueryOrderById(id)
	if err != nil {
		return nil, err
	}
	//step3 set cache
	if order != nil {
		dao.cache.SetString(no_prefix+order.OrderNo, order)
		dao.cache.SetString(id_prefix+id, no_prefix+order.OrderNo)
	}
	return order, nil
}

func (dao *MyOrderDao) QueryOrderByNo(OrderNo string) (order *model.DemoOrder, err error) {
	//step1 get from cache
	flag, err := dao.cache.Exist(no_prefix + OrderNo)
	if err != nil {
		return nil, err
	}
	if flag == true {
		if err = dao.cache.GetString(no_prefix+OrderNo, &order); err != nil {
			return nil, err
		}
		if order != nil {
			return order, nil
		}
	}
	//step2 get from mySQL_db
	order, err = dao.db.QueryOrderByNo(OrderNo)
	if err != nil {
		return nil, err
	}
	//step3 set cache
	if order != nil {
		if err := dao.cache.SetString(no_prefix+OrderNo, order); err != nil {
			return nil, err
		}
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
	var cacheNoKey string
	if err := dao.cache.GetString(id_prefix+id, &cacheNoKey); err != nil {
		return nil
	}
	if _, err = dao.cache.Delete(cacheNoKey); err != nil {
		return err
	}
	//step3 cache写入
	if err := dao.cache.SetString(cacheNoKey, order); err != nil {
		return err
	}
	return nil
}
