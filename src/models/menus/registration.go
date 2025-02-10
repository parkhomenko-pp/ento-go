package menus

import (
	"ento-go/src/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const MenuRegistration = "registration"

type Registration struct {
	Message *tgbotapi.Message
	Player  *entities.Player

	ReplyMessage string
	NextMenu     string
}

func (r *Registration) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, "Hello! Please, enter your nickname. It will be shown to other players.")
	return &message
}

func (r *Registration) GetName() string {
	return MenuRegistration
}

func (r *Registration) DoAction() {
	if r.Message.Text == "" {
		r.ReplyMessage = "Please, enter your nickname."
		return
	}
	if len([]rune(r.Message.Text)) < 2 {
		r.ReplyMessage = "Nickname must be 2 characters or more."
		return
	}
	if len([]rune(r.Message.Text)) > 20 {
		r.ReplyMessage = "Nickname must be 20 characters or less."
		return
	}
	if strings.HasPrefix(r.Message.Text, "/") {
		r.ReplyMessage = "Nickname can't start with '/'"
	}
	if strings.Contains(r.Message.Text, " ") {
		r.ReplyMessage = "Nickname can't contain spaces."
		return
	}

	r.Player.Nickname = r.Message.Text
	r.NextMenu = MenuMain
}

func (r *Registration) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, r.ReplyMessage)
	return &message
}

func (r *Registration) ChangeLastMenu() {
	if r.NextMenu != "" {
		r.Player.LastMenu = r.NextMenu
		r.Player.IsMenuVisited = false
	}
}
