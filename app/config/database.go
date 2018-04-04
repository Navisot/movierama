package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

type Database struct {
	DB *gorm.DB
}

type DBConfig struct {
	Dialect string
	Host    string
	DBName  string
	DBUser  string
	DBPass  string
	SSLMode string
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Dialect: "postgres",
		Host:    os.Getenv("CUE_DB_HOST"),
		DBName:  os.Getenv("CUE_DB_NAME"),
		DBUser:  os.Getenv("CUE_DB_USER"),
		DBPass:  os.Getenv("CUE_DB_PASS"),
		SSLMode: "disabled",
	}
}

func NewDB(config *DBConfig) *Database {

	source := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.DBUser, config.DBPass, config.DBName)

	db, err := gorm.Open(config.Dialect, source)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return &Database{DB: db}
}
