package models

import (
	"ento-go/src/entities"
	"ento-go/src/models/menus"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Menu struct {
	Message *tgbotapi.Message
	Player  *entities.Player

	menus.Menuable
}

func (m *Menu) InitMenu() {
	switch m.Player.LastMenu {
	case menus.MenuRegistration:
		m.Menuable = &menus.Registration{Message: m.Message, Player: m.Player}
	case menus.MenuMain:
		m.Menuable = &menus.Main{Message: m.Message, Player: m.Player}
	default:
		if m.Player.Nickname == "" {
			m.Menuable = &menus.Registration{Message: m.Message, Player: m.Player}
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
