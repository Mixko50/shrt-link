package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shrt-server/utilities/config"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open(config.C.MySqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if config.C.AllowAutoMigrate {
		// Create all tables
		err := db.AutoMigrate(Shrts{})
		if err != nil {
			panic(err)
		}
	}

	DB = db
}
