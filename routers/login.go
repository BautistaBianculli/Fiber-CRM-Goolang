package routers

import(
	"time"
	"github.com/BautistaBianculli/HelloFiber/bd"
	"github.com/BautistaBianculli/HelloFiber/models"
	"github.com/BautistaBianculli/HelloFiber/jwt"
	"github.com/gofiber/fiber/v2"

)


func Login(c *fiber.Ctx)error {
	user := new(models.User)
	err := c.BodyParser(&user)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"cannot parse json",
			"message": err.Error(),
		})
	}
	status, err := bd.Login(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status"	: "error",
			"message"	: err.Error(),
		})
	}
	if !status{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status"	: "error",
			"message"	: "Unauthorized",
		})
	}

	jwtKey, err := jwt.GenerateJWT(*user)
	if err !=nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status"	: "error",
			"message"	: err.Error(),
		})
	}
	c.Cookie(&fiber.Cookie{
		Name: "token",
		Value: jwtKey,
		Path: "/",
		Expires: time.Now().Add(time.Hour * 24),
		Secure: false,
		HTTPOnly: true,
		Domain: "localhost",
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status" 	: "OK",
		"message" 	: jwtKey,
	})
}