package handlers

import tele "gopkg.in/telebot.v3"

func Start(c tele.Context) error {
	return c.Send("Hello. Command for get qr code /qr YOUR TEXT. Max len your text 4 096 symbols.")
}
