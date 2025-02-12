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

func (m *Main) CheckReply() bool {
	switch m.Message.Text {
	case "New game":
		return true
	case "My games":
		return true
	case "Info":
		return true
	default:
		return false
	}
}

func (m *Main) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(
		0,
		fmt.Sprintf(
			"Hello, %s! This is the main menu.\nGames played: %d\nWins: %d (Win rate: %.2f%%)",
			m.Player.Nickname,
			m.Player.GamesCount,
			m.Player.WinsCount,
			m.Player.GetWinRate(),
		),
	)
	message.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("New game"),
			tgbotapi.NewKeyboardButton("My games"),
			tgbotapi.NewKeyboardButton("Info"),
		),
	)
	return &message
}

func (m *Main) GetName() string {
	return MenuMain
}

func (m *Main) GetReplyMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(
		0,
		fmt.Sprintf(
			"%s, this is the main menu.\n\nGames played: %d\nWins: %d (Win rate: %.2f%%)",
			m.Player.Nickname,
			m.Player.GamesCount,
			m.Player.WinsCount,
			m.Player.GetWinRate(),
		),
	)
	message.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("New game"),
			tgbotapi.NewKeyboardButton("My games"),
			tgbotapi.NewKeyboardButton("Info"),
		),
	)
	return &message
}

func (m *Main) DoAction() {
	//TODO implement me
}
