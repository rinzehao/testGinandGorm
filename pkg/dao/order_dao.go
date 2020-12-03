package dao

import (
	"github.com/jinzhu/gorm"
	_ "testGinandGorm/db"
	"testGinandGorm/pkg/model"
)

type OrderDao struct {
	db *gorm.DB
}

// todo NewOrderDao
func NewOrderDao(Db *gorm.DB) *OrderDao {
	return &OrderDao{db: Db}
}

// todo CreateOrder
func (orderDao *OrderDao) CreateOrder(s *model.DemoOrder) (err error) {
	if err = orderDao.db.Create(&s).Error; err != nil {
		return err
	}
	return err
}

// todo DeleteById
func (orderDao *OrderDao) DeleteById(id string) error {
	orderDao.db.LogMode(true)
	return orderDao.db.Where("id = ?", id).Delete(&model.DemoOrder{}, id).Error
}

// todo UpdateByOrderNo
//func (orderDao *OrderDao) UpdateByOrderNo(orderNo string, s *model.DemoOrder) error {
//	orderDao.db.LogMode(true)
//	return orderDao.db.Model(&model.DemoOrder{}).Where("order_no = ?", orderNo).Update(&s).Error
//}

// todo UpdateByParam
func (orderDao *OrderDao ) UpdateByParam(ordermap map[string]string,paramName string , s *model.DemoOrder) error  {
	orderDao.db.LogMode(true)
	return orderDao.db.Model(&model.DemoOrder{}).Where(paramName+" =?",ordermap[paramName]).Update(&s).Error
}


// todo QueryOrderById
func (orderDao *OrderDao) QueryOrderById(id string) (order *model.DemoOrder, err error) {
	order = &model.DemoOrder{}
	if err = orderDao.db.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return order, err
}

// todo QueryOrderIsExist
func (orderDao *OrderDao) QueryOrderIsExist(m map[string]string ,queryParam string,order *model.DemoOrder) (isExit bool, err error) {
	if err = orderDao.db.Where(queryParam+ " = ?", m[queryParam]).First(&order).Error; err == gorm.ErrRecordNotFound {
		return false, err
	}
	return true, err
}

// todo QueryOrders
func (orderDao *OrderDao) QueryOrders(page, pageSize int) (orders []*model.DemoOrder, err error) {
	orderDao.db.LogMode(true)
	if err = orderDao.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, err
}

// todo QuerySortedOrdersByUserName
func (orderDao *OrderDao) QuerySortedOrdersByUserName(UserName, orderBy, desc string) (orders []*model.DemoOrder, err error) {

	orderDao.db.LogMode(true)

	if err = orderDao.db.Where("user_name LIKE ?", "%"+UserName+"%").Order(orderBy + " " + desc).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, err
}

// todo TranscationUpdateById
func (orderDao *OrderDao) TransactionUpdateById(updateMap map[string]string, id string) error {
	var k,v string
	for k, v = range updateMap{}
	tx := orderDao.db.Begin()
	defer orderDao.db.Rollback()
	orderDao.db.LogMode(true)
	if err := orderDao.db.Model(model.DemoOrder{}).Where("id = ?", id).Update(k, v).Error; err != nil {
		return err
	}
	return tx.Commit().Error
}
