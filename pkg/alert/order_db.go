package alert

import (
	"github.com/jinzhu/gorm"
	"testGinandGorm/pkg/model"
)

type MyOrderDB interface {
	CreateOrder(s *model.OrderMould) error
	DeleteById(id string) error
	UpdateByNo(no string,m map[string]interface{}) error
	QueryOrderById(id string) (*model.OrderMould, error)
	QueryOrderByNo(no string) (*model.OrderMould, error)
	QueryOrders(page, pageSize int) (orders []*model.OrderMould, err error)
	QueryOrdersByName(userName, orderBy, desc string) (orders []*model.OrderMould, err error)
	UpdateById(m map[string]interface{}, id string) error
}

type OrderDB struct {
	db * gorm.DB
}

func NewOrderDB (Db *gorm.DB) *OrderDB {
	return &OrderDB{db: Db}
}

func (d *OrderDB) CreateOrder(s *model.OrderMould) error {
	return d.db.Create(s).Error
}

func (d *OrderDB) DeleteById(id string) error {
	return d.db.Where("id = ?", id).Delete(&model.OrderMould{}).Error
}

func (d *OrderDB ) UpdateByNo(no string,m map[string]interface{}) error  {
	return d.db.Model(&model.OrderMould{}).Where("order_no = ?", no).Update(m).Error
}

func (d *OrderDB) QueryOrderById(id string) (*model.OrderMould, error) {
	var order model.OrderMould
	return &order, d.db.Where("id = ?", id).First(&order).Error
}


func (d *OrderDB) QueryOrderByNo(no string) (*model.OrderMould, error) {
	var order model.OrderMould
	return &order, d.db.Where("order_no = ?", no).First(&order).Error
}


func (d *OrderDB) QueryOrders(page, pageSize int) (orders []*model.OrderMould, err error) {
	return orders, d.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&orders).Error
}


func (d *OrderDB) QueryOrdersByName(userName, orderBy, desc string) (orders []*model.OrderMould, err error) {
	return orders, d.db.Where("user_name LIKE ?", "%"+userName+"%").Order(orderBy + " " + desc).Find(&orders).Error
}

func (d *OrderDB) UpdateById(m map[string]interface{}, id string) error {
	tx := d.db.Begin()
	defer tx.Rollback()
	if err := d.db.Model(model.OrderMould{}).Where("id = ?", id).Updates(m).Error; err != nil {
		return err
	}
	tx.Commit()
	return nil
}

