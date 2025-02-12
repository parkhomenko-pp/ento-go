package menus

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type NotFound struct {
}

func (n *NotFound) CheckReply() bool {
	return true
}

func (n *NotFound) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Menu not found1")
	return &message
}

func (n *NotFound) GetName() string {
	return "not_found"
}

func (n *NotFound) DoAction() {
	// TODO:
	// 	1. change menu to main or registration
	// 	2. send message to user from main or registration
}

func (n *NotFound) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Menu not found2")
	return &message
}
