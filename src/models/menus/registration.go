package menus

import (
	"ento-go/src/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const MenuNameRegistration = "registration"

type MenuRegistration struct {
	Message *tgbotapi.Message
	Player  *entities.Player

	ReplyMessage string
}

func (m *MenuRegistration) GetNavigation() map[string]string {
	return nil
}

func (m *MenuRegistration) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Hello! Please, enter your nickname. It will be shown to other players.")
	return &message
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

func (m *MenuRegistration) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, m.ReplyMessage)
	return &message
}

func (m *MenuRegistration) GetOpponentMessage() *tgbotapi.MessageConfig {
	return nil
}
