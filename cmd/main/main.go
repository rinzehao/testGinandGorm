package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/common/logger"
	"testGinandGorm/pkg/server"
	"testGinandGorm/pkg/server/builder"
)

func main() {
	orderHandler :=builder.Construct()
	logger.InitLogger()
	if err := server.BindRoute(orderHandler); err != nil {
		logger.SugarLogger.Errorf("Fail to Route OrderHandler : InputID =%s , Error = %s", err)
		panic(err)
	}
}


//func main() {
//	sqlDB := mySQL.DbInit()
//	defer sqlDB.Close()
//	orderDB := mysql.NewOrderDB(sqlDB)
//	orderCache := redis.NewRedisCache(redis.DEFAULT)
//	orderDao := dao.NewOrderDao(orderDB, &orderCache)
//	orderService := profile.NewOrderService(orderDao)
//	runtime := service.NewProfileRuntime(orderService)
//	orderHandler := handler.NewOrderHandler(runtime)
//	logger.InitLogger()
//	if err := router.BindRoute(orderHandler); err != nil {
//		logger.SugarLogger.Errorf("Fail to Route OrderHandler : InputID =%s , Error = %s", err)
//		panic(err)
//	}
//}

