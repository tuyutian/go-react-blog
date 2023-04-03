package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"tomaxut/config"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	conf := config.Get()
	db, err := gorm.Open(sqlite.Open(conf.SQLLite), &gorm.Config{})

	if err == nil {
		if conf.SQLDebug {
			DB = db.Debug()
		} else {
			DB = db
		}

		return db, err
	}
	return nil, err
}
