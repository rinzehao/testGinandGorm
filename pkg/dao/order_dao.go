package dao

import (
	"log"
	"strconv"
	"testGinandGorm/common/redis"
	"testGinandGorm/pkg/model"
)

type OrderDB interface {
	CreateOrder(s *model.Order) error
	DeleteById(id string) error
	UpdateByNo(no string, m map[string]interface{}) error
	QueryOrderById(id string) (*model.Order, error)
	QueryOrderByNo(no string) (*model.Order, error)
	QueryOrders(page, pageSize int) (orders []*model.Order, err error)
	QueryOrdersByName(userName, orderBy, desc string) (orders []*model.Order, err error)
	UpdateById(id string, m map[string]interface{}) error
}

type OrderDao struct {
	db    OrderDB
	cache *redis.Cache
}

const (
	idPrefix = "OrderID:"
	noPrefix = "OrderNo:"
)

func NewOrderDao(db OrderDB, cache *redis.Cache) *OrderDao {
	return &OrderDao{db: db, cache: cache}
}

func (dao *OrderDao) CreateOrder(s *model.Order) error {
	// step1 写入db
	if err := dao.db.CreateOrder(s); err != nil {
		return err
	}
	// step2 写入cache
	// ID -> OrderNo OrderNo->OrderEntity
	if err := dao.cache.SetString(idPrefix+strconv.Itoa(s.ID), noPrefix+s.OrderNo); err != nil {
		return err
	}
	if err := dao.cache.SetString(noPrefix+s.OrderNo, s); err != nil {
		return err
	}
	return nil
}

func (dao *OrderDao) DeleteOrderById(id string) error {
	//step1 cache
	var cacheNoKey string
	dao.cache.GetString(idPrefix+id, &cacheNoKey)
	if _, err := dao.cache.Delete(cacheNoKey); err != nil {
		return err
	}
	if _, err := dao.cache.Delete(idPrefix + id); err != nil {
		return err
	}

	//step2 mySQL
	if err := dao.db.DeleteById(id); err != nil {
		log.Println("delete order from mySQL failed ")
		return err
	}
	return nil
}

func (dao *OrderDao) UpdateByNo(orderNo string, m map[string]interface{}) error {
	//step1 cache淘汰
	if _, err := dao.cache.Delete(noPrefix + orderNo); err != nil {
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
	if err := dao.cache.SetString(noPrefix+orderNo, order); err != nil {
		return err
	}
	return nil
}

func (dao *OrderDao) QueryOrderById(id string) (order *model.Order, err error) {
	//step1 get from cache
	flag, err := dao.cache.Exist(idPrefix + id)
	if err != nil {
		return nil, err
	}
	var cacheNoKey string
	if flag == true {
		if err := dao.cache.GetString(idPrefix+id, &cacheNoKey); err != nil {
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
	//step2 get from mySQL
	order, err = dao.db.QueryOrderById(id)
	if err != nil {
		return nil, err
	}
	//step3 set cache
	if order != nil {
		dao.cache.SetString(noPrefix+order.OrderNo, order)
		dao.cache.SetString(idPrefix+id, noPrefix+order.OrderNo)
	}
	return order, nil
}

func (dao *OrderDao) QueryOrderByNo(OrderNo string) (order *model.Order, err error) {
	//step1 get from cache
	flag, err := dao.cache.Exist(noPrefix + OrderNo)
	if err != nil {
		return nil, err
	}
	if flag == true {
		if err = dao.cache.GetString(noPrefix+OrderNo, &order); err != nil {
			return nil, err
		}
		if order != nil {
			return order, nil
		}
	}
	//step2 get from mySQL
	order, err = dao.db.QueryOrderByNo(OrderNo)
	if err != nil {
		return nil, err
	}
	//step3 set cache
	if order != nil {
		if err := dao.cache.SetString(noPrefix+OrderNo, order); err != nil {
			return nil, err
		}
	}
	return order, nil
}

func (dao *OrderDao) QueryOrders(page, pageSize int) (orders []*model.Order, err error) {
	return dao.db.QueryOrders(page, pageSize)
}

func (dao *OrderDao) QueryOrdersByName(userName, orderBy, desc string) (orders []*model.Order, err error) {
	return dao.db.QueryOrdersByName(userName, orderBy, desc)
}

func (dao *OrderDao) UpdateById(id string, m map[string]interface{}) error {
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
	if err := dao.cache.GetString(idPrefix+id, &cacheNoKey); err != nil {
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
