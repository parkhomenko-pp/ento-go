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
	validReplies := []string{"New game", "My games", "Info"}
	for _, reply := range validReplies {
		if m.Message.Text == reply {
			return true
		}
	}
	return false
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
		m.Player.ChangeMenu(MenuNameNewGame)
	}
}

func (m *MenuMain) GetOpponentMessage() *tgbotapi.MessageConfig {
	return nil
}
