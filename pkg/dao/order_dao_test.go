package dao


import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testGinandGorm/common/mysql"
	"testGinandGorm/common/redis"
	mysql2"testGinandGorm/pkg/dao/mysql"
	"testGinandGorm/pkg/model"
	"testing"
	"time"
)

func initial() (dao *OrderDao, sample *model.Order) {
	sqlDb := mysql.DbInit()
	orderDb := mysql2.NewOrderDB(sqlDb)
	cache := redis.NewRedisCache(1e10 * 6 * 20)
	dao = NewOrderDao(orderDb, &cache)
	sqlDb = sqlDb.LogMode(true)
	timeNow := time.Now()
	hour:=timeNow.Hour()        //小时
	minute:=timeNow.Minute()      //分钟
	second:=timeNow.Second()      //秒
	nanoSecond:=timeNow.Nanosecond()  //纳秒
	str := strconv.Itoa(hour)+":"+strconv.Itoa(minute)+":"+strconv.Itoa(second)+":"+strconv.Itoa(nanoSecond)
	sample = &model.Order{OrderNo: str, UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
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
	assert.NoError(t, dao.DeleteOrderById(strconv.Itoa(sample.ID)))
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
	assert.NoError(t, dao.DeleteOrderById(strconv.Itoa(sample.ID)))
}

func TestQueryOrderById(t *testing.T) {
	dao, sample := initial()
	assert.NoError(t, dao.CreateOrder(sample))
	sample, err := dao.QueryOrderById(strconv.Itoa(sample.ID))
	assert.NoError(t, err)
	assert.NotEmpty(t, sample)
	assert.NoError(t, dao.DeleteOrderById(strconv.Itoa(sample.ID)))
}

func TestQueryOrdersByName(t *testing.T) {
	dao, sample := initial()
	assert.NoError(t, dao.CreateOrder(sample))
	order, err := dao.QueryOrdersByName(sample.UserName, "amount", "desc")
	assert.NoError(t, err)
	assert.NotEmpty(t, order)
	assert.NoError(t, dao.DeleteOrderById(strconv.Itoa(sample.ID)))
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
	assert.NoError(t, dao.DeleteOrderById(strconv.Itoa(sample.ID)))
}

func TestUpdateById(t *testing.T) {
	dao, sample := initial()
	assert.NoError(t, dao.CreateOrder(sample))
	assert.NoError(t, dao.UpdateById(strconv.Itoa(sample.ID),
		map[string]interface{}{
			"file_url": ".././test",
		}))
	assert.NoError(t, dao.DeleteOrderById(strconv.Itoa(sample.ID)))
}
