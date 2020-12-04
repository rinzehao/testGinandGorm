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

func TestDeleteOrderById(t *testing.T) {
	testService :=Init()
	id :="16"
	err := testService.DeleteOrderById(id)
	assert.NoError(t,err)
}

func TestQueryOrderById(t *testing.T) {
	testService :=Init()
	id :="15"
	order,err := testService.QueryOrderById(id)
	assert.NotEmpty(t,order)
	assert.NoError(t,err)
}

func TestUpdateByOrderNo(t *testing.T) {
	testService :=Init()
	sample := model.DemoOrder{OrderNo:"16", UserName: "raious", Amount: 444, Status: "over"}
	err := testService.UpdateByOrderNo(&sample)
	assert.NoError(t,err)
}


func TestCreateOrderByOrderNo(t *testing.T) {
	testService :=Init()
	sample := model.DemoOrder{ID: 16, OrderNo:"16",UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	err := testService.CreateOrder(&sample)
	assert.NoError(t,err)
}

func TestQueryOrders(t *testing.T) {
	testService :=Init()
	orders,err := testService.QueryOrders()
	assert.NoError(t,err)
	assert.NotEmpty(t,orders)
}

func TestSortedOrdersByUserName(t *testing.T) {
	testService :=Init()
	sample := model.DemoOrder{ID: 16, UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	userName:= sample.UserName
	orders, err := testService.QueryOrdersByName(userName)
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
	err := testService.UpdateUrlById(id,url)
	assert.NoError(t,err)
}