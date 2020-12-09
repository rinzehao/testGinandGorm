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

//func init() {
//	var Db *gorm.DB
//	var err error
//	Db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test_gorm?charset=utf8&parseTime=True&loc=Local")
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	Db.SingularTable(true)
//	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Demo_order{})
//}
func main() {
	Db := db.DbInit()
	orderDao := dao.NewOrderDao(Db)
	orderService := service.NewService(orderDao)
	orderHander := handler.NewHandler(orderService)
	fmt.Println(orderHander)
	router.BindRoute(orderHander)
	db.DbClose(Db)
}


//在分支内进行修改