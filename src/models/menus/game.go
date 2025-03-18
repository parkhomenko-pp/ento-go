package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

const MenuNameGame = "game"

type MenuGame struct {
	Message *tgbotapi.Message
	Player  *entities.Player

	GameId int
}

func (m MenuGame) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMyGames},
	}
}

func (m MenuGame) GetReplyText() string {
	return "game #" + strconv.Itoa(m.GameId)
}

func (m MenuGame) IsConcatReply() bool {
	return false
}

func (m MenuGame) GetName() string {
	return MenuNameGame
}

func (m MenuGame) DoAction() {

}

func (m MenuGame) GetOpponentMessage() *tgbotapi.MessageConfig {
	return nil
}
