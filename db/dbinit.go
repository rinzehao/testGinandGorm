package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testGinandGorm/pkg/model"
)

var Db *gorm.DB
var err error

func DbInit() error {
	Db, _ = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test_gorm?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		return err
	}
	Db.SingularTable(true)

	Db.AutoMigrate(&model.Demo_order{})

	return err
}

func DbClose() {
	defer Db.Close()
}
