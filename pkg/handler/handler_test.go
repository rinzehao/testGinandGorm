package handler

import (
	"github.com/gin-gonic/gin"
	"testGinandGorm/db"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/service"
	"testing"
)

func Init() *OrderHandler {
	db := db.DbInit()
	testDao := dao.NewOrderDao(db)
	testService := service.NewService(testDao)
	testHandler := NewHandler(testService)
	return testHandler
}

func TestDeleteOrderById(t *testing.T) {
	handler := Init()
	var c *gin.Context
	handler.CreateOrder(c)
}

//
//func TestCreateOrder(t *testing.T) {
//	testService := Init()
//	orderSample := model.DemoOrder{ID: 16, UserName: "test", OrderNo: "16", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
//	err := testService.CreateOrder(&orderSample)
//	assert.NoError(t,err)
//	orderSample = model.DemoOrder{ID: 20, UserName: "test", OrderNo: "20", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
//	err = testService.CreateOrder(&orderSample)
//	assert.NoError(t,err)
//}
//
//func TestDeleteOrder(t *testing.T) {
//	id := "16"
//	testService := Init()
//	err := testService.DeleteOrderById(id)
//	assert.NoError(t,err)
//}
//
//func TestGetOrder(t *testing.T) {
//	id := "16"
//	testService := Init()
//	//orderSample := model.DemoOrder{ID: id, UserName: "test", OrderNo: "test", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
//	orderSample,err := testService.QueryOrderById(id)
//	assert.NoError(t,err)
//	assert.NotEmpty(t,orderSample)
//}
//
//func TestGetOrderList(t *testing.T) {
//	testService := Init()
//	orders, err := testService.QueryOrders()
//	assert.NotEmpty(t,orders)
//	assert.NoError(t,err)
//}
//
//func TestUpdateOrder(t *testing.T) {
//	testService := Init()
//	var order = model.DemoOrder{ID: 1, UserName: "test", OrderNo: "test", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
//	err := testService.UpdateByOrderNo(&order)
//	assert.NoError(t,err)
//}
//
//func TestDownLoadExcel(t *testing.T) {
//	testService := Init()
//	file := xlsx.NewFile()
//	err := testService.DownLoadExcel(file)
//	assert.NoError(t,err)
//}
//
//func TestGetSortedOrders(t *testing.T) {
//	testService := Init()
//	username := "test"
//	orders, err := testService.QueryOrdersByName(username)
//	assert.NotEmpty(t,orders)
//	assert.NoError(t,err)
//}
//
//func TestGetUploadUrl(t *testing.T) {
//	testService := Init()
//	id := "1"
//	url := "www.test.com"
//	err := testService.UpdateUrlById(id, url)
//	assert.NoError(t,err)
//}
