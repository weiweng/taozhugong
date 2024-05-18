package handler

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

var (
	menu        = &tele.ReplyMarkup{ResizeKeyboard: true}
	selector    = &tele.ReplyMarkup{}
	BtnHelp     = menu.Text("ℹ Help")
	BtnSettings = menu.Text("⚙ Settings")
	BtnPrev     = selector.Data("⬅", "上一页")
	BtnNext     = selector.Data("➡", "下一页")
)

func Start(c tele.Context) error {
	cJSON, _ := json.Marshal(c.Bot())
	log.Infof("请求Start接口 req[%v]", string(cJSON))
	menu.Reply(
		menu.Row(BtnHelp),
		menu.Row(BtnSettings),
	)
	selector.Inline(
		selector.Row(BtnPrev, BtnNext),
	)
	log.Info("请求Start结束")
	return c.Send("Hello!", menu)
}

// On reply button pressed (message)
func Help(c tele.Context) error {
	cJSON, _ := json.Marshal(c)
	log.Infof("请求Help接口 req[%v]", string(cJSON))
	return c.Edit("Here is some help: ...")
}

// On inline button pressed (callback)
func Settings(c tele.Context) error {
	cJSON, _ := json.Marshal(c)
	log.Infof("请求Settings接口 req[%v]", string(cJSON))
	return c.Respond()
}
