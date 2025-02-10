package menus

import (
	"ento-go/src/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const MenuMain = "main"

type Main struct {
	Message *tgbotapi.Message
	Player  *entities.Player
}

func (m *Main) ChangeLastMenu() {
	//TODO implement me
}

func (m *Main) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Main1")
	return &message
}

func (m *Main) GetName() string {
	return MenuMain
}

func (m *Main) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Main2")
	return &message
}

func (m *Main) DoAction() {
	//TODO implement me
}
