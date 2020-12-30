package builder

import (
	"testGinandGorm/common/mySQL"
	"testGinandGorm/common/redis"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/db"
	"testGinandGorm/pkg/handler"
	"testGinandGorm/pkg/service"
	"testGinandGorm/pkg/service/profile"
)


func Construct() *handler.OrderHandler  {
	sqlDB := mySQL.DbInit()
	orderDB := db.NewOrderDB(sqlDB)
	orderCache := redis.NewRedisCache(redis.DEFAULT)
	orderDao := dao.NewOrderDao(orderDB, &orderCache)
	orderService := profile.NewOrderService(orderDao)
	profileRuntime := service.NewProfileRuntime(orderService)
	return handler.NewOrderHandler(profileRuntime)
}
