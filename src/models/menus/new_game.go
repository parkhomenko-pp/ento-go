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

	ReplyMessage    string
	OpponentMessage *tgbotapi.MessageConfig
}

func (m *MenuNewGame) GetName() string {
	return MenuNameNewGame
}

func (m *MenuNewGame) DoAction() {
	if m.Message.Text == "Cancel" {
		m.Player.ChangeMenu(MenuNameMain)
		return
	}

	if m.Message.Text == m.Player.Nickname {
		m.ReplyMessage = "You can't play with yourself"
		return
	}

	var opponent *entities.Player
	result := m.Db.First(&opponent, "nickname = ?", m.Message.Text)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		m.ReplyMessage = "User not found"
		return
	}

	var game *entities.Game
	result = m.Db.First(&game, "player_chat_id = ? AND opponent_chat_id = ?", m.Player.ChatID, opponent.ChatID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		game = &entities.Game{
			PlayerChatId:   m.Player.ChatID,
			OpponentChatId: opponent.ChatID,
			Status:         entities.GameStatusWaitingForAccept,
		}
		m.Db.Create(&game)
	}

	// TODO: message to opponent
	m.OpponentMessage = tgbotapi.NewMessage( // TODO: fix
		opponent.ChatID,
		"User "+m.Player.Nickname+" invited you to the game",
	)

	// TODO: change menu to waiting for accept
	m.Player.ChangeMenu(MenuNameMain)

}

func (m *MenuNewGame) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(
		0,
		"Please, send me Nickname of your opponent to invite him to the game",
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
	return true
}

func (m *MenuNewGame) GetOpponentMessage() *tgbotapi.MessageConfig {
	return m.OpponentMessage
}
