package dao

import (
	"fmt"
	"strconv"
	"testGinandGorm/db"
	"testGinandGorm/pkg/model"
	"testing"
)

func initial() (dao OrderDao, sample model.DemoOrder) {
	dao = OrderDao{db: db.DbInit()}
	sample = model.DemoOrder{ID: 16, UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	return dao, sample
}

func TestOrderDao_Create(t *testing.T) {
	dao, sample := initial()
	if err := dao.CreateOrder(&sample); err != nil {
		fmt.Println(err)
	}
}

func TestOrderDao_DeleteById(t *testing.T) {
	dao, sample := initial()
	id := sample.ID
	if err := dao.DeleteById(strconv.Itoa(id)); err != nil {
		fmt.Println(err)
	}
}

func TestOrderDao_UpdateByOrderNo(t *testing.T) {
	dao, sample := initial()
	orderNo := sample.OrderNo
	if err := dao.UpdateByOrderNo(orderNo, &sample); err != nil {
		fmt.Print(err)
	}
}

func TestQueryOrderById(t *testing.T) {
	dao, sample := initial()
	orderId := sample.ID
	var err error
	var samples *model.DemoOrder
	if samples, err = dao.QueryOrderById(strconv.Itoa(orderId)); err != nil {
		fmt.Println(samples)
		fmt.Print(err)
	}
}

func TestQueryOrderIsExistByOrderNo(t *testing.T) {
	dao, sample := initial()
	orderNo := sample.OrderNo
	if isExit, err := dao.QueryOrderIsExistByOrderNo(orderNo); err != nil {
		fmt.Println(isExit)
		fmt.Print(err)
	}
}

func TestQueryOrders(t *testing.T) {
	dao, sample := initial()
	page, pageSize := 0, 100
	if isExit, err := dao.QueryOrders(page, pageSize); err != nil {
		fmt.Println(sample)
		fmt.Print(isExit)
	}
}

func TestQuerySortedOrdersByUserName(t *testing.T) {
	dao, sample := initial()
	userName, orderBy, desc := sample.UserName, "amount", "desc"
	if orders, err := dao.QuerySortedOrdersByUserName(userName, orderBy, desc); err != nil {
		fmt.Println(sample)
		fmt.Print(orders)
	}
}

func TestQueryTranscationUpdateById(t *testing.T) {
	dao, sample := initial()
	url, id := ".././test", "16"
	if err := dao.QueryTranscationUpdateById(url, id); err != nil {
		fmt.Print(err)
		fmt.Print(sample)
	}
}
