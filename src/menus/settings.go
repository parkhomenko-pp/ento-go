package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

const MenuNameSettings = "settings"

type MenuSettings struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	replyText string
}

func (m *MenuSettings) GetName() string {
	return MenuNameSettings
}

func (m *MenuSettings) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameMain},
		{Text: "Change theme", Destination: MenuNameSettingsChangeTheme},
	}
}

func (m *MenuSettings) DoAction() {

}

func (m *MenuSettings) GetReplyText() string {
	if m.replyText == "" {
		m.replyText = "Current settings:\n\n" +
			"Theme: " + m.Player.GetThemeName()
	}

	return m.replyText
}

func (m *MenuSettings) GetReplyImage() *tgbotapi.FileBytes {
	return nil
}

func (m *MenuSettings) GetOpponentMessage() tgbotapi.Chattable {
	return nil
}

func (m *MenuSettings) IsConcatReply() bool {
	return false
}
