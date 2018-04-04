package main

import (
	"github.com/jinzhu/gorm"
	"project/database"
)

func main() {
	db, err := database.Database{}
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
}