package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/common/mySQL_db"
	"testGinandGorm/common/redis_utils"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/db"
	"testGinandGorm/pkg/handler"
	"testGinandGorm/pkg/router"
	"testGinandGorm/pkg/service"
)

func main() {
	sqlDB := mySQL_db.DbInit()
	orderDB := db.NewMyOrderDB(sqlDB)
	orderCache := redis_utils.NewRedisCache(1e10 * 6 * 20)
	orderDao := dao.NewMyOrderDao(orderDB, &orderCache)
	orderService := service.NewOrderService(orderDao)
	ordeHandler := handler.NewOrderHandler(orderService)
	router.BindRoute(ordeHandler)
	sqlDB.Close()
}
