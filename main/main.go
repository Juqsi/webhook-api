package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"time"
)

func main() {
	app := fiber.New()

	fmt.Println("▶️ test database connection ...")
	dbError := errors.New("")
	var db *sql.DB
	for dbError != nil {
		db, dbError = dbConncetion()
		defer db.Close()
		if dbError != nil {
			fmt.Println("❌  no database connection -> " + dbError.Error())
		} else {
			fmt.Println("✅  connected")
		}
		time.Sleep(time.Second)
	}

	//start DB connection routine
	go dbReachable(db)

	app.Use(cors.New())

	app.Use(requestid.New())

	//Monitor
	app.Get("/monitor", monitor.New(monitor.Config{Title: "Justus template"}))

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

	//all functions below needs an Authentication e.g. Bearer u can use Libraries like Gocloak or Firebase
	app.Use(authMiddleware)

	//manage POST-Request
	app.Post("/pfad")

	//manage GET-Request
	//Example with parameter
	app.Get("/pfad/:para")

	fmt.Println("▶️ start server")
	//start Server with HTTP localhost or 127.0.0.1
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("❌ cant Start Server probably port 3000 in use")
		panic(err)
	}
}
