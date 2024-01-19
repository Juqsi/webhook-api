package main

import (
	_ "context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"runtime/debug"
)

func recovery(ctx *fiber.Ctx) error {
	response := Response{
		ctx:     ctx,
		Access:  true,
		Error:   []string{},
		Content: nil,
		Msg:     "Es ist ein Interner Fehler aufgetreten, wenn es häufig passiert wende dich an den Support",
	}

	defer func() {
		if r := recover(); r != nil {
			var err error
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
			response.Error = append(response.Error, "Panic: Recovery Done")
			response.Error = append(response.Error, err.Error())
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
			response.send(fiber.StatusInternalServerError)
		}
	}()
	return ctx.Next()
}

func logging(app *fiber.App) {
	//Logger
	app.Use(logger.New(logger.Config{
		Format:     "${time} -- ${status} -- ${method} ${path} ${queryParams} ${latency} \n",
		TimeFormat: "2006-01-02 15:04:05.00000",
	}))
	file, err := os.OpenFile("./Logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	app.Use(logger.New(logger.Config{
		Output:     file,
		Format:     "${time} -- ${status} -- ${method} ${path} ${queryParams} ${latency} \n",
		TimeFormat: "2006-01-02 15:04:05.00000",
	}))
}
