package handler

import (
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
	menu.Reply(
		menu.Row(BtnHelp),
		menu.Row(BtnSettings),
	)
	selector.Inline(
		selector.Row(BtnPrev, BtnNext),
	)
	return c.Send("Hello!", menu)
}
