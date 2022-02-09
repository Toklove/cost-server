package BillModel

type Bill struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	PayType  int    `json:"pay_type" form:"pay_type"`
	Amount   string `json:"amount" form:"amount"`
	Date     int64  `json:"date" form:"date"`
	TypeId   string `json:"type_id" form:"type_id"`
	TypeName string `json:"type_name" form:"type_name"`
	UserId   int    `json:"user_id" form:"user_id"`
	Remark   string `json:"remark" form:"remark"`
}
type List struct {
	Date     string `json:"date" query:"date"`
	Page     int    `json:"page" query:"page"`
	PageSize int    `json:"page_size" query:"page_size"`
	TypeId   string `json:"type_id" query:"type_id"`
}

type SelectList struct {
	Date  string `json:"date"`
	Bills []Bill `json:"bills"`
}

type Data struct {
	Date string `json:"date" form:"date"`
}
type DataItem struct {
	Number   int    `json:"number"`
	PayType  int    `json:"pay_type"`
	TypeId   string `json:"type_id"`
	TypeName string `json:"type_name"`
}
