package menus

import (
	"ento-go/src/entities"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const MenuMain = "main"

type Main struct {
	Message *tgbotapi.Message
	Player  *entities.Player
}

func (m *Main) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, fmt.Sprintf("Hello, %s! This is the main menu.", m.Player.Nickname))
	return &message
}

func (m *Main) GetName() string {
	return MenuMain
}

func (m *Main) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(0, fmt.Sprintf("Hello, %s! This is the main menu.", m.Player.Nickname))
	return &message
}

func (m *Main) DoAction() {
	//TODO implement me
}
