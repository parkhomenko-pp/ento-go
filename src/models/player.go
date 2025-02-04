package models

type Player struct {
	ChatID     int64 `gorm:"primaryKey;unique"`
	LastState  string
	Nickname   string
	ThemeId    uint8
	GamesCount int
	WinsCount  int
}

func NewPlayer(chatID int64, nickname string) Player {
	return Player{
		ChatID:     chatID,
		LastState:  StateMainMenu,
		Nickname:   nickname,
		ThemeId:    light,
		GamesCount: 0,
		WinsCount:  0,
	}
}
