package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/BautistaBianculli/HelloFiber/routers"
	// jwtware "github.com/gofiber/jwt"
)


func AuthCheck(ctx * fiber.Ctx) error{
	// var tokenString string
	err := routers.VerifyToken(ctx.Get("Authorization"))
	if err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status" : "error",
			"message": err.Error(),
		})
	}
	return ctx.Next()
}