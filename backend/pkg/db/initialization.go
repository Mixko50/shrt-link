package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shrt-server/types/entity"
	"shrt-server/utilities/configuration"
)

func Init(config configuration.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.MySqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if config.AllowAutoMigrate {
		// Create all tables
		err := db.AutoMigrate(entity.Shrt{})
		if err != nil {
			panic(err)
		}
	}

	return db
}
