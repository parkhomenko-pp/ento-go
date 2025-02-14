package menus

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type MenuNotFound struct {
}

func (m *MenuNotFound) CheckReply() bool {
	return true
}

func (m *MenuNotFound) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Menu not found1")
	return &message
}

func (m *MenuNotFound) GetName() string {
	return "not_found"
}

func (m *MenuNotFound) DoAction() {
	// TODO:
	// 	1. change menu to main or registration
	// 	2. send message to user from main or registration
}

func (m *MenuNotFound) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Menu not found2")
	return &message
}
