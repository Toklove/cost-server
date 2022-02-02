package middleware

import (
	"fiber/core"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Middleware() {

	app := core.AppCore
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.JSON(&fiber.Map{
				"code": 500,
				"msg":  "Token错误或缺失",
			})
		},
	}))

}
