package models

import (
	"ento-go/src/entities"
	"ento-go/src/models/menus"
	"ento-go/src/models/menus/interfaces"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

type Menu struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	interfaces.Menuable
	returnText      string
	opponentMessage *tgbotapi.MessageConfig
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
		m.Menuable = &menus.MenuNewGame{Message: m.Message, Player: m.Player, Db: m.Db}
	case menus.MenuNameMyGames:
		m.Menuable = menus.NewMenuMyGames(m.Message, m.Player, m.Db)
	default:
		if m.Player.Nickname == "" {
			m.Menuable = &menus.MenuRegistration{Message: m.Message, Player: m.Player}
		} else {
			m.Menuable = &menus.MenuNotFound{Message: m.Message, Player: m.Player}
		}
	}
}

func (m *Menu) DoAction() {
	//TODO: навигация
	if !m.NavigateToMenu() {
		m.Menuable.DoAction()
	}

	m.opponentMessage = m.Menuable.GetOpponentMessage()

	// если меню изменилось, то отправить первое сообщение из следующего меню
	if m.Player.LastMenu != m.Menuable.GetName() {
		// TODO: NEW GAME return message to player
		m.InitMenu()
		m.returnText = m.Menuable.GetReplyText()
	}
}

func (m *Menu) GetMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(m.Message.Chat.ID, "")

	// fill message text
	if m.returnText == "" { // проверка вдруг это 1 сообщение @see DoAction
		message.Text = m.returnText
	} else {
		message.Text = m.Menuable.GetReplyText()
	}

	// get navigation buttons
	navigationButtons := []tgbotapi.KeyboardButton{}
	navigation := m.Menuable.GetNavigation()
	if len(navigation) > 0 {
		for buttonText := range navigation {
			navigationButtons = append(navigationButtons, tgbotapi.NewKeyboardButton(buttonText))
		}
		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(navigationButtons...),
		)
	} else {
		message.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	// set chat id
	message.ChatID = m.Message.Chat.ID
	return &message
}

func (m *Menu) GetOpponentMessage() *tgbotapi.MessageConfig {

	return m.opponentMessage
}

func (m *Menu) NavigateToMenu() bool {
	navigation := m.Menuable.GetNavigation()
	if nextMenu, exists := navigation[m.Message.Text]; exists {
		m.Player.ChangeMenu(nextMenu)
		return true
	}
	return false
}
