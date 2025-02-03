package models

type Player struct {
	ID         int `gorm:"primaryKey"`
	Nickname   string
	ChatID     int64
	ThemeId    uint8
	GamesCount int
	WinsCount  int
}
