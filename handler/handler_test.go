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
	id := 1
	testService := Init()
	orderSample := model.Demo_order{ID: id, UserName: "test", OrderNo: "test", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
	if err := testService.CreateOrder(&orderSample); err != nil {
		fmt.Print(err)
	}
}

func TestDeleteOrder(t *testing.T) {
	id := "1"
	testService := Init()
	if err := testService.DeleteOrder(id); err != nil {
		fmt.Print(err)
	}
}

func TestGetOrder(t *testing.T) {
	id := 1
	testService := Init()
	orderSample := model.Demo_order{ID: id, UserName: "test", OrderNo: "test", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
	if err := testService.GetOrder(strconv.Itoa(id), &orderSample); err != nil {
		fmt.Print(err)
	}
}

func TestGetOrderList(t *testing.T) {
	testService := Init()
	var err error
	var list []model.Demo_order
	if err, list = testService.GetOrderList(list); err != nil {
		fmt.Print(err)
	}
	if list == nil {
		panic("表中无条目")
	}
}

func TestUpdateOrder(t *testing.T) {
	testService := Init()
	id := "1"
	var order model.Demo_order
	var err error
	if err = testService.UpdateOrder(id, &order); err != nil {
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
	var orderSampleArray []model.Demo_order
	orderSample := model.Demo_order{ID: 1, UserName: "枣糕", OrderNo: "test", Amount: 14.6, Status: "Test", FileUrl: "www.test.com"}
	if err, orderSampleArray = testService.GetSortedOrderList(&orderSample, orderSampleArray); err != nil {
		fmt.Println(err)
	}
	if orderSampleArray == nil {
		panic("数组为空")
	}
}

func TestGetUploadUrl(t *testing.T) {
	testService := Init()
	id := "1"
	str := "www.test.com"
	if err := testService.GetUploadUrl(id, str); err != nil {
		fmt.Println(err)
	}
}
