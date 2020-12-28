package db

import (
	"github.com/jinzhu/gorm"
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

type MyOrderDB struct {
	db *gorm.DB
}

func NewMyOrderDB(Db *gorm.DB) *MyOrderDB {
	return &MyOrderDB{db: Db}
}

func (d *MyOrderDB) CreateOrder(s *model.Order) error {
	return d.db.Create(s).Error
}

func (d *MyOrderDB) DeleteById(id string) error {
	return d.db.Where("id = ?", id).Delete(&model.Order{}).Error
}

func (d *MyOrderDB) UpdateByNo(no string, m map[string]interface{}) error {
	return d.db.Model(&model.Order{}).Where("order_no = ?", no).Update(m).Error
}

func (d *MyOrderDB) QueryOrderById(id string) (*model.Order, error) {
	var order model.Order
	return &order, d.db.Where("id = ?", id).First(&order).Error
}

func (d *MyOrderDB) QueryOrderByNo(no string) (*model.Order, error) {
	var order model.Order
	return &order, d.db.Where("order_no = ?", no).First(&order).Error
}

func (d *MyOrderDB) QueryOrders(page, pageSize int) (orders []*model.Order, err error) {
	return orders, d.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&orders).Error
}

func (d *MyOrderDB) QueryOrdersByName(userName, orderBy, desc string) (orders []*model.Order, err error) {
	return orders, d.db.Where("user_name LIKE ?", "%"+userName+"%").Order(orderBy + " " + desc).Find(&orders).Error
}

func (d *MyOrderDB) UpdateById(id string, m map[string]interface{}) error {
	tx := d.db.Begin()
	defer tx.Rollback()
	if err := d.db.Model(model.Order{}).Where("id = ?", id).Updates(m).Error; err != nil {
		return err
	}
	tx.Commit()
	return nil
}
