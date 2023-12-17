package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shrt-server/types/entity"
	"shrt-server/utilities/config"
)

func Init() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.C.MySqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if config.C.AllowAutoMigrate {
		// Create all tables
		err := db.AutoMigrate(entity.Shrt{})
		if err != nil {
			panic(err)
		}
	}

	return db
}
