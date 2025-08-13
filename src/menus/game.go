package menus

import (
	"ento-go/src/common"
	"ento-go/src/entities"
	"ento-go/src/models"
	"ento-go/src/models/types"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"log"
	"strconv"
)

const MenuNameGame = "game"

type MenuGame struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	Game      *entities.Game
	ReplyText string

	goban *models.Goban
}

func (m *MenuGame) GetName() string {
	return MenuNameGame
}

func (m *MenuGame) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMyGames},
		{Text: "Help"}, // TODO: add help message
	}
}

func NemMenuGame(message *tgbotapi.Message, player *entities.Player, db *gorm.DB, additional string) *MenuGame {
	gameId := 0
	gameId, _ = strconv.Atoi(additional)

	menu := MenuGame{
		Message: message,
		Player:  player,
		Db:      db,
	}

	menu.Db.
		Where("id = ?", gameId).
		First(&menu.Game)

	menu.goban = models.NewGobanBySize(menu.Game.Size)
	menu.goban.SetDots(menu.Game.GetDots())
	menu.goban.SetLastColor(1)

	return &menu
}

func (m *MenuGame) GetReplyText() string {
	return m.ReplyText
}

func (m *MenuGame) GetReplyImage() *tgbotapi.FileBytes {
	img := m.goban.GetImage()

	byteImage, err := common.EncodeImageToPNGBytes(*img)

	if err != nil {
		return nil
	}

	fileImage := &tgbotapi.FileBytes{
		Name:  "goban.png",
		Bytes: byteImage,
	}

	return fileImage
}

func (m *MenuGame) IsConcatReply() bool {
	return false
}

func (m *MenuGame) DoAction() {
	if m.Message.Text == "Help" {
		m.ReplyText = "Help message"
		return
	}

	if m.isNotMyTurn() {
		m.ReplyText = "Now is opponent's turn"
		return
	}

	runeRow, intColumn, err := m.validateMove()
	if err != nil {
		m.ReplyText = "Wrong move"
		return
	}

	if m.isPlaceBlack() {
		err = m.goban.PlaceBlack(runeRow, intColumn)
	} else {
		err = m.goban.PlaceWhite(runeRow, intColumn)
	}

	if err != nil {
		m.ReplyText = "Wrong move"
		return
	}

	err = m.Game.SetDots(m.goban.GetDots())
	if err != nil {
		m.ReplyText = "Cannot take your move"
		return
	}
	m.Game.ToggleIsPlayerTurn()
	m.Db.Save(m.Game)
	m.ReplyText = "Successfully placed stone. Now it's opponent's turn."
}

func (m *MenuGame) GetOpponentMessage() *tgbotapi.MessageConfig {
	return nil
}

func (m *MenuGame) isNotMyTurn() bool {
	if m.Game.PlayerChatID == m.Message.Chat.ID {
		if m.Game.IsPlayerTurn {
			return true
		} else {
			return false
		}
	} else {
		if m.Game.IsPlayerTurn {
			return false
		} else {
			return true
		}
	}
}

func (m *MenuGame) validateMove() (rune, uint8, error) {
	messageText := m.Message.Text
	if messageText == "" {
		return 0, 0, fmt.Errorf("message is empty")
	}

	runeRow := []rune(messageText)[0]
	intColumn, err := strconv.Atoi(messageText[1:])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid column: %v", err)
	}

	return runeRow, uint8(intColumn), nil
}

func (m *MenuGame) isPlaceBlack() bool {
	log.Println(m.Game.PlayerChatID, m.Message.Chat.ID, m.Game.IsPlayerBlack)
	if m.Game.PlayerChatID == m.Message.Chat.ID {
		if m.Game.IsPlayerBlack {
			return true
		} else {
			return false
		}
	} else {
		if m.Game.IsPlayerBlack {
			return false
		} else {
			return true
		}
	}
}
