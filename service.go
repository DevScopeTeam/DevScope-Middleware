package main

import (
	_ "DevScope-Middleware/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	handler "DevScope-Middleware/handler"
)

func regiserService(app *fiber.App) {
	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// 注册路由
	github := app.Group("/github")

	// Github User
	github_user := github.Group("/user")
	github_user.Get("/info", handler.GetDeveloperInfoHandler)
}
