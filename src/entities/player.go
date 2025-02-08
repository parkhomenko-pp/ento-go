package entities

type Player struct {
	ChatID     int64 `gorm:"primaryKey;unique"`
	LastMenu   string
	Nickname   string
	ThemeId    uint8
	GamesCount int
	WinsCount  int
}

func (p Player) isNew() bool {
	return p.Nickname == ""
}

func NewPlayer(chatID int64) *Player {
	return &Player{
		ChatID:     chatID,
		LastMenu:   "",
		Nickname:   "",
		ThemeId:    1,
		GamesCount: 0,
		WinsCount:  0,
	}
}
