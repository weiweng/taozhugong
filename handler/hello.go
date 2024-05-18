package handler

import (
	tele "gopkg.in/telebot.v3"
)

type HelloHandler struct {
}

func NewHelloHandler() HelloHandler {
	return HelloHandler{}
	// return tele.HandlerFunc(HelloHandler{}.Handle)
}

func (h HelloHandler) Job(c tele.Context) error {
	return c.Send("Hello!")
}
