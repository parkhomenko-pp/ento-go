package menus

import (
	"ento-go/src/common"
	"ento-go/src/entities"
	"ento-go/src/models"
	"ento-go/src/models/types"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strconv"
)

const MenuNameGame = "game"

type MenuGame struct {
	Message              *tgbotapi.Message
	Player               *entities.Player
	Db                   *gorm.DB
	Game                 *entities.Game
	ReplyText            string
	OpponentReplyMessage tgbotapi.Chattable
	goban                *models.Goban
	replyImage           *tgbotapi.FileBytes
}

// Interface/Basic Info
func (m *MenuGame) GetName() string { return MenuNameGame }

func (m *MenuGame) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMyGames},
		{Text: "Surrender"},
		{Text: "Pass"},
		{Text: "Help"},
	}
}

// Factory/Constructor
func NemMenuGame(message *tgbotapi.Message, player *entities.Player, db *gorm.DB, additional string) *MenuGame {
	gameId, _ := strconv.Atoi(additional)
	menu := &MenuGame{Message: message, Player: player, Db: db}
	db.Preload("Opponent").Preload("Player").Where("id = ?", gameId).First(&menu.Game)
	menu.goban = models.NewGobanBySize(menu.Game.Size)
	menu.goban.SetDots(menu.Game.GetDots())
	menu.goban.SetLast(menu.Game.LastStonePosition)
	menu.goban.ChangeTheme(models.CreateGobanThemeById(player.ThemeId))
	return menu
}

// Reply Helpers
func (m *MenuGame) GetReplyText() string { return m.ReplyText }

func (m *MenuGame) GetReplyImage() *tgbotapi.FileBytes {
	if m.replyImage == nil {
		img := m.getImageForGoban(*m.goban)
		m.replyImage = &img
	}
	return m.replyImage
}

func (m *MenuGame) getImageForGoban(goban models.Goban) tgbotapi.FileBytes {
	img := goban.GetImage()
	byteImage, _ := common.EncodeImageToPNGBytes(*img)
	return tgbotapi.FileBytes{Name: "goban.png", Bytes: byteImage}
}

func (m *MenuGame) IsConcatReply() bool { return false }

func (m *MenuGame) GetOpponentMessage() tgbotapi.Chattable { return m.OpponentReplyMessage }

// Game Logic
func (m *MenuGame) DoAction() {
	// ... (same as before)
}

func (m *MenuGame) placeStone(runeRow rune, intColumn uint8) error {
	if m.isPlaceBlack() {
		return m.goban.PlaceBlack(runeRow, intColumn)
	}
	return m.goban.PlaceWhite(runeRow, intColumn)
}

func (m *MenuGame) validateMove() (rune, uint8, error) {
	if m.Message.Text == "" {
		return 0, 0, fmt.Errorf("message is empty")
	}
	runes := []rune(m.Message.Text)
	if len(runes) < 2 {
		return 0, 0, fmt.Errorf("invalid move format")
	}
	intColumn, err := strconv.Atoi(m.Message.Text[1:])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid column")
	}
	return runes[0], uint8(intColumn), nil
}

// Turn/Player Helpers
func (m *MenuGame) isNotMyTurn() bool {
	isPlayer := m.Game.PlayerChatID == m.Message.Chat.ID
	return (isPlayer && m.Game.IsPlayerTurn) || (!isPlayer && !m.Game.IsPlayerTurn)
}

func (m *MenuGame) isPlaceBlack() bool {
	isPlayer := m.Game.PlayerChatID == m.Message.Chat.ID
	return (isPlayer && m.Game.IsPlayerBlack) || (!isPlayer && !m.Game.IsPlayerBlack)
}

func (m *MenuGame) getRealOpponent() *entities.Player {
	if m.Game.PlayerChatID == m.Message.Chat.ID {
		return &m.Game.Opponent
	}
	return &m.Game.Player
}

func (m *MenuGame) getPlacedDotEmoji(inverted bool) string {
	if m.isPlaceBlack() != inverted {
		return "⚫"
	}
	return "⚪"
}
