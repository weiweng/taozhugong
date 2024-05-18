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

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/hello", handler.Hello)

	b.Handle("/start", handler.Start)

	log.Info("陶朱公启动................................................................")
	b.Start()
}
