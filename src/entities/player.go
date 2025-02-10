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

func (p Player) isNew() bool {
	return p.Nickname == ""
}

func NewPlayer(chatID int64) *Player {
	return &Player{
		ChatID:        chatID,
		LastMenu:      "",
		IsMenuVisited: false,
		Nickname:      "",
		ThemeId:       0,
		GamesCount:    0,
		WinsCount:     0,
	}
}
