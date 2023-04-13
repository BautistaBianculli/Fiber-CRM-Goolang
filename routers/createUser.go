package routers

import (

	"github.com/BautistaBianculli/HelloFiber/bd"
	"github.com/BautistaBianculli/HelloFiber/models"
	"github.com/gofiber/fiber/v2"
)


func CreateUser(c *fiber.Ctx) error{
	user := new(models.User)
	err := c.BodyParser(&user)
	if err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
			"mesaje": err.Error(),
		})
	}
	if len(user.User) == 0 || len(user.Password) == 0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot create a user",
			"mesaje": "user or password invalidate",
		})
	}
	status, err := bd.CreateUser(*user)
	if !status || err != nil{
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "cannot create the user",
			"messaje": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).SendString("Usuario creado !")
}