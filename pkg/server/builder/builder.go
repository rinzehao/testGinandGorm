package builder

import (
	"testGinandGorm/common/mysql"
	"testGinandGorm/common/redis"
	"testGinandGorm/pkg/dao"
	mysql2"testGinandGorm/pkg/dao/mysql"
	"testGinandGorm/pkg/server"
	"testGinandGorm/pkg/service"
	"testGinandGorm/pkg/service/profile"
	"testGinandGorm/pkg/service/profile/profile-item"
)

type Router struct {
	Handler *server.Handler
}

func NewRouter() *Router {
	r := &Router{}
	build(r)
	return r
}

func build(r *Router) {
	sqlDB := mysql.DbInit()
	orderDB := mysql2.NewOrderDB(sqlDB)
	orderCache := redis.NewRedisCache(redis.DEFAULT)
	orderDao := dao.NewOrderDao(orderDB, &orderCache)
	orderService := profile_item.NewOrderService(orderDao)
	profileRuntime := profile.NewProfileRuntime(orderService)
	profileManager := service.NewProfileManager(profileRuntime)
	r.Handler = server.NewHandler(profileManager)
}

//func Construct() *handler.OrderHandler {
//	sqlDB := mysql.DbInit()
//	orderDB := mysql.NewOrderDB(sqlDB)
//	orderCache := redis.NewRedisCache(redis.DEFAULT)
//	orderDao := dao.NewOrderDao(orderDB, &orderCache)
//	orderService := profile.NewOrderService(orderDao)
//	profileRuntime := service.NewProfileRuntime(orderService)
//	return handler.NewOrderHandler(profileRuntime)
//}
