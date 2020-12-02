package service

import (
	"fmt"
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
	if err := testService.DeleteOrderById(id); err != nil {
		fmt.Println(err)
	}
}

func TestOrderService_QueryOrderById(t *testing.T) {
	testService :=Init()
	id :="16"
	if order,err := testService.QueryOrderById(id); err != nil {
		fmt.Println(order)
		fmt.Println(err)
	}
}

func TestOrderService_UpdateByOrderNo(t *testing.T) {
	testService :=Init()
	sample := model.DemoOrder{ID: 16, UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	if err := testService.UpdateByOrderNo(&sample); err != nil {
		fmt.Println(err)
	}
}


func TestOrderService_CreateOrderByOrderNo(t *testing.T) {
	testService :=Init()
	sample := model.DemoOrder{ID: 16, UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	if err := testService.CreateOrderByOrderNo(&sample); err != nil {
		fmt.Println(err)
	}
}

func TestOrderService_QueryOrders(t *testing.T) {
	testService :=Init()
	if orders,err := testService.QueryOrders(); err != nil {
		fmt.Println(err)
		fmt.Println(orders)
	}
}

func TestOrderService_QuerySortedOrdersByUserName(t *testing.T) {
	testService :=Init()
	sample := model.DemoOrder{ID: 16, UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	userName:= sample.UserName
	if orders, err := testService.QuerySortedOrdersByUserName(userName); err != nil {
		fmt.Println(sample)
		fmt.Print(orders)
	}
}

func TestOrderService_DownLoadExcel(t *testing.T) {
	testService :=Init()
	var file =xlsx.NewFile()
	if err := testService.DownLoadExcel(file); err != nil {
		fmt.Println(err)
	}
}

func TestOrderService_GetUploadUrlAndSave(t *testing.T) {
	testService :=Init()
	id,url :="16",".././test"
	if err := testService.GetUploadUrlAndSave(id,url); err != nil {
		fmt.Println(err)
	}
}