package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "sanskar:Sanskar@123/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != null {
		panic(err)
	}
	db = d
}

func getDB() *gorm.DB {
	return db
}