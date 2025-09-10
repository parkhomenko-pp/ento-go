package entities

type Player struct {
	ChatID     int64 `gorm:"primaryKey;unique"`
	LastMenu   string
	Nickname   string
	ThemeId    uint8
	GamesCount int
	WinsCount  int
}

func (p *Player) isNew() bool {
	return p.Nickname == ""
}

func NewPlayer(chatID int64) *Player {
	return &Player{
		ChatID:     chatID,
		LastMenu:   "",
		Nickname:   "",
		ThemeId:    0,
		GamesCount: 0,
		WinsCount:  0,
	}
}

func (p *Player) GetWinRate() float64 {
	if p.GamesCount == 0 {
		return 0
	}
	return float64(p.WinsCount) / float64(p.GamesCount) * 100
}

func (p *Player) GetAvailableThemes() []struct {
	ID   uint8
	Name string
} {
	themes := []struct {
		ID   uint8
		Name string
	}{
		{0, "light"},
		{1, "dark"},
		{2, "wood_light"},
		{3, "wood_dark"},
		{4, "blue_light"},
		{5, "blue_dark"},
	}
	return themes
}

func (p *Player) GetAvailableThemesIds() []uint8 {
	themes := p.GetAvailableThemes()
	ids := make([]uint8, 0, len(themes))
	for _, theme := range themes {
		ids = append(ids, theme.ID)
	}
	return ids
}

func (p *Player) SetThemeByName(themeName string) bool {
	themes := p.GetAvailableThemes()
	for _, theme := range themes {
		if theme.Name == themeName {
			p.ThemeId = theme.ID
			return true
		}
	}
	return false
}

func (p *Player) GetThemeName() string {
	themes := p.GetAvailableThemes()
	for _, theme := range themes {
		if theme.ID == p.ThemeId {
			return theme.Name
		}
	}
	return "error"
}

func (p *Player) ChangeMenu(menu string) {
	p.LastMenu = menu
}

func (p *Player) ChangeMenuWithAdditional(menu string, additional string) {
	menu = menu + ":" + additional
	p.ChangeMenu(menu)
}
