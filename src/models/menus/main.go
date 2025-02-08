package menus

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const MenuMain = "main"

type Main struct {
}

func (m *Main) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Main")
	return &message
}

func (m *Main) DoAction() {
	//TODO implement me
}
