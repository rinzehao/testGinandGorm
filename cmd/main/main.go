package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go.uber.org/zap"
	"testGinandGorm/common/log"
	"testGinandGorm/pkg/server"
	"testGinandGorm/pkg/server/builder"
)

func main() {
	router := builder.NewRouter()
	log.InitLogger()
	if err := server.BindRoute(router.Handler); err != nil {
		log.Logger.Error("Fail to Route OrderHandler", zap.Error(err))
		panic(err)
	}
}

//func main() {
//	sqlDB := mysql.DbInit()
//	defer sqlDB.Close()
//	orderDB := mysql.NewOrderDB(sqlDB)
//	orderCache := redis.NewRedisCache(redis.DEFAULT)
//	orderDao := dao.NewOrderDao(orderDB, &orderCache)
//	orderService := profile.NewOrderService(orderDao)
//	runtime := service.NewProfileRuntime(orderService)
//	orderHandler := handler.NewOrderHandler(runtime)
//	log.InitLogger()
//	if err := router.BindRoute(orderHandler); err != nil {
//		log.SugarLogger.Errorf("Fail to Route OrderHandler : InputID =%s , Error = %s", err)
//		panic(err)
//	}
//}
