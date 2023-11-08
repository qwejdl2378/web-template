package store

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsnTmpl = `%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`
)

var db *gorm.DB

func InitDB() error {
	var err error
	if db == nil {
		dsn := fmt.Sprintf(
			dsnTmpl,
			"your_user",
			"your_password",
			"your_host",
			"your_port",
			"your_db",
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		return err
	}
	return nil
}

func GetDBInstance() *gorm.DB {
	return db
}
