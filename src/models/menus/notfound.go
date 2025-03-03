package menus

import (
	"ento-go/src/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MenuNotFound struct {
	Message *tgbotapi.Message
	Player  *entities.Player
}

func (m *MenuNotFound) GetNavigation() map[string]string {
	//TODO implement me
	panic("implement me")
}

func (m *MenuNotFound) CheckReply() bool {
	return true
}

func (m *MenuNotFound) GetName() string {
	return "not_found"
}

func (m *MenuNotFound) DoAction() {
	if m.Player.Nickname == "" {
		m.Player.ChangeMenu(MenuNameRegistration)
	} else {
		m.Player.ChangeMenu(MenuNameMain)
	}
}

func (m *MenuNotFound) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Menu not found")
	return &message
}

func (m *MenuNotFound) GetOpponentMessage() *tgbotapi.MessageConfig {
	return nil
}
