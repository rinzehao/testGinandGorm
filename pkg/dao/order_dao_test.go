package dao

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testGinandGorm/db"
	"testGinandGorm/pkg/model"
	"testing"
)

func initial() (dao OrderDao, sample model.DemoOrder) {
	dao = OrderDao{db: db.DbInit()}
	sample = model.DemoOrder{ID: 16, OrderNo:"16",UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	return dao, sample
}

func TestCreateOrder(t *testing.T) {
	dao, sample := initial()
	err := dao.CreateOrder(&sample)
	assert.Error(t,err)
	sample = model.DemoOrder{ID: 3, OrderNo:"3",UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	err = dao.CreateOrder(&sample)
	assert.Error(t,err)
}

func TestDeleteById(t *testing.T) {
	dao, sample := initial()
	dao.db.LogMode(true)
	err:=dao.DeleteById(strconv.Itoa(sample.ID))
	assert.NoError(t,err)
}

func TestUpdateByNoStruct(t *testing.T) {
	dao, sample := initial()
	dao.db.LogMode(true)
	err := dao.UpdateByNoStruct(&sample)
	assert.NoError(t,err)
	sample.ID =0
	sample.Amount=0.00
	err = dao.UpdateByNoStruct(&sample)
	assert.NoError(t,err)
	dao.db.Model(&model.DemoOrder{}).Where("order_no = ?",sample.OrderNo).Update(sample)
	dao.db.Where("order_no = ?",sample.OrderNo).Update(sample)
	sample.ID =16
	dao.db.Update(sample)
}

func TestUpdateByNo(t *testing.T) {
	dao, sample := initial()
	dao.db.LogMode(true)
	orderNo := sample.OrderNo
	m:=map[string]interface{}{
			"Id" :sample.ID,
			"order_No":sample.OrderNo,
			"user_name" :sample.UserName,
			"amount" :sample.Amount,
			"status" :sample.Status,
			"file_url":sample.FileUrl,
	}
	err := dao.UpdateByNo(orderNo, m)
	assert.NoError(t,err)
	m = map[string]interface{}{
		"order_No":sample.OrderNo,
		"user_name" :sample.UserName,
		"amount" :sample.Amount,
		"file_url":sample.FileUrl,
	}
	err = dao.UpdateByNo(orderNo, m)
	assert.NoError(t,err)
}

func TestQueryOrderById(t *testing.T) {
	dao, sample := initial()
	orderId := sample.ID
	samples,err :=dao.QueryOrderById(strconv.Itoa(orderId))
	assert.Error(t,err)
	orderId =17
	samples,err =dao.QueryOrderById(strconv.Itoa(orderId))
	t.Log(samples)
	assert.NoError(t,err)
}

func TestQueryOrdersByName(t *testing.T) {
	dao, sample := initial()
	dao.db.LogMode(true)
	userName, orderBy, desc := sample.UserName, "amount", "desc"
	orders, err := dao.QueryOrdersByName(userName, orderBy, desc)
	assert.NotEmpty(t,orders)
	assert.NoError(t,err)
}

func TestQueryOrders(t *testing.T) {
	dao, sample := initial()
	log.Println(sample)
	dao.db.LogMode(true)
	page, pageSize := 0, 10
	orders, err := dao.QueryOrders(page, pageSize)
	assert.NoError(t,err)
	assert.NotEmpty(t,orders)
	page, pageSize = 0, 100
	orders, err = dao.QueryOrders(page, pageSize)
	assert.NoError(t,err)
	assert.NotEmpty(t,orders)
	page, pageSize = -1, 100
	orders, err = dao.QueryOrders(page,pageSize)
	assert.NoError(t,err)
	assert.NotEmpty(t,orders)
	page, pageSize = 1, 10
	orders, err = dao.QueryOrders(page,pageSize)
	assert.NoError(t,err)
	assert.NotEmpty(t,orders)
}

func TestUpdateById(t *testing.T) {
	dao, sample := initial()
	url, id := ".././test", "16"
	m := map[string]interface{}{
		"file_url": url,
	}
	err := dao.UpdateById(m, id)
	log.Println(sample)
	assert.NoError(t,err)
}

