package menus

import (
	"ento-go/src/entities"
	"ento-go/src/models/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"regexp"
)

const MenuNameSettingsChangeTheme = "settings-change_theme"

type MenuSettingsChangeTheme struct {
	Message *tgbotapi.Message
	Player  *entities.Player
	Db      *gorm.DB

	replyText string
	isConcat  bool
}

func (m *MenuSettingsChangeTheme) GetName() string {
	return MenuNameSettingsChangeTheme
}

func (m MenuSettingsChangeTheme) GetNavigation() []types.KeyboardButton {
	return []types.KeyboardButton{
		{Text: "< Back", Destination: MenuNameSettings},
	}
}

func (m *MenuSettingsChangeTheme) DoAction() {
	re := regexp.MustCompile(`/([a-zA-Z_]+)`)
	match := re.FindStringSubmatch(m.Message.Text)
	if len(match) < 2 {
		m.replyText = "Please select a theme from the list below"
		return
	}

	themeName := match[1]

	if m.Player.SetThemeByName(themeName) {
		m.Player.ChangeMenu(MenuNameSettings)
		m.isConcat = true
		m.Db.Save(m.Player)
		m.replyText = "Theme changed to " + m.Player.GetThemeName()
	} else {
		m.replyText = "Selected theme does not exist. Please choose a valid theme."
	}
}

func (m *MenuSettingsChangeTheme) GetReplyText() string {
	if m.replyText == "" {
		m.replyText = "Available themes:\n\n"
		for _, s := range m.Player.GetAvailableThemes() {
			m.replyText += "/" + s.Name

			if s.ID == m.Player.ThemeId {
				m.replyText += " (current)"
			}

			m.replyText = m.replyText + "\n"
		}

		m.replyText += "\nTo change theme, send (or tap) the command with the theme name, e.g. /dark"
	}

	return m.replyText
}

func (m *MenuSettingsChangeTheme) GetReplyImage() *tgbotapi.FileBytes {
	return nil
}

func (m *MenuSettingsChangeTheme) GetOpponentMessage() tgbotapi.Chattable {
	return nil
}

func (m *MenuSettingsChangeTheme) IsConcatReply() bool {
	return m.isConcat
}
