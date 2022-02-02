package userService

import (
	db "fiber/application/database"
	UserModel "fiber/application/model/User"
	MD5 "fiber/application/utils/md5"
	"gorm.io/gorm"
)

func QueryUser(user *UserModel.User) *gorm.DB {
	return db.DB.Where(&UserModel.User{Username: user.Username}).First(&user)
}
func UserLogin(user *UserModel.User) *gorm.DB {
	return db.DB.Where(&UserModel.User{Username: user.Username, Password: MD5.Create(user.Password)}).First(&user)
}

func CreateUser(user *UserModel.User) *gorm.DB {
	return db.DB.Create(&user)
}
func UpdateUser(user *UserModel.User) *gorm.DB {
	return db.DB.Model(&user).Where("username", &user.Username).Updates(&user)
}
