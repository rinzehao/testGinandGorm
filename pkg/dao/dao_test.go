package dao

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testGinandGorm/common/db"
	"testGinandGorm/pkg/model"
	"testing"
	"time"
)

func initial() (dao OrderDao, sample *model.DemoOrder) {
	dao = OrderDao{db: db.DbInit()}
	dao.db = dao.db.LogMode(true)
	sample = &model.DemoOrder{OrderNo: time.Now().Format("2006-01-02 15:04:05")+queryRandomString(10), UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	return dao, sample
}

func TestCreateOrder(t *testing.T) {
	dao, sample := initial()
	assert.NoError(t, dao.CreateOrder(sample))
	assert.Error(t, dao.CreateOrder(sample))
}

func TestDeleteById(t *testing.T) {
	dao, sample := initial()
	assert.NoError(t, dao.CreateOrder(sample))
	assert.NoError(t, dao.DeleteById(strconv.Itoa(sample.ID)))
}

func TestUpdateByNo(t *testing.T) {
	dao, sample := initial()
	assert.NoError(t, dao.CreateOrder(sample))
	m := map[string]interface{}{
		"Id":        sample.ID,
		"order_No":  sample.OrderNo,
		"user_name": sample.UserName,
		"amount":    sample.Amount,
		"status":    "testAlert",
		"file_url":  sample.FileUrl,
	}
	assert.NoError(t, dao.UpdateByNo(sample.OrderNo, m))
	m = map[string]interface{}{
		"order_No":  sample.OrderNo,
		"user_name": "testAlert",
		"amount":    sample.Amount,
		"file_url":  sample.FileUrl,
	}
	assert.NoError(t, dao.UpdateByNo(sample.OrderNo, m))
	assert.NoError(t, dao.DeleteById(strconv.Itoa(sample.ID)))
}

func TestQueryOrderById(t *testing.T) {
	dao, sample := initial()
	assert.NoError(t, dao.CreateOrder(sample))
	sample, err := dao.QueryOrderById(strconv.Itoa(sample.ID))
	assert.NoError(t, err)
	assert.NotEmpty(t, sample)
	assert.NoError(t, dao.DeleteById(strconv.Itoa(sample.ID)))
}

func TestQueryOrdersByName(t *testing.T) {
	dao, sample := initial()
	assert.NoError(t, dao.CreateOrder(sample))
	order, err := dao.QueryOrdersByName(sample.UserName, "amount", "desc")
	assert.NoError(t, err)
	assert.NotEmpty(t, order)
	assert.NoError(t, dao.DeleteById(strconv.Itoa(sample.ID)))
}

func TestQueryOrders(t *testing.T) {
	dao, sample := initial()
	assert.NoError(t, dao.CreateOrder(sample))
	page, pageSize := 0, 10
	orders, err := dao.QueryOrders(page, pageSize)
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
	page, pageSize = 0, 100
	orders, err = dao.QueryOrders(page, pageSize)
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
	page, pageSize = -1, 100
	orders, err = dao.QueryOrders(page, pageSize)
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
	page, pageSize = 1, 10
	orders, err = dao.QueryOrders(page, pageSize)
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
	assert.NoError(t, dao.DeleteById(strconv.Itoa(sample.ID)))
}

func TestUpdateById(t *testing.T) {
	dao, sample := initial()
	assert.NoError(t, dao.CreateOrder(sample))
	assert.NoError(t, dao.UpdateById(map[string]interface{}{
		"file_url": ".././test",
	}, strconv.Itoa(sample.ID)))
	assert.NoError(t, dao.DeleteById(strconv.Itoa(sample.ID)))
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