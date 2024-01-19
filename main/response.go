package main

import (
	"github.com/gofiber/fiber/v2"
)

const MSG_DEFAULT = ""
const MSG_NO_RETRY = ""

type Response struct {
	Access  bool        `json:"access"`
	Msg     string      `json:"msg"`
	Error   []string    `json:"error"`
	Content interface{} `json:"content"`
	ctx     *fiber.Ctx
}

func (response *Response) send(HTTPCode int) {
	_ = response.ctx.Status(HTTPCode).JSON(response)
}
