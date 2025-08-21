package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const MenuNameMain = "main"

type MenuMain struct {
	Message *tgbotapi.Message
	Player  *entities.Player
}

func (m *MenuMain) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "New game", Destination: MenuNameNewGame},
		{Text: "My games", Destination: MenuNameMyGames},
	}
}

func (m *MenuMain) IsConcatReply() bool {
	return false
}

func (m *MenuMain) GetName() string {
	return MenuNameMain
}

func (m *MenuMain) GetReplyText() string {
	return fmt.Sprintf(
		"%s, this is the main menu.\n\nGames played: %d\nWins: %d (Win rate: %.2f%%)",
		m.Player.Nickname,
		m.Player.GamesCount,
		m.Player.WinsCount,
		m.Player.GetWinRate(),
	)
}

func (m *MenuMain) DoAction() {

}

func (m *MenuMain) GetOpponentMessage() tgbotapi.Chattable {
	return nil
}

func (m *MenuMain) GetReplyImage() *tgbotapi.FileBytes {
	return nil
}
