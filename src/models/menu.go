package models

import (
	"ento-go/src/entities"
	"ento-go/src/models/menus"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Menu struct {
	Message *tgbotapi.Message
	Player  *entities.Player

	menus.Menuable
	returnMessage *tgbotapi.MessageConfig
}

func (m *Menu) String() string {
	nickname := ""
	if m.Player.Nickname != "" {
		nickname = m.Player.Nickname
	} else {
		nickname = "[Anonymous]"
	}

	return fmt.Sprintf(
		"Player: %v\t Menu: %v",
		nickname,
		m.Menuable.GetName(),
	)
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
	// если это первый раз, то отправить первое сообщение меню
	if m.Player.IsMenuVisited == false {
		m.Player.IsMenuVisited = true
		m.returnMessage = m.Menuable.GetFirstTimeMessage()
		return
	}

	m.Menuable.DoAction()

	// если меню изменилось, то отправить первое сообщение из следующего меню
	if m.Player.LastMenu != m.Menuable.GetName() {
		m.InitMenu()
		m.returnMessage = m.Menuable.GetFirstTimeMessage()
		return
	}
}

func (m *Menu) GetMessage() *tgbotapi.MessageConfig {
	var message *tgbotapi.MessageConfig

	if m.returnMessage != nil {
		message = m.returnMessage
	} else {
		message = m.Menuable.GetReplyMessage()
	}

	message.ChatID = m.Message.Chat.ID
	return message
}
