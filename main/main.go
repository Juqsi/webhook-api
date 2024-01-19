package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(requestid.New())

	//logging - Middleware
	logging(app)

	//Recovers panics
	app.Use(recovery)

	//publish swagger infos
	//swag init in main
	app.Get("/docs", func(c *fiber.Ctx) error {
		filePath := "./main/docs/swagger.json"

		return c.SendFile(filePath)
	})

	//make swagger environment
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          "/docs",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	//manage POST-Request
	app.Post("/deploy-from-github", deployFromGithub)
	app.Post("/", deployFromGitea)

	fmt.Println("▶️ start server")
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("❌ cant Start Server probably port 3000 in use")
		panic(err)
	}
}
