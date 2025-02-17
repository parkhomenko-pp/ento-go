package menus

import (
	"ento-go/src/entities"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

const MenuNameNewGame = "new_game"

type MenuNewGame struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	Opponent     *entities.Player
	ReplyMessage string
}

func (m *MenuNewGame) GetName() string {
	return MenuNameNewGame
}

func (m *MenuNewGame) DoAction() {
	if m.Message.Text == "Cancel" {
		m.Player.ChangeMenu(MenuNameMain)
		return
	}

	if m.Opponent == nil {

	}
}

func (m *MenuNewGame) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(
		0,
		"Please, send me Nickneme of your opponent to invite him to the game",
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
		m.ReplyMessage,
	)
	message.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Cancel"),
		),
	)

	return &message
}

func (m *MenuNewGame) CheckReply() bool {
	if m.Message.Text == "Cancel" {
		return true
	}

	if m.Message.Text == m.Player.Nickname {
		m.ReplyMessage = "You can't play with yourself"
		return false
	}

	var opponent *entities.Player

	result := m.Db.First(&opponent, "nickname = ?", m.Message.Text)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		m.Opponent = opponent
		return true
	} else {
		m.ReplyMessage = "User not found"
		return false
	}
}
