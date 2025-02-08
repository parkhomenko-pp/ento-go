package models

import (
	"ento-go/src/models/menus"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Menu struct {
	Message *tgbotapi.Message
	Player  *Player

	menus.Menuable
}

func (m *Menu) DoAction() {
	// TODO
}

func (m *Menu) GetReplyMessage() *tgbotapi.MessageConfig {
	message := m.Menuable.GetReplyMessage()
	message.ChatID = m.Message.Chat.ID
	return message
}
