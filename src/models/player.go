package models

type Player struct {
	ChatID     int64 `gorm:"primaryKey;unique"`
	Nickname   string
	ThemeId    uint8
	GamesCount int
	WinsCount  int
}
