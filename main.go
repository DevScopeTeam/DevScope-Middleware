package main

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"

// 	_ "DevScope-Middleware/docs"
// )

// //	@title						DevScope中间件
// //	@version					v1.0.0
// //	@description				DevScope中间件（DevScope-Middleware）是一个基于 Fiber 的 RESTful API 服务，用于提供DevScope的后端服务。
// //	@description				注意，有 🦸 标识的接口需要管理员权限才能访问。
// //	@host						127.0.0.1:9000
// //	@BasePath					/
// //	@securityDefinitions.apikey	ApiKeyAuth
// //	@in							header
// //	@name						Authorization

// func main() {
// 	// 创建 Fiber 实例
// 	app := fiber.New(fiber.Config{
// 		BodyLimit: 1024 * 1024 * 1024,
// 	})

// 	// 启动 CORS
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins:     "*",
// 		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
// 		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
// 		AllowCredentials: true,
// 	}))

// 	// 注册服务
// 	regiserService(app)

// 	// 启动服务
// 	app.Listen(":9000")
// }
