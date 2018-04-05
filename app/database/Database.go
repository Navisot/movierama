package database

import (
	"github.com/jinzhu/gorm"
)

var (
	// Global DB Connection
	DB *gorm.DB
	Err error
)
