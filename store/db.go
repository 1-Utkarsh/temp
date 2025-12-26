package db

import (
	"log"
	"sync"

	"gorm.io/gorm"
	"github.com/1-Utkarsh/temp/conf"
	"gorm.io/driver/postgres"
)

var Db *gorm.DB
var once sync.Once

func DbConnect() {
	logger := log.Default()
	config := conf.Get()

	once.Do(func() {
		dsn := "host=" + config.DbAddress +
			" user=" + config.DbUser +
			" password=" + config.DbPass +
			" dbname=" + config.DbName +
			" port=" + config.DbPort +
			" sslmode=disable TimeZone=Asia/Shanghai"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Fatal("Failed to connect to database:", err)
		}

		Db = db
	})

	logger.Println("Database connection established")
}

func GetDB() *gorm.DB {
	return Db
}
