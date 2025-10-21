package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

const MenuNameNewGameSize = "new_game_size"

type MenuNewGameSize struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	ReplyText string
}

func NewMenuNewGameSize(message *tgbotapi.Message, player *entities.Player, db *gorm.DB) *MenuNewGameSize {
	return &MenuNewGameSize{
		Message: message,
		Player:  player,
		Db:      db,
		ReplyText: "Select the board size for the new game:\n" +
			"7x7 - /size_7\n" +
			"9x9 - /size_9\n" +
			"13x13 - /size_13\n" +
			"19x19 - /size_19",
	}
}

func (m MenuNewGameSize) GetName() string {
	return MenuNameNewGameSize
}

func (m MenuNewGameSize) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMain},
	}
}

func (m MenuNewGameSize) DoAction() {
	switch m.Message.Text {
	case "/size_7":
		m.Player.ChangeMenuWithAdditional(MenuNameNewGame, "7")
		m.Db.Save(m.Player)
	case "/size_9":
		m.Player.ChangeMenuWithAdditional(MenuNameNewGame, "9")
		m.Db.Save(m.Player)
	case "/size_13":
		m.Player.ChangeMenuWithAdditional(MenuNameNewGame, "13")
		m.Db.Save(m.Player)
	case "/size_19":
		m.Player.ChangeMenuWithAdditional(MenuNameNewGame, "19")
		m.Db.Save(m.Player)
	default:
		m.ReplyText = "Please select a valid board size:\n" +
			"7x7 - /size_7\n" +
			"9x9 - /size_9\n" +
			"13x13 - /size_13\n" +
			"19x19 - /size_19"
	}
}

func (m MenuNewGameSize) GetReplyText() string {
	return m.ReplyText
}

func (m MenuNewGameSize) GetReplyImage() *tgbotapi.FileBytes {
	return nil
}

func (m MenuNewGameSize) GetOpponentMessage() tgbotapi.Chattable {
	return nil
}

func (m MenuNewGameSize) IsConcatReply() bool {
	return false
}
