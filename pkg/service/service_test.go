package service

//
//import (
//	"github.com/stretchr/testify/assert"
//	"math/rand"
//	"strconv"
//	"testGinandGorm/common/db"
//	"testGinandGorm/pkg/dao"
//	"testGinandGorm/pkg/model"
//	"testing"
//	"time"
//)
//
//func initial() (testService *OrderService, sample *model.DemoOrder) {
//	db := db.DbInit()
//	db.LogMode(true)
//	dao := dao.NewOrderDao(db)
//	sample = &model.DemoOrder{OrderNo:time.Now().Format("2006-01-02 15:04:05")+queryRandomString(10),UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
//	return NewService(dao) ,sample
//}
//
//func TestCreateOrderByOrderNo(t *testing.T) {
//	testService, sample:=initial()
//	assert.NoError(t,testService.CreateOrder(sample))
//}
//
//func TestDeleteOrderById(t *testing.T) {
//	testService, sample:=initial()
//	assert.NoError(t,testService.CreateOrder(sample))
//	assert.NoError(t,testService.DeleteOrderById(strconv.Itoa(sample.ID)))
//}
//
//func TestQueryOrderById(t *testing.T) {
//	testService, sample:=initial()
//	assert.NoError(t,testService.CreateOrder(sample))
//	order,err := testService.QueryOrderById(strconv.Itoa(sample.ID))
//	assert.NoError(t,err)
//	assert.NotEmpty(t,order)
//}
//
//func TestUpdateByOrderNo(t *testing.T) {
//	testService, sample:=initial()
//	assert.NoError(t,testService.CreateOrder(sample))
//	m:=map[string]interface{}{
//		"Id" :sample.ID,
//		"order_No":sample.OrderNo,
//		"user_name" :sample.UserName,
//		"amount" :sample.Amount,
//		"status" :"testAlter",
//		"file_url":sample.FileUrl,
//	}
//	assert.NoError(t,testService.UpdateByOrderNo(m, sample.OrderNo))
//	m = map[string]interface{}{
//		"order_No":sample.OrderNo,
//		"user_name" :"testAlter",
//		"amount" :sample.Amount,
//		"file_url":sample.FileUrl,
//	}
//	assert.NoError(t,testService.UpdateByOrderNo(m, sample.OrderNo))
//}
//
//func TestQueryOrders(t *testing.T) {
//	testService, sample:=initial()
//	assert.NoError(t,testService.CreateOrder(sample))
//	orders,err := testService.QueryOrders(1,100)
//	assert.NoError(t,err)
//	assert.NotEmpty(t,orders)
//}
//
//func TestQueryOrdersByName(t *testing.T) {
//	testService, sample:=initial()
//	assert.NoError(t,testService.CreateOrder(sample))
//	order, err := testService.QueryOrdersByName(sample.UserName,"amount","desc")
//	assert.NoError(t,err)
//	assert.NotEmpty(t,order)
//}
//
//func TestUpdateUrlById(t *testing.T) {
//	testService, sample:=initial()
//	assert.NoError(t,testService.CreateOrder(sample))
//	m := map[string]interface{}{
//		"file_url": ".././test",
//	}
//	err := testService.UpdateById(m,strconv.Itoa(sample.ID))
//	assert.NoError(t,err)
//}
//
//func  queryRandomString(l int) string {
//	str := "0123456789abcdefghijklmnopqrstuvwxyz"
//	bytes := []byte(str)
//	result := []byte{}
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	for i := 0; i < l; i++ {
//		result = append(result, bytes[r.Intn(len(bytes))])
//	}
//	return string(result)
//}
