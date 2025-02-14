package menus

import (
	"ento-go/src/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const MenuNewGame = "new_game"

type NewGame struct {
	Message *tgbotapi.Message
	Player  *entities.Player

	Opponent *entities.Player
}

func (n *NewGame) GetName() string {
	return MenuNewGame
}

func (n *NewGame) DoAction() {
	if n.Message.Text == "Cancel" {
		n.Player.ChangeMenu(MenuMain)
		return
	}
}

func (n *NewGame) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(
		0,
		"Please, send me username or contact of your opponent to invite him to the game",
	)
	message.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Cancel"),
		),
	)

	return &message
}

func (n *NewGame) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(
		0,
		"new game reply message",
	)
	message.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Cancel"),
		),
	)

	return &message
}

func (n *NewGame) CheckReply() bool {
	// Cancel -> true

	// проверить что это контакт или юзернейм

	// найти пользователя

	// если найден, то сохранить в поле Opponent

	return false
}
