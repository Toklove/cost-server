package BillService

import (
	db "fiber/application/database"
	BillModel "fiber/application/model/Bill"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddBill(bill *BillModel.Bill) *gorm.DB {
	return db.DB.Create(&bill)
}
func GetAllBill(user int) []BillModel.Bill {
	var list []BillModel.Bill
	db.DB.Where("user_id", user).Find(&list)
	return list
}
func GetBill(bill *BillModel.Bill) (*BillModel.Bill, error) {
	r := db.DB.Where("id", bill.Id).Where("user_id", bill.UserId).Find(&bill)
	if r.RowsAffected != 0 {
		return bill, nil
	} else {
		return bill, &fiber.Error{Code: 500, Message: "无此账单,请重试"}
	}
}

func UpdateBillDetail(bill *BillModel.Bill) (*BillModel.Bill, error) {
	r := db.DB.Where("id", bill.Id).Where("user_id", bill.UserId).Updates(&bill).Find(&bill)
	if r.RowsAffected != 0 {
		return bill, nil
	} else {
		return bill, &fiber.Error{Code: 500, Message: "更新失败,请重试"}
	}
}
func RemoveBillDetail(bill *BillModel.Bill) *gorm.DB {
	return db.DB.Where("id", bill.Id).Where("user_id", bill.UserId).Delete(&bill)
}
