package db

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testGinandGorm/common/mySQL_db"
	"testGinandGorm/pkg/model"
	"testing"
	"time"
)

func initial() (orderDB OrderDB, sample *model.DemoOrder) {
	db :=mySQL_db.DbInit()
	orderDB = NewMyOrderDB(db)
	db = db.LogMode(true)
	sample = &model.DemoOrder{OrderNo: time.Now().Format("2006-01-02 15:04:05")+queryRandomString(10), UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/mySQL_db"}
	return orderDB, sample
}

func TestCreateOrder(t *testing.T) {
	db, sample := initial()
	assert.NoError(t, db.CreateOrder(sample))
	assert.Error(t, db.CreateOrder(sample))
}

func TestDeleteById(t *testing.T) {
	db, sample := initial()
	assert.NoError(t, db.CreateOrder(sample))
	assert.NoError(t, db.DeleteById(strconv.Itoa(sample.ID)))
}

func TestUpdateByNo(t *testing.T) {
	db, sample := initial()
	assert.NoError(t, db.CreateOrder(sample))
	m := map[string]interface{}{
		"Id":        sample.ID,
		"order_No":  sample.OrderNo,
		"user_name": sample.UserName,
		"amount":    sample.Amount,
		"status":    "testAlert",
		"file_url":  sample.FileUrl,
	}
	assert.NoError(t, db.UpdateByNo(sample.OrderNo, m))
	m = map[string]interface{}{
		"order_No":  sample.OrderNo,
		"user_name": "testAlert",
		"amount":    sample.Amount,
		"file_url":  sample.FileUrl,
	}
	assert.NoError(t, db.UpdateByNo(sample.OrderNo, m))
	assert.NoError(t, db.DeleteById(strconv.Itoa(sample.ID)))
}

func TestQueryOrderById(t *testing.T) {
	db, sample := initial()
	assert.NoError(t, db.CreateOrder(sample))
	sample, err := db.QueryOrderById(strconv.Itoa(sample.ID))
	assert.NoError(t, err)
	assert.NotEmpty(t, sample)
	assert.NoError(t, db.DeleteById(strconv.Itoa(sample.ID)))
}

func TestQueryOrdersByName(t *testing.T) {
	db, sample := initial()
	assert.NoError(t, db.CreateOrder(sample))
	order, err := db.QueryOrdersByName(sample.UserName, "amount", "desc")
	assert.NoError(t, err)
	assert.NotEmpty(t, order)
	assert.NoError(t, db.DeleteById(strconv.Itoa(sample.ID)))
}

func TestQueryOrders(t *testing.T) {
	db, sample := initial()
	assert.NoError(t, db.CreateOrder(sample))
	page, pageSize := 0, 10
	orders, err := db.QueryOrders(page, pageSize)
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
	page, pageSize = 0, 100
	orders, err = db.QueryOrders(page, pageSize)
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
	page, pageSize = -1, 100
	orders, err = db.QueryOrders(page, pageSize)
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
	page, pageSize = 1, 10
	orders, err = db.QueryOrders(page, pageSize)
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
	assert.NoError(t, db.DeleteById(strconv.Itoa(sample.ID)))
}

func TestUpdateById(t *testing.T) {
	db, sample := initial()
	assert.NoError(t, db.CreateOrder(sample))
	assert.NoError(t, db.UpdateById(strconv.Itoa(sample.ID), map[string]interface{}{
		"file_url": ".././test",
	}))
	assert.NoError(t, db.DeleteById(strconv.Itoa(sample.ID)))
}


func  queryRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

