package result

import "github.com/gofiber/fiber/v2"

func Success(msg string, data *fiber.Map, c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}
func Error(msg string, c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"code": 500,
		"msg":  msg,
	})
}
