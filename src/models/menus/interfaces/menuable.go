package interfaces

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Menuable interface {
	GetName() string
	GetNavigation() map[string]string
	DoAction()
	GetReplyMessage() *tgbotapi.MessageConfig
	GetOpponentMessage() *tgbotapi.MessageConfig
}
