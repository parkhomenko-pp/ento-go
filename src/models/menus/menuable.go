package menus

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Menuable interface {
	GetName() string
	DoAction()
	GetFirstTimeMessage() *tgbotapi.MessageConfig
	GetReplyMessage() *tgbotapi.MessageConfig
	CheckReply() bool
}
