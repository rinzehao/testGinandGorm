package profile

//import (
//	"github.com/stretchr/testify/assert"
//	"math/rand"
//	"strconv"
//	"testGinandGorm/common/mySQL"
//	"testGinandGorm/common/redis"
//	"testGinandGorm/pkg/dao"
//	"testGinandGorm/pkg/db"
//	"testGinandGorm/pkg/model"
//	"testing"
//	"time"
//)
//
//func initial() (service OrderService, sample model.OrderMade) {
//	sqlDb := mySQL.DbInit()
//	orderDb := db.NewMyOrderDB(sqlDb)
//	cache := redis.NewRedisCache(1e10 * 6 * 20)
//	dao := dao.NewMyOrderDao(orderDb, &cache)
//	orderService := NewOrderService(dao)
//	//orderSample := &model.DemoOrder{OrderNo: time.Now().Format("2006-01-02 15:04:05")+queryRandomString(5), UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
//	timeNow := time.Now()
//	hour:=timeNow.Hour()        //小时
//	minute:=timeNow.Minute()      //分钟
//	second:=timeNow.Second()      //秒
//	nanoSecond:=timeNow.Nanosecond()  //纳秒
//	str := strconv.Itoa(hour)+":"+strconv.Itoa(minute)+":"+strconv.Itoa(second)+":"+strconv.Itoa(nanoSecond)
//	orderSample := &model.DemoOrder{OrderNo: str, UserName: "raious", Amount: 444, Status: "over", FileUrl: ".././pkg/dao"}
//	sample = model.OrderMade{
//		Order:   orderSample,
//		OrderID: strconv.Itoa(orderSample.ID),
//		OrderNo: orderSample.OrderNo,
//	}
//	return orderService, sample
//}
//
//func TestCreateOrderByOrderNo(t *testing.T) {
//	service, sample := initial()
//	assert.NoError(t, service.CreateOrder(&sample))
//}
//
//func TestDeleteOrderById(t *testing.T) {
//	service, sample := initial()
//	assert.NoError(t, service.CreateOrder(&sample))
//	assert.NoError(t, service.DeleteOrderById(&sample))
//}
//
//func TestQueryOrderById(t *testing.T) {
//	service, sample := initial()
//	assert.NoError(t, service.CreateOrder(&sample))
//	order :=sample.Order.(*model.DemoOrder)
//	sample.OrderID =strconv.Itoa(order.ID)
//	err := service.QueryOrderById(&sample)
//	assert.NoError(t, err)
//	assert.NotEmpty(t, sample.Order)
//}
//
//func TestUpdateByOrderNo(t *testing.T) {
//	service, sample := initial()
//	assert.NoError(t, service.CreateOrder(&sample))
//	m := map[string]interface{}{
//		"Id":        sample.Order,
//		"order_No":  sample.OrderNo,
//		"user_name": sample.UserName+ queryRandomString(5),
//	}
//	sample.UpdateMap =m
//	assert.NoError(t, service.UpdateByOrderNo(&sample))
//}
//
//func TestQueryOrders(t *testing.T) {
//	service, sample := initial()
//	assert.NoError(t, service.CreateOrder(&sample))
//	sample.Page=1
//	sample.PageSize=100
//	err := service.QueryOrders(&sample)
//	assert.NoError(t, err)
//	assert.NotEmpty(t, sample.Group)
//}
//
//func TestQueryOrdersByName(t *testing.T) {
//	service, sample := initial()
//	assert.NoError(t, service.CreateOrder(&sample))
//	sample.OrderBy="amount"
//	sample.Desc ="desc"
//	err := service.QueryOrdersByName(&sample)
//	assert.NoError(t, err)
//	assert.NotEmpty(t, sample.Group)
//}
//
//func TestUpdateUrlById(t *testing.T) {
//	service, sample := initial()
//	assert.NoError(t, service.CreateOrder(&sample))
//	m := map[string]interface{}{
//		"file_url": ".././test",
//	}
//	sample.UpdateMap =m
//	err := service.UpdateById(&sample)
//	assert.NoError(t, err)
//}
//
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
