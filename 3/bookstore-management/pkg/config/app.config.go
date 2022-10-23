package config

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectDB() {
	d, err := gorm.Open("mysql", "hr:0000@tcp(localhost:3306)/testinggo?parseTime=true")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
