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
	score := app.Group("/score")

	// Github User
	github_user := github.Group("/user")
	github_user.Get("/info", handler.GetDeveloperInfoHandler)
	github_user.Get("/repos", handler.GetUserReposHandler)
	github_user.Get("/events", handler.GetUserEventsHandler)
	github_user.Get("/nation", handler.GetUserNationalityHandler)

	// Score
	score.Get("/get", handler.GetUserScoreHandler)
}
