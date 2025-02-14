package menus

import (
	"ento-go/src/entities"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const MenuNameMain = "main"

type MenuMain struct {
	Message *tgbotapi.Message
	Player  *entities.Player
}

func (m *MenuMain) CheckReply() bool {
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

func (m *MenuMain) GetFirstTimeMessage() *tgbotapi.MessageConfig {
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

func (m *MenuMain) GetName() string {
	return MenuNameMain
}

func (m *MenuMain) GetReplyMessage() *tgbotapi.MessageConfig {
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

func (m *MenuMain) DoAction() {
	if m.Message.Text == "New game" {
		m.Player.LastMenu = MenuNameNewGame
	}
}
