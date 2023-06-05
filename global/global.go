package global

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	GLOBAL_DB *gorm.DB
)

func Init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	GLOBAL_DB = db

	if err != nil {
		panic("fail to connect db")
	}
}
