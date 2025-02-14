package menus

import (
	"ento-go/src/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const MenuNameNewGame = "new_game"

type MenuNewGame struct {
	Message *tgbotapi.Message
	Player  *entities.Player

	Opponent *entities.Player
}

func (m *MenuNewGame) GetName() string {
	return MenuNameNewGame
}

func (m *MenuNewGame) DoAction() {
	if m.Message.Text == "Cancel" {
		m.Player.ChangeMenu(MenuNameMain)
		return
	}
}

func (m *MenuNewGame) GetFirstTimeMessage() *tgbotapi.MessageConfig {
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

func (m *MenuNewGame) GetReplyMessage() *tgbotapi.MessageConfig {
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

func (m *MenuNewGame) CheckReply() bool {
	// Cancel -> true

	// проверить что это контакт или юзернейм

	// найти пользователя

	// если найден, то сохранить в поле Opponent

	return false
}
