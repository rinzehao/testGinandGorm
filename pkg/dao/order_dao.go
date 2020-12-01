package dao

import (
	"github.com/jinzhu/gorm"
	_ "testGinandGorm/db"
	"testGinandGorm/pkg/model"
)

type OrderDao struct {
	db *gorm.DB
}

func NewOrderDao(Db *gorm.DB) *OrderDao {
	return &OrderDao{db: Db}
}

func (orderDao *OrderDao) Create(s *model.Demo_order) error {
	return orderDao.db.Create(&s).Error
}

func (orderDao *OrderDao) Delete(id string) error {
	return orderDao.db.Where("id = ?", id).Delete(&model.Demo_order{}, id).Error
}

func (orderDao *OrderDao) Update(s *model.Demo_order) error {
	return orderDao.db.Save(&s).Error
}

func (orderDao *OrderDao) GetOrder(id string, s *model.Demo_order) error {
	return orderDao.db.Where("id = ?", id).First(&s).Error
}

func (OrderDao *OrderDao) FindOrder(id string) error {
	return OrderDao.db.Raw("select * from demo_order where id = ? " + id).Error
}

func (orderDao *OrderDao) GetOrderList(list *[]model.Demo_order) error {
	return orderDao.db.Find(&list).Error
}

func (orderDao *OrderDao) GetSortedOrderList(likeName string, orderList *[]model.Demo_order) error {
	return orderDao.db.Raw("select * from demo_order where user_name like ? ORDER BY amount DESC",
		"%"+likeName+"%").Scan(&orderList).Error
}

func (orderDao *OrderDao) GetDownLoadList(likeName string, orderList *[]model.Demo_order) error {
	return orderDao.db.Raw("select * from demo_order where user_name like ? ORDER BY amount DESC",
		"%"+likeName+"%").Scan(&orderList).Error
}

func (OrderDao *OrderDao) GetSessionBegin() (error, *gorm.DB) {
	tx := OrderDao.db.Begin()
	return OrderDao.db.Begin().Error, tx
}

func (OrderDao *OrderDao) UpdateUrl(url string, id string, tx *gorm.DB) error {
	sql := "UPDATE demo_order SET file_url=?  WHERE id=?"
	return tx.Exec(sql, url, id).Error
}
