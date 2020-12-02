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

// todo Create
func (orderDao *OrderDao) CreateOrder(s *model.DemoOrder) (err error) {
	if err = orderDao.db.Create(&s).Error; err != nil {
		return err
	}
	return err
}

// todo DeleteById
func (orderDao *OrderDao) DeleteById(id string) error {
	return orderDao.db.Where("id = ?", id).Delete(&model.DemoOrder{}, id).Error
}

// todo UpdateByOrderNo
func (orderDao *OrderDao) UpdateByOrderNo(orderNo string, s *model.DemoOrder) error {
	orderDao.db.LogMode(true)
	return orderDao.db.Model(&model.DemoOrder{}).Where("order_no = ?", orderNo).Update(&s).Error
}

// todo QueryOrderById
func (orderDao *OrderDao) QueryOrderById(id string) (order *model.DemoOrder, err error) {
	order = &model.DemoOrder{}
	if err = orderDao.db.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return order, err
}

// todo QueryOrderIsExistByOrderNo
func (OrderDao *OrderDao) QueryOrderIsExistByOrderNo(orderNo string) (isExit bool, err error) {
	if err = OrderDao.db.Where("order_no = ?", orderNo).Error; err != nil {
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

// todo QueryTranscationUpdateById
func (OrderDao *OrderDao) QueryTranscationUpdateById(url string, id string) error {

	var err error
	tx := OrderDao.db.Begin()
	defer tx.Rollback()
	if err = OrderDao.db.Model(model.DemoOrder{}).Where("id = ?", id).Update("file_url", url).Error; err != nil {
		return err
	} else if err = tx.Commit().Error; err != nil {
		return err
	}
	return err
}
