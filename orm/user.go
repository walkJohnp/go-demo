package orm

import (
	"github.com/walkjohnp/go-demo/global"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint
	Username string
	Password string
}

func ListUser() ([]User, error) {
	db := global.GLOBAL_DB.Model(&User{})

	db.AutoMigrate(&User{})

	var userList []User

	err := db.Limit(10).Offset(0).Find(&userList).Error

	return userList, err
}
