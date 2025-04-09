package menus

import (
	"ento-go/src/models/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Menuable interface {
	GetName() string
	GetNavigation() []types.KeyboardButton
	DoAction()
	GetReplyText() string
	GetReplyImage() *tgbotapi.FileBytes
	GetOpponentMessage() *tgbotapi.MessageConfig
	IsConcatReply() bool
}
