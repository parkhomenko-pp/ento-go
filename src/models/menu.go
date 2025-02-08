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

func (m *Menu) InitMenu() {
	switch m.Player.LastMenu {
	case menus.MenuRegistration:
		m.Menuable = &menus.Registration{}
	case menus.MenuMain:
		m.Menuable = &menus.Main{}

	default:
		if m.Player.Nickname == "" {
			m.Menuable = &menus.Registration{}
		} else {
			m.Menuable = &menus.NotFound{}
		}
	}
}

func (m *Menu) DoAction() {
	// TODO
}

func (m *Menu) GetReplyMessage() *tgbotapi.MessageConfig {
	message := m.Menuable.GetReplyMessage()
	message.ChatID = m.Message.Chat.ID
	return message
}
