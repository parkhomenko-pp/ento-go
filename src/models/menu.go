package models

import (
	"ento-go/src/entities"
	"ento-go/src/menus"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strings"
)

type Menu struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	menus.Menuable

	replyText            string
	replyOpponentMessage *tgbotapi.MessageConfig
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
	lastMenu := ""
	additional := ""

	if strings.Contains(m.Player.LastMenu, ":") {
		splitted := strings.Split(m.Player.LastMenu, ":")
		lastMenu = splitted[0]
		additional = splitted[1]
	} else {
		lastMenu = m.Player.LastMenu
	}

	switch lastMenu {
	case menus.MenuNameRegistration:
		m.Menuable = &menus.MenuRegistration{Message: m.Message, Player: m.Player, Db: m.Db}
	case menus.MenuNameMain:
		m.Menuable = &menus.MenuMain{Message: m.Message, Player: m.Player}
	case menus.MenuNameNewGame:
		m.Menuable = &menus.MenuNewGame{Message: m.Message, Player: m.Player, Db: m.Db}
	case menus.MenuNameMyGames:
		m.Menuable = menus.NewMenuMyGames(m.Message, m.Player, m.Db)
	case menus.MenuNameGame:
		m.Menuable = menus.NemMenuGame(m.Message, m.Player, m.Db, additional)
	default:
		if m.Player.Nickname == "" {
			m.Menuable = &menus.MenuRegistration{Message: m.Message, Player: m.Player}
		} else {
			m.Menuable = &menus.MenuNotFound{Message: m.Message, Player: m.Player}
		}
	}
}

func (m *Menu) DoAction() {
	if !m.NavigateToMenu() {
		m.Menuable.DoAction()
	}

	m.replyOpponentMessage = m.Menuable.GetOpponentMessage()

	// если меню изменилось, то отправить первое сообщение из следующего меню
	if m.Player.LastMenu != m.Menuable.GetName() {
		oldMenuConcat := m.Menuable.IsConcatReply()
		oldMessageText := m.Menuable.GetReplyText()
		m.InitMenu()
		m.replyText = m.Menuable.GetReplyText()
		if oldMenuConcat {
			m.replyText = oldMessageText + "\n\n----\n" + m.replyText
		}
	}
}

func (m *Menu) GetReplyMessage() tgbotapi.Chattable {
	message := tgbotapi.NewMessage(m.Message.Chat.ID, "")

	// fill message text
	if m.replyText != "" { // проверка вдруг это 1 сообщение @see DoAction
		message.Text = m.replyText
	} else {
		message.Text = m.Menuable.GetReplyText()
	}

	// get navigation buttons
	navigationButtons := []tgbotapi.KeyboardButton{}
	navigation := m.Menuable.GetNavigation()
	if len(navigation) > 0 {
		for _, keyboardButton := range navigation {
			navigationButtons = append(navigationButtons, tgbotapi.NewKeyboardButton(keyboardButton.Text))
		}
		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(navigationButtons...),
		)
	} else {
		message.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	// set chat id
	message.ChatID = m.Message.Chat.ID

	// check for reply image
	replyImage := m.Menuable.GetReplyImage()
	if replyImage != nil {
		photoMessage := tgbotapi.NewPhoto(m.Message.Chat.ID, replyImage)
		photoMessage.Caption = message.Text
		photoMessage.ReplyMarkup = message.ReplyMarkup
		return &photoMessage
	}

	return message
}

func (m *Menu) GetOpponentMessage() *tgbotapi.MessageConfig {
	return m.replyOpponentMessage
}

func (m *Menu) NavigateToMenu() bool {
	navigation := m.Menuable.GetNavigation()
	for _, button := range navigation {
		if button.Text == m.Message.Text && button.Destination != "" {
			m.Player.ChangeMenu(button.Destination)
			return true
		}
	}
	return false
}
