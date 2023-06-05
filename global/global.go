package global

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func Init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	Db = db

	if err != nil {
		panic("fail to connect db")
	}
}
