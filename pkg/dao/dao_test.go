package dao

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testGinandGorm/common/db"
	"testGinandGorm/pkg/model"
	"testing"
	"time"
)
var sample = &model.DemoOrder{OrderNo:time.Now().Format("2006-01-02 15:04:05"),UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}

func initial() (dao OrderDao) {
	dao = OrderDao{db: db.DbInit()}
	return dao
}

func TestCreateOrder(t *testing.T) {
	dao := initial()
	assert.NoError(t, dao.CreateOrder(sample))
	assert.Error(t, dao.CreateOrder(sample))
}

func TestDeleteById(t *testing.T) {
	dao := initial()
	assert.NoError(t,dao.DeleteById(strconv.Itoa(sample.ID)))
}

func TestUpdateByNo(t *testing.T) {
	dao:= initial()
	m:=map[string]interface{}{
			"Id" :sample.ID,
			"order_No":sample.OrderNo,
			"user_name" :sample.UserName,
			"amount" :sample.Amount,
			"status" :sample.Status,
			"file_url":sample.FileUrl,
	}
	assert.NoError(t,dao.UpdateByNo(sample.OrderNo, m))
	m = map[string]interface{}{
		"order_No":sample.OrderNo,
		"user_name" :sample.UserName,
		"amount" :sample.Amount,
		"file_url":sample.FileUrl,
	}
	assert.NoError(t,dao.UpdateByNo(sample.OrderNo, m))
}

func TestQueryOrderById(t *testing.T) {
	dao := initial()
	sample,err :=dao.QueryOrderById(strconv.Itoa(sample.ID))
	assert.Error(t,err)
	assert.Empty(t,sample)
}

func TestQueryOrdersByName(t *testing.T) {
	dao := initial()
	_, err := dao.QueryOrdersByName(sample.UserName, "amount", "desc")
	assert.NoError(t,err)
}

func TestQueryOrders(t *testing.T) {
	dao := initial()
	assert.NoError(t, dao.CreateOrder(sample))
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
	assert.NoError(t,dao.DeleteById(strconv.Itoa(sample.ID)))
}

func TestUpdateById(t *testing.T) {
	dao := initial()
	assert.NoError(t,dao.UpdateById(map[string]interface{}{
		"file_url": ".././test",
	}, "16"))
}


