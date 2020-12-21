package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/common/db"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/handler"
	"testGinandGorm/pkg/router"
	"testGinandGorm/pkg/service"
)

func main() {
	Db := db.DbInit()
	orderDao := dao.NewOrderDao(Db)
	orderService := service.NewService(orderDao)
	orderHander := handler.NewHandler(orderService)
	fmt.Println(orderHander)
	router.BindRoute(orderHander)
	db.DbClose(Db)
}

//func main2() {
//	redisConn := db.RedisInit()
//	Db := db.DbInit()
//	orderDao := dao.NewOrderDao(Db)
//	orderService := service.NewService(orderDao)
//	orderHander := handler.NewHandler(orderService)
//	fmt.Println(orderHander)
//	router.BindRoute(orderHander)
//
//
//	myDB := alert.NewOrderDB(db.DbInit())
//	myDao := alert.NewMyOrderDao(*myDB,&redisConn)
//	myService := alert.NewOrderService(myDao)
//	myHandler := alert.NewOrderHandler(myService)
//	fmt.Println(myHandler)
//	//router.BindRoute(myHandler)
//	db.DbClose(Db)
//}
