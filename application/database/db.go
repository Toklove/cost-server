package db

import (
	"fiber/application/config/profile"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Main() {
	conf := profile.Profile.Database
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Hostname, conf.Port, conf.Name)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
}
