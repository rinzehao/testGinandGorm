package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/common/logger"
	"testGinandGorm/common/mySQL_db"
	"testGinandGorm/common/redis"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/db"
	"testGinandGorm/pkg/handler"
	"testGinandGorm/pkg/router"
	"testGinandGorm/pkg/service"
)

func main() {
	sqlDB := mySQL_db.DbInit()
	defer sqlDB.Close()
	orderDB := db.NewMyOrderDB(sqlDB)
	orderCache := redis.NewRedisCache(1e10 * 6 * 20)
	orderDao := dao.NewMyOrderDao(orderDB, &orderCache)
	orderService := service.NewOrderService(orderDao)
	runtime := service.NewProfileRuntime(orderService)
	ordeHandler := handler.NewOrderHandler(runtime)
	logger.InitLogger()
	if err := router.BindRoute(ordeHandler); err != nil {
		logger.SugarLogger.Errorf("Fail to Route OrderHandler : InputID =%s , Error = %s", err)
		panic(err)
	}
}
