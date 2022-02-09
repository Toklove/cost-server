package router

import (
	"fiber/application/controller/Bill"
	"fiber/application/controller/Type"
	user "fiber/application/controller/User"
	"fiber/application/controller/upload"
	"fiber/core"
)

var app = core.AppCore
var api = app.Group("/api")
var uApi = api.Group("/user")
var tApi = api.Group("/type")

func BeforeRouter() {
	uApi.Post("/login", user.Login)
	uApi.Post("/register", user.Register)
	app.Static("/uploads", "./uploads")
}
func AfterRouter() {
	uApi.Get("/userinfo", user.GetUserInfo)
	uApi.Put("/userinfo", user.EditUserInfo)
	api.Post("/upload", upload.UpImg)
	api.Post("/bill", bill.AddBill)
	api.Get("/bill", bill.GetBill)
	api.Get("/bill/data", bill.GetEchartsData)
	api.Get("/bill/:id", bill.GetBillDetail)
	api.Put("/bill/:id", bill.UpdateBillDetail)
	api.Delete("/bill/:id", bill.RemoveBillDetail)
	tApi.Get("/list", Type.GetAllType)
}
