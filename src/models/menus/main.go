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

func (m *Main) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Main")
	return &message
}

func (m *Main) DoAction() {
	//TODO implement me
}
