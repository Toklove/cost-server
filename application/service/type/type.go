package typeService

import (
	db "fiber/application/database"
	TypeModel "fiber/application/model/Type"
)

func GetAll() []TypeModel.Types {
	var list []TypeModel.Types
	db.DB.Where("user_id", 0).Find(&list)
	return list
}
