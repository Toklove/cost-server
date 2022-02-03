package main

import (
	"fiber/application/config/profile"
	db "fiber/application/database"
	"fiber/application/middleware"
	"fiber/application/router"
	"fiber/core"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

var app = core.AppCore

func main() {
	app.Use(cors.New())
	profile.Main()          //加载配置文件
	db.Main()               //加载数据库
	router.BeforeRouter()   //加载前置路由
	middleware.Middleware() //加载中间件
	router.AfterRouter()    //加载后置路由
	log.Fatal(app.Listen(fmt.Sprintf(":%v", profile.Profile.Server.Port)))
}
