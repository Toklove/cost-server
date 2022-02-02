package fileService

import (
	db "fiber/application/database"
	FileModel "fiber/application/model/File"
	"gorm.io/gorm"
)

func SaveData(resources *FileModel.Resources) *gorm.DB {
	return db.DB.Table("resources").Create(&resources)
}
