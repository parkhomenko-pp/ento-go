package models

import (
	"ento-go/src/entities"
	"ento-go/src/models/menus"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

type Menu struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

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
	case menus.MenuNameRegistration:
		m.Menuable = &menus.MenuRegistration{Message: m.Message, Player: m.Player}
	case menus.MenuNameMain:
		m.Menuable = &menus.MenuMain{Message: m.Message, Player: m.Player}
	case menus.MenuNameNewGame:
		m.Menuable = &menus.MenuNewGame{Message: m.Message, Player: m.Player}
	default:
		if m.Player.Nickname == "" {
			m.Menuable = &menus.MenuRegistration{Message: m.Message, Player: m.Player}
		} else {
			m.Menuable = &menus.MenuNotFound{}
		}
	}
}

func (m *Menu) DoAction() {
	// если ответ не тот, который ожидается, то отправить сообщение об ошибке
	if m.Message.Text != "/menu" && !m.Menuable.CheckReply() {
		message := m.Menuable.GetFirstTimeMessage()
		message.Text = "Sorry, I don't understand you 😔\n\n" + message.Text

		m.returnMessage = message
		return
	}

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

	if m.returnMessage != nil { // проверка вдруг это 1 сообщение @see DoAction
		message = m.returnMessage
	} else {
		message = m.Menuable.GetReplyMessage()
	}

	message.ChatID = m.Message.Chat.ID
	return message
}
