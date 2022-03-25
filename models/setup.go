package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func SetUpModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/webgolanggdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error koneksi kedalam database")
	}

	db.AutoMigrate(&MataKuliah{})
	return db
}

