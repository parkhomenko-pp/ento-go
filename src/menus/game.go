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
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	Game                 *entities.Game
	ReplyText            string
	OpponentReplyMessage tgbotapi.Chattable

	goban      *models.Goban
	replyImage *tgbotapi.FileBytes
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
		Preload("Opponent").
		Preload("Player").
		Where("id = ?", gameId).
		First(&menu.Game)

	menu.goban = models.NewGobanBySize(menu.Game.Size)
	menu.goban.SetDots(menu.Game.GetDots())
	menu.goban.SetLast(menu.Game.LastStonePosition)

	return &menu
}

func (m *MenuGame) GetReplyText() string {
	return m.ReplyText
}

func (m *MenuGame) GetReplyImage() *tgbotapi.FileBytes {
	if m.replyImage != nil {
		return m.replyImage
	}

	img := m.goban.GetImage()

	byteImage, err := common.EncodeImageToPNGBytes(*img)

	if err != nil {
		return nil
	}

	fileImage := &tgbotapi.FileBytes{
		Name:  "goban.png",
		Bytes: byteImage,
	}

	m.replyImage = fileImage

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
		m.ReplyText = "Wrong move: " + err.Error()
		return
	}

	if m.isPlaceBlack() {
		err = m.goban.PlaceBlack(runeRow, intColumn)
	} else {
		err = m.goban.PlaceWhite(runeRow, intColumn)
	}

	if err != nil {
		m.ReplyText = "Wrong move: " + err.Error()
		return
	}

	err = m.Game.SetDots(m.goban.GetDots())
	if err != nil {
		m.ReplyText = "Cannot take your move"
		return
	}
	m.Game.ToggleIsPlayerTurn()
	m.Game.LastStonePosition = m.goban.GetLast()
	m.Db.Save(m.Game)
	m.ReplyText = "Successfully placed stone. Now it's opponent's turn."

	realOpponent := m.getRealOpponent()

	if realOpponent.LastMenu == MenuNameGame+":"+strconv.Itoa(int(m.Game.ID)) {
		photoMessage := tgbotapi.NewPhoto(realOpponent.ChatID, m.GetReplyImage())
		photoMessage.Caption = "Now your turn"
		m.OpponentReplyMessage = photoMessage
	} else {
		m.OpponentReplyMessage = tgbotapi.NewMessage(
			realOpponent.ChatID,
			"Your opponent made a move in game "+strconv.Itoa(int(m.Game.ID)),
		)
	}
}

func (m *MenuGame) GetOpponentMessage() tgbotapi.Chattable {
	return m.OpponentReplyMessage
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
		return 0, 0, fmt.Errorf("invalid column")
	}

	return runeRow, uint8(intColumn), nil
}

func (m *MenuGame) isPlaceBlack() bool {
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

func (m *MenuGame) getRealOpponent() *entities.Player {
	if m.Game.PlayerChatID == m.Message.Chat.ID {
		return &m.Game.Opponent
	} else {
		return &m.Game.Player
	}
}
