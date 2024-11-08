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
	rank := app.Group("/rank")
	tag := app.Group("/tag")
	domain := app.Group("/domain")

	// Github User
	github_user := github.Group("/user")
	github_user.Get("/info", handler.GetDeveloperInfoHandler)
	github_user.Get("/repos", handler.GetUserReposHandler)
	github_user.Get("/repo", handler.GetUserRepoHandler)
	github_user.Get("/events", handler.GetUserEventsHandler)
	github_user.Get("/nation", handler.GetUserNationHandler)

	// Rank
	rank.Get("/score", handler.GetUserScoreHandler)
	rank.Get("/list", handler.GetUserRankListHandler)

	// Tag
	tag.Get("/list", handler.GetTagListHandler)
	tag.Get("/get", handler.GetTagHandler)

	// Domain
	domain.Get("/users", handler.GetUsernameListByTagHandler)
	domain.Get("/list", handler.GetUserRankListByDomainHandler)
	domain.Get("/user", handler.GetUserDoaminsHandler)
}
