package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strconv"
)

const MenuNameInvited = "invited"

type MenuInvited struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB
	Game    *entities.Game

	ReplyText string
}

func (m *MenuInvited) GetName() string {
	return MenuNameInvited
}

func (m *MenuInvited) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMyGames},
	}
}

func NewMenuInvited(message *tgbotapi.Message, player *entities.Player, db *gorm.DB, additional string) *MenuInvited {
	menu := &MenuInvited{
		Message: message,
		Player:  player,
		Db:      db,
	}
	gameId, _ := strconv.Atoi(additional)
	db.Preload("Opponent").Preload("Player").Where("id = ?", gameId).First(&menu.Game)
	if menu.Game == nil || menu.Game.ID == 0 {
		menu.ReplyText = "You have no invitations"
		player.ChangeMenu(MenuNameMyGames)
		db.Save(player)
		return menu
	}
	menu.ReplyText = "You have an invitation from " + menu.Game.Player.Nickname + "\n" +
		"To accept, type /accept\n" +
		"To decline, type /decline"
	return menu
}

func (m *MenuInvited) DoAction() {
	if m.Message.Text == "/accept" {
		m.Game.Status = entities.GameStatusPlaying
		m.Db.Save(m.Game)
		m.ReplyText = "You accepted the invitation from " + m.Game.Player.Nickname
		m.Player.ChangeMenuWithAdditional(MenuNameGame, strconv.Itoa(int(m.Game.ID)))
		m.Db.Save(m.Player)
	} else if m.Message.Text == "/decline" {
		m.Game.Status = entities.GameStatusDeclined
		m.Db.Save(m.Game)
		m.ReplyText = "You declined the invitation from " + m.Game.Player.Nickname
		m.Player.ChangeMenu(MenuNameMyGames)
		m.Db.Save(m.Player)
	}
}

func (m *MenuInvited) GetReplyText() string {
	return m.ReplyText
}

func (m *MenuInvited) GetReplyImage() *tgbotapi.FileBytes {
	return nil
}

func (m *MenuInvited) GetOpponentMessage() tgbotapi.Chattable {
	return nil
}

func (m *MenuInvited) IsConcatReply() bool {
	return false
}
