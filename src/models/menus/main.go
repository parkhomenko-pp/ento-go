package menus

import (
	"ento-go/src/entities"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const MenuNameMain = "main"

type MenuMain struct {
	Message *tgbotapi.Message
	Player  *entities.Player
}

func (m *MenuMain) GetName() string {
	return MenuNameMain
}

var menuMainNavigation = map[string]string{
	"New game": MenuNameNewGame,
	"My games": MenuNameMyGames,
}

func (m *MenuMain) GetNavigation() map[string]string {
	return menuMainNavigation
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

func (m *MenuMain) GetOpponentMessage() *tgbotapi.MessageConfig {
	return nil
}
