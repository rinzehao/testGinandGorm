package dao

import (
	"github.com/jinzhu/gorm"
	_ "testGinandGorm/common/db"
	"testGinandGorm/pkg/model"
)

type OrderDao struct {
	db *gorm.DB
}

func NewOrderDao(Db *gorm.DB) *OrderDao {
	return &OrderDao{db: Db}
}

func (dao *OrderDao) CreateOrder(s *model.DemoOrder) error {
	return dao.db.Create(s).Error
}

func (dao *OrderDao) DeleteById(id string) error {
	return dao.db.Where("id = ?", id).Delete(&model.DemoOrder{}).Error
}

//传入结构体跟map的updata方法的区别
//传入结构体->索引用的主键的话，且结构提内的主键不为zero，则可以省略where 直接使用 db.Update（）来进行更新
//传入结构体->索引用的主键的话，但结构提内的主键为zero，则不可以省略where 声明使用的字段 db.WHere().Update（）来进行更新
//传入结构体->更新字段为结构体内不为zero的字段

//传入map->不可省略Model(&model.DemoOrder{})  来声明修改的表名
//传入map ->不可省略Where().否则会对所有条目进行更新
//传入map ->不可省略Where().否则会对所有条目进行更新
//传入map ->更新的字段为map内存在的字段

func (dao *OrderDao ) UpdateByNo(no string,m map[string]interface{}) error  {
	return dao.db.Model(&model.DemoOrder{}).Where("order_no = ?", no).Update(m).Error
}

func (dao *OrderDao) QueryOrderById(id string) (*model.DemoOrder, error) {
	var order model.DemoOrder
	return &order, dao.db.Where("id = ?", id).First(&order).Error
}


func (dao *OrderDao) QueryOrderByNo(no string) (*model.DemoOrder, error) {
	var order model.DemoOrder
	return &order, dao.db.Where("order_no = ?", no).First(&order).Error
}


func (dao *OrderDao) QueryOrders(page, pageSize int) (orders []*model.DemoOrder, err error) {
	return orders, dao.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&orders).Error
}


func (dao *OrderDao) QueryOrdersByName(userName, orderBy, desc string) (orders []*model.DemoOrder, err error) {
	return orders, dao.db.Where("user_name LIKE ?", "%"+userName+"%").Order(orderBy + " " + desc).Find(&orders).Error
}


//Update Updates Save的区别
func (dao *OrderDao) UpdateById(m map[string]interface{}, id string) error {
	tx := dao.db.Begin()
	defer tx.Rollback()
	if err := dao.db.Model(model.DemoOrder{}).Where("id = ?", id).Updates(m).Error; err != nil {
		return err
	}
	tx.Commit()
	return nil
}
