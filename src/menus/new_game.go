package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

const MenuNameNewGame = "new_game"

type MenuNewGame struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	ReplyText       string
	OpponentMessage *tgbotapi.MessageConfig

	concat bool
}

func (m *MenuNewGame) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMain},
	}
}

func (m *MenuNewGame) IsConcatReply() bool {
	return m.concat
}

func (m *MenuNewGame) GetName() string {
	return MenuNameNewGame
}

func (m *MenuNewGame) DoAction() {
	if m.Message.Text == m.Player.Nickname {
		m.ReplyText = "You can't play with yourself"
		return
	}

	var opponent *entities.Player
	result := m.Db.First(&opponent, "nickname = ?", m.Message.Text)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		m.ReplyText = "User with	this nickname not found"
		return
	}

	var game *entities.Game
	result = m.Db.First(
		&game,
		"player_chat_id = ? AND opponent_chat_id = ? OR player_chat_id = ? AND opponent_chat_id = ?",
		m.Player.ChatID, opponent.ChatID, opponent.ChatID, m.Player.ChatID,
	)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		game = &entities.Game{
			PlayerChatID:   m.Player.ChatID,
			OpponentChatID: opponent.ChatID,
			Status:         entities.GameStatusWaitingForAccept,
		}
		m.Db.Create(&game)

		newOpponentMessage := tgbotapi.NewMessage(
			opponent.ChatID,
			"User "+m.Player.Nickname+" invited you to a game",
		)
		m.OpponentMessage = &newOpponentMessage

		m.ReplyText = "Invitation sent"
	} else {
		m.ReplyText = "You already have game with this user"
		return
	}

	m.concat = true
	m.Player.ChangeMenu(MenuNameMain)
}

func (m *MenuNewGame) GetReplyText() string {
	if m.ReplyText == "" {
		return "Please, send me Nickname of your opponent to invite him to the game"
	}
	return m.ReplyText
}

func (m *MenuNewGame) CheckReply() bool {
	return true
}

func (m *MenuNewGame) GetOpponentMessage() tgbotapi.Chattable {
	return m.OpponentMessage
}

func (m *MenuNewGame) GetReplyImage() *tgbotapi.FileBytes {
	return nil
}
