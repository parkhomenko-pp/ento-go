package entities

type Player struct {
	ChatID        int64 `gorm:"primaryKey;unique"`
	LastMenu      string
	IsMenuVisited bool
	Nickname      string
	ThemeId       uint8
	GamesCount    int
	WinsCount     int
}

func (p *Player) isNew() bool {
	return p.Nickname == ""
}

func NewPlayer(chatID int64) *Player {
	return &Player{
		ChatID:        chatID,
		LastMenu:      "registration",
		IsMenuVisited: false,
		Nickname:      "",
		ThemeId:       0,
		GamesCount:    0,
		WinsCount:     0,
	}
}

func (p *Player) GetWinRate() float64 {
	if p.GamesCount == 0 {
		return 0
	}
	return float64(p.WinsCount) / float64(p.GamesCount) * 100
}

func (p *Player) ChangeMenu(menu string) {
	p.LastMenu = menu
	p.IsMenuVisited = false
}
