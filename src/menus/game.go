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

func (m *MenuGame) GetName() string { return MenuNameGame }

func (m *MenuGame) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMyGames},
		{Text: "Surrender"},
		{Text: "Pass"},
		{Text: "Help"},
	}
}

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

func (m *MenuGame) GetReplyText() string { return m.ReplyText }

func (m *MenuGame) getImageForGoban(goban models.Goban) tgbotapi.FileBytes {
	img := goban.GetImage()
	byteImage, _ := common.EncodeImageToPNGBytes(*img)
	return tgbotapi.FileBytes{Name: "goban.png", Bytes: byteImage}
}

func (m *MenuGame) GetReplyImage() *tgbotapi.FileBytes {
	if m.replyImage == nil {
		img := m.getImageForGoban(*m.goban)
		m.replyImage = &img
	}
	return m.replyImage
}

func (m *MenuGame) IsConcatReply() bool { return false }

func (m *MenuGame) DoAction() {
	switch m.Message.Text {
	case "Help":
		m.ReplyText = "Help message"
		return
	case "Surrender":
		m.Game.Status = entities.GameStatusFinished
		m.Db.Save(&m.Game)
		m.ReplyText = "You surrendered. Game over."
		m.OpponentReplyMessage = tgbotapi.NewMessage(m.getRealOpponent().ChatID, "Your opponent surrendered. You win!")
		return
	case "Pass":
		m.Game.PassCount++
		if m.Game.PassCount >= 3 {
			m.Game.Status = entities.GameStatusFinished
			m.ReplyText = "Game over. Both players passed 3 times."
			m.OpponentReplyMessage = tgbotapi.NewMessage(m.getRealOpponent().ChatID, "Game over. Both players passed 3 times.")
		} else {
			m.ReplyText = "You passed your turn. Now it's opponent's turn."
			m.Game.ToggleIsPlayerTurn()
			m.Db.Save(&m.Game)
		}
		return
	}

	if m.Game.Status == entities.GameStatusFinished {
		if m.Message.Text == "/delete" {
			m.Db.Delete(&m.Game)
			m.ReplyText = "Game deleted."
			m.Player.ChangeMenu(MenuNameMyGames)
		} else {
			m.ReplyText = "Game is already finished.\n\n/delete to delete the game."
		}
		return
	}

	if m.isNotMyTurn() {
		m.ReplyText = "Now is opponent's turn"
		return
	}

	m.Game.PassCount = 0
	runeRow, intColumn, err := m.validateMove()
	if err != nil {
		m.ReplyText = "Wrong move: " + err.Error()
		return
	}

	if err = m.placeStone(runeRow, intColumn); err != nil {
		m.ReplyText = "Wrong move: " + err.Error()
		return
	}

	if err = m.Game.SetDots(m.goban.GetDots()); err != nil {
		m.ReplyText = "Cannot take your move"
		return
	}

	m.Game.ToggleIsPlayerTurn()
	m.Game.LastStonePosition = m.goban.GetLast()
	m.Db.Save(m.Game)
	m.ReplyText = m.getPlacedDotEmoji(false) + " Successfully placed stone. Now it's opponent's turn."

	realOpponent := m.getRealOpponent()
	if realOpponent.LastMenu == MenuNameGame+":"+strconv.Itoa(int(m.Game.ID)) {
		opponentGoban := m.goban.Clone()
		opponentGoban.ChangeTheme(models.CreateGobanThemeById(realOpponent.ThemeId))
		photoMessage := tgbotapi.NewPhoto(realOpponent.ChatID, m.getImageForGoban(opponentGoban))
		photoMessage.Caption = m.getPlacedDotEmoji(true) + " Now your turn"
		m.OpponentReplyMessage = photoMessage
	} else {
		m.OpponentReplyMessage = tgbotapi.NewMessage(realOpponent.ChatID, "Your opponent made a move in game "+strconv.Itoa(int(m.Game.ID)))
	}
}

func (m *MenuGame) placeStone(runeRow rune, intColumn uint8) error {
	if m.isPlaceBlack() {
		return m.goban.PlaceBlack(runeRow, intColumn)
	}
	return m.goban.PlaceWhite(runeRow, intColumn)
}

func (m *MenuGame) GetOpponentMessage() tgbotapi.Chattable { return m.OpponentReplyMessage }

func (m *MenuGame) isNotMyTurn() bool {
	isPlayer := m.Game.PlayerChatID == m.Message.Chat.ID
	return (isPlayer && m.Game.IsPlayerTurn) || (!isPlayer && !m.Game.IsPlayerTurn)
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
