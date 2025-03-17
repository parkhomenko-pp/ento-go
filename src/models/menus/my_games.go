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

	menu.Db.Where("player_chat_id = ? OR opponent_chat_id = ?", menu.Player.ChatID, menu.Player.ChatID).Find(&menu.Games)
	return &menu
}

func (m MenuMyGames) GetName() string {
	return MenuNameMyGames
}

func (m MenuMyGames) DoAction() {}

func (m MenuMyGames) GetReplyText() string {
	replyMessage := fmt.Sprintf("You have %d game(s)\n\n", len(m.Games))

	replyMessage = concatGamesByStatus(m.Games, entities.GameStatusWaitingForAccept, "Invites", replyMessage)
	replyMessage = concatGamesByStatus(m.Games, entities.GameStatusPlaying, "Playing", replyMessage)
	replyMessage = concatGamesByStatus(m.Games, entities.GameStatusFinished, "Finished", replyMessage)

	return replyMessage
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

func concatGamesByStatus(games []*entities.Game, status int8, label string, replyMessage string) string {
	filtered := []*entities.Game{}
	for _, game := range games {
		if game.Status == status {
			filtered = append(filtered, game)
		}
	}
	if len(filtered) > 0 {
		replyMessage += label + ":\n"
		for _, game := range filtered {
			replyMessage += fmt.Sprintf("/g_%d - \n", game.ID) // TODO: добавить никнеймы игроков
		}
		replyMessage += "\n"
	}
	return replyMessage
}
