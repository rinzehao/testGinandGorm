package service

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testGinandGorm/common/db"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
	"testing"
	"time"
)

var sample = &model.DemoOrder{OrderNo:time.Now().Format("2006-01-02 15:04:05"),UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}

func Init() *OrderService {
	db := db.DbInit()
	db.LogMode(true)
	testDao := dao.NewOrderDao(db)
	return NewService(testDao)
}

func TestCreateOrderByOrderNo(t *testing.T) {
	testService :=Init()
	assert.NoError(t,testService.CreateOrder(sample))
}

func TestDeleteOrderById(t *testing.T) {
	testService :=Init()
	assert.NoError(t,testService.DeleteOrderById(strconv.Itoa(sample.ID)))
}

func TestQueryOrderById(t *testing.T) {
	testService :=Init()
	order,err := testService.QueryOrderById(strconv.Itoa(sample.ID))
	assert.Error(t,err)
	assert.Empty(t,order)
}

func TestUpdateByOrderNo(t *testing.T) {
	testService :=Init()
	m:=map[string]interface{}{
		"Id" :sample.ID,
		"order_No":sample.OrderNo,
		"user_name" :sample.UserName,
		"amount" :sample.Amount,
		"status" :sample.Status,
		"file_url":sample.FileUrl,
	}
	assert.NoError(t,testService.UpdateByOrderNo(m, sample.OrderNo))
	m = map[string]interface{}{
		"order_No":sample.OrderNo,
		"user_name" :sample.UserName,
		"amount" :sample.Amount,
		"file_url":sample.FileUrl,
	}
	assert.NoError(t,testService.UpdateByOrderNo(m, sample.OrderNo))
}

func TestQueryOrders(t *testing.T) {
	testService :=Init()
	_,err := testService.QueryOrders(1,100)
	assert.NoError(t,err)
}

func TestQueryOrdersByName(t *testing.T) {
	testService :=Init()
	_, err := testService.QueryOrdersByName("raious","amount","desc")
	assert.NoError(t,err)
}

func TestUpdateUrlById(t *testing.T) {
	testService :=Init()
	m := map[string]interface{}{
		"file_url": ".././test",
	}
	err := testService.UpdateById(m,"16")
	assert.NoError(t,err)
}