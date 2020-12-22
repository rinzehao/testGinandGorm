package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/common/db"
	"testGinandGorm/pkg/alert"
	"testGinandGorm/pkg/router"
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
	orderDB := alert.NewOrderDB(db)
	orderCache := alert.NewCache()
	orderDao := alert.NewMyOrderDao(*orderDB,orderCache)
	orderService := alert.NewOrderService(*orderDao)
	ordeHandler := alert.NewOrderHandler(orderService)
	router.BindRoute(ordeHandler)
	db.Close()
}
