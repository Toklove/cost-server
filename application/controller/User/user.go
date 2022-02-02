package user

import (
	"fiber/application/config/result"
	"fiber/application/model/User"
	userService "fiber/application/service/user"
	"fiber/application/utils/decode"
	MD5 "fiber/application/utils/md5"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func Login(c *fiber.Ctx) error {
	user := new(UserModel.User)
	if err := c.BodyParser(user); err != nil {
		return result.Error("服务器错误", c)
	}
	res := userService.UserLogin(user)
	if res.Error != nil {
		return result.Error("请检查账号密码", c)
	}
	claims := jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return result.Success("success", &fiber.Map{"token": t}, c)
}

func Register(c *fiber.Ctx) error {
	user := new(UserModel.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	user.Password = MD5.Create(user.Password)
	user.Avatar = "https://q1.qlogo.cn/g?b=qq&nk=2586438083&s=640"
	user.CreateAt = time.Now().Unix()
	user.Slogan = "冰冻三尺非一日之寒"
	if user.Username == "" || user.Password == "" {
		return result.Error("请完善用户信息", c)
	}

	if err := userService.QueryUser(user); err.Error == nil {
		return result.Error("当前用户名已存在", c)
	}

	if err := userService.CreateUser(user); err.Error != nil {
		return result.Error("注册失败请检查您的信息", c)
	}
	return result.Success("success", &fiber.Map{"userInfo": &fiber.Map{
		"username": user.Username,
	}}, c)

}

func GetUserInfo(c *fiber.Ctx) error {
	user := new(UserModel.User)

	user.Username = decode.UName(c)
	userService.QueryUser(user)
	return result.Success("获取成功", &fiber.Map{
		"userInfo": &fiber.Map{
			"username": user.Username,
			"avatar":   user.Avatar,
			"slogan":   user.Slogan,
		},
	}, c)
}

func EditUserInfo(c *fiber.Ctx) error {
	user := new(UserModel.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	user.Username = decode.UName(c)
	if user.Password != "" {
		user.Password = MD5.Create(user.Password)
	}
	res := userService.UpdateUser(user)
	if res.RowsAffected == 0 {
		return result.Error("未修改成功", c)
	}
	return result.Success("修改成功", &fiber.Map{}, c)
}
