package menus

import (
	"ento-go/src/common"
	"ento-go/src/entities"
	"ento-go/src/models"
	"ento-go/src/models/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strconv"
)

const MenuNameGame = "game"

type MenuGame struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	Game      *entities.Game
	ReplyText string
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

	return &menu
}

func (m *MenuGame) GetReplyText() string {
	return "game #" + strconv.Itoa(int(m.Game.ID))
}

func (m *MenuGame) GetReplyImage() *tgbotapi.FileBytes {
	goban := models.NewGobanBySize(m.Game.Size)
	goban.SetDots(m.Game.GetDots())

	img := goban.GetImage()

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

	if m.isMyTurn() {
		m.ReplyText = "You are my turn"
	} else {
		m.ReplyText = "It is not your turn"
	}
}

func (m *MenuGame) GetOpponentMessage() *tgbotapi.MessageConfig {
	return nil
}

func (m *MenuGame) isMyTurn() bool {
	if m.Game.Player.ChatID == m.Game.PlayerChatID {
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
