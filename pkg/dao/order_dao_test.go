package dao

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testGinandGorm/db"
	"testGinandGorm/pkg/model"
	"testing"
)

func initial() (dao OrderDao, sample model.DemoOrder) {
	dao = OrderDao{db: db.DbInit()}
	sample = model.DemoOrder{ID: 16, OrderNo:"16",UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	return dao, sample
}

func TestOrderDao_Create(t *testing.T) {
	dao, sample := initial()
	err := dao.CreateOrder(&sample)
	assert.Error(t,err)
	sample = model.DemoOrder{ID: 3, OrderNo:"3",UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	err = dao.CreateOrder(&sample)
	assert.Error(t,err)
}

func TestOrderDao_DeleteById(t *testing.T) {
	dao, sample := initial()
	err:=dao.DeleteById(strconv.Itoa(sample.ID))
	assert.NoError(t,err)
}

func TestOrderDao_UpdateByOrderNo(t *testing.T) {
	dao, sample := initial()
	orderNo := sample.OrderNo
	err := dao.UpdateByOrderNo(orderNo, &sample)
	assert.NoError(t,err)
}

func TestQueryOrderById(t *testing.T) {
	dao, sample := initial()
	orderId := sample.ID
	samples,err :=dao.QueryOrderById(strconv.Itoa(orderId))
	assert.Error(t,err)
	orderId =17
	samples,err =dao.QueryOrderById(strconv.Itoa(orderId))
	t.Log(samples)
	assert.NoError(t,err)
}

func TestQueryOrderIsExist(t *testing.T) {
	dao, sample := initial()
	m:=map[string]string{
		"Id" :strconv.Itoa(sample.ID),
		"order_No":sample.OrderNo,
		"user_name" :sample.UserName,
		"amount" :strconv.FormatFloat(sample.Amount, 'E', -1, 64),
		"status" :sample.Status,
		"file_url":sample.FileUrl,
	}
	paramName :="user_name"
	isExit, err := dao.QueryOrderIsExist(m,paramName,&sample)
	assert.True(t,isExit)
	assert.NoError(t,err)
}

func TestQueryOrders(t *testing.T) {
	dao, sample := initial()
	page, pageSize := 0, 100
	log.Print(sample)
	orders, err := dao.QueryOrders(page, pageSize)
	assert.NoError(t,err)
	assert.NotEmpty(t,orders)
}

func TestQuerySortedOrdersByUserName(t *testing.T) {
	dao, sample := initial()
	userName, orderBy, desc := sample.UserName, "amount", "desc"
	orders, err := dao.QuerySortedOrdersByUserName(userName, orderBy, desc)
	assert.NotEmpty(t,orders)
	assert.NoError(t,err)
}

func TestQueryTranscationUpdateById(t *testing.T) {
	dao, sample := initial()
	url, id := ".././test", "16"
	m := map[string]string{
		"file_url": url,
	}
	err := dao.TransactionUpdateById(m, id)
	log.Println(sample)
	assert.NoError(t,err)
}