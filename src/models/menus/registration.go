package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const MenuNameRegistration = "registration"

type MenuRegistration struct {
	Message *tgbotapi.Message
	Player  *entities.Player

	ReplyMessage string
}

func (m *MenuRegistration) GetNavigation() []types.KeyboardButton {
	return nil
}

func (m *MenuRegistration) IsConcatReply() bool {
	return false
}

func (m *MenuRegistration) GetName() string {
	return MenuNameRegistration
}

func (m *MenuRegistration) CheckReply() bool {
	if m.Message.Text == "/menu" {
		return false
	}
	return true
}

func (m *MenuRegistration) DoAction() {
	if m.Player.LastMenu == "" {
		m.Player.LastMenu = MenuNameRegistration
		m.ReplyMessage = "Hello! Please, enter your nickname. It will be shown to other players"
		return
	}

	if m.Message.Text == "" {
		m.ReplyMessage = "Please, enter your nickname."
		return
	}
	if len([]rune(m.Message.Text)) < 2 {
		m.ReplyMessage = "Nickname must be 2 characters or more."
		return
	}
	if len([]rune(m.Message.Text)) > 20 {
		m.ReplyMessage = "Nickname must be 20 characters or less."
		return
	}
	if strings.HasPrefix(m.Message.Text, "/") {
		m.ReplyMessage = "Nickname can't start with '/'"
		return
	}
	if strings.Contains(m.Message.Text, " ") {
		m.ReplyMessage = "Nickname can't contain spaces."
		return
	}

	m.Player.Nickname = m.Message.Text
	m.Player.LastMenu = MenuNameMain
}

func (m *MenuRegistration) GetReplyText() string {
	message := ""

	if m.ReplyMessage == "" {
		message = "Hello! Please, enter your nickname. It will be shown to other players."
	} else {
		message = m.ReplyMessage
	}

	return message
}

func (m *MenuRegistration) GetOpponentMessage() *tgbotapi.MessageConfig {
	return nil
}
