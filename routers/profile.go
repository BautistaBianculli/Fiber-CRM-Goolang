package routers

import (
	"github.com/BautistaBianculli/HelloFiber/bd"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)



func GetProfile(c *fiber.Ctx)error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)
	users,status, err := bd.Profile(id)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status"	:	"error",
			"message"	: 	err.Error(),
		})	
	}
	if !status{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status"	:	"error",
			"message"	: 	"Invalid user",
		})	
	}
	users.Password = ""
	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"status"	:	"ok",
	// 	"message"	:	map[string]interface{}{
	// 		"id"		: users.ID,
	// 		"user"		: users.User,
	// 		"FirstName" : users.FirstName,
	// 		"LastName"	: users.LastName,
	// 	},
	// })
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status"	:	"ok",
			"value"		:	users,
	})
}