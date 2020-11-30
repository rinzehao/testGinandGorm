package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/db"
	"testGinandGorm/pkg/router"
)

func main() {

	db.DbInit()
	r := gin.Default()
	router.BindRoute(r)
	db.DbClose()
}
