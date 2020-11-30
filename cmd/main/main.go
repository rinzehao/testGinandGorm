package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/db"
	"testGinandGorm/pkg/router"
)

func main() {

	db.DbInit()
	router.BindRoute()
	db.DbClose()
}
