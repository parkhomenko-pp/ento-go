package menus

import (
	"ento-go/src/entities"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strings"
)

type Menu struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	Menuable

	replyText            string
	replyOpponentMessage tgbotapi.Chattable
}

func (m *Menu) String() string {
	nickname := ""
	if m.Player.Nickname != "" {
		nickname = m.Player.Nickname
	} else {
		nickname = "[Anonymous]"
	}

	return fmt.Sprintf(
		"Player: %v, Menu: %v, Reply: %v",
		nickname,
		m.Menuable.GetName(),
		m.Menuable.GetReplyText(),
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
	case MenuNameRegistration:
		m.Menuable = &MenuRegistration{Message: m.Message, Player: m.Player, Db: m.Db}
	case MenuNameMain:
		m.Menuable = &MenuMain{Message: m.Message, Player: m.Player}
	case MenuNameNewGame:
		m.Menuable = &MenuNewGame{Message: m.Message, Player: m.Player, Db: m.Db}
	case MenuNameMyGames:
		m.Menuable = NewMenuMyGames(m.Message, m.Player, m.Db)
	case MenuNameGame:
		m.Menuable = NemMenuGame(m.Message, m.Player, m.Db, additional)
	case MenuNameSettings:
		m.Menuable = &MenuSettings{Message: m.Message, Player: m.Player, Db: m.Db}
	case MenuNameSettingsChangeTheme:
		m.Menuable = &MenuSettingsChangeTheme{Message: m.Message, Player: m.Player, Db: m.Db}
	default:
		if m.Player.Nickname == "" {
			m.Menuable = &MenuRegistration{Message: m.Message, Player: m.Player}
		} else {
			m.Menuable = &MenuNotFound{Message: m.Message, Player: m.Player}
		}
	}
}

func (m *Menu) isMenuChanged() bool {
	lastPlayerMenu := m.Player.LastMenu
	menuName := m.Menuable.GetName()

	if strings.Contains(lastPlayerMenu, ":") {
		lastPlayerMenu = strings.Split(lastPlayerMenu, ":")[0]
	}
	return lastPlayerMenu != menuName
}

func (m *Menu) DoAction() {
	if !m.NavigateToMenu() {
		m.Menuable.DoAction()
	}

	m.replyOpponentMessage = m.Menuable.GetOpponentMessage()

	// если меню изменилось, то отправить первое сообщение из следующего меню
	if m.isMenuChanged() {
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

func (m *Menu) GetOpponentMessage() tgbotapi.Chattable {
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
