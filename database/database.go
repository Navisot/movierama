package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Database struct {
	DB	*gorm.DB
}
