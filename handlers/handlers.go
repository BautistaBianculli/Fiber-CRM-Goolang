package handlers

import (
	"github.com/BautistaBianculli/HelloFiber/routers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Handlers(){
	app := fiber.New()

	app.Post("/" , routers.CreateUser)

	login := app.Group("/login")
	login.Post("", routers.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("adruid123"),
	}))

	profile:= app.Group("/profile")
	profile.Get("",routers.GetProfile)
	//profile.Get("",middleware.AuthCheck,routers.GetProfile)

	articles:= app.Group("/articles")
	articles.Post("/create", routers.CreateArticle)
	articles.Get("/get",routers.GetArticles)

	app.Listen(":3000")
}