package menus

import (
	"ento-go/src/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const MenuRegistration = "registration"

type Registration struct {
	Message *tgbotapi.Message
	Player  *entities.Player
}

func (r *Registration) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Registration1")
	return &message
}

func (r *Registration) GetName() string {
	return MenuRegistration
}

func (r *Registration) DoAction() {
	// TODO: implement
}

func (r *Registration) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Registration2")
	return &message
}
