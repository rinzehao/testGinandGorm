package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db gorm.DB

func Run(Type string, dsn string) error {
	db, err := gorm.Open(Type, dsn)
	if err != nil {
		return err
	}

	db.SingularTable(true)

	Db = *db
	return nil
}
