package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shrt-server/types/entity"
	"shrt-server/utility/configuration"
	"time"
)

func Init(config configuration.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.MySqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(17000)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if config.AllowAutoMigrate {
		// Create all tables
		err := db.AutoMigrate(entity.Shrt{})
		if err != nil {
			panic(err)
		}
	}

	return db
}
