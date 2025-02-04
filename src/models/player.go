package models

type Player struct {
	ChatID     int64 `gorm:"primaryKey;unique"`
	Nickname   string
	ThemeId    uint8
	GamesCount int
	WinsCount  int
}

func NewPlayer(chatID int64, nickname string) Player {
	return Player{ChatID: chatID, Nickname: nickname, ThemeId: light, GamesCount: 0, WinsCount: 0}
}
