package Type

import (
	"fiber/application/config/result"
	typeService "fiber/application/service/type"
	"github.com/gofiber/fiber/v2"
)

func GetAllType(c *fiber.Ctx) error {
	list := typeService.GetAll()
	return result.Success("success", &fiber.Map{"list": list}, c)
}
