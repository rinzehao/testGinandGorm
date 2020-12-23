package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/common/db"
	"testGinandGorm/common/redis_utils"
	"testGinandGorm/pkg/dao"
	db2 "testGinandGorm/pkg/db"
	"testGinandGorm/pkg/handler"
	"testGinandGorm/pkg/router"
	"testGinandGorm/pkg/service"
)

//func main() {
//	Db := db.DbInit()
//	orderDao := dao.NewOrderDao(Db)
//	orderService := service.NewService(orderDao)
//	orderHander := handler.NewHandler(orderService)
//	fmt.Println(orderHander)
//	router.BindRoute(orderHander)
//	db.DbClose(Db)
//}

func main() {
	db := db.DbInit()
	orderDB := db2.NewMyOrderDB(db)
	orderCache := redis_utils.NewRedisCache(1e10 * 6 * 20)
	orderDao := dao.NewMyOrderDao(orderDB, &orderCache)
	orderService := service.NewOrderService(orderDao)
	ordeHandler := handler.NewOrderHandler(orderService)
	router.BindRoute(ordeHandler)
	db.Close()
}
