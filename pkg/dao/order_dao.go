package dao

import (
	redigo "github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"strconv"
	"testGinandGorm/common/logger"
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
	// ID -> OrderEntity
	// OrderNo -> OrderEntity
	if err := dao.cache.SetString(idPrefix+strconv.Itoa(s.ID), s); err != nil {
		return err
	}
	if err := dao.cache.SetString(noPrefix+s.OrderNo, s); err != nil {
		return err
	}
	return nil
}

func (dao *OrderDao) DeleteOrderById(id string) error {
	// step1 delete from db
	order, err := dao.db.QueryOrderById(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound || order == nil {
		return nil
	}
	if err := dao.db.DeleteById(id); err != nil {
		return err
	}
	// step2 delete from cache
	if _, err := dao.cache.Delete(idPrefix + id); err != nil {
		logger.Logger.Warn("delete target cache fail", zap.String("id", idPrefix+id), zap.Error(err))
	}
	if _, err := dao.cache.Delete(noPrefix + order.OrderNo); err != nil {
		logger.Logger.Warn("delete target cache fail", zap.String("orderNo", noPrefix+order.OrderNo), zap.Error(err))
	}
	return nil
}

func (dao *OrderDao) UpdateByNo(orderNo string, m map[string]interface{}) error {
	//step1 查询
	order, err := dao.db.QueryOrderByNo(orderNo)
	if err != nil {
		return err
	}
	//step2 删除cache
	if _, err := dao.cache.Delete(idPrefix + strconv.Itoa(order.ID)); err != nil {
		logger.Logger.Warn("delete target cache fail", zap.String("id", idPrefix+strconv.Itoa(order.ID)), zap.Error(err))
	}
	if _, err := dao.cache.Delete(noPrefix + orderNo); err != nil {
		logger.Logger.Warn("delete target cache fail", zap.String("orderNo", noPrefix+order.OrderNo), zap.Error(err))
	}
	//step3 DB更新
	if err := dao.db.UpdateByNo(orderNo, m); err != nil {
		return err
	}
	//step4 设置cache
	order, err = dao.db.QueryOrderByNo(orderNo)
	if err != nil {
		return err
	}
	if err := dao.cache.SetString(idPrefix+strconv.Itoa(order.ID), order); err != nil {
		return err
	}
	if err := dao.cache.SetString(noPrefix+orderNo, order); err != nil {
		return err
	}
	return nil
}

func (dao *OrderDao) QueryOrderById(id string) (order *model.Order, err error) {
	//step1 查询cache
	err = dao.cache.GetString(idPrefix+id, &order)
	if err != nil && err != redigo.ErrNil {
		logger.Logger.Warn("Query Order From Cache Fail", zap.String("id", idPrefix+id), zap.Error(err))
	}
	if order != nil {
		return order, nil
	}
	//step2 DB查询
	order, err = dao.db.QueryOrderById(id)
	if err != nil {
		return nil, err
	}
	//step3 写入cache
	if order != nil {
		if err := dao.cache.SetString(idPrefix+strconv.Itoa(order.ID), order); err != nil {
			logger.Logger.Warn("Set Order Cache Fail", zap.String("id", idPrefix+strconv.Itoa(order.ID)), zap.Error(err))
			return order, err
		}
		if err := dao.cache.SetString(noPrefix+order.OrderNo, order); err != nil {
			logger.Logger.Warn("Set Order Cache Fail", zap.String("orderNo", noPrefix+order.OrderNo), zap.Error(err))
			return order, err
		}
	}
	return order, nil
}

func (dao *OrderDao) QueryOrderByNo(orderNo string) (order *model.Order, err error) {
	//step1 查询cache
	err = dao.cache.GetString(noPrefix+orderNo, &order)
	if err != nil && err != redigo.ErrNil {
		logger.Logger.Warn("Query Order From Cache Fail", zap.String("orderNo", noPrefix+orderNo), zap.Error(err))
	}
	if order != nil {
		return order, nil
	}
	//step2 DB查询
	order, err = dao.db.QueryOrderByNo(orderNo)
	if err != nil {
		return nil, err
	}
	//step3 写入cache
	if order != nil {
		if err := dao.cache.SetString(idPrefix+strconv.Itoa(order.ID), order); err != nil {
			logger.Logger.Warn("Set Order Cache Fail", zap.String("id", idPrefix+strconv.Itoa(order.ID)), zap.Error(err))
			return order, err
		}
		if err := dao.cache.SetString(noPrefix+order.OrderNo, order); err != nil {
			logger.Logger.Warn("Set Order Cache Fail", zap.String("orderNo", noPrefix+order.OrderNo), zap.Error(err))
			return order, err
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
	//step1 查询
	order, err := dao.db.QueryOrderById(id)
	if err != nil {
		return err
	}
	//step2 删除cache
	if _, err := dao.cache.Delete(idPrefix + id); err != nil {
		logger.Logger.Warn("delete target cache fail", zap.String("id", idPrefix+id), zap.Error(err))
	}
	if _, err := dao.cache.Delete(noPrefix + order.OrderNo); err != nil {
		logger.Logger.Warn("delete target cache fail", zap.String("orderNo", noPrefix+order.OrderNo), zap.Error(err))
	}
	//step3 DB更新
	if err := dao.db.UpdateById(id, m); err != nil {
		return err
	}
	//step4 设置cache
	order, err = dao.db.QueryOrderById(id)
	if err != nil {
		return err
	}
	if err := dao.cache.SetString(idPrefix+id, order); err != nil {
		logger.Logger.Warn("Set Order Cache Fail", zap.String("id", idPrefix+strconv.Itoa(order.ID)), zap.Error(err))
		return err
	}
	if err := dao.cache.SetString(noPrefix+order.OrderNo, order); err != nil {
		logger.Logger.Warn("Set Order Cache Fail", zap.String("orderNo", noPrefix+order.OrderNo), zap.Error(err))
		return err
	}
	return nil
}
