package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models"
	"ento-go/src/models/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strconv"
)

const MenuNameGameFinished = "game-finished"

type MenuGameFinished struct {
	Message   *tgbotapi.Message
	Player    *entities.Player
	Db        *gorm.DB
	Game      *entities.Game
	ReplyText string
}

func (m *MenuGameFinished) GetName() string {
	return MenuNameGameFinished
}

func NemMenuGameFinished(message *tgbotapi.Message, player *entities.Player, db *gorm.DB, additional string) *MenuGameFinished {
	gameId, _ := strconv.Atoi(additional)
	menu := &MenuGameFinished{Message: message, Player: player, Db: db}
	db.Preload("Opponent").Preload("Player").Where("id = ?", gameId).First(&menu.Game)

	realOponent := menu.getRealOpponent()
	menu.ReplyText = "Finished game with " + realOponent.Nickname + "\n" +
		"⚫️Stone - " + menu.Game.Player.Nickname + "\n" +
		"⚪️Stone - " + menu.Game.Opponent.Nickname + "\n\n" +
		"/replay - Replay with color swap\n" +
		"/replay_no_swap - Replay with same stones\n"

	return menu
}

func (m *MenuGameFinished) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMyGames},
	}
}

func (m *MenuGameFinished) DoAction() {
	switch m.Message.Text {
	case "/replay":
		m.Game.Status = entities.GameStatusPlaying
		m.Game.IsPlayerBlack = !m.Game.IsPlayerBlack
		m.ReplyText = "Game restarted with color swap"
	case "/replay_no_swap":
		m.Game.Status = entities.GameStatusPlaying
		m.ReplyText = "Game restarted without color swap"
	default:
		m.ReplyText = "Unknown command"
		return
	}

	m.Game.PlayerCaptureDotsCount = 0
	m.Game.OpponentCaptureDotsCount = 0
	m.Game.SetDots(models.NewGoban7().GetDots())
	m.Game.LastStonePosition = ""
	m.Db.Save(&m.Game)

	m.Player.ChangeMenuWithAdditional(MenuNameGame, strconv.Itoa(int(m.Game.ID)))
	m.Db.Save(&m.Player)
}

func (m *MenuGameFinished) GetReplyText() string {
	return m.ReplyText
}

func (m *MenuGameFinished) GetReplyImage() *tgbotapi.FileBytes {
	return nil
}

func (m *MenuGameFinished) GetOpponentMessage() tgbotapi.Chattable {
	return nil
}

func (m *MenuGameFinished) IsConcatReply() bool {
	return false
}

func (m *MenuGameFinished) getRealOpponent() *entities.Player {
	if m.Game.PlayerChatID == m.Message.Chat.ID {
		return &m.Game.Opponent
	}
	return &m.Game.Player
}
