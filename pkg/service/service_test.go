package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/tealeg/xlsx"
	"testGinandGorm/db"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
	"testing"
)

func Init() *OrderService {
	db := db.DbInit()
	testDao := dao.NewOrderDao(db)
	testService := NewService(testDao)
	return testService
}

func TestOrderService_DeleteOrderById(t *testing.T) {
	testService :=Init()
	id :="16"
	err := testService.DeleteOrderById(id)
	assert.NoError(t,err)
}

func TestOrderService_QueryOrderById(t *testing.T) {
	testService :=Init()
	id :="15"
	order,err := testService.QueryOrderById(id)
	assert.NotEmpty(t,order)
	assert.NoError(t,err)
}

func TestOrderService_UpdateByOrderNo(t *testing.T) {
	testService :=Init()
	sample := model.DemoOrder{ID: 16, UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	err := testService.UpdateByOrderNo(&sample)
	assert.NoError(t,err)
}


func TestOrderService_CreateOrderByOrderNo(t *testing.T) {
	testService :=Init()
	sample := model.DemoOrder{ID: 16, OrderNo:"16",UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	err := testService.CreateOrderByOrderNo(&sample)
	assert.NoError(t,err)
}

func TestOrderService_QueryOrders(t *testing.T) {
	testService :=Init()
	orders,err := testService.QueryOrders()
	assert.NoError(t,err)
	assert.NotEmpty(t,orders)
}

func TestOrderService_QuerySortedOrdersByUserName(t *testing.T) {
	testService :=Init()
	sample := model.DemoOrder{ID: 16, UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	userName:= sample.UserName
	orders, err := testService.QuerySortedOrdersByUserName(userName)
	assert.NotEmpty(t,orders)
	assert.NoError(t,err)
}

func TestOrderService_DownLoadExcel(t *testing.T) {
	testService :=Init()
	var file =xlsx.NewFile()
	err := testService.DownLoadExcel(file)
	assert.NoError(t,err)
}

func TestOrderService_GetUploadUrlAndSave(t *testing.T) {
	testService :=Init()
	id,url :="16",".././test"
	err := testService.GetUploadUrlAndSave(id,url)
	assert.NoError(t,err)
}