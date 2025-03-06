package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

const MenuNameMyGames = "my-games"

type MenuMyGames struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	Games []*entities.Game
}

func (m MenuMyGames) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMain},
	}
}

func (m MenuMyGames) IsConcatReply() bool {
	return false
}

func NewMenuMyGames(message *tgbotapi.Message, player *entities.Player, db *gorm.DB) *MenuMyGames {
	menu := MenuMyGames{
		Message: message,
		Player:  player,
		Db:      db,
	}

	if err := menu.Db.Where("player_id = ? OR opponent_id = ?", menu.Player.ChatID, menu.Player.ChatID).Find(&menu.Games).Error; err != nil {
		menu.Games = []*entities.Game{}
	}

	return &menu
}

func (m MenuMyGames) GetName() string {
	return MenuNameMyGames
}

func (m MenuMyGames) DoAction() {}

func (m MenuMyGames) GetReplyText() string {
	return fmt.Sprintf("You have %d games", len(m.Games))
}

func (m MenuMyGames) CheckReply() bool {
	if m.Message.Text == "< Back" {
		return true
	}

	return false
}

func (m MenuMyGames) GetOpponentMessage() *tgbotapi.MessageConfig {
	return nil
}
