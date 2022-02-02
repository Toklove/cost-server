package upload

import (
	"fiber/application/config/result"
	FileModel "fiber/application/model/File"
	fileService "fiber/application/service/file"
	"fiber/application/utils/decode"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func UpImg(c *fiber.Ctx) error {
	saveFile := new(FileModel.Resources)
	file, _ := c.FormFile("file")
	user := decode.UName(c)
	fileInfo := fmt.Sprintf("./uploads/%v-%v.png", user, time.Now().Unix())
	err := c.SaveFile(file, fileInfo)
	if err != nil {
		return err
	}
	saveFile.Url = fileInfo
	saveFile.UserName = user
	go fileService.SaveData(saveFile)
	return result.Success("上传成功", &fiber.Map{
		"res": fileInfo,
	}, c)
}
