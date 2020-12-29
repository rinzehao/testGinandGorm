package profile

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testGinandGorm/common/mySQL"
	"testGinandGorm/common/redis"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/db"
	"testGinandGorm/pkg/model"
	model2 "testGinandGorm/pkg/server/model"
	"testing"
	"time"
)

func initial() (*OrderService, model.Order) {
	sqlDb := mySQL.DbInit()
	sqlDb =sqlDb.LogMode(true)
	orderDb := db.NewOrderDB(sqlDb)
	cache := redis.NewRedisCache(1e10 * 6 * 20)
	dao := dao.NewOrderDao(orderDb, &cache)
	orderService := NewOrderService(dao)
	//orderSample := &model.DemoOrder{OrderNo: time.Now().Format("2006-01-02 15:04:05")+queryRandomString(5), UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
	timeNow := time.Now()
	hour:=timeNow.Hour()        //小时
	minute:=timeNow.Minute()      //分钟
	second:=timeNow.Second()      //秒
	nanoSecond:=timeNow.Nanosecond()  //纳秒
	str := strconv.Itoa(hour)+":"+strconv.Itoa(minute)+":"+strconv.Itoa(second)+":"+strconv.Itoa(nanoSecond)
	orderSample := &model.Order{
		ID:       0,
		OrderNo:  str,
		UserName: "raious",
		Amount:   444,
		Status:   "over",
		FileUrl:  ".././pkg/dao",
	}
	return orderService, *orderSample
}

func TestCreateOrderByOrderNo(t *testing.T) {
	orderService, order := initial()
	sample := &model2.QueryCtx{
		ItemTyp: "order",
		Req:     order,
	}
	orderService.Create(sample)
	assert.NoError(t, orderService.Create(sample))
	id :=strconv.Itoa(sample.GetResult().(model.Order).ID)
	delSample := &model2.DeleteCtx{
		ItemTyp: "order",
		Req:     id,
	}
	assert.NoError(t, orderService.Delete(delSample))
}

func TestDeleteOrderById(t *testing.T) {
	orderService, order := initial()
	createSample := &model2.QueryCtx{
		ItemTyp: "order",
		Req:     order,
	}
	assert.NoError(t, orderService.Create(createSample))
	id :=strconv.Itoa(createSample.GetResult().(model.Order).ID)
	delSample := &model2.DeleteCtx{
		ItemTyp: "order",
		Req:     id,
	}
	assert.NoError(t, orderService.Delete(delSample))
}

func TestQueryOrderById(t *testing.T) {
	orderService, order := initial()
	createSample := &model2.QueryCtx{
		ItemTyp: "order",
		Req:     order,
	}
	assert.NoError(t, orderService.Create(createSample))
	id :=strconv.Itoa(createSample.GetResult().(model.Order).ID)
	querySample := &model2.QueryCtx{
		ItemTyp: "order",
		Req:     id,
	}
	assert.NoError(t, orderService.Query(querySample))
}

func TestUpdateByOrderNo(t *testing.T) {
	orderService, order := initial()
	createSample := &model2.QueryCtx{
		ItemTyp: "order",
		Req:     order,
	}
	assert.NoError(t, orderService.Create(createSample))
	m :=map[string]interface{} {
		"user_name": "maruhire",
		"amount":    17.8,
		"status":    "已修改",
	}
	updateSample := &model2.UpdateCtx{
		ItemTyp:  "order",
		Identify: order.OrderNo,
		Req:     m,
	}
	assert.NoError(t, orderService.UpdateByNo(updateSample))
}

func TestQueryOrders(t *testing.T) {
	orderService, _ := initial()
	sample := model2.QueryCtxs{
		ItemTyp: "order",
		ReqPage: 1,
		ReqSize: 200,
	}
	assert.NoError(t, orderService.QueryOrders(&sample))
	assert.NotEmpty(t, sample.GetResult())
}

func TestQueryOrdersByName(t *testing.T) {
	orderService, order := initial()
	sample := model2.QueryByNameCtx{
		ItemTyp:     "order",
		Req:         order.UserName,
		OrderOption: "amount",
		DescOrder:   true,
	}
	assert.NoError(t, orderService.QueryOrdersByName(&sample))
	assert.NotEmpty(t, sample.GetResult())
}

func TestUpdateUrlById(t *testing.T) {
	orderService, order := initial()
	createSample := &model2.QueryCtx{
		ItemTyp: "order",
		Req:     order,
	}
	assert.NoError(t, orderService.Create(createSample))
	id :=strconv.Itoa(createSample.GetResult().(model.Order).ID)
	sample := model2.UpdateCtx{
		ItemTyp:  "order",
		Identify: id,
		Req:      map[string]interface{}{
			"file_url": ".././test",
		},
	}
	assert.NoError(t, orderService.UpdateById(&sample))
}

//func queryRandomString(l int) string {
//	str := "0123456789abcdefghijklmnopqrstuvwxyz"
//	bytes := []byte(str)
//	result := []byte{}
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	for i := 0; i < l; i++ {
//		result = append(result, bytes[r.Intn(len(bytes))])
//	}
//	return string(result)
//}

