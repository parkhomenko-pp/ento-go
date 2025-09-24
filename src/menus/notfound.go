package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MenuNotFound struct {
	Message *tgbotapi.Message
	Player  *entities.Player
}

func (m *MenuNotFound) GetNavigation() []types.KeyboardButton {
	return nil
}

func (m *MenuNotFound) IsConcatReply() bool {
	return false
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

func (m *MenuNotFound) GetReplyText() string {
	return "Menu not found"
}

func (m *MenuNotFound) GetOpponentMessage() tgbotapi.Chattable {
	return nil
}

func (m *MenuNotFound) GetReplyImage() *tgbotapi.FileBytes {
	return nil
}
