package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strconv"
)

const MenuNameGame = "game"

type MenuGame struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	Game      *entities.Game
	ReplyText string
}

func (m *MenuGame) GetName() string {
	return MenuNameGame
}

func (m *MenuGame) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMyGames},
		{Text: "Help"}, // TODO: add help message
	}
}

func NemMenuGame(message *tgbotapi.Message, player *entities.Player, db *gorm.DB, additional string) *MenuGame {
	gameId := 0
	gameId, _ = strconv.Atoi(additional)

	menu := MenuGame{
		Message: message,
		Player:  player,
		Db:      db,
	}

	menu.Db.
		Where("id = ?", gameId).
		First(&menu.Game)

	return &menu
}

func (m *MenuGame) GetReplyText() string {
	return "game #" + strconv.Itoa(int(m.Game.ID))
}

func (m *MenuGame) GetReplyImage() *tgbotapi.FileBytes {
	return nil
}

func (m *MenuGame) IsConcatReply() bool {
	return false
}

func (m *MenuGame) DoAction() {
	if m.Message.Text == "Help" {
		m.ReplyText = "Help message"
		return
	}
}

func (m *MenuGame) GetOpponentMessage() *tgbotapi.MessageConfig {
	return nil
}
