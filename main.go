package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/weiweng/taozhugong/handler"
	"github.com/weiweng/taozhugong/repo/logs"
	tele "gopkg.in/telebot.v3"
)

func initFuncs() {
	logs.Init()
}

func main() {
	initFuncs()

	pref := tele.Settings{
		Token:  os.Getenv("TG_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	log.Info("陶朱公bot链接测试................................................................")
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("陶朱公bot链接正常................................................................")

	b.Handle("/hello", handler.Hello)
	b.Handle("/start", handler.Start)
	b.Handle(&handler.BtnHelp, handler.Help)
	b.Handle(&handler.BtnSettings, handler.Settings)

	b.Handle(tele.OnText, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.

		var (
			user = c.Sender()
			text = c.Text()
		)
		log.Info(text)

		// Use full-fledged bot's functions
		// only if you need a result:
		msg, err := b.Send(user, text)
		if err != nil {
			return err
		}
		log.Info(msg)

		// Instead, prefer a context short-hand:
		return c.Send(text)
	})

	b.Handle(tele.OnChannelPost, func(c tele.Context) error {
		// Channel posts only.
		msg := c.Message()
		log.Info(msg)
		return nil
	})

	b.Handle(tele.OnPhoto, func(c tele.Context) error {
		// Photos only.
		// photo := c.Message().Photo
		return nil
	})

	b.Handle(tele.OnQuery, func(c tele.Context) error {
		// Incoming inline queries.
		return c.Answer(&tele.QueryResponse{})
	})

	log.Info("陶朱公启动................................................................")
	b.Start()
}
