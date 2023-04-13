package routers

import (
	"github.com/BautistaBianculli/HelloFiber/bd"
	"github.com/BautistaBianculli/HelloFiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	var articles []models.Article
	articles, err :=bd.GetArticles()
	if err!= nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status" : "error",
			"message": err.Error(),
		})
	}
	// json , err := json.Marshal(articles)
	// if err != nil{
	// 	c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"status"	:	"error",
	// 		"message"	:	err.Error(),
	// 	})
	// }
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status" 	:	"ok",
		"value"	: articles,
		},
	)
	return nil
}

