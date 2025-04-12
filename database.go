package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&Product{}, &Category{}, &Cart{}); err != nil {
		panic("failed to auto-migrate database: " + err.Error())
	}
}
