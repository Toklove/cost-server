package bill

import (
	"fiber/application/config/result"
	BillModel "fiber/application/model/Bill"
	BillService "fiber/application/service/Bill"
	"fiber/application/utils/decode"
	"fiber/application/utils/page"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func AddBill(c *fiber.Ctx) error {

	bill := new(BillModel.Bill)

	if err := c.BodyParser(bill); err != nil {
		return result.Error("服务器错误", c)
	}
	if bill.Amount == "" || bill.TypeId == "" || bill.TypeName == "" || bill.Date == 0 || bill.PayType == 0 {
		return result.Error("参数错误", c)
	}
	bill.UserId = int(decode.UId(c))

	res := BillService.AddBill(bill)
	if res.RowsAffected != 1 {
		return result.Error("添加错误请重试", c)
	}

	return result.Success("success", &fiber.Map{}, c)
}
func GetBill(c *fiber.Ctx) error {
	listItem := new(BillModel.List)
	if err := c.QueryParser(listItem); err != nil {
		return err
	}
	uid := decode.UId(c)
	list := BillService.GetAllBill(int(uid))
	totalExpense, totalIncome := 0, 0
	if listItem.PageSize != 0 && listItem.Page != 0 {
		sliceStart, sliceEnd := page.SlicePage(listItem.Page, listItem.PageSize, len(list))
		list = list[sliceStart:sliceEnd]
	}
	var resList []BillModel.SelectList
	hasMap := make(map[string]string)
	if listItem.TypeId == "all" {
		for _, v := range list {
			r, _ := strconv.Atoi(v.Amount)
			if v.PayType == 1 {
				totalExpense += r
			} else {
				totalIncome += r
			}
			var showList []BillModel.Bill
			date := time.Unix(v.Date, 0).Format("2006-01-02")
			if hasMap[date] == date {
				continue
			}
			for _, v := range list {
				if time.Unix(v.Date, 0).Format("2006-01-02") == date {
					showList = append(showList, v)
				}
			}
			resList = append(resList, BillModel.SelectList{Date: date, Bills: showList})
			hasMap[date] = date
		}
	} else {
		for _, v := range list {
			var showList []BillModel.Bill
			date := time.Unix(v.Date, 0).Format("2006-01-02")
			if hasMap[date] == date || v.TypeId != listItem.TypeId {
				continue
			}
			for _, v := range list {
				if time.Unix(v.Date, 0).Format("2006-01-02") == date && v.TypeId == listItem.TypeId {
					showList = append(showList, v)
				}
			}
			resList = append(resList, BillModel.SelectList{Date: date, Bills: showList})
			hasMap[date] = date
		}
	}
	return result.Success("获取成功", &fiber.Map{
		"total_expense": totalExpense,
		"total_income":  totalIncome,
		"list":          resList}, c)
}
func GetBillDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return result.Error("账单ID不能为空", c)
	}
	bill := new(BillModel.Bill)
	bill.Id, _ = strconv.Atoi(id)
	bill.UserId = int(decode.UId(c))
	r, err := BillService.GetBill(bill)
	if err != nil {
		return result.Error(err.Error(), c)
	}
	return result.Success("success", &fiber.Map{
		"bill": r,
	}, c)
}
func UpdateBillDetail(c *fiber.Ctx) error {
	detail := new(BillModel.Bill)
	if err := c.BodyParser(detail); err != nil {
		return err
	}
	detail.Id, _ = strconv.Atoi(c.Params("id"))
	detail.UserId = int(decode.UId(c))
	r, err := BillService.UpdateBillDetail(detail)
	if err != nil {
		return result.Error(err.Error(), c)
	}
	return result.Success("success", &fiber.Map{"bill": r}, c)
}
func RemoveBillDetail(c *fiber.Ctx) error {
	detail := new(BillModel.Bill)
	detail.Id, _ = strconv.Atoi(c.Params("id"))
	detail.UserId = int(decode.UId(c))
	r := BillService.RemoveBillDetail(detail)
	if r.RowsAffected == 0 {
		return result.Error("删除失败请重试", c)
	} else {
		return result.Success("success", &fiber.Map{"bill": r.RowsAffected}, c)
	}
}
func GetEchartsData(c *fiber.Ctx) error {
	dataTime := new(BillModel.Data)
	err := c.BodyParser(dataTime)
	if err != nil {
		return err
	}
	var resultList []BillModel.Bill
	var res []BillModel.DataItem
	hasMap := make(map[BillModel.DataItem]BillModel.DataItem)
	list := BillService.GetAllBill(int(decode.UId(c)))
	totalExpense, totalIncome := 0, 0
	//筛选出时间范围内的数据
	for _, v := range list {
		data := int(v.Date)
		if dataTime.Start <= data && data <= dataTime.End {
			resultList = append(resultList, v)
		}
	}
	for _, v := range resultList {
		r, _ := strconv.Atoi(v.Amount)
		if v.PayType == 1 {
			totalExpense += r
		} else {
			totalIncome += r
		}
		var nowItem BillModel.DataItem
		nowItem.TypeName = v.TypeName
		nowItem.PayType = v.PayType
		nowItem.TypeId = v.TypeId
		nowItem.Number = 0
		for _, v := range resultList {
			if v.PayType == nowItem.PayType && v.TypeId == nowItem.TypeId && v.TypeName == nowItem.TypeName {
				r, _ := strconv.Atoi(v.Amount)
				nowItem.Number += r
			}
		}
		if hasMap[nowItem] == nowItem {
			continue
		}
		res = append(res, nowItem)
		hasMap[nowItem] = nowItem
	}

	return result.Success("success", &fiber.Map{
		"total_expense": totalExpense,
		"total_income":  totalIncome,
		"total_data":    res,
	}, c)
}
