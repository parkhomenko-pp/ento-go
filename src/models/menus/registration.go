package menus

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const MenuRegistration = "registration"

type Registration struct {
}

func (r *Registration) DoAction() {
	//TODO implement me
	panic("implement me")
}

func (r *Registration) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Registration")
	return &message
}
