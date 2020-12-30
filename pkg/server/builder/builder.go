package builder

import (
	"testGinandGorm/common/mySQL"
	"testGinandGorm/common/redis"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/dao/mysql"
	"testGinandGorm/pkg/server/handler"
	"testGinandGorm/pkg/service"
	"testGinandGorm/pkg/service/profile"
)

type Router struct {
	Handler *handler.OrderHandler
}

func NewRouter () *Router{
	r := &Router{}
	build(r)
	return r
}

func build(r *Router)  {
	sqlDB := mySQL.DbInit()
	orderDB := mysql.NewOrderDB(sqlDB)
	orderCache := redis.NewRedisCache(redis.DEFAULT)
	orderDao := dao.NewOrderDao(orderDB, &orderCache)
	orderService := profile.NewOrderService(orderDao)
	profileRuntime := service.NewProfileRuntime(orderService)
	r.Handler = handler.NewOrderHandler(profileRuntime)
}



//func Construct() *handler.OrderHandler {
//	sqlDB := mySQL.DbInit()
//	orderDB := mysql.NewOrderDB(sqlDB)
//	orderCache := redis.NewRedisCache(redis.DEFAULT)
//	orderDao := dao.NewOrderDao(orderDB, &orderCache)
//	orderService := profile.NewOrderService(orderDao)
//	profileRuntime := service.NewProfileRuntime(orderService)
//	return handler.NewOrderHandler(profileRuntime)
//}
