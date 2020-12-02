package handler

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
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
	orderSample := model.DemoOrder{ID: 16, UserName: "test", OrderNo: "test", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
	if err := testService.CreateOrderByOrderNo(&orderSample); err != nil {
		fmt.Print(err)
	}
}

func TestDeleteOrder(t *testing.T) {
	id := "16"
	testService := Init()
	if err := testService.DeleteOrderById(id); err != nil {
		fmt.Print(err)
	}
}

func TestGetOrder(t *testing.T) {
	id := 16
	testService := Init()
	//orderSample := model.DemoOrder{ID: id, UserName: "test", OrderNo: "test", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
	if err, orderSample := testService.QueryOrderById(strconv.Itoa(id)); err != nil {
		fmt.Print(err)
	} else if orderSample == nil {
		panic("quzhiweikong")
	}
}

func TestGetOrderList(t *testing.T) {
	testService := Init()
	var err error
	var orders []*model.DemoOrder
	if orders, err = testService.QueryOrders(); err != nil {
		fmt.Print(err)
	}
	if orders == nil {
		panic("无条目")
	}
}

func TestUpdateOrder(t *testing.T) {
	testService := Init()
	var order = model.DemoOrder{ID: 1, UserName: "test", OrderNo: "test", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
	var err error
	if err = testService.UpdateByOrderNo(&order); err != nil {
		fmt.Println(err)
	}
}

func TestDownLoadExcel(t *testing.T) {
	testService := Init()
	file := xlsx.NewFile()
	if err := testService.DownLoadExcel(file); err != nil {
		panic(err)
	}
}

func TestGetSortedOrderList(t *testing.T) {
	testService := Init()
	var err error
	var orderSampleArray []*model.DemoOrder
	username := "test"
	if orderSampleArray, err = testService.QuerySortedOrdersByUserName(username); err != nil {
		fmt.Println(err)
	}
	if orderSampleArray == nil {
		panic("数组为空")
	}
}

func TestGetUploadUrl(t *testing.T) {
	testService := Init()
	id := "1"
	url := "www.test.com"
	if err := testService.GetUploadUrlAndSave(id, url); err != nil {
		fmt.Println(err)
	}
}
