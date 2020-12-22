package alert

import (
	"log"
	"strconv"
	"testGinandGorm/common/redis_utils"
	"testGinandGorm/pkg/model"
)

type OrderDao interface {
	CreateOrder(s *model.DemoOrder) error
	Delete(id string) error
	UpdateByNo(no string, m map[string]interface{}) error
	QueryOrderById(id string) (*model.DemoOrder, error)
	QueryOrderByNo(no string) (*model.DemoOrder, error)
	QueryOrders(page, pageSize int) (orders []*model.DemoOrder, err error)
	QueryOrdersByName(userName, orderBy, desc string) (orders []*model.DemoOrder, err error)
	UpdateById(id string,m map[string]interface{}) error
}

type MyOrderDao struct {
	db    OrderDB
	cache *redis_utils.Cache
}

func NewMyOrderDao(db OrderDB, cache *redis_utils.Cache) *MyOrderDao {
	return &MyOrderDao{db: db, cache: cache}
}

func NewCache() *redis_utils.Cache {
	return &redis_utils.Cache{}
}

func (dao *MyOrderDao) CreateOrder(s *model.DemoOrder) error {
	// step1 cache
	// ID -> Hash Hash:orderNo->order
	if err := dao.cache.HSet("orderHash",strconv.Itoa(s.ID),s.OrderNo);err != nil{
		log.Println("cache setting failed")
		return err
	}
	if err := dao.cache.SetStruct(s.OrderNo,s); err !=nil {
		log.Println("cache setting failed")
		return err
	}
	// step2 db
	if err := dao.db.CreateOrder(s);err != nil{
		log.Println("db insert failed")
		return err
	}
	return nil
}

func (dao *MyOrderDao) Delete(id string) error {
	//step1 cache
	flag, err := dao.cache.Hdel("orderHash",id)
	if err != nil{
		log.Println("delete cache failed")
		return err
	}
	if flag == false {
		log.Println("delete cache failed, cache might not exist")
	}
	//step2 db
	if err =dao.db.DeleteById(id); err != nil{
		log.Println("delete order from db failed ")
		return err
	}
	return nil
}

func (dao *MyOrderDao) UpdateByNo(no string, m map[string]interface{}) error {
	//step1 cache淘汰
	flag, err := dao.cache.Hdel("orderHash",no)
	if err != nil {
		log.Println("cache delete failed")
		return err
	}
	if flag == true {
		log.Println("cache delete successfully")
	}
	//step2 db写入
	if err := dao.UpdateByNo(no, m) ;err != nil{
		log.Println("updateByNo at DB failed")
		return err
	}
	//step3 cache写入
	order, err :=dao.db.QueryOrderByNo(no)
	if err !=nil {
		log.Println("QueryOrder Failed")
		return err
	}
	if err := dao.cache.SetStruct(no,order) ; err != nil{
		log.Println("Update cache Failed")
		return err
	}
	return nil
}

func (dao *MyOrderDao) QueryOrderById(id string) (order *model.DemoOrder, err error) {
	//step1 cache
	var flag bool
	if flag, err = dao.cache.HExists("orderHash",id); err != nil {
		return nil, err
	}
	if flag == false {
		log.Println("queryTarget form cache fail")
	}
	if flag == true {
		if err = dao.cache.GetStruct(id, order); err != nil {
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
	if flag, err = dao.cache.Exist(no); err != nil {
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
	return dao.db.QueryOrders(page,pageSize)
}

func (dao *MyOrderDao) QueryOrdersByName(userName, orderBy, desc string) (orders []*model.DemoOrder, err error) {
	return dao.db.QueryOrdersByName(userName,orderBy,desc)
}

func (dao *MyOrderDao) UpdateById(id string, m map[string]interface{}) error {
	//step1 db写入
	if err := dao.db.UpdateById(id, m) ;err != nil{
		log.Println("updateByNo at DB failed")
		return err
	}
	//step2 cache淘汰
	order, err := dao.db.QueryOrderById(id)
	if err !=nil {
		log.Println("QueryOrder Failed")
		return err
	}
	flag, err := dao.cache.Delete(order)
	if err != nil {
		log.Println("cache delete failed")
		return err
	}
	if flag == true {
		log.Println("cache delete successfully")
	}
	flag, err = dao.cache.Hdel("orderHash",id)
	if err != nil {
		log.Println("cache delete failed")
		return err
	}
	if flag == true {
		log.Println("cache delete successfully")
	}
	//step3 cache写入
	if err := dao.cache.SetStruct(order.OrderNo,order) ; err != nil{
		log.Println("Update cache Failed")
		return err
	}
	if err := dao.cache.SetStruct(id,order.OrderNo) ; err != nil{
		log.Println("Update cache Failed")
		return err
	}
	return nil
}
