package handler

import (
	"github.com/stretchr/testify/assert"
	"github.com/tealeg/xlsx"
	"testGinandGorm/db"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
	"testGinandGorm/pkg/service"
	"testing"
)

func Init() *service.OrderService {
	db := db.DbInit()
	testDao := dao.NewOrderDao(db)
	testService := service.NewService(testDao)
	return testService
}

func TestCreateOrder(t *testing.T) {
	testService := Init()
	orderSample := model.DemoOrder{ID: 16, UserName: "test", OrderNo: "16", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
	err := testService.CreateOrderByOrderNo(&orderSample)
	assert.NoError(t,err)
	orderSample = model.DemoOrder{ID: 20, UserName: "test", OrderNo: "20", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
	err = testService.CreateOrderByOrderNo(&orderSample)
	assert.NoError(t,err)
}

func TestDeleteOrder(t *testing.T) {
	id := "16"
	testService := Init()
	err := testService.DeleteOrderById(id)
	assert.NoError(t,err)
}

func TestGetOrder(t *testing.T) {
	id := "16"
	testService := Init()
	//orderSample := model.DemoOrder{ID: id, UserName: "test", OrderNo: "test", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
	orderSample,err := testService.QueryOrderById(id)
	assert.NoError(t,err)
	assert.NotEmpty(t,orderSample)
}

func TestGetOrderList(t *testing.T) {
	testService := Init()
	orders, err := testService.QueryOrders()
	assert.NotEmpty(t,orders)
	assert.NoError(t,err)
}

func TestUpdateOrder(t *testing.T) {
	testService := Init()
	var order = model.DemoOrder{ID: 1, UserName: "test", OrderNo: "test", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
	err := testService.UpdateByOrderNo(&order)
	assert.NoError(t,err)
}

func TestDownLoadExcel(t *testing.T) {
	testService := Init()
	file := xlsx.NewFile()
	err := testService.DownLoadExcel(file)
	assert.NoError(t,err)
}

func TestGetSortedOrderList(t *testing.T) {
	testService := Init()
	username := "test"
	orders, err := testService.QuerySortedOrdersByUserName(username)
	assert.NotEmpty(t,orders)
	assert.NoError(t,err)
}

func TestGetUploadUrl(t *testing.T) {
	testService := Init()
	id := "1"
	url := "www.test.com"
	err := testService.GetUploadUrlAndSave(id, url)
	assert.NoError(t,err)
}
