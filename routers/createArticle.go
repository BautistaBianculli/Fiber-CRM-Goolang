package routers

import (
	// "github.com/BautistaBianculli/HelloFiber/bd"
	"strconv"

	"github.com/BautistaBianculli/HelloFiber/models"
	"github.com/BautistaBianculli/HelloFiber/bd"
	"github.com/gofiber/fiber/v2"
)



func CreateArticle(c *fiber.Ctx)error {
	article := new(models.Article)
	if len(c.FormValue("codart")) == 0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status" : "error",
			"message": "Send codart please",
		})
	}
	article.Codart = c.FormValue("codart")
	if len(c.FormValue("descripcion")) == 0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status"	: "error",
			"message"	: "Send descripcion for the article",
		})
	}
	article.Descripcion = c.FormValue("descripcion")
	article.PrecioLista, _  = strconv.ParseFloat(c.FormValue("precioLista", "0"),64)
	article.Bonif1, _  = strconv.ParseFloat(c.FormValue("bonif1", "0"),64)
	article.Bonif2, _  = strconv.ParseFloat(c.FormValue("bonif2", "0"),64)
	article.Prventa, _  = strconv.ParseFloat(c.FormValue("prventa", "0"),64)
	status, err := bd.CreateArticle(article)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status"	: "error",
			"message"	: err.Error(),
		})
	}
	if !status{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status"	: "error",
			"message"	: "can't create the article",
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status"	: "error",
		"message"	: article,
	})
}