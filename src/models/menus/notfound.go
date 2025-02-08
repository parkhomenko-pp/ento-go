package menus

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type NotFound struct {
}

func (n *NotFound) DoAction() {
	// TODO:
	// 	1. change menu to main or registration
	// 	2. send message to user from main or registration
}

func (n *NotFound) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Menu not found")
	return &message
}
