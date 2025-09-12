package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

const MenuNameDeclined = "declined"

type MenuDeclined struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	Games []*entities.Game

	ReplyText string
}

func (m *MenuDeclined) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMyGames},
	}
}

func (m *MenuDeclined) IsConcatReply() bool {
	return false
}

func NewMenuDeclined(message *tgbotapi.Message, player *entities.Player, db *gorm.DB) *MenuDeclined {
	menu := MenuDeclined{
		Message: message,
		Player:  player,
		Db:      db,
	}

	menu.Db.
		Preload("Player").
		Preload("Opponent").
		Where("status = ? AND (player_chat_id = ? OR opponent_chat_id = ?)", entities.GameStatusDeclined, menu.Player.ChatID, menu.Player.ChatID).
		Find(&menu.Games)

	return &menu
}

func (m *MenuDeclined) GetName() string {
	return MenuNameDeclined
}

func (m *MenuDeclined) DoAction() {
	gameIDStr := strings.TrimPrefix(m.Message.Text, "/g_")
	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		m.ReplyText = "Invalid game ID"
		return
	}

	m.Player.ChangeMenuWithAdditional(MenuNameInvited, strconv.Itoa(gameID))

	m.ReplyText = "Game not found"
}

func (m *MenuDeclined) GetReplyText() string {
	if m.ReplyText != "" {
		return m.ReplyText
	}

	replyMessage := fmt.Sprintf("You have %d declined game(s)\n\n", len(m.Games))

	replyMessage = m.concatGames(replyMessage)

	return replyMessage
}

func (m *MenuDeclined) CheckReply() bool {
	if strings.HasPrefix(m.Message.Text, "/g_") {
		return true
	}

	return false
}

func (m *MenuDeclined) GetOpponentMessage() tgbotapi.Chattable {
	return nil
}

func (m *MenuDeclined) concatGames(replyMessage string) string {
	if len(m.Games) > 0 {
		for _, game := range m.Games {
			replyMessage += fmt.Sprintf(
				"/g_%d - %s\n",
				game.ID,
				game.GetOpponentChatIdForPlayer(m.Player).Nickname,
			)
		}
		replyMessage += "\n"
	}
	return replyMessage
}

func (m *MenuDeclined) GetReplyImage() *tgbotapi.FileBytes {
	return nil
}
