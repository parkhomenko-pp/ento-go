package menus

import (
	"ento-go/src/entities"
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

func (m MenuMyGames) GetNavigation() map[string]string {
	return map[string]string{
		"< Back": MenuNameMain,
	}
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

func (m MenuMyGames) DoAction() {
	if m.Message.Text == "< Back" {
		m.Player.ChangeMenu(MenuNameMain)
		return
	}
}

func (m MenuMyGames) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(
		0,
		fmt.Sprintf("You have %d games", len(m.Games)),
	)
	return &message
}

func (m MenuMyGames) GetReplyMessage() *tgbotapi.MessageConfig {
	return m.GetFirstTimeMessage()
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
